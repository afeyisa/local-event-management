<template>
  <div>
    <header class="bg-blue-500  sticky top-0 z-10 shadow-md text-white p-4 dark:bg-gray-800">
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
                to="/events"
                class="hover:underline"
              >
                My Events
              </NuxtLink>
            </li>
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
    <div class="flex flex-col">
      <!-- Content area: either default content or search results -->
      <div class="flex-grow flex overflow-hidden">
        <!-- Default content displayed if no search results yet -->
        <div
          v-if="!isSearching"
          class="w-full p-4 overflow-y-auto"
        >
          <h2 class="text-2xl font-bold mb-4">
            Featured Events
          </h2>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-6">
            <EventCard
              v-for="event in events"
              :key="event.id"
              :event="event"
            />
          </div>
        </div>

        <!-- Search result layout when searching -->
        <div
          v-if="isSearching"
          class="flex w-full"
        >
          <!-- Filters on the left (scrollable) -->
          <div
            class="flex justify-between resize-handle w-72 bg-gray-100 p-4 overflow-y-auto flex-shrink-0"
            @mousedown="initResize"
          >
            <div>
              <h2 class="text-xl font-bold mb-4">
                Filters
              </h2>
              <ul>
                <li class="mb-2">
                  Category: Music
                </li>
                <li class="mb-2">
                  Location: New York
                </li>
                <li class="mb-2">
                  Date: Upcoming
                </li>
                <!-- Add more filter options as needed -->
              </ul>
            </div>
            <button class="w-1 cursor-ew-resize bg-slate-200" />
          </div>

          <!-- Search results on the right (scrollable) -->
          <div class="flex-grow p-4 overflow-y-auto">
            <h2 class="text-2xl font-bold mb-4">
              Search Results
            </h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-1 gap-6">
              <EventCard
                v-for="event in events"
                :key="event.id"
                :event="event"
              />
            </div>
          </div>
        </div>

        <!-- Default content in grid layout if not searching -->
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import EventCard from '~/components/dashboard/EventCard.vue'

// Example event data
const events = [
  {
    id: 1,
    title: 'Vue.js Conference 2024',
    description: 'A global Vue.js conference bringing together developers from around the world.',
    date: '2024-11-15',
    location: 'San Francisco, CA',
    image: 'https://via.placeholder.com/400x300.png?text=Vue+Conference',
  },
  {
    id: 2,
    title: 'Nuxt Meetup 2024',
    description: 'Join the Nuxt community for a day of talks, workshops, and networking.',
    date: '2024-12-05',
    location: 'New York, NY',
    image: 'https://via.placeholder.com/400x300.png?text=Nuxt+Meetup',
  },
  {
    id: 3,
    title: 'TailwindCSS Workshop',
    description: 'A hands-on workshop to master TailwindCSS and build beautiful UIs.',
    date: '2024-10-25',
    location: 'Online',
    image: 'https://via.placeholder.com/400x300.png?text=Tailwind+Workshop',
  },
]

const defaultEvents = [
  { id: 2, title: 'Rock Concert 2024', description: 'The best rock bands performing live this summer.' },
]

// Dummy search results to display after search
const searchResults = ref([
  { id: 1, title: 'Jazz Night in New York', description: 'An evening of live jazz in a cozy venue.' },
])

// Resizing behavior for the filters panel
const initResize = (event) => {
  const div = event.target.parentElement
  const startX = event.clientX
  const startWidth = div.offsetWidth

  const doDrag = (e) => {
    div.style.width = `${startWidth + e.clientX - startX}px`
  }

  const stopDrag = () => {
    document.removeEventListener('mousemove', doDrag)
    document.removeEventListener('mouseup', stopDrag)
  }

  document.addEventListener('mousemove', doDrag)
  document.addEventListener('mouseup', stopDrag)
}

// Search state
const isSearching = ref(false)

// Simulate a search action
function onSearch() {
  isSearching.value = true // Display search results after search
}

// Simulate initial search call
// onSearch()
</script>
