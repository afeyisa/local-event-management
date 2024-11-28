<script setup>
import { ref } from 'vue'
import OrganizationCard from '~/components/dashboard/OrganizationCard.vue'
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

  const { data, loading, error } = await apolloClient.query({ query: GET_ORGANIZATIONS, variables: { where: { organizes: { organizer_id: { _eq: myId } } } } })
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
      <div class="grid sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-6">
        <OrganizationCard
          v-for="organization in organizations"
          :key="organization.organization_name"
          :organization="organization"
        />
      </div>
    </div>
  </div>
</template>
