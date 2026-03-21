import { createRouter, createWebHistory } from 'vue-router';
import PageLogin from "../views/PageLogin.vue";
import PageAdmin from "../views/PageAdmin.vue";
import PagePlayer from "../views/PagePlayer.vue";
import PageHost from "../views/PageHost.vue";

const routes = [
    {
        path: '/',
        name: 'login',
        component: PageLogin
    },
    {
        path: '/admin',
        name: 'admin',
        component: PageAdmin
    },
    {
        path: '/player',
        name: 'player',
        component: PagePlayer
    },
    {
        path: '/host',
        name: 'host',
        component: PageHost
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router