<template>
  <div class="flex flex-col nin-h-screen">
    <!-- Header -->
    <header class="bg-blue-500 text-white p-4 dark:bg-gray-800 sticky top-0 z-10 shadow-md">
      <nav class="container mx-auto flex justify-between items-center flex-wrap">
        <div class="text-lg font-bold">
          <NuxtLink
            to="/"
            class="hover:underline"
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
        <ul class="space-y-4">
          <li>
            <NuxtLink
              to="/"
              class="flex items-center space-x-2 text-gray-700 dark:text-white hover:text-blue-500 dark:hover:text-blue-500"
            >
              <i class="fa fa-home text-blue-500" /> Home
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events"
              class="flex items-center space-x-2 text-gray-700 dark:text-white hover:text-blue-500 dark:hover:text-blue-500"
            >
              <i class="fa fa-calendar text-blue-500" /> My Events
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/bookmarks"
              class="flex items-center space-x-2 text-gray-700 dark:text-white hover:text-blue-500 dark:hover:text-blue-500"
            >
              <i class="fa fa-bookmark text-blue-500" /> Bookmarks
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/following"
              class="flex items-center space-x-2 text-gray-700 dark:text-white hover:text-blue-500 dark:hover:text-blue-500"
            >
              <i class="fa fa-user text-blue-500" /> Following
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/organizations"
              class="flex items-center space-x-2 text-gray-700 dark:text-white hover:text-blue-500 dark:hover:text-blue-500"
            >
              <i class="fa fa-building text-blue-500" /> Organizations
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/create"
              class="flex items-center space-x-2 text-gray-700 dark:text-white hover:text-blue-500 dark:hover:text-blue-500"
            >
              <i class="fa fa-plus text-blue-500" /> Create Event
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/organizations/create"
              class="flex items-center space-x-2 text-gray-700 dark:text-white hover:text-blue-500 dark:hover:text-blue-500"
            >
              <i class="fa fa-plus text-blue-500" /> Create Organizations
            </NuxtLink>
          </li>
          <li>
            <DashboardLogoutButton />
          </li>
          <ThemeButton class="mr-4" />
        </ul>
      </div>
    </aside>

    <main class="flex flex-1 min-h-screen">
      <!-- Left Sidebar -->
      <aside
        class="w-1/4 text-xl  bg-gray-100 p-4 dark:bg-gray-700 shadow-md hidden md:block"
      >
        <ul class="space-y-4  pl-10">
          <li class="flex items-center space-x-2 ">
            <i class="fa fa-home text-blue-500" />

            <NuxtLink
              to="/"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
            >
              Home
            </NuxtLink>
          </li>

          <li class="flex items-center space-x-2 ">
            <i class="fa fa-calendar text-blue-500" />

            <NuxtLink
              to="/events"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
            >
              My Events
            </NuxtLink>
          </li>
          <li class="flex items-center space-x-2 ">
            <i class="fa fa-bookmark text-blue-500" />
            <NuxtLink
              to="/events/bookmarks"
              class=" text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
            >
              Bookmarks
            </NuxtLink>
          </li>
          <li class="flex items-center space-x-2 ">
            <i class="fa fa-user text-blue-500" />
            <NuxtLink
              to="/events/following"
              class=" text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
            >
              Following
            </NuxtLink>
          </li>
          <li class="flex items-center space-x-2 ">
            <i class="fa fa-building text-blue-500" />
            <NuxtLink
              to="/events/organizations"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500"
            >
              Organizations
            </NuxtLink>
          </li>
          <li class="flex items-center space-x-2 ">
            <i class="fa fa-plus text-blue-500" />

            <NuxtLink
              to="/events/create"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500 "
            >
              Create Event
            </NuxtLink>
          </li>
          <li class="flex items-center space-x-2 ">
            <i class="fa fa-plus text-blue-500" />

            <NuxtLink
              to="/events/organizations/create"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500 dark:hover:text-blue-500 "
            >
              Create Organizations
            </NuxtLink>
          </li>
          <li class="flex items-center space-x-2 ">
            <i class="fa fa-sign-out text-blue-500" />

            <DashboardLogoutButton />
          </li>
        </ul>
      </aside>

      <!-- Right Content Area -->
      <section
        class="w-full  md:w-3/4 p-6 bg-white dark:bg-gray-800"
      >
        <NuxtPage />
      </section>
    </main>

    <!-- Footer -->
    <footer class="bg-gray-800 text-white text-center p-4 dark:bg-gray-900">
      <p>&copy; 2024 My Event Site</p>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isSidebarOpen = ref(false)

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}
</script>
