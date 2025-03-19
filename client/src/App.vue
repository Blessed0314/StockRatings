<script setup lang="ts">
  import { onMounted } from "vue";
  import { useStockStore } from "./store/store-stock-selected";

  import PageHeader from "./components/PageHeader.vue";
  import StockList from "./components/StockList.vue";
  import Pagination from "./components/Pagination.vue";
 
  const stockStore = useStockStore();

  onMounted(() => {
    stockStore.isRecomendation = false;
    stockStore.fetchStocks();
  });

</script>

<template>
  <PageHeader />
  <div class="flex justify-center">
    <button class="px-3 py-2 bg-red-300 border rounded hover:bg-red-500" @click="stockStore.toggleRecomendation">
      <span class="font-bold">
        {{stockStore.isRecomendation ? "Return to Stockrating" :"Get Recomendations per Score"}}
      </span>
    </button>
  </div>
  <Pagination 
    :total="stockStore.total" 
    :page="stockStore.page" 
    :pageSize="stockStore.pageSize" 
    @pageChange="stockStore.onPageChange" 
  />
  <StockList :stocks="stockStore.stocks" />
</template>
