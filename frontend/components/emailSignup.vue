<template>
  <div class="userpage">
    <h1>メールアドレスで会員登録</h1>
    <div class="userform">
      <v-form
        ref="form"
        lazy-validation
        class="userform"
      >
        <v-text-field
          v-model="name"
          label="ニックネーム"
          outlined
        ></v-text-field>

        <v-text-field
          v-model="email"
          label="E-mail"
          outlined
          required
        ></v-text-field>

        <v-text-field
          v-model="password"
          label="password"
          type="password"
          outlined
          required
        ></v-text-field>

        <v-text-field
          v-model="password_confirmation"
          label="password（確認用）"
          type="password"
          outlined
          required
        ></v-text-field>

        <v-btn
          color="secondary"
          @click="passwordSignup"
          x-large
          nuxt
          to="/"
        >
          会員登録
        </v-btn>
      </v-form>
      <div class="error" v-if="message">
        {{message}}
      </div>
    </div>
  </div>
</template>

<script>
import firebase, { db } from '@/plugins/firebase'
import { mapState, mapGetters, mapActions } from "vuex";

export default {
  name: "EmailSignup",
  data() {
    return {
      email: '',
      password: '',
      password_confirmation: '',
      name: '',
      message: ''
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
      'fetchUser',
      'login'
    ]),
    getJSONUser() {
      return {
        name: this.name,
        email: this.email
      }
    },
    async passwordSignup () {
      const userData = this.getJSONUser()
      await firebase.auth().createUserWithEmailAndPassword(this.email, this.password)
        .then( () => {
          db.collection('users').add(userData)
          this.fetchUser(userData)
          this.login()
        })
        .catch((error) => {
          alert(error.message)
        })
    }
  }
}
</script>
