<template>
  <div>
    <login v-if="userStauct == null || userStauct == ''"></login>
    <persion-center v-else></persion-center>
  </div>
</template>

<script>
import Login from './Login.vue'
import PersionCenter from './PersionCenter/PersionCenter.vue'
export default {
  components: {
    Login,
    PersionCenter
  },
  data() {
    return {
      userStauct: this.$store.state.user
    }
  }
}
</script>