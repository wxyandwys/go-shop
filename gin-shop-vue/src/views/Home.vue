<template>
  <div>
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
      <van-grid-item v-for="value in 8" :key="value" icon="photo-o" text="文字" />
    </van-grid>

    <van-divider
      :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 16px' }"
    >
      文字
    </van-divider>

    <van-tabs v-model="activeName">
      <van-tab title="标签 1" name="a">
        <van-grid :column-num="2" :gutter="10">
          <van-grid-item v-for="value in 6" :key="value" icon="photo-o" text="文字" />
        </van-grid>
      </van-tab>
      <van-tab title="标签 2" name="b">
        <van-grid :column-num="2" :gutter="10">
          <van-grid-item v-for="value in 5" :key="value" icon="photo-o" text="文字" />
        </van-grid>
      </van-tab>
      <van-tab title="标签 3" name="c">
        <van-grid :column-num="2" :gutter="10">
          <van-grid-item v-for="value in 9" :key="value" icon="photo-o" text="文字" />
        </van-grid>
      </van-tab>
    </van-tabs>

    

    

  </div>
</template>

<script>
export default {
  data () {
    return {
      activeName: 'a',
    }
  },
  methods: {
    getList () {
      this.$API.get("/home").then(res => {
        console.log(11)
      })
    }
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


