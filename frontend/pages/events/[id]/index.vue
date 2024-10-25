<!-- event detail page -->

<template>
  <div class="event-card min-w-md max-h-lg rounded overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300 bg-white dark:bg-gray-900">
    <!-- Event Image -->>
    <div
      v-if="loading"
      class="text-center py-8"
    >
      Loading...
    </div>
    <div
      v-if="error"
      class="text-red-500 text-center py-8"
    >
      Error loading event details
    </div>
    <div
      v-if="event"
      class="max-w-4xl mx-auto p-6 bg-white rounded-lg shadow-lg"
    >
      <h1 class="text-3xl font-bold mb-4">
        {{ event.title }}
      </h1>

      <div class="relative">
        <!-- Left Arrow -->
        <button
          v-if="currentImageIndex > 0"
          class="absolute left-0 top-1/2 transform -translate-y-1/2 bg-gray-800 text-white p-2 rounded-full"
          @click="previousImage"
        >
          &#8592;
        </button>

        <!-- Main Image Display -->
        <img
          :src="event.images[currentImageIndex].image_url"
          alt="Event Image"
          class="w-full h-64 object-cover rounded-md mb-6"
        >

        <!-- Right Arrow -->
        <button
          v-if="currentImageIndex < event.images.length - 1"
          class="absolute right-0 top-1/2 transform -translate-y-1/2 bg-gray-800 text-white p-2 rounded-full"
          @click="nextImage"
        >
          &#8594;
        </button>
      </div>
      <p class="text-gray-700 mb-6">
        {{ event.description }}
      </p>
      <h3 class="text-xl font-semibold mb-2">
        Event Details:
      </h3>
      <ul class="list-disc list-inside mb-6 text-gray-700 space-y-2">
        <li>Price: <span class="font-semibold">{{ event.ticket_price ? `$${event.ticket_price}` : 'Free' }}</span></li>
        <li>Total Tickets: {{ event.total_available_tickets }}</li>
        <li>Date: {{ event.event_date }}</li>
        <li>Category: {{ event.category?.category_name }}</li>
        <li>Venue: {{ event.venue }}</li>
        <li
          v-if="event.is_online"
          class="text-blue-500 font-semibold"
        >
          This event is online
        </li>
        <li v-else>
          Location: {{ event.address?.street_name }}, {{ event.address?.city_name }} {{ event.address?.region_name }}
          <br>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import {GET_EVENT_DETAILS} from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const route = useRoute()
const eventId = route.params.id
// console.log('hello world')


const { data, loading, error } = await apolloClient.query({ query: GET_EVENT_DETAILS, variables: { id: eventId } })
const event = ref(null)
const currentImageIndex = ref(0)
onMounted(() => {
  // console.log(data)
  event.value = data.data_events_by_pk
  const updatedImages = [{ image_url: event.value.thumbnail_image_url }, ...(event.value.images || [])]
  event.value = { ...event.value, images: updatedImages }
})
const previousImage = () => {
  if (currentImageIndex.value > 0) {
    currentImageIndex.value--
  }
}

// Navigate Next Image
const nextImage = () => {
  if (currentImageIndex.value < event.value.images.length - 1) {
    currentImageIndex.value++
  }
}
</script>
