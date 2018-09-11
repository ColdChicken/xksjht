<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col-sm-3 col-md-2 sidebar">
                <ul class="nav nav-sidebar">
                    <li v-bind:class="articleClass">
                        <router-link to="/article">
                            文章管理
                        </router-link>
                    </li>
                    <li v-bind:class="picClass">
                        <router-link to="/pic">
                            图片管理
                        </router-link>
                    </li>
                </ul>
            </div>
            <div class="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
                <div class="row">
                    <router-view></router-view>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
require('jquery')

require('@/assets/bootstrap/css/bootstrap.min.css')
require('@/assets/bootstrap/css/dashboard.css')

require('@/assets/bootstrap/js/bootstrap.min.js')


export default {
  name: 'Content',
  data () {
    return {
        articleClass: "",
        picClass: "",
    }
  },
  watch: {
      '$route': 'fetchData'
  },
  created () {
      this.fetchData()
  },
  methods: {
      fetchData () {
          this.articleClass = ""
          this.picClass = ""

          switch (this.$route.name) {
              case 'Article':
                this.articleClass = "active"
                break
              case 'Pic':
                this.picClass = "active"
                break
              default:
                console.log(`unknown route: ${this.$route.name}`)
                break
          }
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
