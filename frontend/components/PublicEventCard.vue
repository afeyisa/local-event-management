<script setup>
import { useMutation } from '@vue/apollo-composable'
import { apolloClient } from '~/plugins/apollo'
import { UN_BOOK_MARK_EVENT, BOOK_MARK_EVENT, CHECKOUT_TICKET } from '~/graphql/mutation'
import { GET_MY_ID, CHECK_AUTH_QUERY } from '~/graphql/queries'

const emit = defineEmits(['orgid', 'eventid'])
const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
})
const e = ref(props.event)
const myId = ref(null)
const bookmarks = ref([])
const isLoggedIn = ref(false)

bookmarks.value = e.value.bookmarks ? e.value.bookmarks : []
const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
if (data && data.isAuthenticated) {
  isLoggedIn.value = true
}
if (isLoggedIn.value) {
  const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
  if (!loading && !error) {
    myId.value = userData.data_users[0].user_id
  }
}
const formatDate = (date) => {
  const options = { year: 'numeric', month: 'short', day: 'numeric' }
  return new Date(date).toLocaleDateString(undefined, options)
}

const bookOrUnmarkEvent = async () => {
  if (bookmarks.value.length > 0 && myId.value) {
    try {
      const { mutate } = useMutation(UN_BOOK_MARK_EVENT)
      await mutate({ book_marker_user_id: bookmarks.value[0].book_marker_user_id, book_marked_event_id: bookmarks.value[0].book_marked_event_id })
      bookmarks.value = []
    }
    catch {
    /** */
    }
  }
  else if (myId.value) {
    try {
      const { mutate } = useMutation(BOOK_MARK_EVENT)
      const { data } = await mutate({ book_marked_event_id: e.value.event_id, book_marker_user_id: myId.value })
      bookmarks.value = data.insert_data_bookmarks.returning
    }
    catch {
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

const emitEventId = (eventid) => {
  emit('eventid', eventid)
}

const emitOrgId = (orgid) => {
  emit('orgid', orgid)
}
</script>

<template>
  <div
    class="e-card min-w-72 max-h-fit rounded overflow-hidden shadow-lg bg-white dark:bg-gray-900"
  >
    <div class="grid  sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-1 mb-2">
      <img
        :src="e.thumbnail_image_url || 'https://via.placeholder.com/400x300.png?text=Vue+Conference'"
        alt="Event Image"
        class="w-full h-52 rounded-sm  object-cover"
      >
      <div class="px-6 pt-4  pb-2 h-52">
        <div class="text-2xl font-bold text-gray-600 dark:text-gray-300 ">
          {{ e.title }}
        </div>
        <div class="text-sm text-gray-600 dark:text-gray-400">
          <i class="fa fa-tags p-2" /><strong>{{ event.category?.category_name }}</strong>
        </div>

        <div class="flex items-center text-sm text-gray-600 dark:text-gray-400">
          <i class=" p-2 fa fa-calendar" />
          {{ formatDate(e.event_date) }}
        </div>

        <div class="flex items-center text-sm text-gray-600 dark:text-gray-400 mt-2">
          <i class="p-2 fa fa-map-marker" />
          {{ e.address?.street_name || '' }} {{ e.address?.city_name || '' }} {{ e.address?.region_name || '' }}
        </div>
        <div class=" text-gray-600  dark:text-gray-400 ">
          <i class="p-2 fa fa-ticket" />
          Price
          {{ e.ticket_price===0?'free':e.ticket_price +' ETB' }}
        </div>
        <div class=" text-gray-600  dark:text-gray-400 ">
          <i class="p-2 fa fa-ticket" />
          Available Tickets
          {{ e.total_available_tickets }}
        </div>
      </div>
    </div>
    <div class="px-6 py-1" />
    <div class=" items-baseline text-gray-600 dark:text-gray-400 ">
      <div
        v-if="e.organization&&e.organization.organization_name"
        class="px- py w-full"
      >
        <button
          class="w-full  dark:text-gray-200 px-1 text-gray-700"
          @click="emitOrgId(event.organization.organization_id)"
        >
          <PublicOrganizationCard
            class="w-full"
            :organization="event.organization"
            @click="emitOrgId(event.organization.organization_id)"
          />
        </button>
      </div>
      <p
        v-else
        class="ml-2 text-sm"
      >
        {{ e.organization?e.organization.organization_name:'unknown' }}
      </p>
    </div>

    <div class="px-6 py-1 flex space-x-4">
      <div class="px-6 py-4">
        <button
          class="inline-block text-blue-500 hover:text-blue-700 dark:hover:text-blue-500"
          @click="emitEventId(event.event_id)"
        >
          Event Details
        </button>
      </div>
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
        @click="buyTickets(e.event_id)"
      >
        <i class="fa fa-ticket" /> Buy Tickets
      </button>
    </div>
  </div>
</template>
