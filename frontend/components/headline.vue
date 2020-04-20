<template>
  <div class="headline-wrapper">
    <h2>山の帰りに行きたい温泉が見つかるサービス</h2>
    <h1>〇〇〇〇</h1>
    <v-form class="headline-form">
      <v-text-field
          v-model="mountainName"
          label="mountainName"
          placeholder="例:八ヶ岳"
          outlined
          required
          solo
      ></v-text-field>
      <v-btn
        color="primary"
        x-large=""
        @click="searchHotsprings"
        nuxt
        to='#'
      >
        温泉を探す
      </v-btn>
    </v-form>
  </div>
</template>

<script>
  export default {
    name: "Headline",
    data() {
      return {
        mountainName: ''
      }
    },
    methods: {
      searchHotsprings() {
        console.log("温泉を検索")
      }
    }
  }
</script>
