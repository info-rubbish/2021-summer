<template>
    <div>
        <h2>Login</h2>
        <form v-on:submit="login()" onsubmit="return false;">
            <div class="grid grid-rows-3 grid-flow-row grid-cols-2 gap-2 bor">
                <label>Account</label><input type="text" />
                <label>Password</label><input type="password" />
                <label>Submit</label><input type="submit" />
            </div>
        </form>
    </div>
</template>

<script>
export default {
    name: 'Login',
    components: {},
    methods: {
        login() {
            return false;
        },
    }, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style>
form {
    padding: 10px;
}
.bor > label,.bor >input {
    border: 1px solid black;
}
</style>
