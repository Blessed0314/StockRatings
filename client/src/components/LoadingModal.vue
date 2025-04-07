<script setup lang="ts">
  import loaderGif from '@/assets/loader.gif'
</script>

<template>
  <div class="fixed inset-0 flex items-center justify-center bg-black/50 z-50">
    <div
      class="flex flex-col items-center justify-center bg-white p-4 rounded shadow-lg"
    >
      <img :src="loaderGif" alt="Loading" class="w-48 h-48" />
      <p class="mt-2 text-lg font-semibold text-gray-700">Loading...</p>
    </div>
  </div>
</template>
