<template>
  <div class="post-index-wrapper">
    <h1>最近の投稿</h1>
    <v-row no-gutters>
      <v-col
        v-for="(post, index) in posts" :key="post.id"
        cols="12"
        xs="12"
        sm="12"
        md="6"
      >
        <v-card
          class="ma-2"
          flat
        >
          <v-img
            class="white--text align-end"
            height="160px"
            :src="getImageUrl(post.id,index)"
          >
          </v-img>
          <v-btn
            text
            color="success"
          >
            {{ post.mountains }}
          </v-btn>
          <v-card-title class="title">
            <nuxt-link to='#'>{{ post.name }}</nuxt-link>
          </v-card-title>
          <v-card-text>
            {{ '〒' + post.zip_address + ' ' + post.address }}
          </v-card-text>
          <v-list>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>露天風呂：{{ post.has_openbath ? '有' : '無' }}</v-list-item-title>
              </v-list-item-content>
              <v-list-item-content>
                <v-list-item-title>サウナ：{{ post.has_sauna ? '有' : '無' }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <v-card-actions>
            <v-btn
              nuxt
              to='#'
            ><v-icon>fas fa-user</v-icon>〇〇</v-btn>
            <v-spacer />
            <v-btn
              icon
              nuxt
              to='/'
            ><v-icon color="#00aced">fab fa-twitter</v-icon></v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
  import { mapGetters, mapState, mapActions } from 'vuex'
  export default {
    name: "PostIndex",
    mounted () {
      this.fetchPosts()
    },
    methods: {
      ...mapActions({
        fetchPosts: 'posts/fetchPosts'
      }),
      getImageUrl(id, index) {
        return 'images/'+this.posts[index].id+'.jpg'
      }
    },
    computed: {
      ...mapState({
        posts: state => state.posts.posts
      }),
      ...mapGetters('posts', {
        posts: 'posts'
      })
    }
  }
</script>
