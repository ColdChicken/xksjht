package handle

import (
	"be/common"
	xe "be/common/error"
	"be/common/log"
	"be/model"
	"be/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func apiGetArticleById(res http.ResponseWriter, req *http.Request) {
	type Request struct {
		Id int64 `json:"id"`
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
		log.WithFields(log.Fields{
			"err":     err.Error(),
			"request": string(reqContent),
		}).Error("解析模板JSON失败")
		common.ResMsg(res, 400, xe.HandleRequestError().Error())
		return
	}

	article, err := model.Article.GetArticle(request.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, xe.HandleRequestError().Error())
		return
	}

	b, err := json.Marshal(article)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		common.ResMsg(res, 500, xe.HandleRequestError().Error())
		return
	}
	common.ResMsg(res, 200, string(b))
}

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
			"err":     err.Error(),
			"request": string(reqContent),
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

func apiDownloadPic(res http.ResponseWriter, req *http.Request) {
	location := ""

	variables := mux.Vars(req)
	if _, ok := variables["location"]; !ok {
		log.WithFields(log.Fields{}).Error("location不在URL中")
		common.ResMsg(res, 400, xe.HandleRequestError().Error())
		return
	}
	location = variables["location"]

	data, picName, picType, err := model.Pic.DownloadPic(location)
	if err != nil {
		log.Errorln(err.Error())
		common.ResMsg(res, 400, xe.HandleRequestError().Error())
		return
	}
	res.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", picName))
	res.Header().Set("Content-Type", fmt.Sprintf("image/%s", picType))
	res.Write(data)
}
