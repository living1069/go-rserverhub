import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/servers',
      name: 'servers',
      component: () => import('./views/Servers.vue')
    },
    {
      path: '/hosts',
      alias: '/',
      name: 'hosts',
      component: () => import('./views/Hosts.vue')
    },
    {
      path: '/host/:id',
      name: 'host',
      props: true,
      component: () => import('./views/HostView.vue')
    },
    {
      path: '/server/:id',
      name: 'server',
      props: true,
      component: () => import('./views/ServerView.vue'),
      meta: { tab: 1 }
    },
    {
      path: '/configurations',
      name: 'configurations',
      props: true,
      component: () => import('./views/Configurations.vue')
    },
  ]
})
