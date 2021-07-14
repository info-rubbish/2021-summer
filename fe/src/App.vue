<template>
    <Alert
        v-for="(el, i) in alerts"
        :key="i"
        @close="popAlert()"
        :level="el.level"
    >
        <h1 class="text-2xl m-4">{{ el.status }}</h1>
        {{ el.message }}
    </Alert>
    <div class="h-screen justify-between flex flex-col">
        <router-view to="/"></router-view>
        <NavFooter/>
    </div>
</template>

<script>
import NavFooter from '@/components/footer.vue'
import { mapMutations, mapState } from 'vuex'
import Alert from '@/components/alert.vue';
export default {
    name: 'app',
    components: { NavFooter,Alert },
    methods: {
        ...mapMutations(['popAlert']),
    },
    computed: {
        ...mapState(['alerts']),
    },
}
</script>

<style></style>
