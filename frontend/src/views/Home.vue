<template>
    <div>
        <div class="box" align="left">🔔 Beta Version</div>
        <div class="columns">
            <div class="column is-three-quarters">
                <el-tabs v-model="tab" @tab-click="handleClick">
                    <el-tab-pane label="total" name="total">
                        <BlogList/>
                    </el-tab-pane>
                    <el-tab-pane label="follow" name="follow">
                        <BlogList opt="follow"/>
                    </el-tab-pane>
                </el-tabs>
            </div>
            <div class="column">
                <Sidebar/>
            </div>
        </div>
    </div>
</template>

<script>
    import BlogList from "../views/blog/List"
    import Sidebar from '../views/sidebar/Index'

    export default {
        name: "Home",
        components: { BlogList, Sidebar },
        data() {
            return {
                tab: 'total'
            }
        },
        methods: {
            handleClick(tab) {
                this.tab = tab.name
                this.init(tab.name)
            }
        }
    }
</script>

<style scoped>

</style>