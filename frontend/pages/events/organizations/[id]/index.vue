<script setup>
import { apolloClient } from '~/plugins/apollo'
import { ORG_TOTAL_EVENTS } from '~/graphql/querie/getOrgTotalEvents.graphql'
import { CHECK_AUTH_QUERY } from '~/graphql/querie/checkAuthentication.graphql'
import { PUBLIC_GET_ORGANIZATIONS } from '~/graphql/querie/getPublicOrgs.graphql'
import { GET_ORGANIZATIONS } from '~/graphql/querie/getOrganization.graphql'
import { GET_MY_ID } from '~/graphql/querie/getUserId.graphql'
import { formatDate } from '#imports'
import { fetchBase64Image } from '~/composables/fetchImage'

const route = useRoute()
const Id = route.params.id

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})

const follows = ref([])
const myId = ref(null)
const totalEventCreated = ref(0)
const organization = ref(null)
const isLoggedIn = ref(false)
const base64Image = ref(null)

const { data: ch, loading, error } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
if (ch && ch.isAuthenticated) {
  isLoggedIn.value = true
}

if (isLoggedIn.value && Id) {
  const { data: orgData } = await apolloClient.query({ query: GET_ORGANIZATIONS, variables: { where: { organization_id: { _eq: Id } } } })
  const { data_organizations } = orgData
  organization.value = data_organizations[0]
  follows.value = organization.value.follows
  const { data: userData } = await apolloClient.query({ query: GET_MY_ID })

  myId.value = userData.data_users[0].user_id

  const { data: totEvents } = await apolloClient.query({ query: ORG_TOTAL_EVENTS, variables: { organizationId: organization.value.organization_id } })
  totalEventCreated.value = totEvents.data_events_aggregate.aggregate.count
}
else {
  const { data: orgData } = await apolloClient.query({ query: PUBLIC_GET_ORGANIZATIONS, variables: { where: { organization_id: { _eq: organization.value.id } } } })
  const { data_organizations } = orgData
  organization.value = data_organizations[0]
  const { data: totEvents } = await apolloClient.query({ query: ORG_TOTAL_EVENTS, variables: { organizationId: organization.value.organization_id } })
  totalEventCreated.value = totEvents.data_events_aggregate.aggregate.count
}

const prefetchImages = async () => {
  base64Image.value = await fetchBase64Image(organization.value.profile_photo_url)
}
onMounted(() => {
  prefetchImages()
})
</script>

<template>
  <div class="min-w-md max-h-lg rounded overflow-hidden  text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900">
    <GoBack />
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
      Error loading event details
    </div>

    <div
      v-if="organization"
      class="max-w-4xl mx-auto p-6 bg-white rounded-lg s min-w-md max-h-lg  overflow-hidden shadow-lg text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900"
    >
      <img
        v-if="organization.profile_photo_url"
        :src="base64Image"
        alt="Event Image"
        class="w-full h-64 object-cover rounded-md mb-6"
      >
      <h1 class="text-3xl font-bold mb-4 dark:text-gray-300 ">
        {{ organization.organization_name }}
      </h1>
      <p
        v-if="organization.description"
        class="text-gray-700 mb-6 dark:text-gray-300 "
      >
        <i
          class="fa fa-info-circle p-2"
        />  {{ organization.description }}
      </p>

      <ul class="list-none list-inside mb-6 text-gray-700 space-y-2 dark:text-gray-300 ">
        <li>
          <em> bio </em>{{ organization.bio?organization.bio:'' }}
        </li>

        <div class="text-left">
          <p class="text-sm  dark:text-gray-300">
            Since {{ formatDate(organization.created_at) }}
          </p>
          <li>
            <i class=" p-2 fa fa-calendar" />
            {{ totalEventCreated }}created events
          </li>
          <div class="flex ">
            <div
              class="text-blue-500"
            >
              <i class="fa fa-user" />
              {{ organization.followers+' '+'Followers' }}
            </div>
          </div>
        </div>
      </ul>
    </div>
  </div>
</template>
