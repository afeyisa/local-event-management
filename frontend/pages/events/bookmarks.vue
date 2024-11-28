<script setup>
import { ref } from 'vue'
import { BROWSEEVENTS } from '~/graphql/querie/browseEvents.graphql'
import { GET_MY_ID } from '~/graphql/querie/getUserId.graphql'

import { apolloClient } from '~/plugins/apollo'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})

const events = ref([])
const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
if (!loading && !error) {
  const myId = userData.data_users[0].user_id
  const { data } = await apolloClient.query({ query: BROWSEEVENTS, variables: { where: { bookmarks: { book_marker_user_id: { _eq: myId } } } } })
  events.value = data.data_events
}
</script>

<template>
  <div>
    <GoBack />
    <h2 class="text-2xl font-bold mb-4 text-gray-700 dark:text-white">
      Book Marked Events
    </h2>
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
      Error loading bookmarked events
    </div>
    <div class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 gap-6">
      <PublicEventCard
        v-for="event in events"
        :key="event.id"
        :event="event"
      />
    </div>
  </div>
</template>
