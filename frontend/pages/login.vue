<!-- login -->
<template>
  <div>
    <h2 class="text-2xl font-bold text-center text-gray-800 dark:text-white mb-6">
      Login
    </h2>
    <Form
      @submit="onLogin"
    >
      <!-- email input -->
      <div class="mb-4">
        <label
          class="block text-gray-700 dark:text-gray-300 mb-2"
          for="email"
        >Email
        </label>
        <Field
          v-model="email"
          rules="validateEmail"
          type="email"
          name="email"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter your email"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="email"
        />
      </div>
      <!-- password input -->
      <div class="mb-4">
        <label
          for="password"
          class="block text-gray-700 dark:text-gray-300 mb-2"
        >
          Password
        </label>
        <Field
          v-model="password"
          name="password"
          type="password"
          rules="required"
          placeholder="Password"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="Enter your password"
        />
      </div>

      <button
        :disabled="loading"
        class="w-full bg-blue-500 rounded-md text-white py-2 hover:bg-blue-600 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700"
      >
        Login
      </button>
    </Form>
    <div
      v-if="error"
      class=" text-red-500 mb-4"
    >
      {{ error.message }}
    </div>
    <!-- Switch to Login -->
    <div class="mt-6 text-center">
      <NuxtLink
        to="/signup"
        class="text-blue-500 hover:underline dark:text-blue-400"
      >
        Already have an account? Log in
      </NuxtLink>
    </div>
  </div>
</template>

<script setup>
import { useMutation } from '@vue/apollo-composable'
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { LOGIN_MUTATION } from '~/graphql/mutation'

definePageMeta({
  layout: 'auth',
  middleware: 'logedinautoaedirector',
})

defineRule('validateEmail', (value) => {
  return true
  if (!value || !value.length) {
    return 'Email is required'
  }
  const regex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i
  if (!regex.test(value)) {
    return 'Email must be a valid email'
  }
  return true
})

defineRule('required', (value) => {
  if (!value || !value.length) {
    return 'password is required'
  }
  return true
})

const email = ref('')
const password = ref('')
// const router = useRouter()

const { mutate, loading, error } = useMutation(LOGIN_MUTATION)
const onLogin = async () => {
  try {
    await mutate({
      email: email.value,
      password: password.value,
    },
    )
    const redirectPath = localStorage.getItem('redirect') || '/events' // get the redirect path from query or default to root
    localStorage.removeItem('redirect')
    // router.push(redirectPath)
    window.location.href = window.location.origin + redirectPath
  }
  catch {
    console.log('unable to login')
  }
}
</script>
