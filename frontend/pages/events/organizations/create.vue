<template>
  <div>
    <GoBack />
    <div class="flex justify-center items-center min-h-screen dark:bg-gray-800 bg-white">
      <div class="w-full max-w-md sm:max-w-lg md:max-w-xl lg:max-w-2xl p-4">
        <EventOrganizationForm />
      </div>
    </div>
  </div>
</template>

<script setup>
import EventOrganizationForm from '~/components/dashboard/EventOrganizationForm.vue'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
</script>
