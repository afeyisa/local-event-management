<script setup>
import { defineProps } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import { DELETE_ORG } from '~/graphql/mutations/deleteOrganization.graphql'
import { fetchBase64Image } from '~/composables/fetchImage'

const props = defineProps({
  organization: {
    type: Object,
    required: true,
  },
})

const organization = ref(props.organization)
const base64Image = ref(null)
const prefetchImages = async () => {
  base64Image.value = await fetchBase64Image(organization.value.profile_photo_url)
}

// watch(organization.profile_photo_url, prefetchImages, { immediate: true })

onMounted(() => {
  prefetchImages()
})
const deleteOrg = async () => {
  const confirmed = confirm('Are you sure you want to delete this organization? This action cannot be undone.')

  if (confirmed) {
    try {
      const { mutate } = useMutation(DELETE_ORG)
      const { data } = await mutate({ organization_id: organization.value.organization_id })

      if (data) {
        organization.value = null
        alert('Organization deleted successfully.')
      }
    }
    catch (err) {
      console.log(err)
      alert('An error occurred while deleting the organization.')
    }
  }
  else {
    alert('Deletion canceled.')
  }
}
</script>

<template>
  <div
    v-if="organization"
    class="bg-white w-full dark:bg-gray-900 shadow-md rounded-lg p-4"
  >
    <div class="flex items-center dark:text-gray-300 mb-4">
      <img
        :src="base64Image || 'https://via.placeholder.com/400x300.png?text=Vue+Conference'"
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

    <p class="text-gray-600text-base dark:text-gray-300 mb-2">
      <strong>Bio:</strong> {{ organization.bio || 'No bio available.' }}
    </p>

    <p class="text-gray-600 dark:text-gray-400">
      <strong>Description:</strong> {{ organization.description || 'No description available.' }}
    </p>
    <div class="flex">
      <div class="px-6 py-4">
        <NuxtLink
          :to="`/events/organizations/${organization.organization_id}`"
          class="inline-block bg-blue-500 text-white px-3 py-2 rounded hover:bg-blue-600 dark:bg-blue-700 dark:hover:bg-blue-800"
        >
          Details
        </NuxtLink>
      </div>
      <div class="px-6 py-4">
        <NuxtLink
          :to="`/events/organizations/${organization.organization_id}/edit`"
          class="inline-block bg-yellow-500 text-white px-3 py-2 rounded hover:bg-yellow-600 dark:bg-yellow-700 dark:hover:bg-yellow-800"
        >
          Update
        </NuxtLink>
      </div>
      <div class="px-6 py-4">
        <button
          class="inline-block text-red-500  px-3 py-2 rounded"
          @click="deleteOrg"
        >
          Delete
        </button>
      </div>
    </div>
  </div>
</template>
