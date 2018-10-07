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

func (d *ArticleDao) ListDeletedArticles() ([]*structs.Article, error) {
	articles := []*structs.Article{}
	tx := mysql.DB.GetTx()

	sql := `SELECT id, title, createTime, editTime, creater, tags, originalTag, content, rawContent, catalog 
			FROM ARTICLE
			WHERE isDeleted=1
			ORDER BY id DESC`
	if stmt, err := tx.Prepare(sql); err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("ListDeletedArticles prepare错误")
		tx.Rollback()
		return nil, err
	} else {
		if rows, err := stmt.Query(); err != nil {
			log.WithFields(log.Fields{
				"sql": sql,
				"err": err.Error(),
			}).Error("ListDeletedArticles prepare错误")
			stmt.Close()
			tx.Rollback()
			return nil, err
		} else {
			for rows.Next() {
				article := &structs.Article{}
				if err := rows.Scan(&article.Id, &article.Title, &article.CreateTime, &article.EditTime, &article.Creater, &article.Tags, &article.OriginalTag, &article.Content, &article.RawContent, &article.Catalog); err != nil {
					log.WithFields(log.Fields{
						"sql": sql,
						"err": err.Error(),
					}).Error("ListDeletedArticles prepare错误")
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
	tx.Commit()
	return articles, nil
}

func (d *ArticleDao) GetArticleById(id int64) (*structs.Article, error) {
	article := &structs.Article{Id: -1}
	tx := mysql.DB.GetTx()
	sql := `SELECT id, title, createTime, editTime, creater, tags, originalTag, content, rawContent, catalog 
				FROM ARTICLE
				WHERE isDeleted=0
				AND id=?`
	if stmt, err := tx.Prepare(sql); err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("GetArticleById prepare错误")
		tx.Rollback()
		return nil, err
	} else {
		if rows, err := stmt.Query(id); err != nil {
			log.WithFields(log.Fields{
				"sql": sql,
				"err": err.Error(),
			}).Error("GetArticleById prepare错误")
			stmt.Close()
			tx.Rollback()
			return nil, err
		} else {
			for rows.Next() {
				if err := rows.Scan(&article.Id, &article.Title, &article.CreateTime, &article.EditTime, &article.Creater, &article.Tags, &article.OriginalTag, &article.Content, &article.RawContent, &article.Catalog); err != nil {
					log.WithFields(log.Fields{
						"sql": sql,
						"err": err.Error(),
					}).Error("ListArticlesByFilter prepare错误")
					rows.Close()
					stmt.Close()
					tx.Rollback()
					return nil, err
				}
			}
			rows.Close()
			stmt.Close()
		}
	}
	tx.Commit()
	return article, nil
}

func (d *ArticleDao) ListArticlesByFilter(filter *structs.ListArticleFilter) ([]*structs.Article, error) {
	articles := []*structs.Article{}
	// 对tag进行排序
	sort.Sort(sort.StringSlice(filter.Tags))
	// 接着用逗号分隔
	tags := strings.Join(filter.Tags, ",")

	tx := mysql.DB.GetTx()

	if tags != "" && filter.Catalog != "" {
		sql := `SELECT id, title, createTime, editTime, creater, tags, originalTag, content, rawContent, catalog 
				FROM ARTICLE
				WHERE tags=? AND isDeleted=0 AND catalog=?
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
			if rows, err := stmt.Query(tags, filter.Catalog, filter.CurrentPos, filter.RequestCnt); err != nil {
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
					if err := rows.Scan(&article.Id, &article.Title, &article.CreateTime, &article.EditTime, &article.Creater, &article.Tags, &article.OriginalTag, &article.Content, &article.RawContent, &article.Catalog); err != nil {
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
	} else if tags == "" && filter.Catalog != "" {
		sql := `SELECT id, title, createTime, editTime, creater, tags, originalTag, content, rawContent, catalog 
				FROM ARTICLE
				WHERE isDeleted=0 AND catalog=?
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
			if rows, err := stmt.Query(filter.Catalog, filter.CurrentPos, filter.RequestCnt); err != nil {
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
					if err := rows.Scan(&article.Id, &article.Title, &article.CreateTime, &article.EditTime, &article.Creater, &article.Tags, &article.OriginalTag, &article.Content, &article.RawContent, &article.Catalog); err != nil {
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
		sql := `SELECT id, title, createTime, editTime, creater, tags, originalTag, content, rawContent, catalog 
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
					if err := rows.Scan(&article.Id, &article.Title, &article.CreateTime, &article.EditTime, &article.Creater, &article.Tags, &article.OriginalTag, &article.Content, &article.RawContent, &article.Catalog); err != nil {
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

func (d *ArticleDao) CreateArticle(title string, creater string, tags string, originalTag int64, content string, rawContent string, catalog string) error {
	tx := mysql.DB.GetTx()

	sql := `INSERT INTO ARTICLE(title, createTime, editTime, 
								creater, tags, originalTag, 
								content, rawContent, catalog, isDeleted) 
			VALUES(?, NOW(), NOW(), ?, ?, ?, ?, ?, ?, 0)`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("prepare错误")
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(title, creater, tags, originalTag, content, rawContent, catalog)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Exec 错误")
		stmt.Close()
		tx.Rollback()
		return err
	}
	stmt.Close()
	tx.Commit()
	return nil
}

func (d *ArticleDao) DeleteArticle(articleId int64) error {
	tx := mysql.DB.GetTx()

	sql := `UPDATE ARTICLE SET isDeleted=1 WHERE id=?`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("prepare错误")
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(articleId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Exec 错误")
		stmt.Close()
		tx.Rollback()
		return err
	}
	stmt.Close()
	tx.Commit()
	return nil
}

func (d *ArticleDao) EditArticle(articleId int64, title string, tags string, originalTag int64, content string, rawContent string) error {
	tx := mysql.DB.GetTx()

	sql := `UPDATE ARTICLE
			SET title=?, editTime=NOW(), tags=?, originalTag=?, content=?, rawContent=?
			WHERE id=?`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("prepare错误")
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(title, tags, originalTag, content, rawContent, articleId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Exec 错误")
		stmt.Close()
		tx.Rollback()
		return err
	}
	stmt.Close()
	tx.Commit()
	return nil
}
