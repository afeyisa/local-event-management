import { ref } from 'vue'
import { useToast } from 'vue-toastification/dist/index.mjs'
import { isLoggedIn, fetchUseId } from './user'

import { apolloClient } from '~/plugins/apollo'
import { ORG_TOTAL_EVENTS } from '~/graphql/querie/getOrgTotalEvents.graphql'
import { GET_ORGANIZATIONS } from '~/graphql/querie/getOrganization.graphql'
import { PUBLIC_GET_ORGANIZATIONS } from '~/graphql/querie/getPublicOrgs.graphql'

export async function useOrganization(id) {
  const organization = ref(null)
  const follows = ref([])
  const totalEventCreated = ref(0)
  const loading = ref(true)
  const error = ref(null)
  const { myId } = await fetchUseId()

  try {
    const orgQuery = await isLoggedIn() ? GET_ORGANIZATIONS : PUBLIC_GET_ORGANIZATIONS
    const { data: orgData } = await apolloClient.query({ query: orgQuery, variables: { where: { organization_id: { _eq: id } } } })
    organization.value = orgData.data_organizations[0]
    follows.value = organization.value.follows || []
    const totalEventsData = await apolloClient.query({ query: ORG_TOTAL_EVENTS, variables: { organizationId: organization.value.organization_id } })
    totalEventCreated.value = totalEventsData.data.data_events_aggregate.aggregate.count
  }
  catch {
    const toast = useToast()
    toast.error('')
  }
  finally {
    loading.value = false
  }

  return { organization, follows, totalEventCreated, myId, loading, error }
}
