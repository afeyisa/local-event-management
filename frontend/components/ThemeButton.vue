<script setup>
import { ref, onMounted } from 'vue'

const theme = ref('light')

function toggleTheme() {
  theme.value = theme.value === 'dark' ? 'light' : 'dark'
  document.documentElement.classList.toggle('dark', theme.value === 'dark')
  localStorage.setItem('theme', theme.value)
}

onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme) {
    theme.value = savedTheme
    document.documentElement.classList.toggle('dark', savedTheme === 'dark')
  }
})
</script>

<template>
  <div
    :class="theme"
    class="relative flex items-center space-x-4"
  >
    <button
      class="focus:outline-none"
      @click="toggleTheme"
    >
      <span
        v-if="theme === 'dark'"
        class="text-yellow-500"
      ><i
        class="fa fa-sun-o"
        aria-hidden="true"
      /></span>
      <span
        v-else
        class="text-gray-600"
      ><i
        class="fa fa-moon-o"
        aria-hidden="true"
      /></span>
    </button>
  </div>
</template>
