<template>
  <div style="margin-bottom:50px">
    <van-dropdown-menu>
      <van-dropdown-item v-model="value1" :options="option1" />
      <van-dropdown-item v-model="value2" :options="option2" />
    </van-dropdown-menu>
    <van-grid :border="false" :column-num="2">
      <van-grid-item v-for="(item, k) in shops" :key="k">
        <van-image src="https://img01.yzcdn.cn/vant/apple-1.jpg" />
        <div @click="shopData(item)" >
          <div>{{ item.name }}</div>
          <div style="color: red">￥{{ item.price }}</div>
        </div>
      </van-grid-item>
    </van-grid>
  </div>
</template>

<script>
import api from './../api/API'
export default {
  data() {
    return {
      data: this.$route.query.id,
      value1: 0,
      value2: 'a',
      option1: [
        { text: '全部商品', value: 0 },
        { text: '新款商品', value: 1 },
      ],
      option2: [
        { text: '默认排序', value: 'a' },
        { text: '销量排序', value: 'c' },
      ],
      shops: []
    }
  },
  methods: {
    GetShopCategoryListChildrenById() {
      let cid = {
        id: this.data
      }
      this.$API.get(api.SHOP_CATEGORY_SHOP_LIST_BYID, cid).then(res => {
        this.shops = res.data.shops
      })
    },
    shopData (item) {
      this.$router.push({path: "/shopData",query:{data: item.id}})
    }
  },
  created() {
    this.GetShopCategoryListChildrenById()
  }
}
</script>