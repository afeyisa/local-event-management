<script setup>
import { defineProps } from 'vue'

// Props to receive organization data
defineProps({
  organization: {
    type: Object,
    required: true,
  },
})
const formatDate = (date) => {
  const options = { year: 'numeric', month: 'short', day: 'numeric' }
  return new Date(date).toLocaleDateString(undefined, options)
}
</script>

<template>
  <div class="w-full bg-gray-100  dark:bg-gray-800 shadow-sm rounded-lg p-1">
    <!-- Organization Profile Photo -->
    <div class="flex items-center dark:text-gray-300 mb-4">
      <img
        :src="organization.profile_photo_url || 'https://via.placeholder.com/400x300.png?text=Vue+Conference'"
        alt="Organization Profile"
        class="w-16 h-16 rounded-full object-cover mr-4"
      >

      <div class="text-left">
        <h2 class="text-xl font-bold text-gray-600 dark:text-gray-300">
          {{ organization.organization_name }}
        </h2>
        <p class="text-sm  dark:text-gray-300">
          Since {{ formatDate(organization.created_at) }}
        </p>
        <button
          class="text-blue-500 "
          @click="followEvent(e)"
        >
          <i :class="organization.followers>0?'fa fa-user':'fa fa-user-plus'" />
          {{organization.followers+' '+'Followers' }}
        </button>
      </div>
    </div>
  </div>
</template>
