<script setup lang="ts">
import { computed, onMounted } from "vue";
import { useStockStore } from "./store/store-stock-selected";

import PageHeader from "./components/PageHeader.vue";
import StockList from "./components/StockList.vue";
import Pager from "./components/Pager.vue";
import type { Pagination } from "./models/pagination.model";
import LoadingModal from "./components/LoadingModal.vue";

const stockStore = useStockStore();

const pagination = computed<Pagination>(() => ({
  total: stockStore.total,
  page: stockStore.page,
  pageSize: stockStore.pageSize,
}));

onMounted(() => {
  stockStore.isRecomendation = false;
  stockStore.fetchStocks();
});
</script>

<template>
  <PageHeader />
  <div class="flex justify-center gap-4">
    <button
      class="w-70 px-3 py-2 bg-sky-800/70 border rounded hover:bg-green-600/70"
      @click="stockStore.toggleRecomendation"
    >
      <span class="font-bold text-white">
        {{
          stockStore.isRecomendation
            ? "Return to Stockrating"
            : "Get Recomendations per Score"
        }}
      </span>
    </button>
    <button
      v-if="!stockStore.isRecomendation"
      class="w-70 px-3 py-2 bg-sky-800/70 border rounded hover:bg-green-600/70"
      @click="stockStore.updateStocks"
    >
      <span class="font-bold text-white"> Update stock list </span>
    </button>
  </div>
  <Pager :pagination="pagination" @pageChange="stockStore.onPageChange" />
  <StockList :stocks="stockStore.stocks" />
  <LoadingModal v-if="stockStore.isUpdate" />
</template>
