<template>
  <div class="auth">
    <h2>Register</h2>
    <form @submit.prevent="handleRegister">
      <input v-model="username" placeholder="Username" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Register</button>
    </form>
    <p v-if="error" class="error">{{ error }}</p>
    <p v-if="success" class="success">Registration successful! You can now log in.</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { register } from '../api';

const username = ref('');
const password = ref('');
const error = ref('');
const success = ref(false);

const handleRegister = async () => {
  error.value = '';
  success.value = false;
  try {
    await register(username.value, password.value);
    success.value = true;
    username.value = '';
    password.value = '';
  } catch (err) {
    error.value = 'Registration failed';
  }
};
</script>

<style scoped>
.auth { max-width: 300px; margin: 2rem auto; }
.error { color: red; }
.success { color: green; }
</style>
