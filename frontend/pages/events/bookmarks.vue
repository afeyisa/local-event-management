<template>
  <div>
    <template v-if="!isOrgDetail && !isEventDetail">
      <GoBack />
      <h2 class="text-2xl font-bold mb-4 text-gray-700 dark:text-white">
        Book Marked Events
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 gap-6">
        <PublicEventCard
          v-for="event in events"
          :key="event.id"
          :event="event"
          @eventid="handleEventId"
          @orgid="handleOrgid"
        />
      </div>
    </template>
    <div
      v-else-if="isEventDetail"
      class=" w-full   bg-gray-100 p-4 dark:bg-gray-700   overflow-y-auto flex-shrink-0 text-xl min-h-screen"
    >
      <EventDetail
        :id="eventid"
        @goback="setIsevent"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { BROWSEEVENTS, GET_MY_ID } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})

const events = ref([])
const isEventDetail = ref(false)
const isOrgDetail = ref(false)
const orgid = ref(null)
const eventid = ref(null)

const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })

if (!loading && !error) {
  const myId = userData.data_users[0].user_id
  const { data } = await apolloClient.query({ query: BROWSEEVENTS, variables: { where: { bookmarks: { book_marker_user_id: { _eq: myId } } } } })
  events.value = data.data_events
}
const handleEventId = (id) => {
  eventid.value = id
  isEventDetail.value = true
  isOrgDetail.value = false

  console.log(id)
}
const handleOrgid = (id) => {
  orgid.value = id
  isOrgDetail.value = true
  isEventDetail.value = false

  console.log(id)
}
const setIsevent = () => {
  isEventDetail.value = false
}
</script>
