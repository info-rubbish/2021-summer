<template>
    <div>
        <navbar />
        <div class="bg-gray-400">
            <div class="ml-6/12 mr-6/12 pt-12">
                <form class="my-2 p-2" @submit.prevent="query">
                    <input
                        class="
                            focus:outline-none
                            rounded-md
                            lg:w-auto
                            w-32
                            px-2
                        "
                        placeholder="尋找課程"
                    />
                </form>
                <h4 class="text-2xl">{{ $route.params.query }}的搜尋結果</h4>
                <ul>
                    <li
                        v-for="course in courseset"
                        :key="course.id"
                        @click="select(course.id)"
                        class="
                            cursor-pointer
                            bg-gray-100
                            m-1
                            p-2
                            rounded-xl
                            course
                            grid grid-flow-col
                            mt-2
                        "
                    >
                        <span class="hidden">{{ course.id }}</span>
                        <span class="w-1/5">{{ course.title }}</span>
                        <span class="w-1/5">{{
                            ParseTime(course.created)
                        }}</span>
                        <span class="w-1/5">{{ course.author }}</span>
                        <span class="w-1/5">{{ course.description }}</span>
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
    async mounted() {
        if (
            !(this.$store.commit('CheckTTL')) ||
            this.$store.commit('Permission') < 1
        )
            this.$router.push('/login')
    },
    methods: {
        ParseTime(x) {
            var time = new Date(x)
            var list = {
                Month: [
                    'january',
                    'february',
                    'march',
                    'april',
                    'may',
                    'june',
                    'july',
                    'august',
                    'september',
                    'october',
                    'november',
                    'december',
                ],
                DayPostfix: [
                    'st',
                    'nd',
                    'rd',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'st',
                    'nd',
                    'rd',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'st',
                    'nd',
                    'rd',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                    'th',
                ],
            }
            return `${list.Month[time.getMonth() - 1]} ${time.getDay()}${
                list.DayPostfix[time.getDay()]
            }`
        },
        async select(id) {
            // const id = e.target.children[0].innerHTML
            this.$router.push('/article/read/' + id)
        },
        async query(e) {
            if (e.target[0].value.length >= 1) {
                await this.$router.push('/article/find/' + e.target[0].value)
                await this.search()
            }
        },
        async search() {
            const query = this.$route.params.query
            var courseArr = (
                await this.$store.dispatch('QueryCourse', { query })
            ).data.data.courses
            this.$data.courseset = courseArr
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
        await this.search()
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
