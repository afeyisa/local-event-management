<template>
  <div class="loading-indicator fixed inset-0 flex flex-col items-center justify-center bg-white bg-opacity-80 dark:bg-gray-500 z-50">
    <div class="spinner animate-spin  w-16 h-16 border-4 border-t-transparent border-blue-500 rounded-full dark:bg-gray-500" />
    <p class="text-gray-700 mt-4 dark:bg-gray-500 dark:text-white">
      Loading...
    </p>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'auth',
  // middleware: 'logedinautoaedirector',
})
</script>
