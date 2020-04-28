<template>
  <EmailSignin />
</template>

<script>
import EmailSignin from '~/components/emailSignin'
import { mapState, mapGetters, mapActions } from "vuex";
export default {
  components: {
    EmailSignin
  },
  computed: {
    ...mapState({
      user: state => state.user.user,
    })
  }
}
</script>
