<script setup>
import { getTicketInfo } from '~/composables/verifyTicket'

definePageMeta({
  layout: 'auth',
  middleware: 'auth',
})

const route = useRoute()
const id = route.params.id
const { error, loading, ticket } = await getTicketInfo(id)
const base64Image = ref(null)
const prefetchImages = async () => {
  base64Image.value = await fetchBase64Image(ticket.event?.thumbnail_image_url)
}
onMounted(() => {
  prefetchImages()
})
</script>

<template>
  <ClientOnly>
    <div
      v-if="ticket"
      class="min-w-md max-h-lg rounded overflow-hidden  text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900"
    >
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
        Error to getting ticket information
      </div>
      <div
        class="max-w-4xl mx-auto p-6 bg-white rounded-lg s min-w-md max-h-lg  overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900"
      >
        <div>
          <NuxtLink
            :to="'/'"
            class="text-yellow-500 hover:text-yellow-700"
          >
            <i class="fa fa-chevron-left" /> Go Back

          </NuxtLink>
        </div>
        <img
          v-if="base64Image"
          :src="base64Image"
          alt="Event Image"
          class="w-full h-64 object-cover rounded-md mb-6"
        >
        <div class="purchased-card bg-gray-100 dark:bg-gray-900 shadow-md rounded-lg p-4 flex flex-col space-y-4">
          <h2 class="text-lg font-bold text-gray-900 dark:text-gray-100">
            {{ ticket.event?.title }}
          </h2>

          <div class="text-sm text-gray-700 dark:text-gray-300">
            <p>
              <span class="font-semibold">Event Venue</span> {{ ticket.event?.venue }}
            </p>
            <p>
              <span class="font-semibold">Purchase Date</span> {{ formatDate(ticket.purchased_date) }}
            </p>
            <p>
              <span class="font-semibold">Event Date</span> {{ formatDate(ticket.event.event_date) }}
            </p>
            <p>
              <span class="font-semibold">Price</span>
              {{ ticket.payment?.amount > 0 ? `${ticket.payment.amount}${ticket.payment.currency}` : 'Free' }}
            </p>
          </div>

          <div class="flex justify-end">
            <NuxtLink
              class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md"
              :to="`/${ticket.event_id}`"
            >
              View Event Details
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>
