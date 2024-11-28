<script setup>
import { formatDate } from '~/composables/formatDate'

const route = useRoute()
const eventId = route.params.id
const { bookmarks, booker, loading, error, event } = await fetcheEventDetails(eventId)
</script>

<template>
  <ClientOnly>
    <div class="min-w-md max-h-lg rounded overflow-hidden  text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900">
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
        class="max-w-4xl mx-auto p-6 bg-white rounded-lg s min-w-md max-h-lg  overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900"
      >
        <div>
          <GoBack />
        </div>
        <h1 class="text-3xl font-bold mb-4 dark:text-gray-300 ">
          {{ event.title }}
        </h1>
        <DispImages :images="event?.images?event.images:null" />
        <p class="text-gray-700 mb-6 dark:text-gray-300 ">
          <i
            class="fa fa-info-circle p-2"
          />  {{ event.description }}
        </p>
        <h3 class="text-xl text-gray-700 font-semibold mb-2  dark:text-gray-300 ">
          Event Details
        </h3>
        <ul class="list-none list-inside mb-6 text-gray-700 space-y-2 dark:text-gray-300 ">
          <li><i class="fa fa-money p-2" />Price <span class="font-semibold"><strong>{{ event.ticket_price ? `${event.ticket_price} ETB` : 'Free' }}</strong></span></li>
          <li><i class="fa fa-ticket p-2" />Total Tickets <strong>{{ event.total_available_tickets }}</strong></li>
          <li>
            <i class=" p-2 fa fa-calendar" />
            Date <strong>{{ formatDate(event.event_date) }}</strong>
          </li>
          <li><i class="fa fa-tags p-2" />Category  <strong>{{ event.category?.category_name }}</strong></li>
          <li>  <i class="fa fa-map-marker p-2" />Venue <strong>{{ event.venue }}</strong></li>
          <li>
            <i class="fa fa-map-marker p-2" />
            Location <strong>{{ event.address?.street_name }}, {{ event.address?.city_name }} {{ event.address?.region_name }}</strong>
            <br>
          </li>
        </ul>
        <div class="flex justify-center w-full">
          <MapPicker :events="[event]" />
        </div>
        <div class="px-6 py-4 flex space-x-4">
          <button
            class="text-yellow-500 hover:text-yellow-700"
            @click="booker"
          >
            <i
              :class="bookmarks.length>0?'fa fa-bookmark':'fa fa-bookmark-o'"
            /> {{ bookmarks.length>0?`Book Marked`:`Book Mark` }}
          </button>
          <button
            class="text-green-500 hover:text-green-700"
          >
            <i class="fa fa-ticket" />
            <NuxtLink
              :to="`/${event.event_id}/ticket`"
            >
              Buy Tickets
            </NuxtLink>
          </button>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>
