<template>
    <el-card class="box-card" shadow="never">
        <div slot="header">
            <span>🥂 Promotion</span>
        </div>
        <div>
            <p v-for="(item, index) in list" :key="index" class="block">
                <a :href="item.link" target="_blank">{{ item.title }}</a>
            </p>
        </div>
    </el-card>
</template>

<script>
    export default {
        data() {
            return {
                list: [
                    {
                        title: 'GitHub: Where the world builds software',
                        link: 'https://github.com/'
                    },
                    {
                        title: 'Stack Overflow - Where Developers Learn, Share, & Build Careers',
                        link: 'https://stackoverflow.com/'
                    },
                    {
                        title: 'Cloud Services - Amazon Web Services (AWS)',
                        link: 'https://aws.amazon.com/'
                    }
                ]
            }
        }
    }
</script>