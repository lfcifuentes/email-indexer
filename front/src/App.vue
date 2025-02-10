<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue';
</script>

<template>
  <div class="dark:bg-gray-800 h-screen overflow-auto">
    <Navbar />
    <div class="container mx-auto pt-[80px] p-4">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
</style>
