import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import TarbalHeader from '../views/TarbalHeader.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'TarbalHeader',
    component: TarbalHeader,
    children: [
      {
        path: '',
        name: 'Home',
        component: Home,
      }
    ]
  },

  {
    path: '/shopData',
    name: 'shopData',
    component: () => import('../views/Shops/Shops.vue')
  }
  /*
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  }
  */
]

const router = new VueRouter({
  // mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
