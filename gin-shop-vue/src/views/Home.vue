<template>
  <div style="margin-bottom:50px">
    <van-nav-bar
      title="首页"
     >
      <template #right>
        <van-icon name="search" size="18" />
      </template>
    </van-nav-bar>
  
    <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white">
      <van-swipe-item>1</van-swipe-item>
      <van-swipe-item>2</van-swipe-item>
      <van-swipe-item>3</van-swipe-item>
      <van-swipe-item>4</van-swipe-item>
    </van-swipe>

    <van-grid square>
      <van-grid-item v-for="(item,k) in categorys" :key="k" icon="photo-o" :text="item.name" />
    </van-grid>

    <van-divider
      :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 16px' }"
    >
      文字
    </van-divider>

    <van-tabs v-model="activeName">
      <van-tab title="新品" name="a">
        <van-grid :column-num="2" :gutter="10">
          <van-grid-item v-for="(item, k) in shops" :key="k+'a'+item.name">
            <van-image src="https://img01.yzcdn.cn/vant/apple-1.jpg"  width="10rem"
  height="10rem" />
            <div @click="shopData(item)"  class="van-ellipsis">
              {{ item.name }}
              <br>
              <span style="color:red">￥{{ item.price }}</span>  
            </div>
          </van-grid-item>
        </van-grid>
      </van-tab>
      <van-tab title="精品" name="b">
        <van-grid :column-num="2" :gutter="10">
          <van-grid-item v-for="value in 5" :key="value" icon="photo-o" text="文字" />
        </van-grid>
      </van-tab>
      <van-tab title="热门" name="c">
        <van-grid :column-num="2" :gutter="10">
          <van-grid-item v-for="value in 9" :key="value" icon="photo-o" text="文字" />
        </van-grid>
      </van-tab>
    </van-tabs>

    

    

  </div>
</template>

<script>
import api from './../api/API'
export default {
  data () {
    return {
      activeName: 'a',
      categorys: [],
      extensions: [],
      shops: []
    }
  },
  methods: {
    getList () {
      this.$API.get(api.HOME).then(res => {
        console.log(11)
      })
    },
    getCategorys () {
      this.$API.get(api.SHOP_CATEGORY_LIST).then(res => {
        this.categorys = res.data
      })
    },
    getShops () {
      this.$API.get(api.SHOP_SHOPS_LIST).then(res => {
        this.shops = res.data
      })
    },
    shopData (item) {
      this.$router.push({path: "/shopData",query:{data: item.id}})
    }
  },
  created() {
    this.getCategorys()
    this.getList()
    this.getShops()
  }
}
</script>

<style scoped>
 .my-swipe .van-swipe-item {
    color: #fff;
    font-size: 20px;
    line-height: 150px;
    text-align: center;
    background-color: #39a9ed;
  }
</style>


