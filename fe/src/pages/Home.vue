<template>
    <div>
        <h2>Home</h2>
        <HelloWorld />
    </div>
</template>

<script>
import HelloWorld from '../components/HelloWorld.vue'
export default {
    name: 'Home',
    components: { HelloWorld },
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
