import axios from 'axios'
import store from '@/store'

export default {
    install: function (Vue) {
        const endpoint = store.getters.endpoint;
        Object.defineProperty(Vue.prototype, '$api', {
            value: {
                configurations: {
                    list: () => axios.get(`http://${endpoint}/api/configuration`),
                    get: (id) => axios.get(`http://${endpoint}/api/configuration/${id}`),
                    save: (data) => axios.post(`http://${endpoint}/api/configuration`, data),
                },
                servers: {
                    list: () => axios.get(`http://${endpoint}/api/server`),
                    get: (id) => axios.get(`http://${endpoint}/api/server/info/${id}`),
                    add: (data) => axios.post(`http://${endpoint}/api/server`, data),
                    save: (data) => axios.patch(`http://${endpoint}/api/server`, data),
                    start: (id) => axios.post(`http://${endpoint}/api/server/start/${id}`),
                    stop: (id) => axios.post(`http://${endpoint}/api/server/stop/${id}`),
                    logs: (id) => axios.get(`http://${endpoint}/api/server/logs/${id}`),
                    delete: (id) => axios.delete(`http://${endpoint}/api/server/delete/${id}`),
                    command: (id, command) => axios.post(`http://${endpoint}/api/server/command/${id}`, command, { headers: { 'Content-Type': 'text/plain' } }),
                    sessions: {
                        list: (id) => axios.get(`http://${endpoint}/api/server/sessions/${id}`),
                        clear: (server) => axios.delete(`http://${endpoint}/api/server/sessions/${server}`)
                    }
                },
                hosts: {
                    list: () => axios.get(`http://${endpoint}/api/host`),
                    get: (id) => axios.get(`http://${endpoint}/api/host/${id}`),
                    save: (data) => axios.patch(`http://${endpoint}/api/host`, data),
                    add: (data) => axios.post(`http://${endpoint}/api/host`, data),
                    servers: (id) => axios.get(`http://${endpoint}/api/host/${id}/servers`),
                    delete: (id) => axios.delete(`http://${endpoint}/api/host/${id}`)
                }
            }
        });
    },
};