import axios from 'axios'
const port = '3623'
const location=window.location;
export default axios.create({
    baseURL: location.protocol+'//'+location.hostname + ':' + port,
    timeout: 1000,
    headers: {},
})
