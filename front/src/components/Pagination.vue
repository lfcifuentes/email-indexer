<template>
  <nav aria-label="Page navigation" class="flex items-center justify-center mt-4">
    <ul class="inline-flex -space-x-px text-sm">
      <li>
        <button
          :disabled="meta?.current_page === 1"
          @click="clickButton({ label: 1 })"
          class="flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-e-0 border-gray-300 rounded-s-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
          :class="{
            'cursor-not-allowed': meta?.current_page === 1,
            'cursor-pointer': meta?.current_page !== 1
          }"
        >
          <i class="pi pi-angle-double-left"></i>
        </button>
      </li>
      <li>
        <button
          :disabled="meta?.current_page === 1"
          @click="clickButton(prevPage)"
          class="flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-e-0 border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
          :class="{
            'cursor-not-allowed': meta?.current_page === 1,
            'cursor-pointer': meta?.current_page !== 1
          }"
        >
          <i class="pi pi-angle-left"></i>
        </button>
      </li>
      <li v-for="link in pageLinks" :key="link.label">
        <button
          v-if="link.index !== ''"
          class="flex items-center justify-center px-3 h-8 leading-tight border"
          :class="{
            'text-blue-600 border-gray-300 bg-blue-50 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white dark:hover:bg-gray-500 dark:hover:text-white': link.active,
            'text-gray-500 bg-white border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white cursor-pointer' : !link.active
          }"
          @click="clickButton(link)"
        >
          {{ link.label  || '...'  }}
        </button>
      </li>
      <li>
        <button
          :disabled="meta?.current_page === meta.last_page"
          @click="clickButton(nextPage)"
          class="flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-e-0 border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
          :class="{
            'cursor-not-allowed': meta?.current_page === meta.last_page,
            'cursor-pointer': meta?.current_page !== meta.last_page
          }"
        >
          <i class="pi pi-angle-right"></i>
        </button>
      </li>
      <li>
        <button
          :disabled="meta?.current_page === meta.last_page"
          @click="clickButton({ label: meta.last_page })"
          class="flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-e-0 border-gray-300 rounded-e-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
          :class="{
            'cursor-not-allowed': meta?.current_page === meta.last_page,
            'cursor-pointer': meta?.current_page !== meta.last_page
          }"
        >
          <i class="pi pi-angle-double-right"></i>
        </button>
      </li>
    </ul>
  </nav>
</template>

<script setup lang="ts">
import { defineProps, computed, defineEmits } from 'vue';
import {Meta, PageLink} from '@/models/pagination';

const props = defineProps<{ meta: Meta }>();

const emit = defineEmits<{
  (e: 'onPaginate', payload: { label: number }): void;
}>();
const trimPageRange = (pageRange: PageLink[]) => {
  // If it's the first or last page, just get the beginning and end
  if (props.meta.current_page < 3 || props.meta.current_page > pageRange.length - 2) {
    const beginning = pageRange.slice(0, 3);
    beginning.push({ index: "1" });
    const end = pageRange.slice(pageRange.length - 3);

    return beginning.concat(end);
  }

  const first = pageRange.slice(0, 1);
  first.push({ index: "1"});

  const middle = pageRange.slice(props.meta.current_page - 2, props.meta.current_page + 1);
  middle.push({ index:  "1" });

  const last = pageRange.slice(pageRange.length - 1);

  return first.concat(middle, last);
};
const pageLinks = computed<PageLink[]>(() => {
  const links: PageLink[] = [];
  const pageLimit = 7;
  const start = 1;
  const end = props.meta.last_page;
  for (let i = start; i <= end; i++) {
    links.push({
      index: i,
      label: i,
      active: i === props.meta.current_page
    });
  }
  return links.length > pageLimit ? trimPageRange(links) : links;
});

const clickButton = (link: PageLink | { label: number }) => {
  const page = typeof link.label === 'number' ? link.label : parseInt(link.label as string, 10);
  emit('onPaginate', { label: page });
};

const prevPage = computed(() => {
  return { label: Math.max(props.meta.current_page - 1, 1) };
});
const nextPage = computed(() => {
  return { label: Math.min(props.meta.current_page + 1, props.meta.last_page) };
});
</script>
