package handle

import (
	"be/common"
	xe "be/common/error"
	"be/common/log"
	"be/model"
	"be/options"
	"be/session"
	"be/structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func tokenValidation(req *http.Request) error {
	token, err := session.CM.Get("token", req)
	if err != nil || token == "" {
		return xe.AuthError()
	}
	_, err = model.Auth.GetUserInfoByToken(token)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("根据token获取用户信息失败")
		return xe.AuthError()
	}
	return nil
}

func ajaxLogout(res http.ResponseWriter, req *http.Request) {
	session.CM.Remove("token", res)
	http.Redirect(res, req, "/ht", http.StatusTemporaryRedirect)
}

func ajaxGenTokenByUMAndPassword(res http.ResponseWriter, req *http.Request) {
	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		common.ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	request := &Request{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	token, err := model.Auth.GenTokenByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("GenToken失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	// 在session中记录token
	session.CM.Set("token", token, res)
	common.ResSuccessMsg(res, 200, "token生成成功")
}

func ajaxGetUserInfo(res http.ResponseWriter, req *http.Request) {
	token, err := session.CM.Get("token", req)
	if err != nil || token == "" {
		common.ResMsg(res, 400, "请求中未包含token")
		return
	}
	userInfo, err := model.Auth.GetUserInfoByToken(token)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("根据token获取用户信息失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(userInfo)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("根据token获取用户信息失败 JSON生成失败")
		common.ResMsg(res, 500, err.Error())
		return
	}
	common.ResMsg(res, 200, string(b))
}

func ajaxListArticlesByFilter(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}
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
		common.ResMsg(res, 400, err.Error())
		return
	}

	articles, err := model.Article.ListArticles(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	b, err := json.Marshal(articles)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		common.ResMsg(res, 500, err.Error())
		return
	}
	common.ResMsg(res, 200, string(b))
}

func ajaxCreateArticle(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		common.ResInvalidRequestBody(res)
		return
	}

	request := &structs.CreateArticleRequest{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("解析模板JSON失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	err = model.Article.CreateArticle(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	common.ResSuccessMsg(res, 200, "操作成功")
}

func ajaxDeleteArticle(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		common.ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		ArticleId int64 `json:"articleId"`
	}

	request := &Request{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("解析模板JSON失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	err = model.Article.DeleteArticle(request.ArticleId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	common.ResSuccessMsg(res, 200, "操作成功")
}

func ajaxListDeletedArticles(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	articles, err := model.Article.ListDeletedArticles()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	b, err := json.Marshal(articles)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		common.ResMsg(res, 500, err.Error())
		return
	}
	common.ResMsg(res, 200, string(b))
}

func ajaxUpdateArticle(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		common.ResInvalidRequestBody(res)
		return
	}

	request := &structs.UpdateArticleRequest{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("解析模板JSON失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	err = model.Article.UpdateArticle(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	common.ResSuccessMsg(res, 200, "操作成功")
}

func ajaxUploadPic(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	req.ParseMultipartForm(32 << 20)
	filename := req.FormValue("xksj-filename")
	file, handler, err := req.FormFile("xksj-file")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("获取文件失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	defer file.Close()

	// 获取后缀
	oriFileNames := strings.Split(handler.Filename, ".")
	if len(oriFileNames) < 2 {
		log.WithFields(log.Fields{
			"name": handler.Filename,
		}).Error("原文件不包含后缀，无法推测文件类型")
		common.ResMsg(res, 400, "原文件不包含后缀，无法推测文件类型")
		return
	}
	picType := strings.ToLower(oriFileNames[len(oriFileNames)-1])
	if common.StringInSlice(picType, []string{"jpg", "jpeg", "png", "gif"}) == false {
		log.WithFields(log.Fields{
			"picType": picType,
		}).Error("图片类型无效")
		common.ResMsg(res, 400, "图片类型无效")
		return
	}

	if picType == "jpg" {
		picType = "jpeg"
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		common.ResMsg(res, 400, err.Error())
		return
	}

	err = model.Pic.UploadPic(buf.Bytes(), filename, picType)
	if err != nil {
		common.ResMsg(res, 400, err.Error())
		return
	}
	common.ResMsg(res, 200, "操作成功")
}

func ajaxDownloadPic(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	location := ""

	variables := mux.Vars(req)
	if _, ok := variables["location"]; !ok {
		log.WithFields(log.Fields{}).Error("location不在URL中")
		common.ResMsg(res, 400, "未提供location")
		return
	}
	location = variables["location"]

	data, picName, picType, err := model.Pic.DownloadPic(location)
	if err != nil {
		common.ResMsg(res, 400, err.Error())
		return
	}
	res.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", picName))
	res.Header().Set("Content-Type", fmt.Sprintf("image/%s", picType))
	res.Write(data)
}

func ajaxListPics(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	pics, err := model.Pic.ListPics()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	for _, pic := range pics {
		pic.Location, _ = model.Pic.TranslatePic(pic.Name)
		pic.Location = fmt.Sprintf("%s/%s", options.Options.PicExternalRootPath, pic.Location)
	}

	b, err := json.Marshal(pics)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		common.ResMsg(res, 500, err.Error())
		return
	}
	common.ResMsg(res, 200, string(b))
}

func ajaxDeletePic(res http.ResponseWriter, req *http.Request) {
	if err := tokenValidation(req); err != nil {
		common.ResMsg(res, 400, "认证失败")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		common.ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Id int64 `json:"id"`
	}

	request := &Request{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("解析模板JSON失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	err = model.Pic.DeletePic(request.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("model调用失败")
		common.ResMsg(res, 400, err.Error())
		return
	}
	common.ResSuccessMsg(res, 200, "操作成功")
}
