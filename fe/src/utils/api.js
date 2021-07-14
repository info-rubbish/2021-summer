import axios from 'axios'
import { AxiosResponse } from 'axios'
var store = {}
const config = {
    port: ':3623',
    protocol: 'http://',
    hostname: window.location.hostname,
}
function init() {
    var tokenDetail = {}
    // automatically check localstorage
    const tokenStorage = localStorage.getItem('usertoken')
    if (tokenStorage != null) tokenDetail = JSON.parse(tokenStorage)
    const userStorage = localStorage.getItem('user')
    store.user = JSON.parse(userStorage || '{}')
    //store it
    store.tokenDetail = tokenDetail
    store.axios = axios.create({
        baseURL: config.protocol + config.hostname + config.port,
        headers: {
            Authorization: tokenDetail.token,
        },
    })
}
init()
export default {
    state: {},
    mutations: {
        Permission() {
            if (store.user.permission == undefined) throw error('')
            return store.user.permission
        },
    },
    actions: {
        /**
         * 
         */
        async QueryCourse(context, data){
            try {
                const req = await store.axios.get('/courses/search/'+data.query)
                return req
            } catch (error) {
                return error.request
            }
        }
        ,
        /**
         * 
         */
        async CourseInfo(context, data){
            try {
                const req = await store.axios.get('/course/'+data.id)
                return req
            } catch (error) {
                return error.request
            }
        },

        /**
         * 
         */
        async ChangeCourse(context, data){
            try {
                const req = await store.axios.patch('/course', data)
                return req
            } catch (error) {
                return error.request
            }
        },

        /**
         *
         */
        async CreateCourse(context, data) {
            try {
                const req = await store.axios.post('/course', data)
                return req
            } catch (error) {
                return error.request
            }
        },

        /**
         *
         */
        async GetSelfCourse(context) {
            const req = await store.axios.get('/courses/user/' + store.user.id)
            return req
        },

        /**
         *
         * @param { object } data
         * @param { string } data.name
         * @param { string } data.password
         * @typedef {object} resp
         * @property {User} resp.user
         * @returns {AxiosResponse<Resp<resp>>}
         */
        async NewUser(context, data) {
            try {
                const req = await store.axios.post('/user', data)
                return req
            } catch (error) {
                return error.request
            }
        },
        /**
         *
         * @param { object } data
         * @param { string } data.name
         * @param { string } data.password
         *
         * @typedef {object} resp
         * @property {Token} resp.token
         * @returns {AxiosResponse<Resp<resp>>}
         */
        async Login(context, data) {
            try {
                const req = await store.axios.post('/token', data)
                if (req.data.data.error) throw error()
                localStorage.setItem(
                    'usertoken',
                    JSON.stringify(req.data.data.token)
                )
                localStorage.setItem('user', JSON.stringify(req.data.data.user))
                init()
                return req
            } catch (error) {
                return error.request
            }
        },
        /**
         *
         * @returns {AxiosResponse<Resp>}
         */
        async Logout(context) {
            try {
                const req = await store.axios.delete('/token')
                localStorage.removeItem('usertoken')
                localStorage.removeItem('user')
                init();
                return req
            } catch (error) {
                return error.request
            }
        },
        /**
         *
         * @typedef {object} resp
         * @property {Token} resp.token
         * @returns {AxiosResponse<Resp<resp>>}
         */
        async Refresh(context) {
            try {
                const req = await store.axios.put('/token')
                return req
            } catch (error) {
                return error.request
            }
        },
        /**
         *
         * @returns {AxiosResponse<Resp>}
         */
        async DeleteSelf(context) {
            try {
                const req = await store.axios.delete('/user')
                localStorage.removeItem('usertoken')
                localStorage.removeItem('user')
                init()
                return req
            } catch (error) {
                return error.request
            }
        },
        /**
         *
         * @param { object } data
         * @param { string|null } data.name
         * @param { string|null } data.password
         * @returns {AxiosResponse<Resp>}
         */ async ChangeSelfInfo(context, data) {
            try {
                const req = await store.axios.patch('/user', data)
                return req
            } catch (error) {
                return error.request
            }
        },
        /**
         * @typedef {object} resp
         * @property {User} resp.user
         * @returns {AxiosResponse<Resp<resp>>}
         */
        async GetSelfInfo(context) {
            try {
                const req = await store.axios.get('/user')
                return req
            } catch (error) {
                return error.request
            }
        },
        /**
         * @returns {boolean} status of ttl
         */
        CheckTTL(context) {
            var tokenDetail = store.tokenDetail
            if (!tokenDetail.ttl) return false
            if (Date.parse(tokenDetail.created) + tokenDetail.ttl > Date.now())
                return true
            return false
        },
    },
}
