package model

import "be/structs"

type NewsS struct {
}

var News *NewsS

func init() {
	News = &NewsS{}
}

func (m *NewsS) ListNews(pos int64, size int64) ([]*structs.News, error) {
	news := []*structs.News{}

	return news, nil
}
