<script setup lang="ts">
import { computed } from "vue";
import type { Pagination } from "../models/pagination.model";

// Definir los props usando la interfaz Pagination
const props = defineProps<{
  pagination: Pagination;
}>();

const emit = defineEmits(["pageChange"]);

// Calcular el total de páginas
const totalPages = computed(() => Math.ceil(props.pagination.total / props.pagination.pageSize));

// Cambiar de página
const goToPage = (newPage: number) => {
  if (newPage >= 1 && newPage <= totalPages.value) {
    emit("pageChange", newPage);
  }
};
</script>

<template>
  <div class="flex justify-center space-x-2 mt-4 mb-4">
    <button 
      class="px-3 py-1 border rounded bg-sky-200 disabled:opacity-50 enabled:hover:bg-sky-300" 
      :disabled="pagination.page <= 1" 
      @click="goToPage(pagination.page - 1)"
    >
      ⬅️ Prev
    </button>

    <span class="px-3 py-1 border rounded">{{ pagination.page }} / {{ totalPages }}</span>

    <button 
      class="px-3 py-1 border rounded bg-sky-200 disabled:opacity-50 enabled:hover:bg-sky-300" 
      :disabled="pagination.page >= totalPages" 
      @click="goToPage(pagination.page + 1)"
    >
      Next ➡️
    </button>
  </div>
</template>