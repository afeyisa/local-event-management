<!-- Home page with event list -->
<template>
  <div class="flex flex-col ">
    <!-- Content area: either default content or search results -->
    <div class="flex-grow flex overflow-hidden">
      <!-- Default content displayed if no search results yet -->
      <div
        v-if="!isSearching"
        class="w-full p-4 overflow-y-auto"
      >
        <!-- Default home content (e.g., featured events) -->
        <h2 class="text-2xl font-bold mb-4">
          Featured Events
        </h2>
        <ul class="space-y-4">
          <li
            v-for="event in defaultEvents"
            :key="event.id"
            class="p-4 bg-gray-100 rounded-md shadow-md"
          >
            <h3 class="text-lg font-semibold">
              {{ event.title }}
            </h3>
            <p>{{ event.description }}</p>
          </li>
        </ul>
      </div>

      <!-- Search result layout when searching -->
      <div
        v-if="isSearching"
        class="flex w-full"
      >
        <!-- Filters on the left (scrollable) -->
        <div
          class="flex justify-between  resize-handle  w-72 bg-gray-100 p-4 overflow-y-auto flex-shrink-0"
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
              <li
                v-for="item in searchResults"
                :key="item.id"
              >
                {{ item.title }}
              </li>
            <!-- Add more filter options as needed -->
            </ul>
          </div>
          <button class=" w-1 cursor-ew-resize bg-slate-200" />
        </div>

        <!-- Search results on the right (scrollable) -->
        <div class=" flex-grow p-4 overflow-y-auto">
          <h2 class="text-2xl font-bold mb-4">
            Search Results
          </h2>
          <ul class="space-y-4">
            <li
              v-for="result in searchResults"
              :key="result.id"
              class="p-4 bg-white rounded-md shadow-md"
            >
              <h3 class="text-lg font-semibold">
                {{ result.title }}
              </h3>
              <p>{{ result.description }}</p>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
// Dummy default events to display on the home page (before search)
const defaultEvents = [
  { id: 2, title: 'Rock Concert 2024', description: 'The best rock bands performing live this summer.' },
]

// Dummy search results to display after search
const searchResults = ref([
  { id: 1, title: 'Jazz Night in New York', description: 'An evening of live jazz in a cozy venue.' },
])
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
// Reactive state for managing search behavior
const isSearching = ref(true)

// Simulate a search action
function onSearch() {
  isSearching.value = true // Display search results
}
onSearch()
</script>
