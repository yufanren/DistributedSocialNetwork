<template>
    <div class="member">
        <div class="columns">
            <div class="column is-one-quarter">
                <el-card shadow="never">
                    <div slot="header" class="has-text-centered">
                        <el-avatar :size="64" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                        <p class="mt-3">{{ username }}</p>
                    </div>
                    <div>
                        <p class="content">Following：{{ followed_num }}</p>
                        <p class="content">Follower：{{ follower_num }}</p>
                    </div>
                </el-card>
            </div>

            <div class="column">
                <BlogList opt="username" :val="this.username"/>
            </div>
        </div>
    </div>
</template>

<script>
import BlogList from "../blog/List"

export default {
    components: { BlogList },
    data() {
        return {
            username: '',
            follower_num: 0,
            followed_num: 0
        }
    },
    created() {
        this.username = this.$route.params.username
        this.$axios.get("/follower_list/" + this.username).then(res => {
            this.follower_num = res.data.data.total
        })
        this.$axios.get("/follow_list/" + this.username).then(res => {
            this.followed_num = res.data.data.total
        })
        console.log(this.follower_num)
    }
}
</script>