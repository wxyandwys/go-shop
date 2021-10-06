<template>
  <div id="app">
    <router-view></router-view>
    <van-tabbar v-model="active">
      <van-tabbar-item icon="home-o" @click="GetView(0)">首页</van-tabbar-item>
      <van-tabbar-item icon="apps-o" @click="GetView(1)">分类</van-tabbar-item>
      <van-tabbar-item icon="cart-o" @click="GetView(2)">购物车</van-tabbar-item>
      <van-tabbar-item icon="user-o" @click="GetView(3)">用户</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      active: 0
    }
  },
  methods: {
    GetView(index) {
      if (index == 0) {
        this.$router.push({name: "Home"})
      } else if (index == 1) {
        this.$router.push({name:"Category"})
      } else if (index == 2) {
        this.$router.push({name:"Cart"})
      } else if (index == 3) {
        this.$router.push({name:"Persion"})
      }
    }
  }
}
</script>
