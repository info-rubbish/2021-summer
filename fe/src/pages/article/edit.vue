<template>
    <div>
        <navbar />
        <ul>
            <li>article 1</li>
            <li>article 2</li>
            <li>article 3</li>
        </ul>
    </div>
</template>
<script>
import navbar from '@/components/navBar.vue'
export default {
    components: {
        navbar,
    },
    data: function () {
        return {}
    },
    methods: {
    },
}
</script>
<style scoped>
</style>