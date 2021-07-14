<template>
    <div>
        <navbar />
        正在閱讀{{ title }}#{{ $route.params.id }} {{ description }}
        {{ authorId }}
        <div v-html="content"></div>
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
        const course = (await this.$store.dispatch('CourseInfo', { id })).data
            .data
        this.$data.description = course.course.description
        this.$data.title = course.course.title
        this.$data.time = course.course.created
        this.$data.authorId = course.course.author
        this.$data.content = md.MutiLine(course.content)
    },
    methods: {},
}
</script>
<style scoped></style>
