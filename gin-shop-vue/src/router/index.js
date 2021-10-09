import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Category from '../views/Category.vue'
import Cart from '../views/Cart.vue'
import Persion from '../views/Persion.vue'
import CategoryChildren from '../views/CategoryChildren.vue'

import Collection from '../views/PersionCenter/Collection.vue'
import Address from '../views/PersionCenter/Address.vue'
import AddressAdd from '../views/PersionCenter/AddressAdd.vue'
import AddressEdit from '../views/PersionCenter/AddressEdit.vue'

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
      },
      {
        path: '/category',
        name: 'Category',
        component: Category,
      },
      {
        path: '/cart',
        name: 'Cart',
        component: Cart
      },
      {
        path: '/persion',
        name: 'Persion',
        component: Persion
      },

      {
        path: '/categoryChildren',
        name: 'CategoryChildren',
        component: CategoryChildren,
      },

      
    ]
  },

  {
    path: '/shopData',
    name: 'shopData',
    component: () => import('../views/Shops/Shops.vue')
  },

  {
    path: '/collection',
    name: 'Collection',
    component: Collection
  },
  {
    path: '/address',
    name: 'Address',
    component: Address
  },
  {
    path: '/addressAdd',
    name: 'AddressAdd',
    component: AddressAdd
  },
  {
    path: '/addressEdit',
    name: 'AddressEdit',
    component: AddressEdit
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
