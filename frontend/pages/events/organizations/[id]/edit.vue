<template>
  <div>
    <GoBack />
    <div class="flex justify-center items-left min-h-screen dark:bg-gray-800 bg-white">
      <div class="w-full max-w-md sm:max-w-lg md:max-w-xl lg:max-w-2xl p-4">
        <DashboardEventOrganizationForm
          :id="Id"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const route = useRoute()
const Id = route.params.id
</script>
