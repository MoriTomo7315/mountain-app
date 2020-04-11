<template>
  <div>
    <v-app-bar color="primary">
      <v-toolbar-title>
        <img src="~/static/weblogo.png">
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        icon
        color="navfont"
        nuxt
        to='/'
      >
        <v-icon>fas fa-home</v-icon>
      </v-btn>
      <v-btn
        color="navfont"
        text
        nuxt
        to='/login'
      >
        ログイン
      </v-btn>
      <v-btn
        color="navfont"
        text
        nuxt
        to='/signup'
      >
        会員登録
      </v-btn>
    </v-app-bar>
  </div>
</template>
