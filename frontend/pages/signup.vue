<template>
  <div>
    <h2 class="text-2xl font-bold text-center text-gray-800 dark:text-white mb-6">
      Signup
    </h2>
    <Form
      @submit="onSummit"
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
          r
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
          rules="passordValidate"
          placeholder="Password"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="password"
        />
      </div>
      <div class="mb-4">
        <label
          class="block text-gray-700 dark:text-gray-300 mb-2"
          for="confirm password"
        >
          Confirm Password
        </label>
        <Field
          name="confirmation"
          type="password"
          placeholder="Confirm Password"
          rules="confirmed:password"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        />
      </div>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="confirmation"
      />

      <button
        :disabled="loading"
        class="w-full bg-blue-500 rounded-md text-white py-2 hover:bg-blue-600 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700"
      >
        Signup
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
        to="/login"
        class="text-blue-500 hover:underline dark:text-blue-400"
      >
        Already have an account? Log in
      </NuxtLink>
    </div>
  </div>
</template>

<script setup>
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { useMutation } from '@vue/apollo-composable'
import { SIGNUP_MUTATION } from '~/graphql/mutation'

const limit = 5

definePageMeta({
  layout: 'auth',
  middleware: 'auth',
})
//* ***********validation ********************************/
defineRule('validateEmail', (value) => {
  // console.log(value)
  if (!value) {
    return 'This field is required'
  }

  const regex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i
  if (!regex.test(value)) {
    return 'This field must be a valid email'
  }

  if (value.length < limit) {
    return `This must be at least ${limit} character`
  }

  return true
})

defineRule('confirmed', (value, [target], ctx) => {
  if (value === ctx.form[target]) {
    return true
  }
  return 'Passwords must match'
})

defineRule('passordValidate', (value) => {
  if (!value || !value.length) {
    return 'password is required'
  }

  if (!/[0-9]/.test(value)) {
    return 'Password must contain number'
  }
  if (!/[A-Z]/.test(value)) {
    return 'Password must contain at least one uppercase letter'
  }

  if (!/[a-z]/.test(value)) {
    return 'Password must contain at least one lowercase letter'
  }

  if (!/[!@#$%^&*(),.?":{}|<>]/.test(value)) {
    return 'Password must contain at least one special character'
  }
  if (value.length < 5) {
    return 'password must contain 5 characters'
  }
  if (value.length > 128) {
    return 'Password cannot exceed 128 characters'
  }

  return true
})
//* *********************************************************************** */

const email = ref('')
const password = ref('')
const router = useRouter()

const { mutate, loading, error } = useMutation(SIGNUP_MUTATION)
const onSummit = async () => {
  try {
    await mutate({
      email: email.value,
      password: password.value,
    },
    )
    router.push('/')
  }
  catch {
    console.log('unable to signup')
  }
}
</script>
