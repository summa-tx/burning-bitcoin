<template>
  <v-app>
    <v-content>
      <Animations />
      <Main/>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  components: {
    Main: () => import(/* webpackChunkName: 'Main' */ './components/Main'),
    Animations: () => import(/* webpackChunkName: 'Animations' */ './components/Animations'),
  },
  data: () => ({
    //
  }),
};
</script>
