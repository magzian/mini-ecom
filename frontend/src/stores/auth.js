// frontend/src/stores/auth.js
import { defineStore } from 'pinia'
import { authAPI } from '../api/index'

export const auth = defineStore('auth', {
  state: () => ({
    user: null,
    token: null,
    isAuthenticated: false,
    loading: false,
    error: null
  }),

  getters: {
    getUser: (state) => state.user,
    isLoggedIn: (state) => state.isAuthenticated,
    getError: (state) => state.error
  },

  actions: {
    // Initialize auth state from localStorage
    initializeAuth() {
      const token = localStorage.getItem('auth_token')
      const userData = authAPI.getUserData()
      
      if (token && userData) {
        this.token = token
        this.user = userData
        this.isAuthenticated = true
      }
    },

    // Register new user
async register(userData) {
  this.loading = true
  this.error = null
  
  try {
    const response = await authAPI.register(userData)
    
    console.log('Registration response:', response)
    
    // Fix: Check for the actual response structure
    if (response.message === "Success") {
      return response
    } else {
      throw new Error(response.message || 'Registration failed')
    }
  } catch (error) {
    this.error = error.message || 'Registration failed'
    throw error
  } finally {
    this.loading = false
  }
},

    // Login user
    async login(credentials) {
      this.loading = true
      this.error = null
      
      try {
        const response = await authAPI.login(credentials)

        
        
        if (response.token && response.user_id) {
          this.token = response.token
          this.user = {
            user_id: response.user_id,
            username: response.username
          }
          this.isAuthenticated = true
          
          return response
        } else {
          throw new Error(response.message || 'Login failed')
        }
      } catch (error) { 
        this.error = error.message || 'Login failed'
        throw error
      } finally {
        this.loading = false
      }
    },

    // Logout user
    logout() {
      authAPI.logout()
      this.user = null
      this.token = null
      this.isAuthenticated = false
      this.error = null
    },

    // Clear error
    clearError() {
      this.error = null
    }
  }
})