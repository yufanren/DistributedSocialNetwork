<template>
    <section>
        <div class="columns">
            <div class="column is-full">
                <el-card class="box-card" shadow="never">
                    <div slot="header" class="clearfix">
                        <span><i class="fa fa fa-book">Create / Edit</i></span>
                    </div>
                    <div>
                        <el-form :model="blog" ref="topic" class="demo-topic">
                            <el-form-item prop="title">
                                <el-input
                                    v-model="blog.title"
                                    placeholder="Title"
                                ></el-input>
                            </el-form-item>

                            <mavon-editor v-model="blog.content" ref="md" language="en"></mavon-editor>

                            <el-form-item class="mt-3">
                                <el-button type="primary" @click="handleUpdate()">
                                    Update
                                </el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </el-card>
            </div>
        </div>
    </section>
</template>

<script>

    export default {
        data() {
            return {
                blog: {
                    blogId: null,
                    title: "",
                    content: ""
                }
            };
        },
        created() {
            this.blog.blogId = this.$route.params.id
            if(this.blog.blogId == null) return
            this.$axios.get("/blog/" + this.blog.blogId).then(res => {
                const data = res.data.data
                this.blog.title = data.title
                this.blog.content = data.content
            })
        },
        methods: {
            handleUpdate: function () {
                const formData = new FormData()
                formData.append('title', this.blog.title)
                formData.append('content', this.blog.content)
                this.$axios.post("/addBlog", formData).then(res => {
                    this.$router.push({
                        name: "BlogDetail",
                        params: {id : res.data.data.blogId}
                    })
                })
            },
        }
    };
</script>