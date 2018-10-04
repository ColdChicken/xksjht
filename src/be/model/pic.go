package model

import (
	"be/pic"
	"be/structs"
)

type PicMgr struct {
	mgr PicMgrBase
}

type PicMgrBase interface {
	// 上传图片
	UploadPic(data []byte, name string, picType string) error
	// 下载图片
	DownloadPic(location string) ([]byte, string, string, error)
	// 列出图片
	ListPics() ([]*structs.PicInfo, error)
	// 转换图片，将图片名转换为图片地址
	TranslatePic(name string) (location string, err error)
	// 获取图片信息
	GetPic(location string) (*structs.PicInfo, error)
	// 删除图片
	DeletePic(id int64) error
}

var Pic *PicMgr

func init() {
	Pic = &PicMgr{
		// 使用本地图片管理器
		mgr: pic.NewLocalPicMgr(),
	}
}

func (m *PicMgr) UploadPic(data []byte, name string, picType string) error {
	return m.mgr.UploadPic(data, name, picType)
}

func (m *PicMgr) DownloadPic(location string) ([]byte, string, string, error) {
	return m.mgr.DownloadPic(location)
}

func (m *PicMgr) ListPics() ([]*structs.PicInfo, error) {
	return m.mgr.ListPics()
}

func (m *PicMgr) TranslatePic(name string) (string, error) {
	return m.mgr.TranslatePic(name)
}

func (m *PicMgr) GetPic(location string) (*structs.PicInfo, error) {
	return m.mgr.GetPic(location)
}

func (m *PicMgr) DeletePic(id int64) error {
	return m.mgr.DeletePic(id)
}
