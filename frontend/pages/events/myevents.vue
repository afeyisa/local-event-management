<template>
  <div>
    <div class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-6">
      <!-- Loop through the events array and pass each event as a prop to EventCard -->
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
import { GET_EVENT_SAMPLE } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

const { data } = await apolloClient.query({ query: GET_EVENT_SAMPLE })
// console.log(data)
const events = data.data_events
definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
</script>
