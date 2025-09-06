import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';

import { auth } from './stores/auth'
import DashboardView from '../views/DashboardView.vue';



const routes = [
  { path: '/login', component: LoginView },
  { path: '/register', component: RegisterView },
  {  path:'/dashboard', component:DashboardView }
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
