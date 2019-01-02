<template>

<div>
  <template v-if="inLoading === true">
    <div style="height: 100%; width: 100%; text-align: center;">
      <img src="../../assets/loading.gif"/>&nbsp; 加载中，请稍等...
    </div>
  </template>
  <template v-else>
    <div class="tabbable">
      <ul class="nav nav-tabs">
        <li class="active">
          <a data-toggle="tab" href="#article_list">
            文章列表
          </a>
        </li>
        <li>
          <a data-toggle="tab" href="#garbage_list">
            垃圾箱
          </a>
        </li>
      </ul>
      <div class="tab-content">
        <div id="article_list" class="tab-pane fade in active">
          <button class="btn btn-success btn-xs" style="float:right;padding:2px;z-index:99999; margin: 2px;" @click="createArticle">新增文章</button>
          <table class="table table-striped table-bordered table-hover" style="table-layout:fixed;word-wrap:break-word;">
            <thead>
              <tr>
                <th style="width: 5%">ID</th>
                <th style="width: 25%">标题</th>
                <th style="width: 10%">创建人</th>
                <th style="width: 10%">创建时间</th>
                <th style="width: 5%">类别</th>
                <th style="width: 20%">标签</th>
                <th style="width: 5%">原创</th>
				        <th style="width: 30%">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="article in articles" :key="article['id']">
                <td>{{article['id']}}</td>
                <td>{{article['title']}}</td>
                <td>{{article['creater']}}</td>
                <td>{{article['createTime']}}</td>
                <td>{{article['catalog']}}</td>
                <td>{{article['tags']}}</td>
				        <td>{{article['originalTag']}}</td>
                <td><a href="javascript:void(0)" @click="showContent(article['id'])" >内容</a>&nbsp;<a href="javascript:void(0)" @click="deleteArticle(article['id'])" >删除</a>&nbsp;<a href="javascript:void(0)" @click="updateArticle(article['id'])" >编辑</a></td>
              </tr>
            </tbody>
          </table>
        </div>
        <div id="garbage_list" class="tab-pane fade">
           <table class="table table-striped table-bordered table-hover" style="table-layout:fixed;word-wrap:break-word;">
            <thead>
              <tr>
                <th style="width: 5%">ID</th>
                <th style="width: 25%">标题</th>
                <th style="width: 10%">创建人</th>
                <th style="width: 10%">创建时间</th>
                <th style="width: 5%">类别</th>
                <th style="width: 20%">标签</th>
                <th style="width: 5%">原创</th>
				        <th style="width: 30%">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="article in deletedArticles" :key="article['id']">
                <td>{{article['id']}}</td>
                <td>{{article['title']}}</td>
                <td>{{article['creater']}}</td>
                <td>{{article['createTime']}}</td>
                <td>{{article['catalog']}}</td>
                <td>{{article['tags']}}</td>
				        <td>{{article['originalTag']}}</td>
                <td><a href="javascript:void(0)" @click="showContent(article['id'])" >内容</a>&nbsp;<a href="javascript:void(0)" @click="updateArticle(article['id'])" >编辑</a></td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </template>


<template v-if="showArticleRawContentEditDiv === true">
    <div class="modal" style="display: block;" tabindex="9999">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<button type="button" class="close" @click="hideContent">&times;</button>
					<h4 class="blue bigger">请输入进行此操作的相关信息:</h4>
				</div>

				<div class="modal-body">
					<div class="row">
            <div class="col-xs-12 col-sm-9">
              <div class="form-group">
                <label>标签:</label>
                <div>
                      <input type="text" class="col-xs-12 col-sm-9" v-model="articleTags"/>
                </div>
              </div>
              <div style='clear:both'></div>

              <div class="form-group">
                <label>类别:</label>
                <div>
                      <select v-model="articleCatalog">
                        <option value='资讯'>水景星空(资讯)</option>
                        <option value='文章'>星空笔记(文章)</option>
                      </select>
                </div>
              </div>
              <div style='clear:both'></div>

              <div class="form-group">
                <label>原创:</label>
                <div>
                      <select v-model="articleOriginalTag">
                        <option value=0>否</option>
                        <option value=1>是</option>
                      </select>
                </div>
              </div>
              <div style='clear:both'></div>

							<div class="form-group">
							  <label>文章内容:</label>
							  <div>
								   <textarea class="col-xs-12 col-sm-9" v-model="articleRawContent"></textarea>
							  </div>
							</div>

						</div>

					</div>
				</div>

				<div class="modal-footer">
					<button class="btn btn-sm" @click="hideContent">
						<i class="ace-icon fa fa-times"></i>
						取消
					</button>

					<button class="btn btn-sm btn-primary" @click="doCreateArticle">
						<i class="ace-icon fa fa-check"></i>
						确定
					</button>
				</div>
			</div>
		</div>
    </div>
 </template>


<template v-if="showArticleEditDiv === true">
    <div class="modal" style="display: block;" tabindex="9999">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<button type="button" class="close" @click="hideAriticleDiv">&times;</button>
					<h4 class="blue bigger">请输入进行此操作的相关信息:</h4>
				</div>

				<div class="modal-body">
					<div class="row">
            <div class="col-xs-12 col-sm-9">
              <div class="form-group">
                <label>标签:</label>
                <div>
                      <input type="text" class="col-xs-12 col-sm-9" v-model="articleEditTags"/>
                </div>
              </div>
              <div style='clear:both'></div>

              <div class="form-group">
                <label>原创:</label>
                <div>
                      <select v-model="articleEditOriginalTag">
                        <option value=0>否</option>
                        <option value=1>是</option>
                      </select>
                </div>
              </div>
              <div style='clear:both'></div>

							<div class="form-group">
							  <label>文章内容:</label>
							  <div>
								   <textarea class="col-xs-12 col-sm-9" v-model="articleEditRawContent"></textarea>
							  </div>
							</div>

						</div>

					</div>
				</div>

				<div class="modal-footer">
					<button class="btn btn-sm" @click="hideAriticleDiv">
						<i class="ace-icon fa fa-times"></i>
						取消
					</button>

					<button class="btn btn-sm btn-primary" @click="doEditArticle">
						<i class="ace-icon fa fa-check"></i>
						确定
					</button>
				</div>
			</div>
		</div>
    </div>
 </template>



</div>

</template>

<script>
require('jquery')

require('@/assets/bootstrap/css/bootstrap.min.css')
require('@/assets/bootstrap/css/dashboard.css')

require('@/assets/bootstrap/js/bootstrap.min.js')

var bootbox = require('bootbox')

import Ajax from "../../common/ajax"
import Config from "../../config"


export default {
  name: 'Article',
  data () {
    return {
      inLoading: true,
      articles: [],
      deletedArticles: [],
      articleFilter: {
        tags: [],
        currentPos: 0,
        requestCnt: 99999,
      },
      articleRawContent: "",
      articleTags: "",
      articleOriginalTag: 1,
      showArticleRawContentEditDiv: false,
      showArticleEditDiv: false,
      articleEditId: -1,
      articleEditTags: "",
      articleEditOriginalTag: 0,
      articleEditRawContent: "",
      articleCatalog: "资讯",
    }
  },
  mounted () {
    this.syncDatas()
  },
  methods: {
    syncDatas: function() {
      var that = this

      that.inLoading = true
      that.articles = []
      that.articleFilter = {
        tags: [],
        currentPos: 0,
        requestCnt: 99999,
        containContent: 1,

      },
      that.articleRawContent = ""
      that.articleTags = ""
      that.articleOriginalTag = 1
      that.showArticleRawContentEditDiv = false
      that.deletedArticles = []
      that.showArticleEditDiv = false
      that.articleEditId = -1
      that.articleEditTags = ""
      that.articleEditOriginalTag = 0
      that.articleEditRawContent = ""
      that.articleCatalog = "资讯"

      that.syncArticles()
      that.syncDeletedArticles()
    },
    syncArticles: function() {
      var that = this
      Ajax.post(
        Config.LIST_ARTICLE_URL,
        that.articleFilter,
        true,
        // success
        (data, textStatus, jqXHR) => {
          that.articles = jqXHR.responseJSON
          console.log(JSON.stringify(that.articles))
          that.inLoading = false
        }
      )
    },
    syncDeletedArticles: function() {
      var that = this
      Ajax.post(
        Config.LIST_DELETED_ARTICLE_URL,
        {},
        true,
        // success
        (data, textStatus, jqXHR) => {
          that.deletedArticles = jqXHR.responseJSON
          console.log(JSON.stringify(that.deletedArticles))
          that.inLoading = false
        }
      )
    },
    createArticle: function() {
      this.showArticleRawContentEditDiv = true
    },
    showContent: function(articleId) {
      var that = this
      var article = null
      for (var article of that.articles) {
        if (article.id == articleId) {
          bootbox.alert({
              message: `<pre>${article['content']}</pre>`,
              size: 'large'
          });
          return true
        }
      }
      for (var article of that.deletedArticles) {
        if (article.id == articleId) {
          bootbox.alert({
              message: `<pre>${article['content']}</pre>`,
              size: 'large'
          });
          return true
        }
      }
      alert("文章不存在")
      return false
    },
    hideContent: function() {
      this.showArticleRawContentEditDiv = false
    },
    doCreateArticle: function() {
      var that = this
      var tags = that.articleTags.split(",").sort().join(",")

      Ajax.post(
        Config.CREATE_ARTICLE_URL,
        {
          creater: "",
          originalTag: parseInt(that.articleOriginalTag),
          rawContent: that.articleRawContent,
          tags: tags,
          catalog: that.articleCatalog,
        },
        true,
        // success
        (data, textStatus, jqXHR) => {
          that.syncDatas()
        }
      )
    },
    deleteArticle: function(articleId) {
      var that = this
      bootbox.confirm("确定要删除文章吗?", function(result){ 
        if (result === true) {
          Ajax.post(
            Config.DELETE_ARTICLE_URL,
            {
              articleId: articleId,
            },
            true,
            // success
            (data, textStatus, jqXHR) => {
              that.syncDatas()
            }
          )
        }
      })
    },
    updateArticle: function(articleId) {
      var that = this
      for (var article of that.articles) {
        if (article['id'] == articleId) {
          that.articleEditTags = article['tags']
          that.articleEditOriginalTag = parseInt(article['originalTag'])
          that.articleEditRawContent = article['rawContent']
          that.articleEditId = articleId
          that.showArticleEditDiv = true
          return true
        }
      }
      for (var article of that.deletedArticles) {
        if (article['id'] == articleId) {
          that.articleEditTags = article['tags']
          that.articleEditOriginalTag = parseInt(article['originalTag'])
          that.articleEditRawContent = article['rawContent']
          that.articleEditId = articleId
          that.showArticleEditDiv = true
          return true
        }
      }
      alert('文章不存在')
      return false
    },
    hideAriticleDiv: function() {
      this.showArticleEditDiv = false
    },
    doEditArticle: function() {
      var that = this
      var tags = that.articleEditTags.split(",").sort().join(",")

      Ajax.post(
        Config.UPDATE_ARTICLE_URL,
        {
          articleId: that.articleEditId,
          originalTag: parseInt(that.articleEditOriginalTag),
          rawContent: that.articleEditRawContent,
          tags: tags,
        },
        true,
        // success
        (data, textStatus, jqXHR) => {
          that.syncDatas()
        }
      )
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
