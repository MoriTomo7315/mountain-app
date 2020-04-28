<template>
  <div v-if="user">
    <ul>
      <li>{{ user.name }}</li>
      <li>{{ user.email }}</li>
    </ul>
  </div>
  <div v-else>
    <p>ログインしてください</p>
  </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from 'vuex';
  export default {
    computed: {
      ...mapState({
        user: state => state.user.user,
      })
    }
  }
</script>
