import Vue from 'vue'
import Router from 'vue-router'

import Article from '@/components/contents/Article'
import Pic from '@/components/contents/Pic'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Article',
      component: Article
    },
    {
      path: '/article',
      name: 'Article',
      component: Article
    },
    {
      path: '/pic',
      name: 'Pic',
      component: Pic
    }
  ]
})
