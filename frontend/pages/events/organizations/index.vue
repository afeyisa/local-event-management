<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols gap-6">
    <p v-if="error">
      Something went wrong...
    </p>
    <p v-if="loading">
      Loading...
    </p>
    <OrganizationCard
      v-for="organization in organizations"
      v-else
      :key="organization.organization_name"
      :organization="organization"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
// import { useQuery } from '@vue/apollo-composable'
import OrganizationCard from '~/components/dashboard/OrganizationCard.vue'
import { GET_ORGANIZATIONS } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const organizations = ref([])

const { data, loading, error } = await apolloClient.query({ query: GET_ORGANIZATIONS })
// const {data,loading,error} = useQuery(query)
if (!loading && !error) {
  organizations.value = data.data_organizations
  // console.log(data.data_organizations)
}
</script>
