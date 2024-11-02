<!-- event detail page -->

<template>
  <div class="min-w-md max-h-lg rounded overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300 bg-white dark:bg-gray-900">
    <!-- Event Image -->
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
      Error loading organization details
    </div>
    <div
      v-if="organization"
      class="max-w-4xl mx-auto p-6 bg-white rounded-lg s min-w-md max-h-lg  overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900"
    >
      <img
        v-if="organization.profile_photo_url"
        :src="organization.profile_photo_url"
        alt="organization Image"
        class="w-full h-64 object-cover rounded-md mb-6"
      >
      <h1 class="text-3xl font-bold mb-4 dark:text-gray-300 ">
        {{ organization.organization_name }}
      </h1>
      <p class="text-gray-700 mb-6 dark:text-gray-300 ">
        {{ organization.description }}
      </p>
      <ul class="list-disc list-inside mb-6 text-gray-700 space-y-2 dark:text-gray-300 ">
        <li>since{{'  '+ formattedDate() }}</li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { useMutation } from '@vue/apollo-composable'
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { GET_ORGANIZATIONS } from '~/graphql/queries'

const route = useRoute()
const organization = ref(null)
const isFollowing = ref(false)

// async function loadOrganization() {
const orgId = route.params.id
const { mutate } = useMutation(GET_ORGANIZATIONS)
const { data } = await mutate({ where: { organization_id: { _eq: orgId } } })
organization.value = data.data_organizations[0]
console.log(data.data_organizations[0].organization_name)
// }

function toggleFollow() {
  isFollowing.value = !isFollowing.value
  // Call follow/unfollow API if needed
}

const formattedDate = () => {
  return new Date(organization.value.created_at).toLocaleDateString()
}
</script>

  <style scoped>
  .organization-detail { /* Styles for the organization detail page */ }
  .organization-header { /* Styles for organization header */ }
  .org-logo { /* Style for org logo */ }
  .organization-card { /* Style for individual organization cards */ }
  .organization-image { /* Style for event images */ }
  </style>
