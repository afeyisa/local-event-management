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
              to="/bookmarks"
              class="hover:underline hover:text-blue-400"
            >
              Bookmarks
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/create"
              class="hover:underline hover:text-blue-400"
            >
              +Create Event
            </NuxtLink>
          </li>
          <li class="relative">
            <UserProfile @click="toggleDropdown" />
            <!-- Profile Dropdown -->
            <div
              v-if="isDropdownOpen"
              class="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg py-2 text-gray-800 z-50 dark:bg-gray-700 dark:text-white"
            >
              <NuxtLink
                to="/profile"
                class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
              >
                Profile
              </NuxtLink>
              <button
                class="block w-full text-left px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
                @click="logout"
              >
                Logout
              </button>
            </div>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { useMutation } from '@vue/apollo-composable'
import { ref } from 'vue'
import ThemeButton from '~/components/ThemeButton.vue'
import { LOGOUT_MUTATION } from '~/graphql/mutation'

const isDropdownOpen = ref(false)

function toggleDropdown() {
  isDropdownOpen.value = !isDropdownOpen.value
}

async function logout() {
  console.log('logging out')
  const { mutate } = useMutation(LOGOUT_MUTATION)
  await mutate()
  // Logic for logging out the user
  console.log('User logged out')
  isDropdownOpen.value = false
}
</script>
