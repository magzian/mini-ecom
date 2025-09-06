<template>
  <div class="auth">
    <h2>Login</h2>
    <form @submit.prevent="handleLogin">
      <input v-model="username" placeholder="Username" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { login } from '../api';
import { store } from '../store';
import { useRouter } from 'vue-router';

const username = ref('');
const password = ref('');
const error = ref('');
const router = useRouter();

const handleLogin = async () => {
  error.value = '';
  try {
    const res = await login(username.value, password.value);
    store.setUser({ username: username.value }, res.data.token);
    router.push('/products');
  } catch (err) {
    error.value = 'Invalid credentials';
  }
};
</script>

<style scoped>
.auth { max-width: 300px; margin: 2rem auto; }
.error { color: red; }
</style>
