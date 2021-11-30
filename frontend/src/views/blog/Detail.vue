<template>
    <div class="columns">
        <div class="column is-three-quarters">
            <el-card
                class="box-card"
                shadow="never"
            >
                <div
                    slot="header"
                    class="has-text-centered"
                >
                    <p class="is-size-5 has-text-weight-bold">{{ blog.Title }}</p>
                    <div class="has-text-grey is-size-7 mt-3">
                        <el-divider direction="vertical" />
                        <span>Blogger: {{ blog.Author }}</span>
                        <el-divider direction="vertical" />
                    </div>
                </div>

                <div class="markdown-body" align="left" v-html="blog.Content" />

                <nav class="level has-text-grey is-size-7 mt-6">
                </nav>
            </el-card>

        </div>
        <div class="column">
            <section>
                <el-card class="" shadow="never">
                    <div slot="header">
                        <span class="has-text-weight-bold">👨‍💻 About Blogger</span>
                    </div>
                    <div class="has-text-centered">
                        <p class="is-size-5 mb-5">
                            <router-link :to="{ path: `/member/${author.username}/home` }">
                                {{ blog.Author }}
                            </router-link>
                        </p>
                        <div>
                            <button
                                v-if="isFollow"
                                class="button is-success button-center is-fullwidth"
                                @click="follow"
                            >
                                Unfollow
                            </button>
                            <button v-else class="button is-link button-center is-fullwidth" @click="follow()">
                                Follow
                            </button>
                        </div>
                    </div>
                </el-card>
            </section>
        </div>
    </div>
</template>

<script>
    import 'github-markdown-css/github-markdown.css'
    import { mapGetters } from 'vuex'
    export default {
        components: { Comment },
        data() {
            return {
                blog: "",
                author: "",
                isFollow: false
            }
        },
        computed: {
            ...mapGetters([
                'token',
                'user'
            ])
        },
        created() {
            console.log(this.$route.params.id)
            this.$axios.get("/blog/" + this.$route.params.id).then(res => {
                this.blog = res.data.data.blog
                var MarkdownIt = require("markdown-it")
                var md = new MarkdownIt()
                this.blog.Content = md.render(this.blog.Content)
                if(this.token != null) {
                    this.$axios.get("/isFollow/" + this.blog.Author).then(res => {
                        this.isFollow = res.data.data.res
                    })
                }
            })
        },
        methods: {
            follow() {
                this.$axios.get("/follow/" + this.blog.Author).then(res => {
                    this.isFollow = 1 - this.isFollow
                })
            }
        }
    }
</script>