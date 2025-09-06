<!-- frontend/src/components/RegisterForm.vue -->
<template>
  <div class="register-container">
    <div class="register-card">
      <h2>Create Account</h2>
      
      <!-- Success Message -->
      <div v-if="successMessage" class="success-message">
        {{ successMessage }}
      </div>
      
      <!-- Error Message -->
      <div v-if="authStore.error" class="error-message">
        {{ authStore.error }}
      </div>
      
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label for="username">Username:</label>
          <input
            id="username"
            v-model="form.username"
            type="text"
            required
            :disabled="authStore.loading"
            placeholder="Choose a username"
            @blur="validateUsername"
          />
          <small v-if="usernameError" class="field-error">{{ usernameError }}</small>
        </div>
        
        <div class="form-group">
          <label for="password">Password:</label>
          <input
            id="password"
            v-model="form.password"
            type="password"
            required
            :disabled="authStore.loading"
            placeholder="Enter your password"
            @blur="validatePassword"
          />
          <small v-if="passwordError" class="field-error">{{ passwordError }}</small>
          <small class="field-help">Password must be at least 6 characters long</small>
        </div>
        
        <div class="form-group">
          <label for="confirmPassword">Confirm Password:</label>
          <input
            id="confirmPassword"
            v-model="form.confirmPassword"
            type="password"
            required
            :disabled="authStore.loading"
            placeholder="Confirm your password"
            @blur="validateConfirmPassword"
          />
          <small v-if="confirmPasswordError" class="field-error">{{ confirmPasswordError }}</small>
        </div>
        
        <button 
          type="submit" 
          :disabled="authStore.loading || !isFormValid"
          class="register-btn"
        >
          {{ authStore.loading ? 'Creating Account...' : 'Create Account' }}
        </button>
      </form>
      
      <p class="login-link">
        Already have an account? 
        <router-link to="/login">Login here</router-link>
      </p>
    </div>
  </div>
</template>

<script setup>
import { reactive, computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { auth } from '../src/stores/auth'

const router = useRouter()
const authStore = auth()

// Form data
const form = reactive({
  username: '',
  password: '',
  confirmPassword: ''
})

// Form validation states
const usernameError = ref('')
const passwordError = ref('')
const confirmPasswordError = ref('')
const successMessage = ref('')

// Validation functions
const validateUsername = () => {
  const username = form.username.trim()
  
  if (!username) {
    usernameError.value = 'Username is required'
    return false
  }
  
  if (username.length < 3) {
    usernameError.value = 'Username must be at least 3 characters long'
    return false
  }
  
  if (!/^[a-zA-Z0-9_]+$/.test(username)) {
    usernameError.value = 'Username can only contain letters, numbers, and underscores'
    return false
  }
  
  usernameError.value = ''
  return true
}

const validatePassword = () => {
  const password = form.password
  
  if (!password) {
    passwordError.value = 'Password is required'
    return false
  }
  
  if (password.length < 6) {
    passwordError.value = 'Password must be at least 6 characters long'
    return false
  }
  
  passwordError.value = ''
  return true
}

const validateConfirmPassword = () => {
  if (!form.confirmPassword) {
    confirmPasswordError.value = 'Please confirm your password'
    return false
  }
  
  if (form.password !== form.confirmPassword) {
    confirmPasswordError.value = 'Passwords do not match'
    return false
  }
  
  confirmPasswordError.value = ''
  return true
}

// Form validation
const isFormValid = computed(() => {
  return form.username.trim() && 
         form.password && 
         form.confirmPassword && 
         !usernameError.value && 
         !passwordError.value && 
         !confirmPasswordError.value
})

// Handle registration
const handleRegister = async () => {
  // Clear previous messages
  authStore.clearError()
  successMessage.value = ''
  
  // Validate all fields
  const isUsernameValid = validateUsername()
  const isPasswordValid = validatePassword()
  const isConfirmPasswordValid = validateConfirmPassword()
  
  if (!isUsernameValid || !isPasswordValid || !isConfirmPasswordValid) {
    return
  }
  
  try {
    await authStore.register({
      username: form.username.trim(),
      password: form.password
    })
    
    // Show success message
    successMessage.value = 'Account created successfully! Redirecting to login...'
    
    // Clear form
    form.username = ''
    form.password = ''
    form.confirmPassword = ''
    
    // Redirect to login after a short delay
    setTimeout(() => {
      router.push('/login')
    }, 2000)
    
  } catch (error) {
    console.error('Registration failed:', error)
    // Error is already handled in the store
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
}

.register-card {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 450px;
}

h2 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #333;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #555;
}

input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

input:focus {
  outline: none;
  border-color: #007bff;
}

input:disabled {
  background-color: #f8f9fa;
  cursor: not-allowed;
}

.register-btn {
  width: 100%;
  padding: 0.75rem;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
}

.register-btn:hover:not(:disabled) {
  background-color: #218838;
}

.register-btn:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}

.success-message {
  background-color: #d4edda;
  color: #155724;
  padding: 0.75rem;
  border-radius: 4px;
  margin-bottom: 1rem;
  border: 1px solid #c3e6cb;
}

.error-message {
  background-color: #f8d7da;
  color: #721c24;
  padding: 0.75rem;
  border-radius: 4px;
  margin-bottom: 1rem;
  border: 1px solid #f5c6cb;
}

.field-error {
  color: #dc3545;
  font-size: 0.875rem;
  margin-top: 0.25rem;
  display: block;
}

.field-help {
  color: #6c757d;
  font-size: 0.875rem;
  margin-top: 0.25rem;
  display: block;
}

.login-link {
  text-align: center;
  margin-top: 1rem;
  color: #666;
}

.login-link a {
  color: #007bff;
  text-decoration: none;
}

.login-link a:hover {
  text-decoration: underline;
}
</style>