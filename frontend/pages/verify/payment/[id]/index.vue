<script setup>
import { verifyPayment } from '~/composables/veifyPayment'

definePageMeta({
  layout: 'auth',
  middleware: 'auth',
})

const route = useRoute()
const tx_ref = route.params.id
const { error, loading, status } = await verifyPayment(tx_ref)
</script>

<template>
  <ClientOnly>
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
        Error to show payment status
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
        <h1 class="text-3xl font-bold mb-4 dark:text-gray-300 ">
          {{ status }}
        </h1>
      </div>
    </div>
  </ClientOnly>
</template>
