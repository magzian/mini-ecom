import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';
import ProductsView from '../views/ProductsView.vue';
import { store } from '../store';

const routes = [
  { path: '/login', component: LoginView },
  { path: '/register', component: RegisterView },
  {
    path: '/products',
    component: ProductsView,
    beforeEnter: (to, from, next) => {
      if (!store.token) return next('/login');
      next();
    },
  },
  { path: '/', redirect: '/products' },
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
