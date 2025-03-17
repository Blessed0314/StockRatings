<script setup lang="ts">
  import { defineProps } from 'vue';
import type { Stock } from '../models/stock.model';
  interface Props {
    stock: Stock;
    columns: string[];
  }

  defineProps<Props>();

  function formatCurrency(value: number): string {
    return `$${value.toFixed(2)}`;
  }
</script>
<template>
  <tr class="border-b border-gray-200 hover:bg-gray-100 transition">
    <td v-for="(col, index) in columns" :key="col" :class="{'font-bold': index === 0}" class="py-3 px-6">
      {{ col === 'target_from' || col === 'target_to' ? formatCurrency(stock[col as keyof Stock] as number) : stock[col as keyof Stock] }}
    </td>
  </tr>
</template>