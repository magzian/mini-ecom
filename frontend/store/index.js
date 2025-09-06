import { reactive } from 'vue';

export const store = reactive({
  user: null,
  token: localStorage.getItem('token') || '',
  setUser(user, token) {
    this.user = user;
    this.token = token;
    localStorage.setItem('token', token);
  },
  logout() {
    this.user = null;
    this.token = '';
    localStorage.removeItem('token');
  },
});
