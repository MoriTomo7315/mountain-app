<template>
  <div class="userpage">
    <h1>ログイン</h1>
    <div class="userform">
      <v-form>
        <v-text-field
          v-model="email"
          label="E-mail"
          outlined
          required
        ></v-text-field>

        <v-text-field
          v-model="password"
          label="password"
          outlined
          type="password"
          required
        ></v-text-field>
        <v-btn
          color="secondary"
          class="mr-4"
          @click="passworLogin"
          x-large
          nuxt
          to="/"
        >
          ログイン
        </v-btn>
        <nuxt-link
          to="#"
        >
          パスワードを忘れた方はこちら...
        </nuxt-link>
      </v-form>
    </div>
  </div>
</template>

<script>
import firebase from "@/plugins/firebase"
import { mapState, mapGetters, mapActions } from "vuex";

export default {
  name: "EmailSignin",
  data() {
    return {
      email: '',
      password: '',
    }
  },
  computed: {
    ...mapState({
      user: state => state.user.user,
      loginStatus: state => state.user.loginStatus
    })
  },
  methods: {
    ...mapActions('user',[
      'login'
    ]),
    passworLogin () {
      firebase.auth().signInWithEmailAndPassword(this.email, this.password)
        .then( () => {
          this.login()
        })
        .catch((error) => {
          alert(error.message)
        })
  }
  }
}
</script>
