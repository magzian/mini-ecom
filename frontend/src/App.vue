<template>
  <div id="app">
    <nav>
     <!--  <router-link to="/products">Products</router-link> | -->
      <router-link to="/login">Login</router-link> |
      <router-link to="/register">Register</router-link>
    </nav>
    <router-view />
  </div>
</template>

<script setup>
// No logic needed here, handled by router
</script>



<style scoped>
#app {
  max-width: 600px;
  margin: 2rem auto;
  font-family: Arial, sans-serif;
}
nav {
  margin-bottom: 2rem;
}
nav a {
  margin-right: 1rem;
  text-decoration: none;
  color: #42b983;
}
nav a.router-link-exact-active {
  font-weight: bold;
}
</style>

<script>
export default {
  name: 'App',
  data() {
    return {
      users: [],
      newUser: {
        name: '',
        email: ''
      },
      connectionMessage: '',
      loading: false
    }
  },
  methods: {
    async testConnection() {
      this.loading = true
      try {
        const response = await fetch('http://localhost:3000/')
        const data = await response.json()
        this.connectionMessage = `✅ ${data.message}`
      } catch (error) {
        this.connectionMessage = `❌ Connection failed: ${error.message}`
      }
      this.loading = false
    },
    
    async fetchUsers() {
      this.loading = true
      try {
        const response = await fetch('http://localhost:3000/api/users')
        const data = await response.json()
        this.users = data
      } catch (error) {
        alert('Failed to fetch users: ' + error.message)
      }
      this.loading = false
    },
    
    async addUser() {
      if (!this.newUser.name || !this.newUser.email) return
      
      this.loading = true
      try {
        const response = await fetch('http://localhost:3000/api/users', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.newUser)
        })
        
        if (response.ok) {
          this.newUser.name = ''
          this.newUser.email = ''
          await this.fetchUsers() // Refresh the list
        }
      } catch (error) {
        alert('Failed to add user: ' + error.message)
      }
      this.loading = false
    }
  }
}
</script>

<style>
#app {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}

.section {
  margin: 30px 0;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
}

.input {
  display: block;
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-sizing: border-box;
}

button {
  background-color: #4CAF50;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin: 5px;
}

button:hover {
  background-color: #45a049;
}

button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.users-list {
  margin-top: 20px;
}

.user-card {
  padding: 10px;
  margin: 5px 0;
  background-color: #f9f9f9;
  border-left: 4px solid #4CAF50;
  border-radius: 4px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 300px;
}
</style>