<template>
  <v-app>
    <v-content>
      <Animations />
      <Main-Form/>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  components: {
    MainForm: () => import(/* webpackChunkName: 'Main-Form' */ './components/Main-Form'),
    Animations: () => import(/* webpackChunkName: 'Animations' */ './components/Animations'),
  },
  data: () => ({
    //
  }),
};
</script>
