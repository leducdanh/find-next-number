<template>
  <v-container>
    <h1>{{nextNum}}</h1>
    <h1>{{countDown}}</h1>
    <div class="d-flex justify-center">
      <div v-for="row in sizeNums" :key="row">
        <v-btn v-for="col in sizeNums" :key="`${row}-${col}`"
          :color="getColorBtn(array[(row * sizeNums) - sizeNums + col - 1])" class="my-1"
          @click="onClick(array[(row * sizeNums) - sizeNums + col - 1])">
          {{ forMatNum(array[(row * sizeNums) - sizeNums + col - 1]) }}
        </v-btn>
      </div>
    </div>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'HelloWorld',
  data() {
    return {
      nextNum: 1,
      sizeNums: 10,
      array: [],
      numbersSelected: [],
      countDown: 10
    }
  },
  created() {
    window.addEventListener("keydown", function (e) {
      if (e.keyCode === 114 || (e.ctrlKey && e.keyCode === 70)) {
        e.preventDefault();
      }
    })
    this.array = Array.from({ length: this.sizeNums * this.sizeNums }, (_, i) => i + 1);
    // this.shuffle();
    this.countDownTimer()

  },
  methods: {
    countDownTimer() {
      if (this.countDown > 0) {
        setTimeout(() => {
          this.countDown -= 1
          this.countDownTimer()
        }, 1000)
      }
    },

    getColorBtn(num) {
      if (this.numbersSelected.includes(num))
        return 'warning'
      return 'success';
    },
    onClick(val) {
      if (val == this.nextNum) {
        this.nextNum++;
        this.numbersSelected.push(val);
      }

    },
    forMatNum(val) {
      return ("000" + val).slice(-3);
    },
    shuffle() {
      let array = this.array;
      let currentIndex = array.length, randomIndex;

      // While there remain elements to shuffle.
      while (currentIndex != 0) {

        // Pick a remaining element.
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex--;

        // And swap it with the current element.
        [array[currentIndex], array[randomIndex]] = [
          array[randomIndex], array[currentIndex]];
      }
      console.log('array', array);

      this.array = array;
    }

  }

})
</script>

<style>
.word-foobar:before {
  content: "Foobar";
}
</style>
