<template>
  <div>
    <AdminLayout>
      <section class="h-[100vh] bg-gray-50 dark:bg-gray-900 py-3 sm:py-5">
        <div class="px-4 mx-auto max-w-screen-2xl lg:px-12">
          <div class="overflow-hidden bg-white shadow-md dark:bg-gray-800 sm:rounded-lg">

            <!-- Loading State -->
            <div v-if="loading" class="flex justify-center items-center h-64">
              <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
            </div>

            <!-- Error State -->
            <div v-else-if="error"
              class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 mx-4 my-4 rounded relative">
              <strong class="font-bold">Error: </strong>
              <span class="block sm:inline">{{ error }}</span>
              <button @click="fetchProducts" class="mt-2 bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700">
                Retry
              </button>
            </div>

            <!-- Main Content -->
            <div v-else>
              <div
                class="flex flex-col flex-shrink-0 space-y-3 md:flex-row md:items-center lg:justify-end md:space-y-0 md:space-x-3">
                <button type="button" @click="openAddModal"
                  class="flex items-center justify-center px-4 py-2 my-4 mx-2 text-sm font-medium text-white rounded-lg bg-green-600 hover:bg-green-700 focus:ring-4 focus:ring-green-300">
                  <svg class="h-3.5 w-3.5 mr-2" fill="currentColor" viewbox="0 0 20 20"
                    xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                    <path clip-rule="evenodd" fill-rule="evenodd"
                      d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" />
                  </svg>
                  Add New Product
                </button>
                <button @click="fetchProducts" :disabled="loading"
                  class="flex items-center justify-center px-4 py-2 my-4 mx-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 disabled:opacity-50">
                  {{ loading ? 'Loading...' : 'Refresh' }}
                </button>
              </div>

              <!-- Products Table -->
              <div v-if="products.length > 0" class="overflow-x-auto">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                  <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                    <tr>
                      <th scope="col" class="px-4 py-3">Product</th>
                      <th scope="col" class="px-4 py-3">Price</th>
                      <th scope="col" class="px-4 py-3">Description</th>
                      
                      <th scope="col" class="px-4 py-3">Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="product in products" :key="product.id"
                      class="border-b dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700">
                      <th scope="row"
                        class="flex items-center px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                        <div
                          class="w-10 h-10 mr-3 bg-gray-200 rounded flex items-center justify-center overflow-hidden">
                          <img v-if="product.image_url" :src="product.image_url" :alt="product.name"
                            class="w-full h-full object-cover" />
                          <svg v-else class="w-6 h-6 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd"
                              d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-8 3 6 2-4 3 6z"
                              clip-rule="evenodd" />
                          </svg>
                        </div>
                        {{ product.name }}
                      </th>
                      <td class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                        {{ formatPrice(product.price) }}
                      </td>
                      <td class="px-4 py-2">
                        <span
                          class="bg-primary-100 text-primary-800 text-xs font-medium px-2 py-0.5 rounded dark:bg-primary-900 dark:text-primary-300">
                          {{ product.description || 'No description' }}
                        </span>
                      </td>
                      
                      <td class="px-4 py-2">
                        <div class="flex space-x-2">
                          <button @click="editProduct(product)"
                            class="text-blue-600 hover:text-blue-900 font-medium text-sm">
                            Edit
                          </button>
                          <button @click="deleteProduct(product.id)"
                            class="text-red-600 hover:text-red-900 font-medium text-sm">
                            Delete
                          </button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div v-else class="text-center py-12">
                <p class="text-gray-500 text-lg">No products found</p>
              </div>

              <!-- Pagination -->
              <nav
                class="flex flex-col items-start justify-between p-4 space-y-3 md:flex-row md:items-center md:space-y-0"
                aria-label="Table navigation">
                <span class="text-sm font-normal text-gray-500 dark:text-gray-400">
                  Showing
                  <span class="font-semibold text-gray-900 dark:text-white">{{ products.length }}</span>
                  products
                </span>
              </nav>
            </div>
          </div>
        </div>
      </section>
    </AdminLayout>

    <!-- Add Product Modal - Simple Version -->
    <div v-if="showAddModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold">Add New Product</h3>
          <button @click="closeAddModal" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form @submit.prevent="createProduct" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Product Name</label>
            <input type="text" v-model="addForm.name" required
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea v-model="addForm.description" rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Price</label>
            <input type="number" v-model.number="addForm.price" step="0.01" min="0" required
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Product Image</label>
            <input type="file" @change="handleImageUpload($event, 'add')" accept="image/*"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" />
            <p class="text-xs text-gray-500 mt-1">Max file size: 5MB. Formats: JPEG, PNG, WebP, GIF</p>
            <p v-if="addForm.image" class="text-xs text-green-600 mt-1">✓ Image selected: {{ addForm.image.name }}</p>
          </div>

          <div class="flex justify-end space-x-3 pt-4">
            <button type="button" @click="closeAddModal"
              class="px-4 py-2 text-gray-700 bg-gray-200 rounded-md hover:bg-gray-300">
              Cancel
            </button>
            <button type="submit" class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700">
              Create Product
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Edit Product Modal - Simple Version -->
    <div v-if="showEditModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold">Edit Product</h3>
          <button @click="closeEditModal" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form @submit.prevent="updateProduct" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Product Name</label>
            <input type="text" v-model="editForm.name" required
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea v-model="editForm.description" rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Price</label>
            <input type="number" v-model.number="editForm.price" step="0.01" min="0" required
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Product Image</label>
            <input type="file" @change="handleImageUpload($event, 'edit')" accept="image/*"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" />
            <p class="text-xs text-gray-500 mt-1">Max file size: 5MB. Formats: JPEG, PNG, WebP, GIF</p>
            <p v-if="editForm.image" class="text-xs text-green-600 mt-1">✓ New image selected: {{ editForm.image.name }}
            </p>
          </div>

          <div class="flex justify-end space-x-3 pt-4">
            <button type="button" @click="closeEditModal"
              class="px-4 py-2 text-gray-700 bg-gray-200 rounded-md hover:bg-gray-300">
              Cancel
            </button>
            <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700">
              Update Product
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import AdminLayout from './Layout/AdminLayout.vue';
import axios from 'axios'
import { ref, onMounted } from 'vue'

// Reactive data
const products = ref([])
const loading = ref(false)
const error = ref(null)

// Modal data
const showEditModal = ref(false)
const showAddModal = ref(false)
const editingProduct = ref({})
const editForm = ref({
  name: '',
  description: '',
  price: 0,
  image: null
})
const addForm = ref({
  name: '',
  description: '',
  price: 0,
  image: null
})

// Methods
const fetchProducts = async () => {
  loading.value = true
  error.value = null

  try {
    // Get token from localStorage
    const token = localStorage.getItem('token')

    const response = await axios.get('http://localhost:3000/api/products/', {
      headers: {
        'Authorization': `${token}`,
        'Content-Type': 'application/json',
      }
    })

    products.value = response.data.products

  } catch (err) {
    error.value = err.response?.data?.message || err.message || 'Failed to fetch products'
    console.error('Error fetching products:', err)
  } finally {
    loading.value = false
  }
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD'
  }).format(price)
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const editProduct = (product) => {
  console.log('Edit product clicked:', product)
  editingProduct.value = { ...product }
  editForm.value = {
    name: product.name,
    description: product.description,
    price: product.price,
    image: null // Reset image field for editing
  }
  showEditModal.value = true
}

const openAddModal = () => {
  console.log('Opening add modal...')
  addForm.value = {
    name: '',
    description: '',
    price: 0,
    image: null
  }
  showAddModal.value = true
  console.log('showAddModal set to:', showAddModal.value)
}

const closeAddModal = () => {
  showAddModal.value = false
  addForm.value = { name: '', description: '', price: 0, image: null }
}

const handleImageUpload = (event, formType = 'edit') => {
  const file = event.target.files[0]
  if (file) {
    // Validate file type
    const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/webp', 'image/gif']
    if (!allowedTypes.includes(file.type)) {
      alert('Please select a valid image file (JPEG, PNG, WebP, or GIF)')
      return
    }

    // Validate file size (5MB limit)
    if (file.size > 5 * 1024 * 1024) {
      alert('Image size should be less than 5MB')
      return
    }

    if (formType === 'edit') {
      editForm.value.image = file
    } else {
      addForm.value.image = file
    }
  }
}

const createProduct = async () => {
  console.log('Creating product with data:', addForm.value)
  try {
    const token = localStorage.getItem('token')

    // Create FormData for file upload
    const formData = new FormData()
    formData.append('name', addForm.value.name)
    formData.append('description', addForm.value.description)
    formData.append('price', addForm.value.price.toString())

    if (addForm.value.image) {
      formData.append('image', addForm.value.image)
    }

    const response = await axios.post('http://localhost:3000/api/products/', formData, {
      headers: {
        'Authorization': `${token}`,
        'Content-Type': 'multipart/form-data'
      }
    })

    // Add new product to local state
    products.value.unshift(response.data.product)

    closeAddModal()
    alert('Product created successfully!')

  } catch (err) {
    console.error('Error creating product:', err)
    alert(err.response?.data?.message || err.message || 'Failed to create product')
  }
}

const updateProduct = async () => {
  try {
    const token = localStorage.getItem('token')

    // Create FormData for file upload
    const formData = new FormData()
    formData.append('name', editForm.value.name)
    formData.append('description', editForm.value.description)
    formData.append('price', editForm.value.price.toString())

    if (editForm.value.image) {
      formData.append('image', editForm.value.image)
    }

    const response = await axios.put(`http://localhost:3000/api/products/${editingProduct.value.id}`, formData, {
      headers: {
        'Authorization': `${token}`,
        'Content-Type': 'multipart/form-data'
      }
    })

    // Update product in local state
    const index = products.value.findIndex(p => p.id === editingProduct.value.id)
    if (index !== -1) {
      products.value[index] = response.data.product || { ...products.value[index], ...editForm.value }
    }

    showEditModal.value = false
    alert('Product updated successfully!')

  } catch (err) {
    console.error('Error updating product:', err)
    alert(err.response?.data?.message || err.message || 'Failed to update product')
  }
}

const closeEditModal = () => {
  showEditModal.value = false
  editForm.value = { name: '', description: '', price: 0, image: null }
  editingProduct.value = {}
}

const deleteProduct = async (productId) => {
  if (confirm('Are you sure you want to delete this product?')) {
    try {
      const token = localStorage.getItem('token')

      await axios.delete(`http://localhost:3000/api/products/${productId}`, {
        headers: {
          'Authorization': `${token}`
        }
      })

      // Remove product from local state
      products.value = products.value.filter(p => p.id !== productId)
      alert('Product deleted successfully!')

    } catch (err) {
      console.error('Error deleting product:', err)
      alert(err.response?.data?.message || err.message || 'Failed to delete product')
    }
  }
}

// Lifecycle
onMounted(() => {
  fetchProducts()
})
</script>

<style scoped>
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}
</style>