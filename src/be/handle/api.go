package handle

import (
	"be/common"
	xe "be/common/error"
	"be/common/log"
	"be/model"
	"be/structs"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func apiListArticles(res http.ResponseWriter, req *http.Request) {
	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		common.ResInvalidRequestBody(res)
		return
	}

	request := &structs.ListArticleFilter{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("解析模板JSON失败")
		common.ResMsg(res, 400, xe.HandleRequestError().Error())
		return
	}

	articles, err := model.Article.ListArticles(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, xe.HandleRequestError().Error())
		return
	}

	b, err := json.Marshal(articles)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		common.ResMsg(res, 500, xe.HandleRequestError().Error())
		return
	}
	common.ResMsg(res, 200, string(b))
}
