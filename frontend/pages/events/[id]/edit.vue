<template>
  <div>
    <GoBack />
    <EventForm
      :id="eventId"
    />
  </div>
</template>

<script setup>
import EventForm from '~/components/dashboard/EventForm.vue'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const route = useRoute()
const eventId = route.params.id
</script>
