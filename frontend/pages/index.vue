<template>
  <v-layout
    justify-center
    align-center
  >
    <v-flex
      xs12
      sm8
      md8
    >
      <Headline />
      <PostIndex />
    </v-flex>
  </v-layout>
</template>

<script>
  import { mapGetters, mapState, mapActions } from 'vuex'
  import Headline from '~/components/headline'
  import PostIndex from '~/components/postIndex'

  export default {
    components: {
      Headline,
      PostIndex
    }
    // async asyncData({ store }) {
    //   if (store.getters['mountain/mountains'].length) {
    //     return
    //   }
    //   await store.dispatch('mountain/fetchMountains')
    // },
    // computed: {
    //   ...mapState({
    //     mountains: state => state.mountain.mountains
    //   }),
    //   ...mapGetters('mountain', {
    //     mountains: 'mountains'
    //   })
    // }
  }
</script>
