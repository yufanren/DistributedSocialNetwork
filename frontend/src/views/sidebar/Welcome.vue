<template>
    <el-card class="box-card" shadow="never">
        <div slot="header">
            <span>💐 Welcome</span>
        </div>
        <div v-if="token != null" class="has-text-centered">
            <b-button type="is-danger" tag="router-link" :to="{ path:'/blog/create' }" outlined>✍ Post Blog</b-button>
        </div>

        <div v-else class="has-text-centered">
            <b-button type="is-primary" tag="router-link" :to="{ path:'/register' }" outlined>Join Us</b-button>
            <b-button type="is-danger" tag="router-link" :to="{ path:'/login' }" outlined class="ml-2">Already User</b-button>
        </div>
    </el-card>
</template>

<script>
    import { mapGetters } from 'vuex'

    export default {
        computed: {
            ...mapGetters([
                'token'
            ])
        },
    }
</script>