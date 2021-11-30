<template>
    <div class="columns py-6">
        <div class="column is-half is-offset-one-quarter">
            <el-card shadow="never">
                <div slot="header" class="has-text-centered has-text-weight-bold">
                    Sign In
                </div>
                <div>
                    <el-form
                        :model="ruleForm"
                        status-icon
                        :rules="rules"
                        ref="ruleForm"
                        label-width="100px"
                        class="demo-ruleForm"
                    >
                        <el-form-item label="Username" prop="name">
                            <el-input v-model="ruleForm.username"></el-input>
                        </el-form-item>

                        <el-form-item label="Password" prop="password">
                            <el-input
                                type="password"
                                v-model="ruleForm.password"
                                autocomplete="off"
                            ></el-input>
                        </el-form-item>

                        <el-form-item>
                            <el-button type="primary" @click="submitForm('ruleForm')">Submit</el-button>
                            <el-button @click="resetForm('ruleForm')">Reset</el-button>
                        </el-form-item>
                    </el-form>
                </div>
            </el-card>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                ruleForm: {
                    username: "",
                    password: "",
                },
                rules: {
                    username: [
                        { required: true, message: "please enter username", trigger: "blur" },
                        {
                            min: 2,
                            max: 15,
                            message: "length should be 2 to 15 characters",
                            trigger: "blur",
                        },
                    ],
                    password: [
                        { required: true, message: "please enter password", trigger: "blur" },
                        {
                            min: 3,
                            max: 20,
                            message: "length should be 3 to 20 characters",
                            trigger: "blur",
                        },
                    ],
                },
            }
        },
        methods: {
            submitForm(formName) {
                const _this = this
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        const formData = new FormData();
                        formData.append('username', this.ruleForm.username)
                        formData.append('password', this.ruleForm.password)
                        this.$axios.post("/login", formData).then(res => {
                            const token = res.data.data.token
                            const username = this.ruleForm.username
                            _this.$store.commit("SET_TOKEN", token)
                            _this.$store.commit("SET_USER", { 'username': username })
                            _this.$router.push("/")
                            location.reload()
                        })
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            }
        }
    }
</script>