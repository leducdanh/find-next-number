<template>
  <v-container>
    <v-btn @click="LogIn()" color="pinary">
      SignIn with google
    </v-btn>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex';

export default Vue.extend({
  name: 'LogIn',
  data() {
    return {
      nextNum: 1,
      sizeNums: 10,
      array: [],
      numbersSelected: [],
      countMin: 0,
      countSecond: 0,
    }
  },
  created() {

  },
  computed: {
    ...mapGetters([
      'logged'
    ])
  },
  methods: {
    sleep(milliseconds) {
      return new Promise((resolve) => setTimeout(resolve, milliseconds));
    },
    async LogIn() {
      try {
        // const authCode = await this.$gAuth.getAuthCode()
        // const response = await this.$http.post('http://your-backend-server-api-to-use-authcode', { code: authCode, redirect_uri: 'postmessage' })
        // console.log(response)
        const googleUser = await this.$gAuth.signIn()
        this.isSignIn = this.$gAuth.isAuthorized
        console.log(googleUser)

        await this.sleep(1000)
        console.log('111')
        this.$store.curUser = googleUser
        this.$store.commit('setCurUser', googleUser)
        this.$store.dispatch('setCurUser')
      } catch (error) {
        console.log('error', error)
      }
      // const authCode = await this.$gAuth.getAuthCode()
      // const response = await this.$http.post('http://your-backend-server-api-to-use-authcode', { code: authCode, redirect_uri: 'postmessage' })

    }
  }

})
</script>

<style>
.word-foobar:before {
  content: "Foobar";
}
</style>
