<template>
  <div>
    <login v-if="getUser"></login>
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
  },
  computed: {
    getUser() {
      if (this.$store.state.user == null || this.$store.state.user == '') {
        return true
      } else {
        return false
      }
    }
  }
}
</script>