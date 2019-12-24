import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Scan from '../views/Scan.vue';

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/scan',
        name: 'scan',
        component: Scan
    }
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
});

export default router
