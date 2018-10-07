package model

import (
	"be/common"
	xe "be/common/error"
	"be/common/log"
	"be/dao"
	"be/options"
	"be/parser"
	"be/structs"
	"encoding/json"
	"fmt"
	"strings"
)

type ArticleMgr struct {
	dao *dao.ArticleDao
}

var Article *ArticleMgr

func init() {
	Article = &ArticleMgr{
		dao: &dao.ArticleDao{},
	}
}

func (m *ArticleMgr) GetArticle(id int64) (*structs.Article, error) {
	article, err := m.dao.GetArticleById(id)
	if err != nil {
		return nil, err
	}
	if article.Id == -1 {
		return nil, xe.New("文章不存在")
	}

	// 解析及处理图片信息
	content := &structs.ParserResult{}
	if err := common.ParseJsonStr(article.Content, &content); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("解析模板JSON失败")
		return nil, err
	}

	refIdx := 1
	for _, c := range content.Contents {
		if c.Type == "ref" && c.Source == "img" {
			picLocation, err := Pic.TranslatePic(c.Value)
			if err != nil {
				log.WithFields(log.Fields{
					"err": err.Error(),
				}).Error("获取图片路径信息失败")
				c.Value = fmt.Sprintf("%s/%s", options.Options.PicExternalRootPath, "404")
			} else {
				c.Value = fmt.Sprintf("%s/%s", options.Options.PicExternalRootPath, picLocation)
			}

		}
		if c.Type == "section" {
			c.SectionLevel = int64(len(strings.Split(c.SectionID, ".")))
		}
		if c.Type == "block" {
			for _, block := range c.Contents {
				if block.Type == "block_ref" {
					block.RefIdx = int64(refIdx)
					refIdx++
				}
			}
		}
	}

	b, err := json.Marshal(content)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		return article, nil
	} else {
		article.Content = string(b)
	}

	article.ParsedContent = content

	return article, nil
}

func (m *ArticleMgr) ListArticles(filter *structs.ListArticleFilter) ([]*structs.Article, error) {
	articles := []*structs.Article{}
	articles, err := m.dao.ListArticlesByFilter(filter)
	if err != nil {
		log.Errorf("ListArticles 失败： %s", err.Error())
		return nil, err
	}
	// 是否要返回文章内容，对于小程序/前台等服务来说不返回这个可以减少网络开销
	if filter.ContainContent == 0 {
		for _, article := range articles {
			article.Content = ""
			article.RawContent = ""
		}
	}
	return articles, nil
}

func (m *ArticleMgr) CreateArticle(request *structs.CreateArticleRequest) error {
	if request.Creater == "" {
		request.Creater = options.Options.DefaultArticleCreater
	}
	if request.OriginalTag != 0 && request.OriginalTag != 1 {
		return xe.New("原创标签无效")
	}
	p := parser.NewParser()
	err := p.Parser(request.RawContent)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	articleInfo := p.GetResult()
	b, err := json.MarshalIndent(articleInfo, "", "    ")
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	title := articleInfo.Title
	if title == "" {
		return xe.New("文章标题不存在")
	}
	err = m.dao.CreateArticle(title, request.Creater, request.Tags, request.OriginalTag, string(b), request.RawContent, request.Catalog)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *ArticleMgr) UpdateArticle(request *structs.UpdateArticleRequest) error {
	if request.OriginalTag != 0 && request.OriginalTag != 1 {
		return xe.New("原创标签无效")
	}
	p := parser.NewParser()
	err := p.Parser(request.RawContent)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	articleInfo := p.GetResult()
	b, err := json.MarshalIndent(articleInfo, "", "    ")
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	title := articleInfo.Title
	if title == "" {
		return xe.New("文章标题不存在")
	}

	err = m.dao.EditArticle(request.ArticleId, title, request.Tags, request.OriginalTag, string(b), request.RawContent)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (m *ArticleMgr) DeleteArticle(articleId int64) error {
	return m.dao.DeleteArticle(articleId)
}

func (m *ArticleMgr) ListDeletedArticles() ([]*structs.Article, error) {
	return m.dao.ListDeletedArticles()
}
