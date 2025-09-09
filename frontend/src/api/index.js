// frontend/src/services/api.js
import axios from 'axios'

// Create axios instance with base configuration
const api = axios.create({
  baseURL: 'http://localhost:3000/api',
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 10000, // 10 seconds timeout
})

// Request interceptor to add auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('auth_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor for error handling
api.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response?.status === 401) {
      // Token expired or invalid - redirect to login
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user_data')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

// Auth API calls
export const authAPI = {
  // Register new user
  register: async (userData) => {
    try {
      const response = await api.post('/signup', userData)
      return response.data
    } catch (error) {
      throw error.response?.data || error.message
    }
  },

  // Login user
  login: async (credentials) => {
    try {
      const response = await api.post('/login', credentials)
      if (response.data.success && response.data.data.token) {
        // Store token and user data
        localStorage.setItem('auth_token', response.data.data.token)
        localStorage.setItem('user_data', JSON.stringify({
          user_id: response.data.data.user_id,
          username: response.data.data.username
        }))
      }
      return response.data
    } catch (error) {
      throw error.response?.data || error.message 
    }
  },

  // Logout user
  logout: () => {
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_data')
  },

  // Get current user profile (protected route example)
  getProfile: async () => {
    try {
      const response = await api.get('/profile')
      return response.data
    } catch (error) {
      throw error.response?.data || error.message
    }
  },

  // Check if user is authenticated
  isAuthenticated: () => {
    const token = localStorage.getItem('auth_token')
    return !!token
  },

  // Get user data from localStorage
  getUserData: () => {
    const userData = localStorage.getItem('user_data')
    return userData ? JSON.parse(userData) : null
  }
}


//product api



// Export the main api instance for other API calls
export default api