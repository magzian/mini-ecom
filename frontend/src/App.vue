<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script setup>
// No logic needed here, handled by router
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* Remove default margin from h1 tags */
h1 {
  margin: 0;
}
</style>
