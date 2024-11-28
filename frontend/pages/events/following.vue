<script setup>
import { ref } from 'vue'
import { GET_ORGANIZATIONS } from '~/graphql/querie/getOrganization.graphql'
import { GET_MY_ID } from '~/graphql/querie/getUserId.graphql'
import { apolloClient } from '~/plugins/apollo'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const organizations = ref([])
const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
if (!loading && !error) {
  const myId = userData.data_users[0].user_id

  const { data, loading, error } = await apolloClient.query({ query: GET_ORGANIZATIONS, variables: { where: { follows: { following_user_id: { _eq: myId } } } } })
  if (!loading && !error) {
    organizations.value = data.data_organizations
  }
}
</script>

<template>
  <div>
    <p v-if="error">
      Something went wrong...
    </p>
    <p v-if="loading">
      Loading...
    </p>
    <div v-else>
      <GoBack />
      <div
        class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-6"
      >
        <div
          v-for="organization in organizations"
          :key="organization.organization_id"
        >
          <NuxtLink
            class="w-full  dark:text-gray-200 px-1 text-gray-700"
            :to="`/org/${organization.organization_id}`"
          >
            <PublicOrganizationCard :organization="organization" />
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>
