<template>
  <div
    tabindex="-1"
    class="flex overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 max-h-full h-screen backdrop-blur-[2px] bg-black/50 "
    @click.self="closeEmailModal"
  >
    <div class="relative p-4 w-full max-w-[700px] max-h-full">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow-sm dark:bg-gray-700">
          <!-- Modal header -->
          <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600 border-gray-200">
              <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
                  Email Details
              </h3>
              <button
                @click="closeEmailModal"
                type="button" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="default-modal">
                  <i class="pi pi-times"></i>
                  <span class="sr-only">Close modal</span>
              </button>
          </div>
          <!-- Modal body -->
          <div class="p-4 md:p-5 space-y-4 overflow-auto max-h-[67vh]">
              <p
              class="text-base leading-relaxed text-gray-500 dark:text-gray-400"
                v-html="
                  highlight(
                    parseBodyToHTML(
                      email.body || 'Sin detalles adicionales'
                    )
                  )
                "
                style="white-space: pre-wrap"
              ></p>
          </div>
          <!-- Modal footer -->
          <div class="flex items-center p-4 md:p-5 border-t border-gray-200 rounded-b dark:border-gray-600">
              <Button
                @click="closeEmailModal"
                type="button"
              >
                  <i class="pi pi-times"></i>
                  Close
              </button>
          </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { Email } from '@/models/email';
import { highlightText } from '@/utils/highlight';
import { parseBodyToHTML } from '@/utils/text';
import Button from '@/components/Button.vue';

// props
interface Props {
  email: Email;
  search: string;
}
const props = defineProps<Props>();

const emit = defineEmits<{
  (e: 'closeEmailModal'): void;
}>();

const highlight = (text: string) => highlightText(text, props.search);

const closeEmailModal = () => {
  // emit close email modal event
  emit('closeEmailModal')
}

</script>

<style scoped>
</style>
