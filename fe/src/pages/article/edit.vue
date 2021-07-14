<template>
    <div>
        <navbar />
        <div class="ml-5 mr-5 mt-16">
            <div class="c" v-show="updated">儲存成功</div>
            <h1 class="text-2xl">正在編輯{{ title }}</h1>
            <form @submit.prevent="save" class="p-0 mt-2">
                <h1 class="text-xl">簡述</h1>
                <div class="clear"></div>
                <textarea
                    v-model="description"
                    class="focus:outline-none border-2 w-full rounded-xl p-2"
                ></textarea>
                <h1 class="text-xl">內容</h1>
                <textarea
                    v-model="content"
                    class="focus:outline-none border-2 p-2 w-full rounded-xl"
                ></textarea>
                <input
                    type="submit"
                    class="rounded-md w-20 hover:bg-gray-300"
                />
            </form>
            <h4 class="mt-4 text-2xl">markdown 檢視</h4>
            <div
                v-bind:innerHTML="compile_content(content)"
                class="md-preview"
            ></div>
        </div>
    </div>
</template>
<script>
import navbar from '@/components/navBar.vue'
import md from '@/utils/markdown.js'
export default {
    components: {
        navbar,
    },
    data: function () {
        return {
            updated: false,
            description: '## loading',
            content: 'content',
            title: 'loading',
        }
    },
    async mounted() {
        var course = await this.$store.dispatch(
            'CourseInfo',
            this.$route.params.id
        )
        if (!course) {
            return
        }
        this.title = course.course.title
        this.description = course.course.description
        this.content = course.content
    },
    methods: {
        async save(e) {
            var description = [] + e.target[0].value
            var content = [] + e.target[1].value
            var id = this.$route.params.id
            this.$store.dispatch('ChangeCourse', {
                description,
                content,
                id,
            })
            this.$store.commit('addAlert', [0, '成功', '儲存成功'])
        },
        compile_content: md.MutiLine,
    },
}
</script>
<style scoped>
.clear {
    clear: both;
}
</style>
