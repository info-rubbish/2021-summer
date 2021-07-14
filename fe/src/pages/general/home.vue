<template>
    <div>
        <navbar />
        <header
            class="
                bg-gray-600
                flex flex-col
                justify-center
                items-center
                h-screen
            "
        >
            <video
                class="
                    absolute
                    h-screen
                    w-screen
                    object-cover
                    z-0
                    lg:top-16
                    opacity-30
                "
                muted
                autoplay
                loop
                src="https://vod-progressive.akamaized.net/exp=1626253275~acl=%2Fvimeo-prod-skyfire-std-us%2F01%2F3716%2F20%2F518580464%2F2414166248.mp4~hmac=a635a8093a4a9af990beafa1f953d382063dfa853caceaa3695d4fa3097e07ba/vimeo-prod-skyfire-std-us/01/3716/20/518580464/2414166248.mp4?download=1&filename=pexels-ron-lach-7013891.mp4"
            ></video>
            <h1 class="p-10 text-white text-3xl bg-gray-800 z-10">
                OO學習網<br />這裡，一個能讓你盡情享受學習的地方
            </h1>
            <a class="p-2 rounded bg-blue-400 absolute bottom-5" href="#main"
                ><ChevronDoubleDownIcon class="w-5 h-5 animate-bounce" />
            </a>
        </header>
        <div class="flex flex-col lg:pt-16" id="main">
            <div class="flex lg:flex-row flex-col-reverse justify-between w-full">
                <lottie-player
                    autoplay
                    loop
                    mode="normal"
                    src="https://assets2.lottiefiles.com/packages/lf20_aljpo8bg.json"
                    class="w-full"
                ></lottie-player>
                <h2 class="text-2xl p-32 text-center w-full my-auto">
                    只要有手機、電腦，這裡的一切都能讓你學習
                </h2>
            </div>
            <!--
            <div class="flex flex-row justify-between w-full">
                <h2 class="text-2xl p-8 text-center w-full my-auto">
                    blalblalblalblalblalblalblalblalblalblalblalb
                </h2>
                <lottie-player
                    autoplay
                    loop
                    mode="normal"
                    src="https://assets3.lottiefiles.com/packages/lf20_oojuetow.json"
                    class="lg:block hidden w-full"
                ></lottie-player>
            </div> --> 
            <div class="flex lg:flex-row flex-col justify-between w-full bg-gray-200">
                <h2 class="text-2xl p-32 text-center w-full my-auto ">
                    累積知識，讓孩子們向上攀爬，是我們的終極目標
                </h2>
                <lottie-player
                    autoplay
                    loop
                    mode="normal"
                    src="https://assets6.lottiefiles.com/private_files/lf30_nIhxTu.json"
                    class="w-full"
                ></lottie-player>
            </div>
            <div class="flex lg:flex-row flex-col-reverse justify-between w-full">
                <lottie-player
                    autoplay
                    loop
                    mode="normal"
                    src="https://assets10.lottiefiles.com/packages/lf20_gditjsml.json"
                    class="w-full"
                ></lottie-player>
                <h2 class="text-2xl p-32 text-center w-full my-auto">
                    不會的嗎？沒關係，有問題提出，我們立即幫你
                </h2>
            </div>
        </div>
        <div class="flex flex-col justify-center items-center" id="main">
            <h1 class="my-3 text-2xl p-10 text-center">緣起</h1>
            <h1 class="text-2xl m-10 md:w-8/12 p-10 bg-gray-100 rounded-xl"> 
                現今我們認為許多貧窮的學生因為沒有適當的學習資源，
                導致知識的缺乏，以致於他們會處於一個貧窮的循環中，
                無法向上發展。我們欲利用線上教學平台推廣教育，
                雖然現今我們在網上有許多的自學平台，
                如coursera、udemy等等，但大多數的平台皆為英文，
                且為較為高難度的學習資源，一般學生無法獲得太多幫助，
                我們欲以平易近人、中文、免費的方式能讓學生更有機會獲得知識，
                才能向上發展，利用所學不只幫助自己、更幫助周遭貧窮的人們擺脫貧窮。
            </h1>
        </div>
    </div>
</template>

<script>
import navbar from '@/components/navBar.vue'
import * as LottiePlayer from '@lottiefiles/lottie-player'
import { ChevronDoubleDownIcon } from '@heroicons/vue/solid'

export default {
    name: 'Home',
    components: { navbar, ChevronDoubleDownIcon },
    methods: {}, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style scoped></style>
