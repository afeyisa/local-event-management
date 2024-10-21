<template>
  <header class="bg-gray-900 text-white p-4 dark:bg-gray-800">
    <!-- Navigation Bar -->
    <nav class="container mx-auto flex justify-between items-center">
      <!-- Logo / Branding -->
      <div class="text-lg font-bold">
        <NuxtLink to="/">
          Event Manager
        </NuxtLink>
      </div>
      <div class="w-1/3">
        <SearchFilters />
      </div>
      <!-- Navigation Links -->
      <div class="relative flex items-center space-x-4">
        <ThemeButton />
        <ul class="flex items-center space-x-4">
          <!-- Conditional Rendering based on user login status -->
          <li>
            <NuxtLink
              to="/login"
              class="hover:underline"
            >
              Login
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/signup"
              class="hover:underline"
            >
              Signup
            </NuxtLink>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script setup>
import ThemeButton from '~/components/ThemeButton.vue'
</script>
