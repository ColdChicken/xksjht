package handle

import (
	"be/common"
	"be/common/log"
	"be/model"
	"be/options"
	"be/structs"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

func templateRealPath(path string) string {
	return filepath.Join(options.Options.TemplateFilePath, path)
}

func showIndexHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles(templateRealPath("index.html"))
	tmpl.ExecuteTemplate(res, "index", nil)
}

func showLoginHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles(templateRealPath("login.html"))
	tmpl.ExecuteTemplate(res, "login", nil)
}

func showZiXunHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(templateRealPath("www_base.html"), templateRealPath("www_articles.html"), templateRealPath("www_article.html"))

	if err != nil {
		log.Errorf("模板解析失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	articles, err := model.Article.ListArticles(&structs.ListArticleFilter{
		CurrentPos:     0,
		RequestCnt:     999999,
		ContainContent: 0,
		Catalog:        "资讯",
	})

	if err != nil {
		log.Errorf("获取文章内容失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	type Data struct {
		PageViewed string
		Articles   []*structs.Article
	}

	data := Data{
		PageViewed: "水景资讯",
		Articles:   articles,
	}
	tmpl.ExecuteTemplate(res, "base", data)
}

func showShuiJingHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(templateRealPath("www_base.html"), templateRealPath("www_articles.html"), templateRealPath("www_article.html"))

	if err != nil {
		log.Errorf("模板解析失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	articles, err := model.Article.ListArticles(&structs.ListArticleFilter{
		CurrentPos:     0,
		RequestCnt:     999999,
		ContainContent: 0,
		Catalog:        "水景",
	})

	if err != nil {
		log.Errorf("获取文章内容失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	type Data struct {
		PageViewed string
		Articles   []*structs.Article
	}

	data := Data{
		PageViewed: "水景欣赏",
		Articles:   articles,
	}
	tmpl.ExecuteTemplate(res, "base", data)
}

func showWenZhangHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(templateRealPath("www_base.html"), templateRealPath("www_articles.html"), templateRealPath("www_article.html"))

	if err != nil {
		log.Errorf("模板解析失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	articles, err := model.Article.ListArticles(&structs.ListArticleFilter{
		CurrentPos:     0,
		RequestCnt:     999999,
		ContainContent: 0,
		Catalog:        "文章",
	})

	if err != nil {
		log.Errorf("获取文章内容失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	type Data struct {
		PageViewed string
		Articles   []*structs.Article
	}

	data := Data{
		PageViewed: "水景文章",
		Articles:   articles,
	}
	tmpl.ExecuteTemplate(res, "base", data)
}

func showQiTaHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(templateRealPath("www_base.html"), templateRealPath("www_articles.html"), templateRealPath("www_article.html"))

	if err != nil {
		log.Errorf("模板解析失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	articles, err := model.Article.ListArticles(&structs.ListArticleFilter{
		CurrentPos:     0,
		RequestCnt:     999999,
		ContainContent: 0,
		Catalog:        "其它",
	})

	if err != nil {
		log.Errorf("获取文章内容失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	type Data struct {
		PageViewed string
		Articles   []*structs.Article
	}

	data := Data{
		PageViewed: "其它文章",
		Articles:   articles,
	}
	tmpl.ExecuteTemplate(res, "base", data)
}

func showArticleHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(templateRealPath("www_base.html"), templateRealPath("www_articles.html"), templateRealPath("www_article.html"))

	if err != nil {
		log.Errorf("模板解析失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	// 获取文章
	variables := mux.Vars(req)
	if _, ok := variables["id"]; !ok {
		log.WithFields(log.Fields{}).Error("id不在URL中")
		common.ResMsg(res, 400, "请提供文章id")
		return
	}
	ids := variables["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		log.WithFields(log.Fields{}).Error("id格式异常")
		common.ResMsg(res, 400, "id格式异常")
		return
	}

	article, err := model.Article.GetArticle(id)
	if err != nil {
		log.Errorf("获取文章内容失败: %s", err.Error())
		common.ResMsg(res, 500, "服务异常")
		return
	}

	type Data struct {
		PageViewed string
		Article    *structs.Article
	}

	data := Data{
		PageViewed: "文章内容",
		Article:    article,
	}
	tmpl.ExecuteTemplate(res, "base", data)
}
