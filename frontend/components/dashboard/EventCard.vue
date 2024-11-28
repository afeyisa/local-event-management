<script setup>
import { useMutation } from '@vue/apollo-composable'
import { DELETE_EVENT } from '~/graphql/mutations/deleteEvent.graphql'
import { formatDate } from '~/composables/formatDate'
import { fetchBase64Image } from '~/composables/fetchImage'

const prorps = defineProps({
  event: {
    type: Object,
    required: true,
  },
})
const event = ref(prorps.event)
const base64Image = ref(null)
const prefetchImages = async () => {
  base64Image.value = await fetchBase64Image(event.value.thumbnail_image_url)
}

onMounted(() => {
  prefetchImages()
})
const deleteEvent = async () => {
  const confirmed = confirm('Are you sure you want to delete this organization? This action cannot be undone.')

  if (confirmed) {
    try {
      const { mutate } = useMutation(DELETE_EVENT)
      const { data } = await mutate({ event_id: event.value.event_id })

      if (data) {
        event.value = null
        alert('event deleted successfully.')
      }
    }
    catch {
      alert('An error occurred while deleting the event.')
    }
  }
  else {
    alert('Deletion canceled.')
  }
}
</script>

<template>
  <div
    v-if="event"
    class="event-card min-w-md max-h-lg rounded overflow-hidden shadow-lg bg-white dark:bg-gray-900"
  >
    <img
      :src="base64Image ||'https://via.placeholder.com/400x300.png?text=Vue+Conference'"
      alt="Event Image"
      class="w-full h-48 object-cover"
    >

    <div class="px-6 py-4">
      <div class="font-bold text-xl mb-2 text-gray-700 dark:text-gray-300">
        {{ event.title }}
      </div>
    </div>

    <div class="px-6 pt-4 pb-2">
      <div class="flex items-center text-sm text-gray-600 dark:text-gray-400">
        <i class=" p-2 fa fa-calendar" />
        {{ formatDate(event.event_date) }}
      </div>
      <div class="flex items-center text-sm text-gray-600 dark:text-gray-400 mt-2">
        <i class="p-2 fa fa-map-marker" />
        {{ event.venue }}
      </div>
    </div>

    <div class="flex flex-wrap  ">
      <div class="px-2 py-2 w-full sm:w-auto ">
        <NuxtLink
          :to="`/events/${event.event_id}`"
          class="inline-block bg-blue-500 text-white px-3 py-2 rounded hover:bg-blue-600 dark:bg-blue-700 dark:hover:bg-blue-800 w-full sm:w-auto text-center"
        >
          Details
        </NuxtLink>
      </div>
      <div class="px-2 py-2 w-full sm:w-auto">
        <NuxtLink
          :to="`/events/${event.event_id}/edit`"
          class="inline-block bg-yellow-500 text-white px-3 py-2 w-full sm:w-auto text-center rounded hover:bg-yellow-600 dark:bg-yellow-700 dark:hover:bg-yellow-800"
        >
          Update
        </NuxtLink>
      </div>
      <div class="px-2 py-2 w-full sm:w-auto">
        <button
          class="inline-block text-red-500  w-full sm:w-auto text-center  border border-red-500 px-3 py-2 rounded"
          @click="deleteEvent"
        >
          Delete
        </button>
      </div>
    </div>
  </div>
</template>

  <style scoped>
  .event-card {
    transition: transform 0.2s;
  }

  .event-card:hover {
    transform: translateY(-5px);
  }
  </style>
