import axios from 'axios'
import store from '@/store'

export default {
    install: function (Vue) {
        const endpoint = store.getters.endpoint;
        Object.defineProperty(Vue.prototype, '$api', {
            value: {
                configurations: {
                    list: () => axios.get(`/api/configuration`),
                    get: (id) => axios.get(`/api/configuration/${id}`),
                    save: (data) => axios.post(`/api/configuration`, data),
                },
                servers: {
                    list: () => axios.get(`/api/server`),
                    get: (id) => axios.get(`/api/server/info/${id}`),
                    add: (data) => axios.post(`/api/server`, data),
                    save: (data) => axios.patch(`/api/server`, data),
                    start: (id) => axios.post(`/api/server/start/${id}`),
                    stop: (id) => axios.post(`/api/server/stop/${id}`),
                    logs: (id) => axios.get(`/api/server/logs/${id}`),
                    delete: (id) => axios.delete(`/api/server/delete/${id}`),
                    command: (id, command) => axios.post(`/api/server/command/${id}`, command, { headers: { 'Content-Type': 'text/plain' } }),
                    sessions: {
                        list: (id) => axios.get(`/api/server/sessions/${id}`),
                        clear: (server) => axios.delete(`/api/server/sessions/${server}`)
                    }
                },
                hosts: {
                    list: () => axios.get(`/api/host`),
                    get: (id) => axios.get(`/api/host/${id}`),
                    save: (data) => axios.patch(`/api/host`, data),
                    add: (data) => axios.post(`/api/host`, data),
                    servers: (id) => axios.get(`/api/host/${id}/servers`),
                    delete: (id) => axios.delete(`/api/host/${id}`)
                }
            }
        });
    },
};