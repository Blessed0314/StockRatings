<script setup lang="ts">
import { computed, defineProps } from "vue";
import StockItem from "./StockItem.vue";

import type { Stock } from "../models/stock.model";
interface Props {
  stocks: Stock[];
}
const props = defineProps<Props>();

const columns = computed(() =>
  props.stocks.length > 0 ? Object.keys(props.stocks[0]) : []
);
</script>

<template>
  <div class="container max-w-7xl mx-auto">
    <div class="overflow-x-auto">
      <table
        class="min-w-full bg-white border border-gray-200 shadow-lg rounded-lg"
      >
        <thead>
          <tr
            class="bg-gray-100 text-gray-700 uppercase text-sm leading-normal"
          >
            <th v-for="col in columns" :key="col" class="py-3 px-6 text-left">
              {{ col.replace(/_/g, " ").toUpperCase() }}
            </th>
          </tr>
        </thead>
        <tbody class="text-gray-600 text-sm">
          <StockItem
            v-for="stock in props.stocks"
            :key="stock.ticker"
            :stock="stock"
            :columns="columns"
          />
        </tbody>
      </table>
    </div>
  </div>
</template>
