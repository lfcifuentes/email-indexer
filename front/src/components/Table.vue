<template>
  <div v-if="emails.length > 0" class="relative overflow-x-auto shadow-md sm:rounded-lg">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
      <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
        <tr>
          <th scope="col" class="px-6 py-3">Id</th>
          <th scope="col" class="px-6 py-3">From</th>
          <th scope="col" class="px-6 py-3">To</th>
          <th scope="col" class="px-6 py-3">Subject</th>
          <th scope="col" class="px-6 py-3">Date</th>
          <th scope="col" class="px-6 py-3">Action</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="email in emails" :key="email.id">
          <!-- Primera fila: información resumida del email -->
          <tr
            class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700 border-gray-200"
          >
            <th
              scope="row"
              class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
              v-html="highlight(email.message_id)"
            ></th>
            <td class="px-6 py-4" v-html="highlight(email.from)"></td>
            <td class="px-6 py-4" v-html="highlight(email.to)"></td>
            <td class="px-6 py-4" v-html="highlight(email.subject)"></td>
            <td class="px-6 py-4">
              {{ email.date }}
            </td>
            <td class="px-6 py-4">
              <a
                @click="handleToggle('item-' + email.message_id)"
                href="#"
                class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
              >
                View
              </a>
            </td>
          </tr>
          <!-- Segunda fila: información adicional o de detalle -->
          <tr
            :id="'item-' + email.message_id"
            class="hidden odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700 border-gray-200"
          >
            <td colspan="6" class="px-6 py-4">
              <!-- Aquí puedes mostrar información extra, por ejemplo, detalles específicos del email -->
              Detalles:
              <br />
              <p
                v-html="
                  highlight(
                    parseBodyToHTML(
                      email.body || 'Sin detalles adicionales'
                    )
                  )
                "
                style="white-space: pre-wrap"
              ></p>
            </td>
          </tr>
        </template>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import type { Email } from '@/models/email';
import { defineProps } from 'vue';
import { highlightText } from '@/utils/highlight';
import { parseBodyToHTML } from '@/utils/text';
import { ref } from 'vue';

// props
interface Props {
  emails: Email[];
  search: string;
}
const props = defineProps<Props>();

const highlight = (text: string) => highlightText(text, props.search);

const handleToggle = (id: string) => {
  const element = document.getElementById(id);
  if (element) {
    element.classList.toggle('hidden');
  }
};
</script>
