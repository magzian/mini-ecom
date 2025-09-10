import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';

import { auth } from './stores/auth'

import Welcome from '../views/Welcome.vue'

import Product from '../views/Admin/Product.vue';
import Dashboard from '../views/Admin/Dashboard.vue';
import Order from '../views/Admin/Order.vue';
import Home from '../views/User/Home.vue';



const routes = [
  { path:'/', component:Welcome },
  { path:'/home', component:Home },
  { path: '/login', component: LoginView },
  { path: '/register', component: RegisterView },
  {  path:'/admin/dashboard', component:Dashboard },

  { path:'/admin/dashboard/products', component:Product },
  { path:'/admin/dashboard/orders', component:Order }
  
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
