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
                <th style="width: 10%">ID</th>
                <th style="width: 25%">标题</th>
                <th style="width: 10%">创建时间</th>
                <th style="width: 25%">标签</th>
                <th style="width: 10%">原创</th>
				        <th style="width: 30%">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="article in articles" :key="article['id']">
                <td>{{article['id']}}</td>
                <td>{{article['title']}}</td>
                <td>{{article['createTime']}}</td>
                <td>{{article['tags']}}</td>
				        <td>{{article['originalTag']}}</td>
                <td><a href="javascript:void(0)" @click="showContent(article['id'])" >内容</a></td>
              </tr>
            </tbody>
          </table>
        </div>
        <div id="garbage_list" class="tab-pane fade">
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
      articleFilter: {
        tags: [],
        currentPos: 0,
        requestCnt: 99999,
      },
      articleRawContent: "",
      articleTags: "",
      articleOriginalTag: 1,
      showArticleRawContentEditDiv: false,
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
      },
      that.articleRawContent = ""
      that.articleTags = ""
      that.articleOriginalTag = 1
      that.showArticleRawContentEditDiv = false

      that.syncArticles()
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
