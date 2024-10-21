<template>
  <div>
    <ThemeButton />
    <li>
      <NuxtLink
        to="/bookmarks"
        class="hover:underline"
      >
        Bookmarks
      </NuxtLink>
    </li>
    <li>
      <NuxtLink
        to="/create-event"
        class="hover:underline"
      >
        Create Event
      </NuxtLink>
    </li>
    <!-- User Profile Dropdown -->
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
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ThemeButton from '~/components/ThemeButton.vue'

definePageMeta({
  middleware: 'auth',
})

const isDropdownOpen = ref(false)

function toggleDropdown() {
  isDropdownOpen.value = !isDropdownOpen.value
}

function logout() {
  // Logic for logging out the user
  console.log('User logged out')
  isDropdownOpen.value = false
}
</script>

<!--
<template>
  <div>
    <h2 class="text-2xl font-bold text-center text-gray-800 dark:text-white mb-6">
      Signup
    </h2>
    <Form
      @submit="onSummit"
    >

      <div class="mb-4">
        <label
          class="block text-gray-700 dark:text-gray-300 mb-2"
          for="email"
        >Email
        </label>
        <Field
          v-model="email"
          rules="validateEmail"
          type="email"
          name="email"
          r
          class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-300 dark:bg-gray-700 dark:text-white dark:border-gray-600"
          placeholder="Enter your email"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="email"
        />
      </div>
      <div class="mb-4">
        <label
          for="password"
          class="block text-gray-700 dark:text-gray-300 mb-2"
        >
          Password
        </label>
        <Field
          v-model="password"
          name="password"
          type="password"
          rules="passordValidate"
          placeholder="Password"
          class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-300 dark:bg-gray-700 dark:text-white dark:border-gray-600"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="password"
        />
      </div>
      <div class="mb-4">
        <label
          class="block text-gray-700 dark:text-gray-300 mb-2"
          for="confirm password"
        >
          Confirm Password
        </label>
        <Field
          name="confirmation"
          type="password"
          placeholder="Confirm Password"
          rules="passordValidate|confirmed:password"
          class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-300 dark:bg-gray-700 dark:text-white dark:border-gray-600"
        />
      </div>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="confirmation"
      />

      <button
        :disabled="loading"
        class="w-full bg-blue-500 text-white py-2 hover:bg-blue-600 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700"
      >
        Signup
      </button>
    </Form>
    <div
      v-if="error"
      class=" text-red-500 mb-4"
    >
      {{ error.message }}
    </div>

    <div class="mt-6 text-center">
      <NuxtLink
        to="/login"
        class="text-blue-500 hover:underline dark:text-blue-400"
      >
        Already have an account? Log in
      </NuxtLink>
    </div>
  </div>
</template> -->
