<template>
  <section>
    <v-layout justify-center class="animation-container">
      <div class="flames">
        <img
          v-for="n in 5"
          :key="n"
          class="flame"
          src="../assets/pixel-flame.gif"
          alt="Flame"
        >
      </div>
      <!-- <img class="btc" src="../assets/pixel-btc.png" alt="BTC"> -->
    </v-layout>
  </section>
</template>

<script>
export default {
  name: 'Animation'
}
</script>

<style lang='stylus' scoped>
.animation-container
  min-height 400px
.btc
  position absolute
  z-index 0
.flames
  position relative
  min-width 400px
  min-height 400px
  background url('../assets/pixel-btc.png') center center no-repeat
  & .flame
    height 150px
    z-index 1
</style>