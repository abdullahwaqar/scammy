import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Scan from '../views/Scan.vue';
import Report from '../views/Report.vue';

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
    },
    {
        path: '/report',
        name: 'report',
        component: Report
    }
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
});

export default router
