<template>
    <div>
        <navbar />
        <div class="ml-5 mr-5 mt-10">
        <div class="c" v-show="updated">儲存成功</div>
        <h1 class="text-2xl">正在編輯{{ title }}</h1>
            <form @submit.prevent="save" class="p-0 mt-2">
                <h1 class="text-xl">簡述</h1>
                <div class="clear"></div>
                <textarea v-model="description" class="focus:outline-none border-2 w-full rounded-xl p-2"></textarea>
                <textarea @change="fetch_content" v-model="content" class="focus:outline-none border-2 p-2 w-full rounded-xl"></textarea>
                <input type="submit" class="rounded-md w-20 hover:bg-gray-300"/>
            </form>
            <h4 class="mt-4 text-2xl">markdown 檢視</h4>
            <div v-html="complied"></div>
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
        return { updated:false, complied: '', description: 'loading', content: 'content' ,title:'loading'}
    },
    async mounted() {
        var id = this.$route.params.id
        var req = await this.$store.dispatch('CourseInfo',{id})
        var course = req.data.data
        this.$data.description = course.course.description
        this.$data.content =course.content
        this.$data.title=course.course.title
    },
    methods: {
        async save(e) {
            var description = []+e.target[0].value
            var content = []+e.target[1].value
            var id = this.$route.params.id
            this.$store.dispatch('ChangeCourse', {
                description,
                content,
                id,
            })
        },
        async fetch_content(e) {
            var txt = e.target.value || 'empty'
            this.$data.complied = md.MutiLine(txt);
        },
    },
}
</script>
<style scoped>
.clear{
    clear: both;
}</style>
