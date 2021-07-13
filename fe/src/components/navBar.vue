<template>
    <div class="sticky top-0" id="app1">
        <div class="lg:hidden">
            <!--lg+ hide-->
            <div class="fixed" v-on:click="switchsidebar()">換成icon吧!</div>
        </div>
        <div class="hidden lg:block">
            <!--lg- hide-->
            <ul class="list">
                <li><router-link to="/home">主頁</router-link></li>
                <li><router-link to="/login">登入/註冊</router-link></li>
            </ul>
        </div>
        <div class="fixed">
            <!--sidebar for lg- -->
            <div
                class="overflow-hidden bg-white w-screen"
                v-bind:class="navClass"
                id="sidebarT"
            >
                <div v-on:click="switchsidebar()">關閉</div>
                <ul>
                    <li><router-link to="/home">主頁</router-link></li>
                    <li><router-link to="/login">登入/註冊</router-link></li>
                </ul>
            </div>
        </div>
    </div>
</template>
<script>
import 'vue'
export default {
    name: 'navBar',
    methods: {
        switchsidebar() {
            if (this.$data.navClass == 'block') this.$data.navClass = 'hidden'
            else this.$data.navClass = 'block'
        },
    },
    data() {
        return { navClass: 'hidden' }
    },
}
// function app1
// w-full
</script>
<style>
ul.list > li {
    display: inline-block;
}
</style>