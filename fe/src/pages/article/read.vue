<template>
    <div>
        <navbar/>
        正在閱讀{{title}}#{{$route.params.id}}
        <div v-html="content" v-on="fetch_content"></div>
    </div>
</template>
<script>
import navbar from '@/components/navBar.vue'
import MarkdownIt from 'markdown-it'
export default {
    components: {
        navbar,
    },
    data: function () {
        return {content:""  }
    },
    methods: {
        async fetch_content() {
            var txt;
            //do something to get content("*.md")

            this.$data.content= new MarkdownIt().render(txt)
        },
    },
}
</script>
<style scoped>

</style>