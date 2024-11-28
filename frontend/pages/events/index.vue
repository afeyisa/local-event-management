<script setup>
import EventCard from '~/components/dashboard/EventCard.vue'
import { GET_MY_EVENTS } from '~/graphql/querie/getUserEvents.graphql'
import { apolloClient } from '~/plugins/apollo'
import { fetchUseId } from '~/composables/user'

const { myId } = await fetchUseId()
const events = ref([])
const loading = ref(null)
const error = ref(null)
if (myId.value) {
  const { data, loading: l, error: e } = await apolloClient.query({ query: GET_MY_EVENTS, variables: { where: { organization: { organizes: { user: { user_id: { _eq: myId.value } } } } } } })
  events.value = data.data_events
  loading.value = l
  error.value = e
}

definePageMeta({
  middleware: 'auth',
  layout: 'mydashboard',
})
</script>

<template>
  <div>
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
      Error loading events
    </div>
    <div class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 gap-6">
      <EventCard
        v-for="event in events"
        :key="event.event_id"
        :event="event"
      />
    </div>
  </div>
</template>
