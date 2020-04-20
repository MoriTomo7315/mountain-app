<template>
  <div>
    <v-app-bar
      color="primary"
    >
      <v-toolbar-title>
        <nuxt-link to='/'>
          <img src="~/static/weblogo.png">
        </nuxt-link>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-menu offset-y>
        <template v-slot:activator="{ on }">
          <v-btn
            icon
            color="transparent"
            v-on="on"
          >
            <v-app-bar-nav-icon></v-app-bar-nav-icon>
          </v-btn>
        </template>
        <v-list v-if="loginStatus">
          <v-list-item
            nuxt
            to='#'
          >
            <v-list-item-title>マイページ</v-list-item-title>
          </v-list-item>
          <v-list-item
              nuxt
              to='/'
              @click="logout"
          >
            <v-list-item-title>ログアウト</v-list-item-title>
          </v-list-item>
        </v-list>
        <v-list v-else>
          <v-list-item
            nuxt
            to='/login'
          >
            <v-list-item-title>ログイン</v-list-item-title>
          </v-list-item>
          <v-list-item
            nuxt
            to='/signup'
          >
            <v-list-item-title>会員登録</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
  </div>
</template>
<script>
import { mapState, mapGetters, mapActions } from "vuex";

export default {
  name: "Header",
  computed: {
    ...mapState({
      user: state => state.user.user,
      loginStatus: state => state.user.loginStatus
    })
  },
  methods: {
    ...mapActions('user',[
      "logout"
    ]),
    logout() {
      this.logout()
    }
  }
}
</script>
