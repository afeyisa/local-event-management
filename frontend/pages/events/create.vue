<template>
  <EventForm />
</template>

<script setup>
import EventForm from '~/components/dashboard/EventForm.vue'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
</script>
