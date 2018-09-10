package dao

import (
	"be/common/log"
	"be/mysql"
	"be/structs"
	"sort"
	"strings"
)

type ArticleDao struct {
}

func (d *ArticleDao) ListArticlesByFilter(filter *structs.ListArticleFilter) ([]*structs.Article, error) {
	articles := []*structs.Article{}

	// 对tag进行排序
	sort.Sort(sort.StringSlice(filter.Tags))
	// 接着用逗号分隔
	tags := strings.Join(filter.Tags, ",")

	tx := mysql.DB.GetTx()

	if tags != "" {
		sql := `SELECT id, title, createTime, editTime, creater, tags, originalTag, content 
				FROM ARTICLE
				WHERE tags=? AND isDeleted=0
				ORDER BY id DESC
				LIMIT ?, ?`
		if stmt, err := tx.Prepare(sql); err != nil {
			log.WithFields(log.Fields{
				"sql": sql,
				"err": err.Error(),
			}).Error("ListArticlesByFilter prepare错误")
			tx.Rollback()
			return nil, err
		} else {
			if rows, err := stmt.Query(tags, filter.CurrentPos, filter.RequestCnt); err != nil {
				log.WithFields(log.Fields{
					"sql": sql,
					"err": err.Error(),
				}).Error("ListArticlesByFilter prepare错误")
				stmt.Close()
				tx.Rollback()
				return nil, err
			} else {
				for rows.Next() {
					article := &structs.Article{}
					if err := rows.Scan(&article.Id, &article.Title, &article.CreateTime, &article.EditTime, &article.Creater, &article.Tags, &article.OriginalTag, &article.Content); err != nil {
						log.WithFields(log.Fields{
							"sql": sql,
							"err": err.Error(),
						}).Error("ListArticlesByFilter prepare错误")
						rows.Close()
						stmt.Close()
						tx.Rollback()
						return nil, err
					} else {
						articles = append(articles, article)
					}
				}
				rows.Close()
				stmt.Close()
			}
		}
	} else {
		sql := `SELECT id, title, createTime, editTime, creater, tags, originalTag, content 
				FROM ARTICLE
				WHERE isDeleted=0
				ORDER BY id DESC
				LIMIT ?, ?`
		if stmt, err := tx.Prepare(sql); err != nil {
			log.WithFields(log.Fields{
				"sql": sql,
				"err": err.Error(),
			}).Error("ListArticlesByFilter prepare错误")
			tx.Rollback()
			return nil, err
		} else {
			if rows, err := stmt.Query(filter.CurrentPos, filter.RequestCnt); err != nil {
				log.WithFields(log.Fields{
					"sql": sql,
					"err": err.Error(),
				}).Error("ListArticlesByFilter prepare错误")
				stmt.Close()
				tx.Rollback()
				return nil, err
			} else {
				for rows.Next() {
					article := &structs.Article{}
					if err := rows.Scan(&article.Id, &article.Title, &article.CreateTime, &article.EditTime, &article.Creater, &article.Tags, &article.OriginalTag, &article.Content); err != nil {
						log.WithFields(log.Fields{
							"sql": sql,
							"err": err.Error(),
						}).Error("ListArticlesByFilter prepare错误")
						rows.Close()
						stmt.Close()
						tx.Rollback()
						return nil, err
					} else {
						articles = append(articles, article)
					}
				}
				rows.Close()
				stmt.Close()
			}
		}
	}

	tx.Commit()
	return articles, nil
}
