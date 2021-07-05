<template>
    <div>
        <h2>Login</h2>
        <form v-on:submit.prevent="login">
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
        login(e) {
            const name = e.target[0].value
            const password = e.target[1].value
            axios.post('/token', {
                name,
                password,
            }).then((data)=>{
                console.log(data)
                //store token and detail...

            })
        },
    }, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style>
form {
    padding: 10px;
}
.bor > label,
.bor > input {
    border: 1px solid black;
}
</style>
