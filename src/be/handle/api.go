package handle

import (
	"be/common"
	"be/common/log"
	"be/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
	获取新闻
*/
func ajaxListNews(res http.ResponseWriter, req *http.Request) {
	type Request struct {
		// 当前请求方的位置
		RequestCurrentPos int64 `json:"pos"`
		// 请求方请求的数目
		Size int64 `json:"size"`
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		common.ResInvalidRequestBody(res)
		return
	}

	request := &Request{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	if request.Size <= 0 {
		request.Size = 10
	}

	if request.RequestCurrentPos < 0 {
		request.RequestCurrentPos = 0
	}

	newsInfo, err := model.News.ListNews(request.RequestCurrentPos, request.Size)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("调用model 失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	b, err := json.Marshal(newsInfo)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		common.ResMsg(res, 500, err.Error())
		return
	}
	common.ResMsg(res, 200, string(b))
}
