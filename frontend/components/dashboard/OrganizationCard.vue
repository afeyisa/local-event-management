<template>
  <div class="bg-white dark:bg-gray-900 shadow-md rounded-lg p-4">
    <!-- Organization Profile Photo -->
    <div class="flex items-center dark:text-gray-300 mb-4">
      <img
        :src="organization.profile_photo_url || 'https://via.placeholder.com/400x300.png?text=Vue+Conference'"
        alt="Organization Profile"
        class="w-16 h-16 rounded-full object-cover mr-4"
      >

      <div>
        <h2 class="text-xl font-bold text-gray-600 dark:text-gray-300">
          {{ organization.organization_name }}
        </h2>
        <p class="text-sm  dark:text-gray-300">
          Created on: {{ new Date(organization.created_at).toLocaleDateString() }}
        </p>
      </div>
    </div>

    <!-- Organization Bio -->
    <p class="text-gray-600text-base dark:text-gray-300 mb-2">
      <strong>Bio:</strong> {{ organization.bio || 'No bio available.' }}
    </p>

    <!-- Organization Description -->
    <p class="text-gray-600 dark:text-gray-400">
      <strong>Description:</strong> {{ organization.description || 'No description available.' }}
    </p>
    <div class="px-6 py-0">
      <button
        :to="`/events/${organization.id}/edit`"
        class="inline-block text-red-500  px-3 py-2 rounded hover:bg-blue-500"
      >
        Delete Event
      </button>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

// Props to receive organization data
defineProps({
  organization: {
    type: Object,
    required: true,
  },
})
</script>
