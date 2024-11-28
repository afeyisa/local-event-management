<script setup>
import { formatDate } from '#imports'
import { getpaymentAmount } from '~/composables/getPaymentAmout'

const { ticket } = defineProps({
  ticket: {
    type: Object,
    required: true,
  },
})

const base64Image = ref(null)
const prefetchImages = async () => {
  base64Image.value = await fetchBase64Image(ticket.qr_link)
}

const payment = ref(null)
const prefetchPayment = async () => {
  if (ticket.tx_ref) {
    const result = await getpaymentAmount(ticket.tx_ref)
    const { paymentAmount, error, loading } = result
    payment.value = paymentAmount.value

    if (!error && !loading) {
      console.log(payment.value)
    }
  }
}

onMounted(() => {
  prefetchImages()
  prefetchPayment()
})
</script>

<template>
  <div class="purchased-card bg-gray-100 dark:bg-gray-900 shadow-md rounded-lg p-4 flex flex-col space-y-4">
    <img
      :src="base64Image"
      alt="ticket qr"
      class="w-52 h-auto rounded-sm  object-cover"
    >
    <h2 class="text-lg font-bold text-gray-900 dark:text-gray-100">
      {{ ticket.event.title }}
    </h2>

    <div class="text-sm text-gray-700 dark:text-gray-300">
      <p>
        <span class="font-semibold">Purchase Date</span> {{ formatDate(ticket.purchased_date) }}
      </p>
      <p>
        <span class="font-semibold">Event Date</span> {{ formatDate(ticket.event.event_date) }}
      </p>
      <p>
        <span class="font-semibold">Price</span>
        {{ payment?.amount > 0 ? `${payment.amount}${payment.currency}` : 'Free' }}
      </p>
    </div>

    <div class="flex justify-end">
      <NuxtLink
        class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md"
        :to="`/events/${ticket.event_id}`"
      >
        View Event Details
      </NuxtLink>
    </div>
  </div>
</template>
