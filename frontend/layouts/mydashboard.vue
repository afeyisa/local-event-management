<template>
  <div class="flex flex-col min-h-screen">
    <!-- Header -->
    <header class="bg-blue-500 text-white p-4 dark:bg-gray-800 sticky top-0 z-10 shadow-md">
      <nav class="container mx-auto flex justify-between items-center">
        <div class="text-lg font-bold">
          <NuxtLink
            to="/"
            class="hover:underline"
          >
            Event Manager
          </NuxtLink>
        </div>
        <div class="flex items-center space-x-4">
          <ThemeButton class="mr-4" />
          <NuxtLink
            to="/events/create"
            class="hover:underline"
          >
            Create Event
          </NuxtLink>
        </div>
      </nav>
    </header>

    <!-- Main Content (Dashboard Layout) -->
    <main class="flex flex-1">
      <!-- Left Sidebar -->
      <aside class="w-1/4 text-xl min-h-screen  bg-gray-100 p-4 dark:bg-gray-700 shadow-md">
        <ul class="space-y-4 pl-10">
          <li>
            <NuxtLink
              to="/"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500"
            >
              Home
            </NuxtLink>
          </li>

          <li>
            <NuxtLink
              to="/events"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500"
            >
              My Events
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/bookmarks"
              class="text-gray-700 dark:text-white font-semibold hover:text-blue-500"
            >
              Bookmarks
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/organizations"
              class="ext-gray-700 dark:text-white font-semibold hover:text-blue-500"
            >
              Organizations
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/create"
              class="ext-gray-700 dark:text-white font-semibold hover:text-blue-500 "
            >
              Create Event
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/events/organizations/create"
              class="ext-gray-700 dark:text-white font-semibold hover:text-blue-500 "
            >
              Create Organizations
            </NuxtLink>
          </li>
          <li>
            <NuxtLink
              to="/settings"
              class="text-gray-700  dark:text-white font-semibold hover:text-blue-500"
            >
              Settings
            </NuxtLink>
          </li>
          <li>
            <button class="w-full text-left px-4 py-2  dark:text-white dark:hover:text-red-600 hover:text-red-600 hover:underline ">
              Logout
            </button>
          </li>
        </ul>
      </aside>

      <!-- Right Content Area -->
      <section class="w-3/4 p-6 bg-white dark:bg-gray-800">
        <NuxtPage />
      </section>
    </main>

    <!-- Footer -->
    <footer class="bg-gray-800 text-white text-center p-4 dark:bg-gray-900">
      <p>&copy; 2024 My Event Site</p>
    </footer>
  </div>
</template>
