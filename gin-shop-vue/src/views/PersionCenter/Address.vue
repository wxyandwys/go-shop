<template>
  <div>
    <van-nav-bar
      title="收藏"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
    />
  </div>
</template>
<script>
export default {
  methods: {
    onClickLeft() {
      this.$router.push({name: "Persion"})
    },
  }
}
</script>