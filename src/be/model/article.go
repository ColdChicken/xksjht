package model

import (
	xe "be/common/error"
	"be/common/log"
	"be/dao"
	"be/options"
	"be/parser"
	"be/structs"
	"encoding/json"
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

func (m *ArticleMgr) ListArticles(filter *structs.ListArticleFilter) ([]*structs.Article, error) {
	articles := []*structs.Article{}
	articles, err := m.dao.ListArticlesByFilter(filter)
	if err != nil {
		log.Errorf("ListArticles 失败： %s", err.Error())
		return nil, err
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
	err = m.dao.CreateArticle(title, request.Creater, request.Tags, request.OriginalTag, string(b), request.RawContent)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}
