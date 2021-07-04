<template>
    <div>
        <h2>Login</h2>
        <form>
            <div class="grid grid-rows-2 grid-flow-col grid-cols-2 gap-2">
                <label class="border-solid border-2">Account</label
                ><input type="text" class="border-solid border-2" /><label
                    class="border-solid border-2"
                    >Password</label
                ><input type="password" class="border-solid border-2" />
            </div>
        </form>
    </div>
</template>

<script>
export default {
    name: 'Login',
    components: {},
    methods: {}, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style>
form {
    padding: 10px;
}
</style>
