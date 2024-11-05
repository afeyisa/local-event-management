<script setup>
import { ref } from 'vue'
import { apolloClient } from '~/plugins/apollo'
import { CHECK_AUTH_QUERY } from '~/graphql/queries'

const isLoggedIn = ref(false)
const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
if (data && data.isAuthenticated) {
  isLoggedIn.value = true
}
</script>

<template>
  <ul class="flex items-center space-x-4">
    <li>
      <NuxtLink
        to="/events"
        class="hover:underline"
      >
        <i class="fa fa-calendar " />
        My Events
      </NuxtLink>
    </li>
    <li v-show="!isLoggedIn">
      <NuxtLink
        to="/login"
        class="hover:underline"
      >
        <i class="fa fa-sign-in" />
        Login
      </NuxtLink>
    </li>
    <li v-show="!isLoggedIn">
      <NuxtLink
        to="/signup"
        class="hover:underline"
      ><i class="fa fa-user-plus" />
        Signup
      </NuxtLink>
    </li>
    <ThemeButton />
  </ul>
</template>
