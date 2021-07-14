<template>
    <router-view to="/"></router-view>
    <Footer/>
</template>

<script>
import Footer from '@/components/footer.vue';
export default { name: 'app', components: {Footer} }
</script>

<style>
</style>
