<template>
  <div class="min-w-md max-h-lg rounded overflow-hidden  text-gray-700 text-base dark:text-gray-300  dark:bg-gray-900">
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
      <div>
        <button
          class="text-yellow-500 hover:text-yellow-700"
          @click="emitGoback"
        >
          <i class="fa fa-chevron-left" /> Go Back
        </button>
      </div>
      <img
        v-if="organization.profile_photo_url"
        :src="organization.profile_photo_url"
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

            <button
              class="text-blue-500 hover:text-blue-700 pl-4"
              @click="followEvent(e)"
            >
              <i :class="follows.length>0?'fa fa-user':'fa fa-user-plus'" />
              {{ follows.length>0?'UnFollow':'Follow' }}
            </button>
          </div>
        </div>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import { apolloClient } from '~/plugins/apollo'
import { UN_FOLLOW_EVENT, FOLLOW_EVENT } from '~/graphql/mutation'
import { GET_MY_ID, GET_ORGANIZATIONS, PUBLIC_GET_ORGANIZATIONS, CHECK_AUTH_QUERY, ORG_TOTAL_EVENTS } from '~/graphql/queries'

const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})
const emit = defineEmits(['goback'])
const follows = ref([])
const myId = ref(null)
const totalEventCreated = ref(0)
const organization = ref(null)
const isLoggedIn = ref(false)

const { data: ch, loading, error } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
if (ch && ch.isAuthenticated) {
  isLoggedIn.value = true
}

if (isLoggedIn.value) {
  const { data: orgData } = await apolloClient.query({ query: GET_ORGANIZATIONS, variables: { where: { organization_id: { _eq: props.id } } } })
  const { data_organizations } = orgData
  organization.value = data_organizations[0]
  follows.value = organization.value.follows
  // console.log(organization.value)
  const { data: userData } = await apolloClient.query({ query: GET_MY_ID })

  myId.value = userData.data_users[0].user_id

  const { data: totEvents } = await apolloClient.query({ query: ORG_TOTAL_EVENTS, variables: { organizationId: organization.value.organization_id } })
  totalEventCreated.value = totEvents.data_events_aggregate.aggregate.count
}
else {
  const { data: orgData } = await apolloClient.query({ query: PUBLIC_GET_ORGANIZATIONS, variables: { where: { organization_id: { _eq: props.id } } } })
  const { data_organizations } = orgData
  organization.value = data_organizations[0]
  const { data: totEvents } = await apolloClient.query({ query: ORG_TOTAL_EVENTS, variables: { organizationId: organization.value.organization_id } })
  totalEventCreated.value = totEvents.data_events_aggregate.aggregate.count
}

const followEvent = async () => {
  console.log(follows.value)
  if (follows.value.length > 0 && myId.value) {
    try {
      const { mutate } = useMutation(UN_FOLLOW_EVENT)
      await mutate({ followed_organization_id: follows.value[0].followed_organization_id, following_user_id: follows.value[0].following_user_id })
      follows.value = []
    }
    catch {
      // console.log(err)
      /** */
    }
  }
  else if (myId.value && organization.value && organization.value.organization_id) {
    try {
      const { mutate } = useMutation(FOLLOW_EVENT)
      const { data } = await mutate({ followed_organization_id: organization.value.organization_id, following_user_id: myId.value })
      follows.value = data.insert_data_follows.returning
    }
    catch {
      /** */
      // console.log(err)
    }
  }
}

const emitGoback = () => {
  emit('goback')
}

const formatDate = (date) => {
  const options = { year: 'numeric', month: 'short', day: 'numeric' }
  return new Date(date).toLocaleDateString(undefined, options)
}
</script>
