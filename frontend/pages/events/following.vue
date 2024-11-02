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
        v-if="!orgId"
        class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-6"
      >
        <PublicOrganizationCard
          v-for="organization in organizations"
          :key="organization.organization_id"
          :organization="organization"
          @click="emitOrgId(organization.organization_id)"
        />
      </div>
      <div v-else>
        <OrgDetail
          :id="orgId" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
// import { useQuery } from '@vue/apollo-composable'
import { GET_ORGANIZATIONS, GET_MY_ID } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

// const emit = defineEmits(['orgid', 'eventid'])

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const organizations = ref([])
const orgId = ref(null)
const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
if (!loading && !error) {
  const myId = userData.data_users[0].user_id

  const { data, loading, error } = await apolloClient.query({ query: GET_ORGANIZATIONS, variables: { where: { follows: { following_user_id: { _eq: myId } } } } })
  if (!loading && !error) {
    organizations.value = data.data_organizations
  }
}
const emitOrgId = (o) => {
  orgId.value = o
}
</script>
