<template>
    <div>
        <navbar />
        <div class="bg-gray-400">
            <div class="ml-12 mr-12 mt-16 pt-4">
                <h4 class="text-2xl mb-8 bg-gray-300 rounded-xl pl-2">
                    點選編輯文章
                </h4>
                <ul>
                    <li
                        @click="changeOld(course.id)"
                        v-for="course in courses"
                        :key="course.id"
                        class="
                            cursor-pointer
                            bg-gray-100
                            m-1
                            p-2
                            rounded-xl
                            hover:bg-gray-200
                            flex flex-row 
                        "
                    >
                        <span class="hidden" >{{ course.id }}</span>
                        <span class="w-1/4 break-words my-1">標題：{{ course.title }} </span> 
                        <span class="w-1/4 break-words my-1">時間：{{
                            course.created
                        }}</span>
                        <span class="w-1/4 break-words my-1">作者：{{ course.author }} </span>
                        <span class="w-1/4 break-words my-1">簡敘：{{
                            course.description
                        }}</span>
                    </li>
                </ul>
            </div>
            <div class="ml-12 mr-12 mt-12">
                <h4
                    @click="show = !show"
                    class="
                        text-2xl
                        bg-gray-300
                        rounded-xl
                        pl-2
                        cursor-pointer
                    "
                >
                    新增文章
                </h4>
                <form
                    v-show="show"
                    @submit.prevent="CreateNew"
                    class="pl-0 pt-4"
                >
                    <label class="text-2xl">標題：</label
                    ><textarea
                        class="outline-none border-0 m-2 p-2 rounded-xl w-full"
                    ></textarea>
                    <label class="text-2xl">敘述：</label
                    ><textarea
                        class="outline-none m-2 p-2 rounded-xl w-full"
                    ></textarea>
                    <input
                        type="submit"
                        class="
                            outline-none
                            border-0
                            w-14
                            leading-relaxed
                            text-white text-xl
                            font-extrabold
                            rounded-md
                            flex
                            items-center
                            justify-center
                            bg-indigo-400
                            hover:bg-indigo-500
                            m-2
                            p-2
                        "
                    />
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/navBar.vue'
import { mapState, mapGetters} from 'vuex'
export default {
    name: 'Login',
    components: { Navbar },
    methods: {
        async CreateNew(e) {
            var title = e.target[0].value
            var description = e.target[1].value
            var res = await this.$store.dispatch('CreateCourse', [
                title,
                description,
                'edit here'
            ])
            this.$router.push('/article/edit/' + res.id)
        },
        async changeOld(id) {
            // const id=e.target.children[0].innerHTML;

            console.log(id)
            this.$router.push('/article/edit/' + id)
        },
    },
    data: function () {
        return {
            courses: [
                // {
                //     id: 'id',
                //     created: 56554,
                //     author: 'auther id',
                //     title: 'title',
                //     description: 'short description',
                // },
            ],
            ...mapState(['user']),
            show: true,
        }
    },
    async mounted() {
        //redirect for permission or ttl flaw
        const resp=await this.$store.dispatch("GetSelfCourses",[])
        if (!resp) {
            return
        }
        this.courses=resp
    },
    computed: { ...mapGetters(['CheckTTL']) },
}
</script>

<style scoped>
form > input {
    border: 2px solid black;
}
</style>
