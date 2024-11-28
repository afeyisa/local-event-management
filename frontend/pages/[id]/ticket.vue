<script setup>
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { useTicket } from '~/composables/useTicket'

definePageMeta({
  layout: 'auth',
  middleware: 'auth',
})

const route = useRoute()
const eventId = route.params.id
const ticket = useTicket()
const buy = async () => {
  await ticket.buyTicket(eventId)
}

defineRule('validateEmail', (value) => {
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
    return 'this field is required'
  }
  return true
})

defineRule('phone', (value) => {
  const phoneRegex = /^(09|07)\d{8}$/
  return phoneRegex.test(value) || 'must start with 09 or 07 and have 10 digits'
})
</script>

<template>
  <div>
    <h2 class="text-2xl font-bold text-center text-gray-800 dark:text-white mb-6">
      Payment Form
    </h2>
    <Form :submit="buy">
      <div class="mb-4">
        <label
          class="flex text-gray-700 dark:text-gray-300 mb-2"
          for="firstName"
        >
          <i class="fa fa-user text-gray-400 dark:text-gray-200 mr-2" />
          First Name<p class="text-red-500">*</p> <ErrorMessage
            class="text-red-500 mb-4 pl-4"
            name="firstName"
          />
        </label>
        <Field
          v-model="ticket.form.value.firstName"
          rules="required"
          type="name"
          name="firstName"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter your name"
        />
      </div>
      <div class="mb-4">
        <label
          class=" text-gray-700 dark:text-gray-300 mb-2 flex"
          for="lastname"
        >
          <i class="fa fa-user text-gray-400 dark:text-gray-200 mr-2" />
          Last Name <p class="text-red-500">*</p>
          <ErrorMessage
            class="text-red-500 mb-4 pl-4"
            name="lastname"
          />
        </label>
        <Field
          v-model="ticket.form.value.lastname"
          rules="required"
          type="name"
          name="lastname"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter your name"
        />
      </div>
      <div class="mb-4">
        <label
          class=" text-gray-700 dark:text-gray-300 mb-2 flex"
          for="phone"
        >
          <i class="fa fa-phone text-gray-400 dark:text-gray-200 mr-2" />
          Phone <p class="text-red-500">*</p> <ErrorMessage
            class="text-red-500 mb-4 pl-4"
            name="phone"
          />
        </label>
        <Field
          v-model="ticket.form.value.phone"
          rules="phone"
          type="tel"
          name="phone"
          class="mt-1 block w-full text-xs border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="phone"
        />
      </div>

      <div class="mb-4">
        <label
          class="flex text-gray-700 dark:text-gray-300 mb-2"
          for="email"
        >
          <i class="fa fa-envelope text-gray-400 dark:text-gray-200 mr-2" />Email
          <p class="text-red-500">*</p>
          <ErrorMessage
            class="text-red-500 mb-4"
            name="email"
          />
        </label>
        <Field
          v-model="ticket.form.value.email"
          rules="validateEmail"
          type="email"
          name="email"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter your email"
        />
      </div>

      <button
        class="flex items-center justify-center w-full bg-blue-700 rounded-md text-white py-2 hover:bg-blue-800 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700"
      >
        Start
      </button>
    </Form>
    <div class="mt-6 text-center" />
  </div>
</template>
