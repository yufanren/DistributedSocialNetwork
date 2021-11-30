<template>
    <div>
        <el-card shadow="never">
            <div slot="header" class="clearfix">
                <article v-for="(blog, index) in blogs" :key="index" class="media">
                    <div class="media-left">
                        <figure class="image is-48x48">
                            <el-avatar style="border-radius: 5px;"> {{ blog.Title[0] }} </el-avatar>
                        </figure>
                    </div>
                    <div class="media-content">
                        <div class="">
                            <p class="ellipsis is-ellipsis-1">
                                <el-tooltip class="item" effect="dark" :content="blog.Title" placement="top">
                                    <router-link :to="{ name:'BlogDetail', params:{ id: blog.Blogid }}">
                                        <span class="level-left">{{ blog.Title }}</span>
                                    </router-link>
                                </el-tooltip>
                            </p>
                        </div>
                        <nav class="level has-text-grey is-mobile  is-size-7 mt-2">
                            <div class="level-left">
                                <div class="level-left">
                                    <router-link class="level-item" :to="{ path: `/user/${blog.Author}` }">
                                        {{ blog.Author }}
                                    </router-link>

                                </div>
                            </div>
                        </nav>
                    </div>
                    <div class="media-right" />
                </article>
            </div>

            <el-pagination
                v-show="page.total > 0"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="page.current"
                :page-size="page.size"
                layout="total, sizes, prev, pager, next, jumper"
                :total="page.total">
            </el-pagination>
        </el-card>
    </div>
</template>

<script>
    import {mapGetters} from "vuex";

    export default {
        props: {
            opt: {
                type: String,
                required: true
            },
            val: {
                type: String,
                required: true
            }
        },
        computed: {
            ...mapGetters([
                'user'
            ])
        },
        data() {
            return {
                blogs: [],
                page: {
                    current: 1,
                    size: 10,
                    total: 0
                }
            }
        },
        created() {
            this.getPage()
        },
        methods: {
            getPage() {
                const formData = new FormData();
                formData.append('pageNum', this.page.current)
                formData.append('pageSize', this.page.size)
                if(this.opt == null) {
                    this.$axios.post("/listBlog", formData).then(res => {
                        const data = res.data.data
                        // this.page.current = data.current
                        this.page.total = data.total
                        // this.page.size = data.size
                        for(let i = 0; i < data.blogs.length; i++) data.blogs[i].Blogid = i.toString()
                        this.blogs = data.blogs
                    })
                } else if(this.opt == "username") {
                    this.$axios.post("/listBlog/" + this.val, formData).then(res => {
                        const data = res.data.data
                        // this.page.current = data.current
                        this.page.total = data.total
                        // this.page.size = data.size
                        for(let i = 0; i < data.blogs.length; i++) data.blogs[i].Blogid = i.toString()
                        this.blogs = data.blogs
                    })
                } else if(this.opt == "follow") {
                    this.$axios.post("/listFollowBlog/" + this.user.username, formData).then(res => {
                        const data = res.data.data
                        // this.page.current = data.current
                        this.page.total = data.total
                        // this.page.size = data.size
                        for(let i = 0; i < data.blogs.length; i++) data.blogs[i].Blogid = i.toString()
                        this.blogs = data.blogs
                    })
                }
            },
            handleSizeChange(size) {
                this.page.size = size
                this.getPage()
            },
            handleCurrentChange(current) {
                this.page.current = current
                this.getPage()
            },
        }
    }
</script>