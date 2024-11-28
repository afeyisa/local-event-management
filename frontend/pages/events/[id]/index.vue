<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import DispImages from '~/components/DispImages.vue'
import { GET_EVENT_DETAILS } from '~/graphql/querie/getEventDetails.graphql'
import { apolloClient } from '~/plugins/apollo'
import { formatDate } from '~/composables/formatDate'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const route = useRoute()
const eventId = route.params.id

const { data, loading, error } = await apolloClient.query({ query: GET_EVENT_DETAILS, variables: { id: eventId } })
const event = ref(null)
onMounted(() => {
  event.value = data.data_events_by_pk
  const updatedImages = [{ image_url: event.value.thumbnail_image_url }, ...(event.value.images || [])]
  event.value = { ...event.value, images: updatedImages }
})
</script>

<template>
  <div class="min-w-md max-h-lg rounded overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300 bg-white dark:bg-gray-900">
    <div
      v-if="loading"
      class="text-center py-8"
    >
      Loading...
    </div>
    <div
      v-if="error"
      class="text-red-500 text-center py-8 dark:text-gray-300 "
    >
      Error loading event details
    </div>
    <div
      v-if="event"
      class=" mx-auto p-6 bg-white rounded-lg s min-w-md max-h-lg  overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900"
    >
      <GoBack />
      <h1 class="text-3xl font-bold mb-4 dark:text-gray-300 ">
        {{ event.title }}
      </h1>
      <DispImages :images="event.images" />
      <p class="text-gray-700 mb-6 dark:text-gray-300 ">
        {{ event.description }}
      </p>
      <h3 class="text-xl text-gray-700 font-semibold mb-2  dark:text-gray-300 ">
        Event Details
      </h3>
      <ul class="list-none list-inside mb-6 text-gray-700 space-y-2 dark:text-gray-300 ">
        <li><i class="fa fa-money p-2" />Price <span class="font-semibold">{{ event.ticket_price ? `$${event.ticket_price}` : 'Free' }}</span></li>
        <li><i class="fa fa-ticket p-2" />Total Tickets {{ event.total_available_tickets }}</li>
        <li><i class=" p-2 fa fa-calendar" />Date {{ formatDate(event.event_date) }}</li>
        <li><i class="fa fa-tags p-2" />Category {{ event.category?.category_name }}</li>
        <li><i class="fa fa-map-marker p-2" />Venue {{ event.venue }}</li>
        <li>
          <i class="fa fa-map-marker p-2" />
          Location {{ event.address?.street_name }}, {{ event.address?.city_name }} {{ event.address?.region_name }}
          <br>
        </li>
      </ul>
      <div
        v-if="!event.is_online"
        class="flex justify-center w-full"
      >
        <MapPicker :events="[event]" />
      </div>
    </div>
  </div>
</template>
