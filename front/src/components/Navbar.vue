<template>
  <nav class="bg-white border-gray-200 dark:bg-gray-900 fixed inset-x-0 top-0 z-50 shadow-md">
    <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
      <a href="#" class="flex items-center space-x-3 rtl:space-x-reverse">
        <img src="@/assets/logo.svg" class="h-8" alt="Flowbite Logo" />
        <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">
          Email Indexer
        </span>
      </a>
      <button
        @click="toggleMenu"
        data-collapse-toggle="navbar-default"
        type="button"
        class="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-t-light-purple  dark:text-gray-400 dark:hover:bg-gray-700"
        aria-controls="navbar-default"
        aria-expanded="false"
      >
        <span class="sr-only">Open main menu</span>
        <i class="pi pi-bars text-2xl"></i>
      </button>
      <div class="hidden w-full md:block md:w-auto" id="navbar-default">
        <ul
          class="font-medium flex items-center flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700"
        >
          <li v-for="item in menuItems" :key="item.name">
            <a
              href="#"
              class="block py-2 px-3 text-gray-900 rounded-sm hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-t-light-purple blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
              v-text="item.name"
            />
          </li>
          <li>
            <button
              id="theme-toggle"
              @click="toggleTheme"
              type="button"
                class="text-gray-900 inline-flex items-center justify-center dark:text-white hover:bg-white-100 w-10 md:hover:text-blue-700 dark:hover:text-white hover:bg-t-light-purple-shadow focus:outline-none focus:ring-4 focus:ring-t-light-purple rounded-lg text-sm p-2.5">
                  <i class="pi pi-sun"></i>
                  <i class="pi pi-moon"></i>
              <span class="sr-only">Toggle dark mode</span>
            </button>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script lang="ts" setup>

import { onMounted, ref } from 'vue';

// menu item list
const menuItems = ref([
  { name: 'Search', url: '#' },
]);

const toggleMenu = () => {
  const navbar = document.getElementById('navbar-default');
  if (navbar) {
    const expanded = navbar.getAttribute('aria-expanded') === 'true';
    navbar.setAttribute('aria-expanded', (!expanded).toString());
    navbar.classList.toggle('hidden', expanded);
  }
};

const toggleTheme = () => {
  const html = document.documentElement;
  const darkMode = html.classList.contains('dark');
  html.classList.toggle('dark', !darkMode);
  localStorage.setItem('dark', (!darkMode).toString());
  // store dark mode preference in local storage
  localStorage.setItem('dark', (!darkMode).toString());
};

onMounted(() => {
  // set dark mode based on local storage
  const darkMode = localStorage.getItem('dark') === 'true';
  const html = document.documentElement;
  html.classList.toggle('dark', darkMode);
});

</script>

<style scoped>
/* En modo light: mostramos la luna por defecto */
.pi.pi-moon {
  display: inline;
}
.pi.pi-sun {
  display: none;
}

/* En modo dark, usando un selector global para .dark */
:global(.dark) {
  .pi.pi-moon {
    display: none;
  }
  .pi.pi-sun {
    display: inline;
  }
}
</style>
