import { createRouter, createWebHistory } from 'vue-router';
import PageLogin from "../views/PageLogin.vue";

const routes = [
    {
        path: '/',
        name: 'login',
        component: PageLogin
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router