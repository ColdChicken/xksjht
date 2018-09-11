<template>

<div>
  <template v-if="inLoading === true">
    <div style="height: 100%; width: 100%; text-align: center;">
      <img src="../../assets/loading.gif"/>&nbsp; 加载中，请稍等...
    </div>
  </template>
  <template v-else>
  </template>
</div>

</template>

<script>
require('jquery')

require('@/assets/bootstrap/css/bootstrap.min.css')
require('@/assets/bootstrap/css/dashboard.css')

require('@/assets/bootstrap/js/bootstrap.min.js')


export default {
  name: 'Pic',
  data () {
    return {
      inLoading: true,
    }
  },
  mounted () {
    this.syncDatas()
  },
  methods: {
    syncDatas: function() {
      var that = this
      that.inLoading = true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
