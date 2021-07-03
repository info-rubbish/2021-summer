<template>
    <div class="sticky top-0" id="app1">
        <div class="lg:hidden">
            <!--lg+ hide-->
            <div class="fixed" v-on:click="switchsidebar()">Open</div>
        </div>
        <div class="hidden lg:block">
            <!--lg- hide-->
            <ul class="list">
                <li>Home</li>
                <li>Home</li>
                <li>Home</li>
            </ul>
        </div>
        <div class="fixed">
            <!--sidebar for lg- -->
            <div
                class="hidden overflow-hidden bg-white w-screen"
                data-close="1"
                id="sidebarT"
            >
                <div v-on:click="switchsidebar()">Close</div>
                <ul>
                    <li>Home</li>
                    <li>Home</li>
                    <li>Home</li>
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
            var ele = document.querySelector('#sidebarT')
            var isClose = !!+ele.getAttribute('data-close')
            if (isClose) ele.classList.replace('hidden', 'block')
            else ele.classList.replace('block', 'hidden')
            ele.setAttribute('data-close', isClose ? '0' : '1')
        },
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