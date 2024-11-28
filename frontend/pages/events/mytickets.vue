<script setup>
import { ref } from 'vue'
import { GET_USER_TICKETS } from '~/graphql/querie/getUserTickets.graphql'
import { GET_MY_ID } from '~/graphql/querie/getUserId.graphql'
import { apolloClient } from '~/plugins/apollo'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const tickets = ref([])
const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
if (!loading && !error) {
  const myId = userData.data_users[0].user_id
  const { data, loading, error } = await apolloClient.query({ query: GET_USER_TICKETS, variables: { id: myId } })
  if (!loading && !error) {
    tickets.value = data.data_tickets
  }
}
</script>

<template>
  <div>
    <p v-if="error">
      Something went wrong...
    </p>
    <p v-if="loading">
      Loading...
    </p>
    <div v-else>
      <GoBack />
      <div
        class="grid sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-3 gap-6"
      >
        <div
          v-for="ticket in tickets"
          :key="ticket.ticket_id"
        >
          <PurchasedTicketCard :ticket="ticket" />
        </div>
      </div>
    </div>
  </div>
</template>
