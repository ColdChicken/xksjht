package model

import (
	"be/common/log"
	"be/dao"
	"be/structs"
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
