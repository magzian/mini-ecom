<template>
  <div>
    <h2>Products</h2>
    <div v-if="error" class="error">{{ error }}</div>
    <ul>
      <li v-for="product in products" :key="product.id">
        <strong>{{ product.name }}</strong> - {{ product.price }}<br />
        <em>{{ product.description }}</em>
        <button @click="order(product.id)">Order</button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getProducts, createOrder } from '../api';
import { store } from '../store';

const products = ref([]);
const error = ref('');

const fetchProducts = async () => {
  try {
    const res = await getProducts(store.token);
    products.value = res.data;
  } catch (err) {
    error.value = 'Failed to load products';
  }
};

const order = async (productId) => {
  try {
    await createOrder({ product_id: productId, quantity: 1 }, store.token);
    alert('Order placed!');
  } catch (err) {
    error.value = 'Order failed';
  }
};

onMounted(fetchProducts);
</script>

<style scoped>
.error { color: red; }
ul { list-style: none; padding: 0; }
li { margin-bottom: 1rem; }
</style>
