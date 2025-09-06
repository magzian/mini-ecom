import axios from 'axios';

const API_URL = 'http://localhost:3000/api';

export const register = async (userData) => {
    try {
        const response = await axios.post(`${API_URL}/register`, userData)
        return response.data
    } catch (error) {
        throw error.response?.data?.error || 'Registration failed'
    }
}

 export const login = async (credentials) => {
    try {
        const response = await axios.post(`${API_URL}/login`, credentials)
        if (response.data.token) {
            localStorage.setItem('token', response.data.token)
            localStorage.setItem('user', JSON.stringify(response.data.user))
        }
        return response.data
    } catch (error) {
        throw error.response?.data?.error || 'Login failed'
    }
} 



export const getProducts = (token) =>
  axios.get(`${API_URL}/products`, {
    headers: { Authorization: `Bearer ${token}` },
  });

export const createOrder = (order, token) =>
  axios.post(`${API_URL}/orders`, order, {
    headers: { Authorization: `Bearer ${token}` },
  });
