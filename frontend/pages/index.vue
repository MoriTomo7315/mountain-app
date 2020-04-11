<template>
  <v-layout
    column
    justify-center
    align-center
  >
    <v-flex
      xs12
      sm8
      md6
    >
      <h1>ここにデフォで山１０こくらい表示されるよ</h1>
      <v-col dense>
        <v-card>
          <v-card-title class="headline">
            〇〇山
          </v-card-title>
          <v-card-text>
            <p>〇〇県のおおきく立派な山</p>
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn
              color="primary"
              nuxt
              to="/inspire"
            >
              詳細
            </v-btn>
          </v-card-actions>
        </v-card>
        <v-card>
          <v-card-title class="headline">
            〇〇山
          </v-card-title>
          <v-card-text>
            <p>〇〇県のおおきく立派な山</p>
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn
              color="primary"
              nuxt
              to="/inspire"
            >
              詳細
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-flex>
  </v-layout>
</template>

<script>

export default {
  components: {

  }
}
</script>
