<template>
    <div>
        <navbar/>
        <h2>Home</h2>
        <HelloWorld />
    </div>
</template>

<script>
import Navbar from '../components/navBar.vue'
import HelloWorld from '../components/HelloWorld.vue'
export default {
    name: 'Home',
    components: { Navbar ,HelloWorld },
    methods: {}, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
}
</style>
