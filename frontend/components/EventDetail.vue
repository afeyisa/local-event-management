<script setup>
import { ref } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import { GET_MY_ID, GET_EVENT_DETAILS, CHECK_AUTH_QUERY, PUBLIC_GET_EVENT_DETAILS } from '~/graphql/queries'
import { UN_BOOK_MARK_EVENT, BOOK_MARK_EVENT } from '~/graphql/mutation'

import { apolloClient } from '~/plugins/apollo'

const emit = defineEmits(['goback'])
const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})
const eventId = props.id
const bookmarks = ref([])
const myId = ref(null)
const isLoggedIn = ref(false)
const currentImageIndex = ref(0)
const event = ref(null)

const { data: ch } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
if (ch && ch.isAuthenticated) {
  isLoggedIn.value = true
}

if (isLoggedIn.value) {
  const { data: userData } = await apolloClient.query({ query: GET_MY_ID })
  myId.value = userData.data_users[0].user_id
}
const loading = ref(null)
const error = ref(null)
if (isLoggedIn.value) {
  const { data, loading: ld, error: er } = await apolloClient.query({ query: GET_EVENT_DETAILS, variables: { id: eventId } })
  loading.value = ld
  error.value = er
  event.value = data.data_events_by_pk
  const updatedImages = [{ image_url: event.value.thumbnail_image_url }, ...(event.value.images || [])]
  event.value = { ...event.value, images: updatedImages }
  bookmarks.value = event.value.bookmarks
}
else {
  const { data, loading: ld, error: er } = await apolloClient.query({ query: PUBLIC_GET_EVENT_DETAILS, variables: { id: eventId } })
  loading.value = ld
  error.value = er
  event.value = data.data_events_by_pk
  const updatedImages = [{ image_url: event.value.thumbnail_image_url }, ...(event.value.images || [])]
  event.value = { ...event.value, images: updatedImages }
}
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

const bookOrUnmarkEvent = async () => {
  // un booking the event
  if (bookmarks.value.length > 0 && myId.value) {
    try {
      const { mutate } = useMutation(UN_BOOK_MARK_EVENT)
      await mutate({ book_marker_user_id: bookmarks.value[0].book_marker_user_id, book_marked_event_id: bookmarks.value[0].book_marked_event_id })
      bookmarks.value = []
    }
    catch (e) {
    /** */
      console.log(e)
    }
  }
  // book marking the event
  else if (myId.value) {
    console.log(event.value.event_id)
    try {
      const { mutate } = useMutation(BOOK_MARK_EVENT)
      const { data } = await mutate({ book_marked_event_id: event.value.event_id, book_marker_user_id: myId.value })
      bookmarks.value = data.insert_data_bookmarks.returning
    }
    catch (e) {
      console.log(e)
      /** */
    }
  }
}

const buyTickets = async (id) => {
  if (myId.value && id) {
    try {
      const { mutate } = useMutation(CHECKOUT_TICKET)
      const { data } = await mutate({ event_id: id, user_id: myId.value })
      console.log(data)
      if (data) {
        window.location.href = data.ticketcheckout.url
      }
    }
    catch (err) {
      console.log(err)
    /** */
    }
  }
}
const emitGoback = () => {
  emit('goback')
}
</script>

<template>
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
        <button
          class="text-yellow-500 hover:text-yellow-700"
          @click="emitGoback"
        >
          <i class="fa fa-chevron-left" /> Go Back
        </button>
      </div>
      <h1 class="text-3xl font-bold mb-4 dark:text-gray-300 ">
        {{ event.title }}
      </h1>

      <div
        class="relative"
      >
        <button
          v-if="currentImageIndex > 0 && Array.isArray(event.images) && event.images.length > 0"
          class="absolute left-0 top-1/2 transform -translate-y-1/2"
          @click="previousImage"
        >
          <i :class="(currentImageIndex > 0 && Array.isArray(event.images) && event.images.length > 0)?'fa fa-chevron-left dark:text-gray-500 text-gray-800 p-2 text-xl':''" />
        </button>

        <img
          v-if="Array.isArray(event.images) && event.images.length > 0 && event.images[currentImageIndex]?.image_url"
          :src="event.images[currentImageIndex].image_url"
          alt="Event Image"
          class="w-full h-64 object-cover rounded-md mb-6"
        >

        <button
          v-if="currentImageIndex < event.images.length - 1 && Array.isArray(event.images) && event.images.length > 0"
          class="absolute right-0 top-1/2 transform -translate-y-1/2"
          @click="nextImage"
        >
          <i :class="(currentImageIndex < event.images.length - 1 && Array.isArray(event.images) && event.images.length > 0) ? 'fa fa-chevron-right dark:text-gray-500 text-gray-800 p-2 text-xl' : ''" />
        </button>
      </div>
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
          Date <strong>{{ event.event_date }}</strong>
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
          @click="bookOrUnmarkEvent"
        >
          <i
            :class="bookmarks.length>0?'fa fa-bookmark':'fa fa-bookmark-o'"
          /> {{ bookmarks.length>0?`Book Marked`:`Book Mark` }}
        </button>
        <button
          class="text-green-500 hover:text-green-700"
          @click="buyTickets(event.event_id)"
        >
          <i class="fa fa-ticket" /> Buy Tickets
        </button>
      </div>
    </div>
  </div>
</template>
