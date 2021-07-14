<template>
    <div>
        <navbar />
        <div class="bg-gray-400">
            <div class="ml-24 mr-48 pt-12">
                <h4 class="text-2xl">正在搜尋{{ $route.params.query }}</h4>
                <ul>
                    <li
                        v-for="course in courseset"
                        :key="course.id"
                        @click="select"
                        class="
                            cursor-pointer
                            bg-gray-100
                            m-1
                            p-2
                            rounded-xl
                            course
                            grid grid-flow-col
                        "
                    >
                        <span class="hidden">{{ course.id }}</span>
                        <span>{{ course.title }}</span>
                        <span>{{ Date.parse(course.created) }}</span>
                        <span>{{ course.author }}</span>
                        <span>{{ course.description }}</span>
                        <!-- 標題： 作者： -->
                        <!-- 作者： {{ course.id }} 簡敘： -->
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/navBar.vue'
export default {
    name: 'Login',
    components: { Navbar },
    methods: {
        async select(e) {
            const id = e.target.children[0].innerHTML
            this.$router.push('/article/read/' + id)
        },
    },
    data: function () {
        return {
            courseset: [
                {
                    id: 'id',
                    created: 56554,
                    author: 'auther id',
                    title: 'title',
                    description: 'short description',
                },
            ],
        }
    },
    async mounted() {
        //redirect for permission or ttl flaw
        if (
            !(await this.$store.dispatch('CheckTTL')) ||
            this.$store.commit('Permission') < 1
        )
            this.$router.push('/login')
        const query = this.$route.params.query
        var courseArr = (await this.$store.dispatch('QueryCourse', { query }))
            .data.data.courses
        this.$data.courseset = courseArr
    },
    computed: {},
}
</script>

<style scoped>
form > input {
    border: 2px solid black;
}
.course {
    grid-template: 30% 40% 50% 60%;
}
</style>
