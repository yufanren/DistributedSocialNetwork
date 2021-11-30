<template>
    <div class="columns py-6">
        <div class="column is-half is-offset-one-quarter">
            <el-card shadow="never">
                <div slot="header" class="has-text-centered has-text-weight-bold">
                    Sign Up
                </div>
                <div>
                    <el-form
                        ref="ruleForm"
                        :model="ruleForm"
                        status-icon
                        :rules="rules"
                        label-width="100px"
                        class="demo-ruleForm"
                    >
                        <el-form-item label="Username" prop="name">
                            <el-input v-model="ruleForm.username" />
                        </el-form-item>

                        <el-form-item label="Password" prop="password">
                            <el-input
                                v-model="ruleForm.password"
                                type="password"
                                autocomplete="off"
                            />
                        </el-form-item>

                        <el-form-item label="Confirm Password" prop="checkPassword">
                            <el-input
                                v-model="ruleForm.checkPassword"
                                type="password"
                                autocomplete="off"
                            />
                        </el-form-item>

                        <el-form-item>
                            <el-button
                                type="primary"
                                @click="submitForm('ruleForm')"
                            >
                                Submit
                            </el-button>
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
            const validatePass = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('please enter password again'))
                } else if (value !== this.ruleForm.password) {
                    callback(new Error('passwords are inconsistent'))
                } else {
                    callback()
                }
            }
            return {
                ruleForm: {
                    username: '',
                    password: '',
                    checkPassword: '',
                },
                rules: {
                    username: [
                        { required: true, message: 'please enter username', trigger: 'blur' },
                        {
                            min: 2,
                            max: 10,
                            message: 'length should be 2 to 15 characters',
                            trigger: 'blur'
                        }
                    ],
                    password: [
                        { required: true, message: 'please enter password', trigger: 'blur' },
                        {
                            min: 3,
                            max: 20,
                            message: 'length should be 3 to 20 characters',
                            trigger: 'blur'
                        }
                    ],
                    checkPassword: [
                        { required: true, message: 'please enter password again', trigger: 'blur' },
                        { validator: validatePass, trigger: 'blur' }
                    ],
                }
            }
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        const formData = new FormData()
                        formData.append('username', this.ruleForm.username)
                        formData.append('password', this.ruleForm.password)
                        this.$axios.post('/register', formData).then(res => {
                            this.$router.push("/login")
                        })
                    } else {
                        return false
                    }
                })
            },
            resetForm(formName) {
                this.$refs[formName].resetFields()
            }
        }
    }
</script>