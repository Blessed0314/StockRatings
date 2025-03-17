<script setup lang="ts">
  import { onMounted, ref } from "vue";

  import PageHeader from "./components/PageHeader.vue";
  import StockList from "./components/StockList.vue";
  import Pagination from "./components/Pagination.vue";
  import apiService from "./utils/api-service";
  import type { Stock } from "./models/stock.model";

  const stocks = ref<Stock[]>([]);
  const total = ref<number>(0);
  const page = ref<number>(1);
  const pageSize = ref<number>(10);
  const isRecomendation = ref<boolean>(false);

  const fetchStocks = async () => {
    try {
      const { data } = isRecomendation.value 
        ? await apiService.get(`/stock/recommendations?page=${page.value}&size=${pageSize.value}`)
        : await apiService.get(`/stock/all?page=${page.value}&size=${pageSize.value}`);
      console.log(data);
        
      stocks.value = data.data.stocks;
      total.value = data.data.total;
    } catch (error) {
      console.log(error);
    }
  };

  onMounted(() => {
    isRecomendation.value = false;
    fetchStocks();
  });

  const onPageChange = (newPage: number) => {
    page.value = newPage;
    fetchStocks();
  };

  const toggleRecomendation = () => {
    isRecomendation.value = !isRecomendation.value;
    page.value = 1;
    fetchStocks();
  };
</script>

<template>
  <PageHeader />
  <div class="flex justify-center">
    <button class="px-3 py-2 bg-red-300 border rounded hover:bg-red-500" @click="toggleRecomendation">
      <span class="font-bold">
        {{isRecomendation ? "Return to Stockrating" :"Get Recomendations per Score"}}
      </span>
    </button>
  </div>
  <Pagination 
    :total="total" 
    :page="page" 
    :pageSize="pageSize" 
    @pageChange="onPageChange" 
  />
  <StockList :stocks="stocks" />
</template>
