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
import axios from 'axios'

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
    async LogIn() {
      try {
        const googleUser = await this.$gAuth.signIn()
        this.isSignIn = this.$gAuth.isAuthorized
        await axios.post("http://localhost:8000/api/sign-in", {
            "name": googleUser.rv.zf,
            "id": googleUser.Ca,
            "email": googleUser.rv.dw,
        })

        this.$store.curUser = googleUser
        this.$store.commit('setCurUser', googleUser)
        this.$store.dispatch('setCurUser')
      } catch (error) {
        console.log('error', error)
      }

    }
  }

})
</script>

<style>
.word-foobar:before {
  content: "Foobar";
}
</style>
