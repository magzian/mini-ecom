<script setup>
import { ref, onMounted } from "vue"
import axios from "axios"

const products = ref([])
const loading = ref(false)
const error = ref(null)

const fetchProducts = async () => {
  loading.value = true
  error.value = null

  try {
    // Get token from localStorage
    const token = localStorage.getItem("token")

    const response = await axios.get("http://localhost:3000/api/products/", {
      headers: {
        Authorization: `${token}`,
        "Content-Type": "application/json",
      },
    })

    products.value = response.data.products || [] // Adjust if your API structure is different
  } catch (err) {
    error.value =
      err.response?.data?.message ||
      err.message ||
      "Failed to fetch products"
    console.error("Error fetching products:", err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProducts()
})
</script>

<template>
  <section class="min-h-screen bg-white py-8 antialiased dark:bg-gray-900 md:py-16">
    <div class="mx-auto grid max-w-screen-xl px-4 pb-8 md:grid-cols-12 lg:gap-12 lg:pb-16 xl:gap-0">
      <div class="content-center justify-self-start md:col-span-7 md:text-start">
        <h1
          class="mb-4 text-4xl font-extrabold leading-none tracking-tight dark:text-white md:max-w-2xl md:text-5xl xl:text-6xl">
          Limited Time Offer!<br />Up to 50% OFF!</h1>
        <p class="mb-4 max-w-2xl text-gray-500 dark:text-gray-400 md:mb-12 md:text-lg mb-3 lg:mb-5 lg:text-xl">Don't
          Wait - Limited Stock at Unbeatable Prices!</p>
      </div>
      <div class="hidden md:col-span-5 md:mt-0 md:flex">
        <img class="dark:hidden" src="https://flowbite.s3.amazonaws.com/blocks/e-commerce/girl-shopping-list.svg"
          alt="shopping illustration" />
        <img class="hidden dark:block"
          src="https://flowbite.s3.amazonaws.com/blocks/e-commerce/girl-shopping-list-dark.svg"
          alt="shopping illustration" />
      </div>
    </div>
    
    <!-- Brand logos section -->
    <div
      class="mx-auto grid max-w-screen-xl grid-cols-2 gap-8 text-gray-500 dark:text-gray-400 sm:grid-cols-3 sm:gap-12 lg:grid-cols-6 px-4">
      <!-- Your existing brand logos here... -->
      <a href="#" class="flex items-center md:justify-center">
        <svg class="h-10 hover:text-gray-900 dark:hover:text-white" viewBox="0 0 106 48" fill="none"
          xmlns="http://www.w3.org/2000/svg">
          <!-- SVG content -->
        </svg>
      </a>
      <!-- Add other brand logos as needed -->
    </div>
  </section>

  <!-- Products Section -->
  <section class="py-16 px-4 mx-auto max-w-screen-xl">
    <div class="text-center mb-12">
      <h1 class="text-4xl font-bold text-gray-900 dark:text-dark mb-4">Featured Products</h1>
      <p class="text-lg text-gray-900 dark:text-gray-900">Discover our most popular items</p>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-8">
      <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      <p class="mt-2 text-gray-600">Loading products...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-8">
      <p class="text-red-600">{{ error }}</p>
      <button @click="fetchProducts" class="mt-2 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
        Try Again
      </button>
    </div>

    <!-- Products Grid -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="product in products"
        :key="product.id"
        class="bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 hover:shadow-lg transition-shadow duration-300"
      >
        <!-- Product Image -->
        <div class="relative overflow-hidden rounded-t-lg">
          <img 
            :src="product.image || 'https://fastly.picsum.photos/id/0/5000/3333.jpg?hmac=_j6ghY5fCfSD6tvtcV74zXivkJSPIfR9B8w34XeQmvU'" 
            :alt="product.name" 
            class="w-full h-48 object-cover hover:scale-105 transition-transform duration-300"
          />
        </div>
        
        <!-- Product Content -->
        <div class="p-5">
          <h5 class="mb-2 text-xl font-bold tracking-tight text-gray-900 dark:text-white line-clamp-2">
            {{ product.name }}
          </h5>
          <p class="mb-3 font-normal text-gray-700 dark:text-gray-400 line-clamp-3">
            {{ product.description }}
          </p>
          <div class="flex items-center justify-between">
            <span class="text-2xl font-bold text-blue-600 dark:text-blue-400">
              Ksh {{ product.price ? product.price.toLocaleString() : 'N/A' }}
            </span>
            <button 
              class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 transition-colors duration-200"
            >
              View Details
              <svg class="rtl:rotate-180 w-3.5 h-3.5 ms-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                viewBox="0 0 14 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M1 5h12m0 0L9 1m4 4L9 9" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-if="!loading && !error && products.length === 0" class="text-center py-12">
      <div class="text-gray-400 mb-4">
        <svg class="mx-auto h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-4m-4 0H8m-4 0h4m0 0v3m0-3h8m0 3v-3" />
        </svg>
      </div>
      <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-1">No products found</h3>
      <p class="text-gray-500 dark:text-gray-400">Check back later for new products.</p>
    </div>
  </section>
</template>