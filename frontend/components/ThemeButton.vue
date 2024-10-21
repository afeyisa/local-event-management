<template>
  <div
    :class="theme"
    class="relative flex items-center space-x-4"
  >
    <!-- Dark/Light Mode Toggle -->
    <button
      class="focus:outline-none"
      @click="toggleTheme"
    >
      <span
        v-if="theme === 'dark'"
        class="text-yellow-500"
      >🌞</span>
      <span
        v-else
        class="text-gray-200"
      >🌜</span>
    </button>
  </div>
</template>

<script setup>
// Importing ref to create reactive state
import { ref, onMounted } from 'vue'

const theme = ref('light') // Default theme

// Toggles the theme between light and dark mode
function toggleTheme() {
  theme.value = theme.value === 'dark' ? 'light' : 'dark'
  document.documentElement.classList.toggle('dark', theme.value === 'dark')
  localStorage.setItem('theme', theme.value) // Persist theme in local storage
}
onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme) {
    theme.value = savedTheme
    document.documentElement.classList.toggle('dark', savedTheme === 'dark')
  }
})
</script>
