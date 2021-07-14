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
            user: JSON.parse(localStorage.getItem('self') || '{}'),
            // token info
            token: token,
        }
    },
    mutations: {
        Permission(state) {
            if (state.self == {}) return 0
            return state.self.permission
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
                localStorage.removeItem('token')
                localStorage.removeItem('self')
                localStorage.removeItem('login')
                refleshAPI('')
            }
        },
        setSelf(state, payload) {
            state.self = payload
            localStorage.setItem('self', JSON.stringify(payload))
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
         *
         */
        async QueryCourse(context, data) {
            try {
                const resp = await send.get('/courses/search/' + data.query)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '查詢失敗',
                        resp.data.message,
                    ])
                    return false
                }
                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '查詢失敗',
                    error.response.data.message,
                ])
                return false
            }
        },
        /**
         *
         */
        async CourseInfo(context, data) {
            try {
                const resp = await send.get('/course/' + data.id)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '查詢失敗',
                        resp.data.message,
                    ])
                    return false
                }
                return true
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '查詢失敗',
                    error.response.data.message,
                ])
                return false
            }
        },

        /**
         *
         */
        async ChangeCourse(context, data) {
            try {
                const req = await send.patch('/course', data)
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
         *
         */
        async CreateCourse(context, data) {
            try {
                const resp = await send.post('/course', data)
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '創建失敗',
                        resp.data.message,
                    ])
                    return false
                }
                return
            } catch (error) {
                context.commit('addAlert', [
                    2,
                    '創建失敗',
                    error.response.data.message,
                ])
                return false
            }
        },

        /**
         *
         */
        async GetSelfCourse(context) {
            try {
                console.log(context.user.id)
                const resp = await (send.get('/courses/user/' + context.user.id))
                console.log('!')
                if (resp.data.error) {
                    context.commit('addAlert', [
                        2,
                        '查詢失敗',
                        resp.data.message,
                    ])
                    return false
                }
                return true
            } catch (error) {
                console.log(error)
                context.commit('addAlert', [
                    2,
                    '查詢失敗',
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
                // const req = await send.get('/user')
                // return req
                return { status: '200', data: { data: { user: state.self } } }
            } catch (error) {
                // return error.request
            }
        },
    },
}
