<script setup>
import { useMutation } from '@vue/apollo-composable'
import { LOGOUT_MUTATION } from '~/graphql/mutations/logout.graphql'

const logout = async () => {
  try {
    const { mutate } = useMutation(LOGOUT_MUTATION)
    await mutate()
    navigateTo('/')
  }
  catch { /* empty */ }
}
</script>

<template>
  <button
    class="w-full text-left px-4 py-2  dark:text-white dark:hover:text-red-600 hover:text-red-600 hover:underline "
    @click="logout"
  >
    Logout
  </button>
</template>
