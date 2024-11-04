<template>
  <div>
    <div class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 gap-6">
      <EventCard
        v-for="event in events"
        :key="event.event_id"
        :event="event"
      />
    </div>
  </div>
</template>

<script setup>
import EventCard from '~/components/dashboard/EventCard.vue'
import { GET_MY_EVENTS, GET_MY_ID } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
const events = ref([])
if (!loading && !error) {
  const myId = userData.data_users[0].user_id
  const { data } = await apolloClient.query({ query: GET_MY_EVENTS, variables: { where: { organization: { organizes: { user: { user_id: { _eq: myId } } } } } } })
  events.value = data.data_events
}

definePageMeta({
  middleware: 'auth',
  layout: 'mydashboard',
})
</script>
