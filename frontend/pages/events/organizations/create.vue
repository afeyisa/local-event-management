<template>
  <EventOrganizationForm />
</template>

<script setup>
import EventOrganizationForm from '~/components/dashboard/EventOrganizationForm.vue'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
</script>
