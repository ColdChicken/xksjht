<template>

<div>
  <template v-if="inLoading === true">
    <div style="height: 100%; width: 100%; text-align: center;">
      <img src="../../assets/loading.gif"/>&nbsp; 加载中，请稍等...
    </div>
  </template>
  <template v-else>
    <button class="btn btn-success btn-xs" style="float:right;padding:2px;z-index:99999; margin: 2px;" @click="uploadPic">上传图片</button>
    <table class="table table-striped table-bordered table-hover" style="table-layout:fixed;word-wrap:break-word;">
      <thead>
        <tr>
          <th style="width: 5%">ID</th>
          <th style="width: 25%">名称</th>
          <th style="width: 10%">类型</th>
          <th style="width: 10%">创建时间</th>
          <th style="width: 30%">路径</th>
          <th style="width: 30%">操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="pic in pics" :key="pic['id']">
          <td>{{pic['id']}}</td>
          <td>{{pic['name']}}</td>
          <td>{{pic['type']}}</td>
          <td>{{pic['createTime']}}</td>
          <td>{{pic['location']}}</td>
          <td><a target="_blank" :href="pic['location']" >查看</a>&nbsp;<a href="javascript:void(0)" @click="deletePic(pic['id'])" >删除</a></td>
        </tr>
      </tbody>
    </table>
  </template>


<template v-if="showUploadPicDiv === true">
    <div class="modal" style="display: block;" tabindex="9999">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<button type="button" class="close" @click="hideUploadPicDiv">&times;</button>
					<h4 class="blue bigger">请输入进行此操作的相关信息:</h4>
				</div>

				<div class="modal-body">
					<div class="row">
            <div class="col-xs-12 col-sm-9">
              <div class="form-group">
                <label>图片名称:</label>
                <div>
                      <input type="text" class="col-xs-12 col-sm-9" v-model="uploadPicName"/>
                </div>
              </div>
              <div style='clear:both'></div>

							<div class="form-group">
							  <label>图片文件:</label>
							  <div>
								  <input type="button" :id="'uploadPicBut'" v-bind:value="uploadButtonLabel" @click="proxyUpload" />
                  <input v-attachment :id="'uploadPicFile'" type="file" style="display:none;" />
							  </div>
							</div>

						</div>

					</div>
				</div>

				<div class="modal-footer">
					<button class="btn btn-sm" @click="hideUploadPicDiv">
						<i class="ace-icon fa fa-times"></i>
						关闭
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

import Vue from 'vue'

import Ajax from "../../common/ajax"
import Config from "../../config"


export default {
  name: 'Pic',
  data () {
    return {
      inLoading: true,
      pics: [],
      showUploadPicDiv: false,
      uploadPicName: "",
      uploadButtonLabel: "点击上传",
    }
  },
  mounted () {
    this.syncDatas()
  },
  created () {
    var that = this
    Vue.directive("attachment",{
      inserted: function(el, binding, vnode){
        $(el).bind("change", function(e) {
          e.preventDefault();
          that.uploadButtonLabel = "上传中..."
          var file = e.target.files[0]
          var data = new FormData()
          data.append('xksj-file', file)
          data.append('xksj-filename', that.uploadPicName)
          Ajax.file(
            Config.UPLOAD_PIC_URL,
            data,
            false,
            // success
            (data, textStatus, jqXHR) => {
              console.log(data)
              that.syncDatas()
            },
            // error
            (jqXHR, textStatus, errorThrown) => {
              console.log(jqXHR)
              var err_msg = JSON.parse(jqXHR.responseText)['msg']
              alert(err_msg);
              that.syncDatas()
            }
          )
        })
      }
    })
  },
  methods: {
    syncDatas: function() {
      var that = this
      that.inLoading = true
      that.pics = []
      that.showUploadPicDiv = false
      that.uploadPicName = ""
      that.uploadButtonLabel = "点击上传"
      that.initPics()
    },
    initPics: function() {
      var that = this
      Ajax.post(
        Config.LIST_PICS_URL,
        {},
        true,
        // success
        (data, textStatus, jqXHR) => {
          that.pics = jqXHR.responseJSON
          console.log(JSON.stringify(that.pics))
          that.inLoading = false
        }
      )
    },
    uploadPic: function() {
      this.showUploadPicDiv = true
    },
    hideUploadPicDiv: function() {
      this.showUploadPicDiv = false
    },
    proxyUpload: function() {
      var that = this
      if (that.uploadPicName == "") {
        alert("请输入文件名")
        return false
      }
      $('#uploadPicFile').click();
    },
    deletePic: function(id) {
      var that = this
      bootbox.confirm("确定要删除图片吗?", function(result){ 
        if (result === true) {
          Ajax.post(
            Config.DELETE_PIC_URL,
            {
              id: id,
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
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
