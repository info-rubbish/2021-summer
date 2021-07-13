<template>
    <div>
        <navbar />
        <header class="bg-gray-400 flex justify-center items-center" style="height:500px">
            <h1 class="p-10 text-white text-3xl bg-gray-800">OO學習網<br>這裡，一個能讓你盡情享受學習的地方</h1>
        </header>
        <h1 class="justify-center my-3 text-2xl">緣起</h1>
        <div class="bg-gray-300 flex justify-center items-center p-6" style="height:400px">
            <h1 class="text-base pt-4">現今我們認為許多貧窮的學生因為沒有適當的學習資源，導致知識的缺乏，以致於他們會處於一個貧窮的循環中，無法向上發展。我們欲利用線上教學平台推廣教育，雖然現今我們在網上有許多的自學平台，如coursera、udemy等等，但大多數的平台皆為英文，且為較為高難度的學習資源，一般學生無法獲得太多幫助，我們欲以平易近人、中文、免費的方式能讓學生更有機會獲得知識，才能向上發展，利用所學不只幫助自己、更幫助周遭貧窮的人們擺脫貧窮。</h1>
        </div>
    </div>
</template>

<script>
import navbar from '@/components/navBar.vue'
import HelloWorld from '@/components/HelloWorld.vue'
export default {
    name: 'Home',
    components: { navbar ,HelloWorld },
    methods: {}, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style scoped>
</style>
