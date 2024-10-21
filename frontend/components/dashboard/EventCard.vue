<template>
  <div class="event-card min-w-md max-h-lg rounded overflow-hidden shadow-lg bg-white dark:bg-gray-900">
    <!-- Event Image -->
    <img
      :src="event.image"
      alt="Event Image"
      class="w-full h-48 object-cover"
    >

    <!-- Event Details -->
    <div class="px-6 py-4">
      <div class="font-bold text-xl mb-2 dark:text-gray-300">
        {{ event.title }}
      </div>
      <p class="text-gray-700 text-base dark:text-gray-300">
        {{ event.description }}
      </p>
    </div>

    <!-- Event Metadata -->
    <div class="px-6 pt-4 pb-2">
      <div class="flex items-center text-sm text-gray-600 dark:text-gray-400">
        <svg
          class="w-4 h-4 mr-2"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M8 7V3m8 4V3m-9 8h10m-9 4h10m-9 4h10"
          />
        </svg>
        {{ formatDate(event.date) }}
      </div>
      <div class="flex items-center text-sm text-gray-600 dark:text-gray-400 mt-2">
        <svg
          class="w-4 h-4 mr-2"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M17.657 16.657L13 12m-4.243 4.243A6 6 0 1121 9a6.007 6.007 0 01-9.757 7.657z"
          />
        </svg>
        {{ event.location }}
      </div>
    </div>

    <!-- Event Actions -->
    <div class="flex">
      <div class="px-6 py-4">
        <NuxtLink
          :to="`/events/${event.id}`"
          class="inline-block bg-blue-500 text-white px-3 py-2 rounded hover:bg-blue-600 dark:bg-blue-700 dark:hover:bg-blue-800"
        >
          Event Details
        </NuxtLink>
      </div>
      <div class="px-6 py-4">
        <NuxtLink
          :to="`/events/${event.id}`"
          class="inline-block bg-blue-500 text-white px-3 py-2 rounded hover:bg-blue-600 dark:bg-blue-700 dark:hover:bg-blue-800"
        >
          Upadte Event
        </NuxtLink>
      </div>
      <div class="px-6 py-4">
        <NuxtLink
          :to="`/events/${event.id}/edit`"
          class="inline-block text-red-500  px-3 py-2 rounded hover:bg-blue-500"
        >
          Delete Event
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
// import { computed } from 'vue'

defineProps({
  event: {
    type: Object,
    required: true,
  },
})

// Function to format the date nicely
const formatDate = (date) => {
  const options = { year: 'numeric', month: 'short', day: 'numeric' }
  return new Date(date).toLocaleDateString(undefined, options)
}
</script>

  <style scoped>
  .event-card {
    transition: transform 0.2s;
  }

  .event-card:hover {
    transform: translateY(-5px);
  }
  </style>
