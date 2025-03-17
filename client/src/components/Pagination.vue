<script setup>
import { computed } from "vue";

const props = defineProps({
  total: Number,  // Total de elementos
  page: Number,   // Página actual
  pageSize: Number, // Tamaño de página
});

const emit = defineEmits(["pageChange"]);

const totalPages = computed(() => Math.ceil(props.total / props.pageSize));


const goToPage = (newPage) => {
  if (newPage >= 1 && newPage <= totalPages.value) {
    emit("pageChange", newPage);
  }
};
</script>

<template>
  <div class="flex justify-center space-x-2 mt-4 mb-4">
    <button 
      class="px-3 py-1 border rounded bg-sky-200 disabled:opacity-50 enabled:hover:bg-sky-300" 
      :disabled="page <= 1" 
      @click="goToPage(page - 1)"
    >
      ⬅️ Prev
    </button>

    <span class="px-3 py-1 border rounded">{{ page }} / {{ totalPages }}</span>

    <button 
      class="px-3 py-1 border rounded bg-sky-200 disabled:opacity-50 enabled:hover:bg-sky-300" 
      :disabled="page >= totalPages" 
      @click="goToPage(page + 1)"
    >
      Next ➡️
    </button>
  </div>
</template>