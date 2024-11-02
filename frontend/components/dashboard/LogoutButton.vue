<template>
  <button
    class="w-full text-left px-4 py-2  dark:text-white dark:hover:text-red-600 hover:text-red-600 hover:underline "
    @click="logout"
  >
    Logout
  </button>
</template>

<script setup>
import { useMutation } from '@vue/apollo-composable'
import { LOGOUT_MUTATION } from '~/graphql/mutation'

const logout = async () => {
  try {
    // const router = useRouter()
    const { mutate } = useMutation(LOGOUT_MUTATION)
    await mutate()
    // router.push('/')
    window.location.href = window.location.origin
  }
  catch { /* empty */ }
}
</script>
