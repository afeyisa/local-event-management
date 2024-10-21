<!-- edit events -->
<template>
  <div>
    hello edit events
    <div class="flex">
      <div class="px-6 py-4">
        <NuxtLink
          :to="`/events`"
          class="inline-block bg-blue-500 text-white px-3 py-2 rounded hover:bg-blue-600 dark:bg-blue-700 dark:hover:bg-blue-800"
        >
          Event Details
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script>
definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
</script>
