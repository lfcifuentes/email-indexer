<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
    <div
      v-for="email in emails"
      :key="email.id"
      class="block p-6 bg-white border border-gray-200 rounded-lg shadow-sm hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
    >
      <div class="mb-2">
        <h2 class="text-lg font-bold whitespace-normal break-words text-gray-900 dark:text-white">
          {{ email.message_id }}
        </h2>
      </div>
      <div class="text-sm text-gray-700 dark:text-gray-300">
        <div class="flex items-center gap-2">
          <i class="pi pi-user"></i>
          <b>From:</b>
          <p v-html="highlight(email.from) || '-Sin Registro-'" class="truncate"></p>
        </div>
        <div class="flex items-center gap-2">
          <i class="pi pi-users"></i>
          <b>To:</b>
          <p v-html="highlight(email.to) || '-Sin Registro-'" class="truncate"></p>
        </div>
        <div class="flex items-center gap-2">
          <i class="pi pi-envelope"></i>
          <b>Subject:</b>
          <p v-html="highlight(email.subject) || '-Sin Registro-'" class="truncate"></p>
        </div>
      </div>
      <div class="mt-3 flex items-center justify-between">
        <div class="text-sm text-gray-700 dark:text-gray-300 flex items-center gap-2">
          <i class="pi pi-calendar"></i>
          <span class="text-gray-500 dark:text-gray-400 text-xs">{{ email.date }}</span>
        </div>
        <button
          class="text-blue-600 dark:text-blue-400 font-medium hover:underline cursor-pointer"
          @click="$emit('viewEmail', { email })"
        >
          View Email Details
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Email } from '@/models/email';
import { defineProps, defineEmits } from 'vue';
import { highlightText } from '@/utils/highlight';

interface Props {
  emails: Email[];
  search: string;
}
const props = defineProps<Props>();

const emit = defineEmits<{
  (e: 'viewEmail', payload: { email: Email }): void;
}>();

const highlight = (text: string): string => highlightText(text, props.search);
</script>
