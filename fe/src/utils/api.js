import axios from 'axios'
import { AxiosStatic, AxiosResponse } from 'axios'

/**
 * @type {AxiosStatic}
 */
var send
const config = {
    port: ':3623',
    protocol: 'http://',
    hostname: window.location.hostname,
}
function refleshAPI(token) {
    send = axios.create({
        baseURL: config.protocol + config.hostname + config.port,
        headers: {
            Authorization: token,
        },
    })
}
export default {
    state() {
        const token = JSON.parse(localStorage.getItem('token') || '{}')
        refleshAPI(token.token)
        return {
            alerts: [],
            login: localStorage.getItem('login') || false,
            // some user info
            user: JSON.parse(localStorage.getItem('user') || '{}'),
            // token info
            token: token,
        }
    },
    mutations: {
        Permission(state) {
            if (state.user == {}) return 0
            return state.user.permission
        },
        popAlert(state) {
            state.alerts.shift()
        },
        addAlert(state, payload) {
            state.alerts.push({
                level: payload[0],
                status: payload[1],
                message: payload[2],
            })
        },

        setLogin(state, payload) {
            state.login = payload
            if (!payload) {
                state.token={}
                state.user={}
                localStorage.removeItem('token')
                localStorage.removeItem('user')
                localStorage.removeItem('login')
                refleshAPI('')
            }
        },
        setSelf(state, payload) {
            state.user = payload
            localStorage.setItem('user', JSON.stringify(payload))
        },
        setToken(state, payload) {
            state.token = payload
            localStorage.setItem('token', JSON.stringify(payload))
            refleshAPI(payload.token)
        },
    },
    getters: {
        /**
         * @returns {boolean} status of ttl
         */
        CheckTTL(state) {
            if (!state.token == {}) return false
            if (Date.parse(state.token.created) + state.token.ttl > Date.now())
                return true
            return false
        },
    },
    actions: {
        /**
         *  id,offset,order
         */
        async QueryCourses(context, payload) {
            try {
                const resp = await send.get('/courses/search/' + payload[0],{
                    params:{
                        offset:payload[1],
                        order:payload[2]
                    }
                })
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '查詢失敗',
                        resp.data.message,
                    ])
                    return null
                }
                return resp.data.data.courses
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '查詢失敗',
                    error.response.data.message,
                ])
                return null
            }
        },
        /**
         *
         */
        async CourseInfo(context, payload) {
            try {
                const resp = await send.get('/course/' + payload)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '查詢失敗',
                        resp.data.message,
                    ])
                    return null
                }
                return resp.data.data
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '查詢失敗',
                    error.response.data.message,
                ])
                return null
            }
        },

        /**
         *
         */
        async ChangeCourse(context, data) {
            try {
                const resp = await send.patch('/course', data)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '更新失敗',
                        resp.data.message,
                    ])
                    return false
                }
                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '更新失敗',
                    error.response.data.message,
                ])
                return false
            }
        },

        /**
         * 0 title 1 des...,2 content
         */
        async CreateCourse(context, payload) {
            try {
                const resp = await send.post('/course', {
                    "title":payload[0],
                    description:payload[1],
                    content:payload[2]
                })
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '創建失敗',
                        resp.data.message,
                    ])
                    return null
                }    
                return resp.data.data.course
            } catch (error) {
                window.e=error
                console.log(error);
                context.commit('addAlert', [
                    2,
                    '創建失敗',
                    error.response.data.message,
                ])
                return null
            }
        },

        /**
         *  paylaod 0 offset,1 order
         */
        async GetSelfCourses(context, payload) {
            try {
                const resp = await send.get(
                    '/courses/user/' + context.state.user.id,
                    {
                        params: {
                            offset: payload[0],
                            order: payload[1],
                        },
                    }
                )
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '查詢失敗',
                        resp.data.message,
                    ])
                    return null
                }
                return resp.data.data.courses
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '查詢失敗',
                    error.response.data.message,
                ])
                return null
            }
        },

        /**
         *
         * @param { object } data
         * @param { string } data.name
         * @param { string } data.password
         *
         * @returns {boolean}
         */
        async NewUserAndLogin(context, data) {
            try {
                const resp = await send.post('/user', data)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '註冊失敗',
                        resp.data.message,
                    ])
                    return false
                }
                return await this.dispatch('Login', data)
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '註冊失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
        /**
         *
         * @param { object } data
         * @param { string } data.name
         * @param { string } data.password
         *
         * @returns {boolean}}
         */
        async Login(context, data) {
            try {
                const resp = await send.post('/token', data)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '登入失敗',
                        resp.data.message,
                    ])
                    return false
                }
                context.commit('setSelf', resp.data.data.user)
                context.commit('setToken', resp.data.data.token)
                context.commit('setLogin', true)
                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '登入失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
        /**
         *
         * @returns {boolean}
         */
        async Logout(context) {
            try {
                const resp = await send.delete('/token')
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '登出失敗',
                        resp.data.message,
                    ])
                    return false
                }
                context.commit('setLogin', false)
                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '登出失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
        /**
         *
         * @typedef {object} resp
         * @property {Token} resp.token
         * @returns {boolean}
         */
        async Refresh(context) {
            try {
                const resp = await send.put('/token')
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '刷新token失敗',
                        resp.data.message,
                    ])
                    context.commit('setLogin', false)
                    return false
                }
                context.commit('setLogin', true)
                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '刷新token失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
        /**
         *
         * @returns {boolean}
         */
        async DeleteSelf(context) {
            try {
                const resp = await send.delete('/user')
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '註銷帳號失敗',
                        resp.data.message,
                    ])
                    context.commit('setLogin', false)
                    return false
                }
                context.commit('setLogin', false)
                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '註銷帳號失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
        /**
         *
         * @param { object } data
         * @param { string|null } data.name
         * @param { string|null } data.password
         * @returns {AxiosResponse<Resp>}
         */
        async ChangeSelfInfo(context, data) {
            try {
                const resp = await send.patch('/user', data)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '更改帳號資訊失敗',
                        resp.data.message,
                    ])
                    context.commit('setLogin', false)
                    return false
                }

                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '更改帳號資訊失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
        /**
         * @typedef {object} resp
         * @property {User} resp.user
         * @returns {AxiosResponse<Resp<resp>>}
         */
        async GetSelfInfo(state) {
            try {
                const resp = await send.get('/user')
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '查詢帳號資訊失敗',
                        resp.data.message,
                    ])
                    context.commit('setLogin', false)
                    return false
                }
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '查詢帳號資訊失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
    },
}
