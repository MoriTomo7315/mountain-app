<template>
  <div class="userpage">
    <h1>ログイン</h1>
    <div class="userform">
      <v-form>
        <v-text-field
          v-model="email"
          label="E-mail"
          outlined
        ></v-text-field>

        <v-text-field
          v-model="password"
          label="password"
          outlined
          type="password"
        ></v-text-field>
        <v-btn
          color="secondary"
          class="mr-4"
          @click.native="login"
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
    })
  },
  methods: {
    ...mapActions('user',[
      "login"
    ]),
    passworLogin () {
      this.login()
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
