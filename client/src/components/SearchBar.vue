<script setup lang="ts">
const props = defineProps<{
  toggleSearching: (ticker: string) => void;
}>();

import { ref } from 'vue';

const searchQuery = ref('');

const handleSearch = () => {
  console.log('Search query:', searchQuery.value);
  
  props.toggleSearching(searchQuery.value);
};

</script>

<template>
  <div class="container max-w-xl mx-auto mt-2 flex items-center bg-gray-100 rounded-lg p-2 shadow-sm">
    <!-- Ícono de búsqueda -->
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5 text-gray-500 mr-2"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M11 4a7 7 0 100 14 7 7 0 000-14zm0 0l6 6"
      />
    </svg>

    <input
      type="text"
      v-model="searchQuery"
      @input="handleSearch"
      placeholder="Search stock by Ticker..."
      class="flex-1 bg-transparent outline-none text-gray-700 placeholder-gray-500"
    />
  </div>
</template>
