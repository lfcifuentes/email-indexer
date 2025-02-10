<script setup lang="ts">
import Button from '@/components/Button.vue';
import EmptyMessage from '@/components/EmptyMessage.vue';
import SearchMessage from '@/components/SearchMessage.vue';
import Table from '@/components/Table.vue';
import Cards from '@/components/Cards.vue';
import Pagination from '@/components/Pagination.vue';
import type { Email } from '@/models/email';
import { getApi } from '@/services/api';
import { ref } from 'vue';
import { PageLink } from '@/models/pagination';
import EmailModal from '@/components/EmailModal.vue';

const search = ref('');
const email = ref<Email>({} as Email);
const emails = ref<Email[]>([]);
const meta = ref({
  current_page: 1,
  last_page: 1,
  total: 0,
});
const view = ref('list');
const loading = ref(false);
const page = ref(1);
const highlight = ref('');

const formHandler = () => {
  page.value = 1
  searchHandler()
}

const paginationHandler = (e: PageLink) => {
  page.value = e.label as number
  searchHandler()
}

const searchHandler = () => {
  loading.value = true
  getApi()
    .get(`/search?query=${search.value}&page=${page.value}`)
    .then(response => {
      highlight.value = search.value
      emails.value = response?.data?.data?.emails
      meta.value = {
        current_page: page.value,
        last_page: response?.data?.data?.last_page,
        total: response?.data?.data?.total,
      }
    })
    .catch(err => {
      console.log(err)
    })
    .finally(() => {
      loading.value = false
    })
}

const openEmailModal = (payload: { email: Email }) => {
  email.value = payload.email
}
const closeEmailModal = () => {
  email.value = {} as Email
}
</script>

<template>
  <main>
    <!-- search bar -->
    <form
      @submit.prevent="formHandler"
      class="max-w-md mx-auto">
      <label for="search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
      <div class="relative">
          <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
              <i class="pi pi-search text-gray-400 dark:text-gray-500"></i>
          </div>
          <input
            v-model="search"
            id="search"
            type="search"
            class="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder="Search Email by Subject, Sender, Receiver, etc."
            required
          />
          <Button
            :loading="loading"
            type="submit"
            class="absolute"
          >Search</button>
      </div>
    </form>
    <!-- email list -->
    <div class="mt-8">
      <EmptyMessage  v-if="emails.length === 0 && highlight !== ''" />
      <SearchMessage v-else-if="emails.length === 0" />
      <div v-if="emails.length > 0"  class="text-sm text-gray-900 dark:text-white mb-2">
        <span class="">
          Showing
          <span class="font-medium">{{ meta.current_page }}</span>
          to
          <span class="font-medium">{{ meta.last_page }}</span>
          pages out
          of <span class="font-medium">{{ meta.total }}</span> results
        </span>
      <div class="flex justify-end gap-2">
        <Button
          @click="view = 'list'"
          class="mt-4 ml-2"
          :class="{
            'ring-2 ring-purple-300': view === 'list',
            'opacity-60': view === 'grid'
          }"
        >
          <i class="pi pi-list"></i>
          List View
        </Button>
        <Button
          @click="view = 'grid'"
          class="mt-4"
          :class="{
            'ring-2 ring-purple-300': view === 'grid',
            'opacity-60': view === 'list'
          }"
        >
          <i class="pi pi-th-large"></i>
        Grid View
        </Button>
      </div>
      </div>
      <Table
        v-if="emails.length > 0 && view === 'list'"
        :emails="emails"
        :search="highlight"
      />
      <Cards
        v-else-if="emails.length > 0 && view === 'grid'"
        :emails="emails"
        :search="highlight"
        @view-email="openEmailModal"
      />
      <pagination
        v-if="emails.length > 0"
        :meta="meta"
        @onPaginate="paginationHandler"
      />
    </div>
    <!-- email modal -->
    <EmailModal
      v-if="Object.keys(email).length > 0"
      :email="email"
      :search="highlight"
      @closeEmailModal="closeEmailModal"
    />
  </main>
</template>
