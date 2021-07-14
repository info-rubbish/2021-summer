<template>
    <div>
        <navbar />
        <div class="bg-gray-400">
            <div class="ml-24 mr-48 pt-12 ">
            <h4 class="text-2xl">變更文章(點選以編輯)</h4>
                <ul>
                    <li v-for="course in courseset" :key="course.id" @click="changeOld" class="cursor-pointer bg-gray-100 m-1 p-2 rounded-xl hover:bg-gray-200">
                        <span class="hidden">{{ course.id }}</span>
                        標題：{{ course.title }} 時間：{{ course.created }} 作者：{{ course.author }}
                        {{ course.id }} 簡敘：{{ course.description }}
                    </li>
                </ul>
            </div>
            <div class="ml-24 mr-48 mt-12">
            <h4 class="text-2xl">新增文章</h4>
                <form @submit.prevent="CreateNew">
                    <label class="text-2xl">標題：</label><textarea class="outline-none border-0 m-2 p-2 rounded-xl w-full"></textarea>
                    <label class="text-2xl">敘述：</label><textarea class="outline-none m-2 p-2 rounded-xl w-full"></textarea>
                    <input type="submit" class="border-0 w-20 leading-relaxed text-white text-2xl font-extrabold rounded-md flex items-center justify-center bg-indigo-400 hover:bg-indigo-500 m-2 p-2"/>
                </form>
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
        async CreateNew(e){
            var title=e.target[0].value;
            var description=e.target[1].value;
            var res=await this.$store.dispatch('CreateCourse',{title,description,content:'edit here'})
            if(res.status=='200')
                this.$router.push('/article/edit/'+res.data.data.course.id)
        },
        async changeOld(e){
            const id=e.target.children[0].innerHTML;
            this.$router.push('/article/edit/'+id)
        }
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
        if(!await this.$store.dispatch("CheckTTL")||this.$store.commit("Permission")<2)
            this.$router.push("/login")
        var courseArr=(await this.$store.dispatch('GetSelfCourse')).data.data.courses;
        this.$data.courseset=courseArr;
    },
    computed: {},
}
</script>

<style scoped>
    form>input{
    border: 2px solid black;
    }
</style>
