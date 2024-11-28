<script setup>
import { useOrganization } from '~/composables/fetchOrgDetail'
import { useFollow } from '~/composables/useFollow'
import { formatDate } from '~/composables/formatDate'

const route = useRoute()
const orgId = route.params.id
// const isFollowingPage = route.path === '/events/following'
// if (isFollowingPage) {
//   definePageMeta({
//     layout: ,
//   })
// }

const { organization, follows, totalEventCreated, myId, loading, error } = await useOrganization(orgId)
const { toggleFollow } = useFollow(myId.value, orgId, follows)
</script>

<template>
  <ClientOnly>
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
          <GoBack />
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
                @click="toggleFollow"
              >
                <i :class="follows.length>0?'fa fa-user':'fa fa-user-plus'" />
                {{ follows.length>0?'UnFollow':'Follow' }}
              </button>
            </div>
          </div>
        </ul>
      </div>
    </div>
  </ClientOnly>
</template>
