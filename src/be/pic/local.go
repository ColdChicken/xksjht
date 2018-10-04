package pic

import (
	"be/common"
	"be/common/log"
	"be/dao"
	"be/options"
	"be/structs"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type LocalPicMgr struct {
	dao *dao.PicDao
}

func NewLocalPicMgr() *LocalPicMgr {
	return &LocalPicMgr{
		dao: &dao.PicDao{},
	}
}

func (m *LocalPicMgr) UploadPic(data []byte, name string, picType string) error {
	// 创建目录
	var subPath = time.Now().Format("20060102150405.999999")
	var fullPath = path.Join(options.Options.LocalPicRootPath, subPath)

	err := common.Mkdir(fullPath)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("创建目录失败")
		return err
	}

	// 存储文件
	f, err := os.OpenFile(path.Join(fullPath, name), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("创建文件失败")
		return nil
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("写入文件失败")
		return nil
	}

	f.Sync()

	// 记录信息
	err = m.dao.SavePicRecord(subPath, name, picType)
	if err != nil {
		return err
	}

	return nil
}

// location就是path
func (m *LocalPicMgr) getNameByLocation(location string) string {
	return location
}

func (m *LocalPicMgr) DownloadPic(location string) ([]byte, string, string, error) {
	name := m.getNameByLocation(location)
	path_, picType, err := m.dao.GetPicInfoByName(name)
	if err != nil {
		return nil, "", "", err
	}
	data, err := ioutil.ReadFile(path.Join(options.Options.LocalPicRootPath, path_, name))
	if err != nil {
		return nil, "", "", err
	}
	return data, name, picType, nil
}

func (m *LocalPicMgr) ListPics() ([]*structs.PicInfo, error) {
	pics, err := m.dao.ListPics()
	if err != nil {
		return nil, err
	}
	return pics, nil
}

func (m *LocalPicMgr) TranslatePic(name string) (string, error) {
	return name, nil
}

func (m *LocalPicMgr) GetPic(location string) (*structs.PicInfo, error) {
	name := m.getNameByLocation(location)
	_, picType, err := m.dao.GetPicInfoByName(name)
	if err != nil {
		return nil, err
	}
	return &structs.PicInfo{
		Name:     name,
		Location: location,
		Type:     picType,
	}, nil
}

func (m *LocalPicMgr) DeletePic(id int64) error {
	return m.dao.DeletePic(id)
}
