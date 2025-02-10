<template>
  <button
    :type="type"
    class="text-white absolute end-2.5 bottom-2.5 bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:outline-none focus:ring-purple-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-800"
    :class="{
      'cursor-not-allowed opacity-80': props.loading
    }">
    <template v-if="props.loading">
      <i class="pi pi-spin pi-spinner m-1 block"></i>
    </template>
    <slot>
      {{ label }}
    </slot>
  </button>

</template>

<script setup lang="ts">
interface Props {
  label?: string;
  loading?: boolean;
  type?: "submit" | "button" | "reset" | undefined;
}

const props = defineProps<Props>()
const type = props.type || "submit"
const label = props.label || ""
</script>
