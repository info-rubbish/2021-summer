<template>
    <div>
        <navbar />
        <div class="ml-4 mr-4 mt-16 bg-gray-200 p-2 rounded-md">
            <h1 class="text-xl mt-2">標題：{{ title }}</h1>
            <h1 class="text-xl mt-2">簡介：{{ description }}</h1>
            <h1 class="text-xl mt-2">作者：{{ authorId }}</h1>
        </div>
        <div
            v-html="content"
            class="ml-4 mr-4 mt-8 mb-4 bg-gray-200 p-2 rounded-md md-preview"
        ></div>
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
        return { content: '', description: '', authorId: '', title: '' }
    },
    async mounted() {
        const id = this.$route.params.id
        const course = await this.$store.dispatch('CourseInfo', id)
        this.$data.description = course.course.description
        this.$data.title = course.course.title
        this.$data.time = course.course.created
        this.$data.authorId = course.course.author
        this.$data.content = md.MutiLine(course.content)
    },
    methods: {},
}
</script>
<style scoped>
</style>
