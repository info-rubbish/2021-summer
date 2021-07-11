import axios from 'axios';
import { AxiosResponse } from 'axios';
/**
 * @typedef {object} Resp
 * @property {boolean} Resp.error
 * @property {string} Resp.message
 * @property {T} Resp.data
 * @template T
 */

/**
 * @typedef {object} Token
 * @property {string} Token.token
 * @property {number} Token.created
 * @property {number} Token.ttl
 */

/**
 * @typedef {object} User
 * @property {string} User.name
 * @property {string} User.id
 * @property {number} User.created
 */

/**
 * @description 要用new
 */
export class api {
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
    async Login(data) {
        try {
            const req = await this.axios.post("/token", data)
            return req
        } catch (error) {
            return error.request
        }
    }

    /**
     * 
     * @returns {AxiosResponse<Resp>}
     */
    async Logout() {
        try {
            const req = await this.axios.delete("/token")
            return req
        } catch (error) {
            return error.request
        }
    }

    /**
     * 
     * @typedef {object} resp
     * @property {Token} resp.token
     * @returns {AxiosResponse<Resp<resp>>}
     */
    async Refresh() {
        try {
            const req = await this.axios.put("/token")
            return req
        } catch (error) {
            return error.request
        }
    }

    /**
     * 
     * @param { object } data
     * @param { string } data.name 
     * @param { string } data.password
     * @typedef {object} resp
     * @property {User} resp.user
     * @returns {AxiosResponse<Resp<resp>>}
     */
    async NewUser(data) {
        try {
            const req = await this.axios.post("/user", data)
            return req
        } catch (error) {
            return error.request
        }
    }

    /**
     * 
     * @returns {AxiosResponse<Resp>}
     */
    async DeleteSelf() {
        try {
            const req = await this.axios.delete("/user")
            return req
        } catch (error) {
            return error.request
        }
    }

    /**
     * 
     * @param { object } data
     * @param { string|null } data.name 
     * @param { string|null } data.password 
     * @returns {AxiosResponse<Resp>}
     */
    async ChangeSelfInfo(data) {
        try {
            const req = await this.axios.patch("/user", data)
            return req
        } catch (error) {
            return error.request
        }
    }

    /**
     * @typedef {object} resp
     * @property {User} resp.user
     * @returns {AxiosResponse<Resp<resp>>}
     */
    async GetSelfInfo() {
        try {
            const req = await this.axios.get("/user")
            return req
        } catch (error) {
            return error.request
        }
    }

    /**
     * 
     * @param {string} url 
     * @param {string} token 
     */
    constructor(url, token) {
        this.axios = axios.create({
            baseURL: url,
            headers: {
                "Authorization": token
            }
        })
    }
}