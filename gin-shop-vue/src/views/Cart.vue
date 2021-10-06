<template>
  <div style="margin-bottom:50px">
    <van-nav-bar 
      title="购物车" style="background: #f2f3f5; "
    />
    <van-empty
      class="custom-image"
      image="https://img01.yzcdn.cn/vant/custom-empty-image.png"
      description="购物车居然空着"
    />

      <van-submit-bar :price="3050" button-text="提交订单" @submit="onSubmit" style="margin-bottom:50px;" />
  </div>
</template>

<script>
export default {
  data() {
    return {

    }
  },
  methods: {
    onSubmit(e) {
      console.log(e)
    }
  }
}
</script>