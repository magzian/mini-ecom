<script setup>
import { ref, onMounted, computed } from "vue"
import axios from "axios"

const products = ref([])
const loading = ref(false)
const error = ref(null)
const currentProductIndex = ref(0)
const favorites = ref(new Set())

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

        products.value = response.data.products || []
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

const currentProduct = computed(() => {
    if (products.value.length === 0) return null
    return products.value[currentProductIndex.value]
})

const nextProduct = () => {
    if (products.value.length > 0) {
        currentProductIndex.value = (currentProductIndex.value + 1) % products.value.length
    }
}

const prevProduct = () => {
    if (products.value.length > 0) {
        currentProductIndex.value = currentProductIndex.value === 0
            ? products.value.length - 1
            : currentProductIndex.value - 1
    }
}

const toggleFavorite = (productId) => {
    if (favorites.value.has(productId)) {
        favorites.value.delete(productId)
    } else {
        favorites.value.add(productId)
    }
}

const isFavorite = (productId) => {
    return favorites.value.has(productId)
}

const formatPrice = (price) => {
    return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD'
    }).format(price)
}

const renderStars = (rating) => {
    const stars = []
    const fullStars = Math.floor(rating)
    const hasHalfStar = rating % 1 !== 0

    for (let i = 0; i < 5; i++) {
        if (i < fullStars) {
            stars.push('full')
        } else if (i === fullStars && hasHalfStar) {
            stars.push('half')
        } else {
            stars.push('empty')
        }
    }
    return stars
}

const addToCart = async (productId) => {
    try {
        const token = localStorage.getItem("token")
        await axios.post(`http://localhost:3000/api/cart/add`,
            { productId, quantity: 1 },
            {
                headers: {
                    Authorization: `${token}`,
                    "Content-Type": "application/json",
                },
            }
        )
        alert('Product added to cart!')
    } catch (err) {
        console.error('Error adding to cart:', err)
        alert('Failed to add product to cart')
    }
}

onMounted(() => {
    fetchProducts()
})
</script>

<template>
    <section class="h-[100vh] py-8 bg-white md:py-16 dark:bg-gray-900 antialiased">
        <div class="max-w-screen-xl px-4 mx-auto 2xl:px-0">

            <!-- Loading State -->
            <div v-if="loading" class="flex justify-center items-center h-64">
                <div class="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600"></div>
            </div>

            <!-- Error State -->
            <div v-else-if="error" class="text-center text-red-600 dark:text-red-400">
                <p class="text-xl font-semibold">{{ error }}</p>
                <button @click="fetchProducts"
                    class="mt-4 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
                    Retry
                </button>
            </div>

            <!-- No Products -->
            <div v-else-if="products.length === 0" class="text-center text-gray-600 dark:text-gray-400">
                <p class="text-xl">No products available</p>
            </div>

            <!-- Product Display -->
            <div v-else class="lg:grid lg:grid-cols-2 lg:gap-8 xl:gap-16">
                <div class="shrink-0 max-w-md lg:max-w-lg mx-auto relative">
                    <!-- Product Image -->
                    <img v-if="currentProduct.image" :src="currentProduct.image" :alt="currentProduct.name"
                        class="w-full rounded-lg shadow-lg object-cover h-96" />
                    <div v-else
                        class="w-full h-96 bg-gray-200 dark:bg-gray-700 rounded-lg flex items-center justify-center">
                        <span class="text-gray-500 dark:text-gray-400">No Image</span>
                    </div>

                    <!-- Product Navigation -->
<!--                     <div v-if="products.length > 1"
                        class="absolute top-1/2 transform -translate-y-1/2 w-full flex justify-between px-4">
                        <button @click="prevProduct"
                            class="bg-white dark:bg-gray-800 p-2 rounded-full shadow-lg hover:bg-gray-50 dark:hover:bg-gray-700">
                            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M15 19l-7-7 7-7"></path>
                            </svg>
                        </button>
                        <button @click="nextProduct"
                            class="bg-white dark:bg-gray-800 p-2 rounded-full shadow-lg hover:bg-gray-50 dark:hover:bg-gray-700">
                            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7">
                                </path>
                            </svg>
                        </button>
                    </div> -->

                    <!-- Product Counter -->
                    <!-- <div v-if="products.length > 1"
                        class="absolute bottom-4 left-1/2 transform -translate-x-1/2 bg-black bg-opacity-50 text-white px-3 py-1 rounded-full text-sm">
                        {{ currentProductIndex + 1 }} / {{ products.length }}
                    </div> -->
                </div>

                <div class="mt-6 sm:mt-8 lg:mt-0">
                    <!-- Product Name -->
                    <h1 class="text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white">
                        {{ currentProduct.name || 'Product Name Not Available' }}
                    </h1>

                    <div class="mt-4 sm:items-center sm:gap-4 sm:flex">
                        <!-- Price -->
                        <p class="text-2xl font-extrabold text-gray-900 sm:text-3xl dark:text-white">
                            {{ currentProduct.price ? formatPrice(currentProduct.price) : 'Price not available' }}
                        </p>

                        <!-- Rating -->
                        <div v-if="currentProduct.rating" class="flex items-center gap-2 mt-2 sm:mt-0">
                            <div class="flex items-center gap-1">
                                <svg v-for="(star, index) in renderStars(currentProduct.rating)" :key="index"
                                    class="w-4 h-4"
                                    :class="star === 'full' ? 'text-yellow-300' : star === 'half' ? 'text-yellow-300' : 'text-gray-300'"
                                    aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                                    fill="currentColor" viewBox="0 0 24 24">
                                    <path
                                        d="M13.849 4.22c-.684-1.626-3.014-1.626-3.698 0L8.397 8.387l-4.552.361c-1.775.14-2.495 2.331-1.142 3.477l3.468 2.937-1.06 4.392c-.413 1.713 1.472 3.067 2.992 2.149L12 19.35l3.897 2.354c1.52.918 3.405-.436 2.992-2.15l-1.06-4.39 3.468-2.938c1.353-1.146.633-3.336-1.142-3.477l-4.552-.36-1.754-4.17Z" />
                                </svg>
                            </div>
                            <p class="text-sm font-medium leading-none text-gray-500 dark:text-gray-400">
                                ({{ currentProduct.rating.toFixed(1) }})
                            </p>
                            <span v-if="currentProduct.reviewCount"
                                class="text-sm font-medium leading-none text-gray-900 underline hover:no-underline dark:text-white cursor-pointer">
                                {{ currentProduct.reviewCount }} Reviews
                            </span>
                        </div>
                    </div>

                    <!-- Stock Status -->
                    <div v-if="currentProduct.stock !== undefined" class="mt-2">
                        <span v-if="currentProduct.stock > 0"
                            class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300">
                            {{ currentProduct.stock }} in stock
                        </span>
                        <span v-else
                            class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300">
                            Out of stock
                        </span>
                    </div>

                    <!-- Action Buttons -->
                    <div class="mt-6 sm:gap-4 sm:items-center sm:flex sm:mt-8">
                        <button @click="toggleFavorite(currentProduct._id || currentProduct.id)" :class="[
                            'flex items-center justify-center py-2.5 px-5 text-sm font-medium focus:outline-none rounded-lg border focus:z-10 focus:ring-4 transition-colors',
                            isFavorite(currentProduct._id || currentProduct.id)
                                ? 'text-red-700 bg-red-50 border-red-200 hover:bg-red-100 focus:ring-red-100 dark:bg-red-900 dark:text-red-300 dark:border-red-600 dark:hover:bg-red-800'
                                : 'text-gray-900 bg-white border-gray-200 hover:bg-gray-100 hover:text-primary-700 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700'
                        ]">
                            <svg class="w-5 h-5 -ms-2 me-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
                                width="24" height="24" fill="none" viewBox="0 0 24 24">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M12.01 6.001C6.5 1 1 8 5.782 13.001L12.011 20l6.23-7C23 8 17.5 1 12.01 6.002Z" />
                            </svg>
                            {{ isFavorite(currentProduct._id || currentProduct.id) ? 'Remove from favorites' : 'Add to favorites'
                             }}
                        </button>

                        <button @click="addToCart(currentProduct._id || currentProduct.id)"
                            :disabled="currentProduct.stock === 0" :class="[
                                'text-white mt-4 sm:mt-0 font-medium rounded-lg text-sm px-5 py-2.5 flex items-center justify-center focus:outline-none transition-colors',
                                currentProduct.stock === 0
                                    ? 'bg-gray-400 cursor-not-allowed'
                                    : 'bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800'
                            ]">
                            <svg class="w-5 h-5 -ms-2 me-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
                                width="24" height="24" fill="none" viewBox="0 0 24 24">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M4 4h1.5L8 16m0 0h8m-8 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4Zm8 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4Zm.75-3H7.5M11 7H6.312M17 4v6m-3-3h6" />
                            </svg>
                            {{ currentProduct.stock === 0 ? 'Out of Stock' : 'Add to cart' }}
                        </button>
                    </div>

                    <hr class="my-6 md:my-8 border-gray-200 dark:border-gray-800" />

                    <!-- Product Description -->
                    <div v-if="currentProduct.description" class="mb-6 text-gray-500 dark:text-gray-400">
                        <p v-for="(paragraph, index) in currentProduct.description.split('\n')" :key="index"
                            class="mb-4 last:mb-0">
                            {{ paragraph }}
                        </p>
                    </div>
                    <p v-else class="mb-6 text-gray-500 dark:text-gray-400">
                        No description available for this product.
                    </p>

                    <!-- Product Specifications -->
                    <div v-if="currentProduct.specifications" class="space-y-2">
                        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-3">Specifications</h3>
                        <div v-for="(value, key) in currentProduct.specifications" :key="key"
                            class="flex justify-between">
                            <span class="text-gray-600 dark:text-gray-400 capitalize">{{ key.replace(/([A-Z])/g, 
                                '$1').trim() }}:</span>
                            <span class="text-gray-900 dark:text-white font-medium">{{ value }}</span>
                        </div>
                    </div>

                    <!-- Product Category -->
                    <div v-if="currentProduct.category" class="mt-4">
                        <span
                            class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300">
                            {{ currentProduct.category }}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<style scoped>
.bg-primary-700 {
    background-color: #1d4ed8;
}

.bg-primary-800 {
    background-color: #1e40af;
}

.bg-primary-600 {
    background-color: #2563eb;
}

.hover\:bg-primary-800:hover {
    background-color: #1e40af;
}

.hover\:bg-primary-700:hover {
    background-color: #1d4ed8;
}

.focus\:ring-primary-300:focus {
    --tw-ring-color: #93c5fd;
}

.dark\:bg-primary-600.dark {
    background-color: #2563eb;
}

.dark\:hover\:bg-primary-700.dark:hover {
    background-color: #1d4ed8;
}

.dark\:focus\:ring-primary-800.dark:focus {
    --tw-ring-color: #1e40af;
}
</style>