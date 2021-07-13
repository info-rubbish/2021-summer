<template>
    <div>
        <navbar />
        <h2>Login</h2>
        <form v-on:submit.prevent="regist">
            <div class="grid grid-rows-3 grid-flow-row grid-cols-2 gap-2 bor">
                <label>帳戶</label><input type="text" />
                <label>密碼</label><input type="password" />
                <label>確認密碼</label><input type="password" />
                <label>確認</label><input type="submit" />
            </div>
        </form>
    </div>
</template>

<script>
import Navbar from '../components/navBar.vue'
export default {
    name: 'Login',
    components: { Navbar },
    methods: {
        regist() {
            const name = e.target[0].value
            const password = e.target[1].value
            const password_ = e.target[2].value
            if(password==password_){
                //send
            }else
                this.alert('密碼無效')
            
        }
        ,
        alert(msg){

        }
        ,
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
