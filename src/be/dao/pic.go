package dao

import (
	xe "be/common/error"
	"be/common/log"
	"be/mysql"
	"be/structs"
)

type PicDao struct {
}

func (d *PicDao) SavePicRecord(path string, name string, picType string) error {
	err := mysql.DB.SimpleInsert("INSERT INTO PIC_LOCAL_RECORD(name, relPath, createTime, isDeleted, picType) VALUES(?, ?, NOW(), 0, ?)", name, path, picType)
	if err != nil {
		log.Errorf("SavePicRecord 失败: %s", err.Error())
		return err
	}
	return nil
}

func (d *PicDao) GetPicInfoByName(name string) (string, string, error) {
	path := ""
	picType := ""
	cnt, err := mysql.DB.SingleRowQuery("SELECT relPath, picType FROM PIC_LOCAL_RECORD WHERE isDeleted=0 AND name=?", []interface{}{name}, &path, &picType)
	if err != nil {
		log.Errorf("GetPicRelPathByName 失败: %s", err.Error())
		return "", "", err
	}
	if cnt == 0 {
		return "", "", xe.New("图片记录信息不存在或已删除")
	}
	return path, picType, nil
}

func (d *PicDao) DeletePic(id int64) error {
	tx := mysql.DB.GetTx()

	sql := `UPDATE PIC_LOCAL_RECORD SET isDeleted=1 WHERE id=?`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("prepare错误")
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id)
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

func (d *PicDao) ListPics() ([]*structs.PicInfo, error) {
	pics := []*structs.PicInfo{}
	tx := mysql.DB.GetTx()
	sql := `SELECT id, name, relPath, createTime, picType 
			FROM PIC_LOCAL_RECORD WHERE isDeleted=0 
			ORDER BY id DESC`
	if stmt, err := tx.Prepare(sql); err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("ListPics prepare错误")
		tx.Rollback()
		return nil, err
	} else {
		if rows, err := stmt.Query(); err != nil {
			log.WithFields(log.Fields{
				"sql": sql,
				"err": err.Error(),
			}).Error("ListPics prepare错误")
			stmt.Close()
			tx.Rollback()
			return nil, err
		} else {
			for rows.Next() {
				pic := &structs.PicInfo{}
				if err := rows.Scan(&pic.Id, &pic.Name, &pic.RelPath, &pic.CreateTime, &pic.Type); err != nil {
					log.WithFields(log.Fields{
						"sql": sql,
						"err": err.Error(),
					}).Error("ListPics prepare错误")
					rows.Close()
					stmt.Close()
					tx.Rollback()
					return nil, err
				} else {
					pics = append(pics, pic)
				}
			}
			rows.Close()
			stmt.Close()
		}
	}
	tx.Commit()
	return pics, nil
}
