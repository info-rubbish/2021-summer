<template>
    <div class="sticky top-0" id="app1">
        <div class="lg:hidden">
            <!--lg+ hide-->
            <div class="fixed" v-on:click="switchsidebar()">
                <button>Open</button>
            </div>
        </div>
        <div class="hidden lg:block">
            <!--lg- hide-->
            <ul>
                <li>1</li>
                <li>2</li>
                <li>3</li>
            </ul>
        </div>
        <div class="fixed">
            <!--sidebar for lg- -->
            <div
                class="transition duration-700 ease-linear w-0"
                data-close="1"
                id="sidebarT"
            ></div>
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
            if (isClose) ele.classList.replace('w-0', 'w-full')
            else ele.classList.replace('w-full', 'w-0')
        },
    },
}
// function app1
// w-full
</script>
<style>
</style>