<template>
  <div
    class="flex flex-col min-h-screen dark:bg-gray-900 shadow-md"
  >
    <header class="bg-gray-900 text-white p-4 dark:bg-gray-800">
      <!-- Navigation Bar -->
      <nav class="container mx-auto flex justify-between items-center">
        <!-- Logo / Branding -->
        <div class="text-lg font-bold">
          <NuxtLink to="/">
            Event Manager
          </NuxtLink>
        </div>
        <ThemeButton />
      </nav>
    </header>

    <main class="flex-grow">
      <div class="p-4 flex flex-col justify-center items-center bg-gray-100 dark:bg-gray-900">
        <div class="w-full max-w-md p-8 bg-white rounded-lg shadow-lg dark:bg-gray-800">
          <!-- login or singup -->
          <NuxtPage />
        </div>
      </div>
    </main>

    <footer class="bg-gray-800 text-white text-center p-4 dark:bg-gray-800">
      <p>&copy; 2024 My Event Site</p>
    </footer>
  </div>
</template>

<script setup>
import ThemeButton from '~/components/ThemeButton.vue'
</script>
