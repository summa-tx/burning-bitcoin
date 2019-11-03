<template>
  <section>
    <v-layout justify-center class="animation-container">
      <div class="flamer">
        <img
          class="flame flame-large"
          src="../assets/pixel-flame.gif"
          alt="Flame"
        >
        <img
          class="flame flame-medium"
          src="../assets/pixel-flame.gif"
          alt="Flame"
        >
         <img
          class="flame flame-small"
          src="../assets/pixel-flame.gif"
          alt="Flame"
        >
      </div>
      <div class="flames" ref="flames">
        <img
          v-for="n in 5"
          :key="n"
          class="flame"
          src="../assets/pixel-flame.gif"
          alt="Flame"
        >
      </div>
    </v-layout>
  </section>
</template>

<script>
import { TimelineMax } from 'gsap'

export default {
  name: 'Animation',

  mounted() {
    const { flames } = this.$refs
    const flame = flames.children

    this.timeline0 = new TimelineMax({
      onComplete: () => this.timeline0.restart()
    })

    this.timeline1 = new TimelineMax({
      onComplete: () => this.timeline1.restart()
    })

    this.timeline2 = new TimelineMax({
      onComplete: () => this.timeline2.restart()
    })

    this.timeline3 = new TimelineMax({
      onComplete: () => this.timeline3.restart()
    })

    this.timeline4 = new TimelineMax({
      onComplete: () => this.timeline4.restart()
    })

    this.timeline0.fromTo(flame[0], 3.5, {scale: 0.4, x: 48, y: 0, opacity: 0.6}, {scale: 1.4, y: -35, opacity: 0.1, yoyo: true})
    this.timeline1.fromTo(flame[1], 4.1, {scale: 0.8, x: 98, y: 0, opacity: 0.8}, {scale: 1.3, y: -35, opacity: 0.2, yoyo: true})
    this.timeline2.fromTo(flame[2], 3.3, {scale: 0.6, x: 188, y: 0, opacity: 0.5}, {scale: 1.2, y: -30, opacity: 0.1, yoyo: true})
    this.timeline3.fromTo(flame[3], 4.6, {scale: 0.5, x: 218, y: 0, opacity: 0.9}, {scale: 1.4, y: -45, opacity: 0.3, yoyo: true})
    this.timeline4.fromTo(flame[4], 2.9, {scale: 0.6, x: 278, y: 0, opacity: 0.7}, {scale: 1.6, y: -60, opacity: 0.2, yoyo: true})
  },

  data: () => {
    return {
      timeline0: null,
      timeline1: null,
      timeline2: null,
      timeline3: null,
      timeline4: null,
    }
  }

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
    position absolute
    bottom 0
    height 150px
    z-index 2
.flamer
  position relative
.flame
  position absolute
  bottom 0
  z-index 1
  opacity 0.6
  &-small
    left 130px
    height 200px
  &-medium
    left 250px
    height 280px
  &-large
    height 330px
</style>