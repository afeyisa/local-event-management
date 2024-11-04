<template>
  <ClientOnly>
    <div
      class="flex flex-col"
    >
      <header class="bg-blue-500 text-white p-4 dark:bg-gray-800 sticky top-0 z-10 shadow-md">
        <nav class="container mx-auto flex justify-between items-center flex-wrap">
          <div class="text-lg font-bold">
            <NuxtLink
              to="/"
              class="hover:underline 900"
            >
              Event Manager
            </NuxtLink>
          </div>
          <button
            class="md:hidden text-white focus:outline-none"
            @click="toggleSidebar"
          >
            <i class="fa fa-bars" />
          </button>

          <div class="hidden md:flex items-center space-x-4">
            <NuxtLink
              to="/events/create"
              class="hover:underline  dark:hover:text-blue-500"
            >
              <i class="fa fa-plus dark:text-blue-500" />
              Create Event
            </NuxtLink>
            <ThemeButton class="mr-4" />
          </div>
        </nav>
      </header>
      <aside
        :class="['fixed inset-0 rounded-sm bg-gray-800 bg-opacity-75 transition-opacity', { hidden: !isSidebarOpen }]"
        @click="toggleSidebar"
      >
        <div class="absolute rounded-md right-0 w-64 bg-white dark:bg-gray-800 shadow-lg p-6">
          <DashboardNavBar />
        </div>
      </aside>

      <main class="flex flex-1 min-h-screen">
        <aside
          class="w-1/4 text-xl  bg-gray-100 p-4 dark:bg-gray-700 shadow-md hidden md:block"
        >
          <DashboardNavBar />
        </aside>

        <section
          class="w-full  md:w-3/4 p-6 bg-white dark:bg-gray-800"
        >
          <NuxtPage />
        </section>
      </main>

      <footer class="bg-gray-800 text-white text-center p-4 dark:bg-gray-900">
        <p>&copy; 2024 My Event Site</p>
      </footer>
    </div>
  </ClientOnly>
</template>

<script setup>
import { ref } from 'vue'

const isSidebarOpen = ref(false)

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}
</script>
