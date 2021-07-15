<template>
    <div>
        <navbar />
        <div class="mt-16 ml-1/12 mr-1/12 pl-14 pr-14">
            <h4 class="text-2xl">最新推薦課程</h4>

            <div class="ml-12 mr-12 mt-16 pt-4">
                <h4 class="text-2xl mb-8 bg-gray-300 rounded-xl pl-2">
                    點選閱讀文章
                </h4>
                <ul>
                    <li
                        @click="read(course.id)"
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
                        <span class="hidden">{{ course.id }}</span>
                        <span class="w-4/12 break-words my-1 px-2"
                            >標題：{{ course.title }}
                        </span>
                        <span class="w-4/12 break-words my-1 px-2"
                            >時間：{{ course.created }}</span
                        >
                        <span class="w-4/12 break-words my-1 px-2"
                            >簡敘：{{ course.description }}</span
                        >
                    </li>
                </ul>
                <div class="w-full flex justify-center">
                    <button class="mx-4" @click="page--" :disabled="!page">
                        <ChevronDoubleLeftIcon class="w-6 h-6" /></button
                    ><button
                        class="mx-4"
                        @click="page++"
                        :disabled="!courses.length"
                    >
                        <ChevronDoubleRightIcon class="w-6 h-6" />
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import navbar from '@/components/navBar.vue'
import { ChevronDoubleRightIcon, ChevronDoubleLeftIcon } from '@heroicons/vue/solid';
export default {
    components: {
        navbar, ChevronDoubleRightIcon, ChevronDoubleLeftIcon
    },
    data: function () {
        return {courses:[],page:0}
    },
    methods: {
        async update() {
            const resp = await this.$store.dispatch('CoursesNew', [
                this.page,
            ])
            if (!resp) {
                return
            }
            this.courses = resp
        },
        async read(id){
            await this.$router.push('/article/read/'+id)
        }
    },
    mounted(){
        this.update()
    },
    watch:{
        page(){
            this.update()
        }
    }
}
</script>
<style scoped></style>
