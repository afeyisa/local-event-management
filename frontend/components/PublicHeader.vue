<script setup>
import { ref } from 'vue'
import { isLoggedIn } from '~/composables/user'

const checkLoggedin = ref(false)
checkLoggedin.value = await isLoggedIn()
</script>

<template>
  <ul class="flex items-end sm:items-center text-sm sm:text-lg  space-x-4">
    <li>
      <NuxtLink
        to="/events"
        class="hover:underline"
      >
        <i class="fa fa-calendar " />
        My Events
      </NuxtLink>
    </li>
    <li v-show="!checkLoggedin">
      <NuxtLink
        to="/login"
        class="hover:underline"
      >
        <i class="fa fa-sign-in" />
        Login
      </NuxtLink>
    </li>
    <li v-show="!checkLoggedin">
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
