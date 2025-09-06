import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';

import { auth } from './stores/auth'



const routes = [
  { path: '/login', component: LoginView },
  { path: '/register', component: RegisterView },
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
