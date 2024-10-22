<template>
  <div class="p-8 dark:bg-gray-800 dark:text-gray-300">
    <h1 class="text-3xl font-bold mb-6">
      {{ formMode === 'create' ? 'Create' : 'Update' }} Organization
    </h1>

    <!-- Form -->
    <Form
      class="space-y-4 "
      @submit.prevent="handleSubmit"
    >
      <div class="mb-4">
        <label
          for="organization_name"
          class="block text-sm font-medium text-gray-700 dark:text-gray-300"
        >Organization Name</label>
        <Field
          v-model="formData.organization_name"
          name="organization_name"
          rules="required"
          type="text"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:border-blue-500 focus:outline-none"
          required
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="organization_name"
        />
      </div>

      <div class="mb-4">
        <label
          for="profile_photo_url"
          class="block text-sm font-medium text-gray-700 dark:text-gray-300"
        >Profile Photo URL</label>
        <Field
          id="profile_photo_url"
          v-model="formData.profile_photo_url"
          rules="required"
          name="profile_photo_url"
          type="url"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:border-blue-500 focus:outline-none"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="profile_photo_url"
        />
      </div>

      <div>
        <label
          for="bio"
          class="block text-sm font-medium text-gray-700 dark:text-gray-300"
        >Bio</label>
        <Field
          id="bio"
          v-model="formData.bio"
          rules="required"
          name="bio"
          type="text"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:border-blue-500 focus:outline-none"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="bio"
        />
      </div>

      <div>
        <label
          for="description"
          class="block text-sm font-medium text-gray-700 dark:text-gray-300"
        >Description</label>
        <textarea
          id="description"
          v-model="formData.description"
          rows="4"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:border-blue-500 focus:outline-none"
        />
      </div>
      <button
        type="submit"
        class="px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700"
      >
        {{ formMode === 'create' ? 'Create' : 'Update' }} Organization
      </button>
    </Form>

    <!-- Success message -->
    <div
      v-if="successMessage"
      class="mt-4 text-green-600"
    >
      {{ successMessage }}
    </div>

    <!-- Error message -->
    <div
      v-if="errorMessage"
      class="mt-4 text-red-600"
    >
      {{ errorMessage }}
    </div>
  </div>
</template>

<script setup>
// Import necessary modules and functions
import { ref } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { CREATE_ORGANIZATION } from '~/graphql/mutation'

// Define the layout and authentication middleware
definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})

defineRule('required', (value) => {
  if (!value || !value.length) {
    return 'This field is required'
  }
  return true
})

// Form data and mutation setup
const formData = ref({
  organization_name: '',
  profile_photo_url: '',
  bio: '',
  description: '', // Replace with the organizer's ID
})

// Form mode: 'create' or 'update'
const formMode = ref('create')

// Success and error messages
const successMessage = ref('')
const errorMessage = ref('')

// Create or update mutation setup
const { mutate: createOrganization } = useMutation(CREATE_ORGANIZATION)
// const { mutate: updateOrganization } = useMutation(UPDATE_ORGANIZATION)
// Handle form submission
const handleSubmit = async () => {
  try {
    if (formMode.value === 'create') {
      // Create new organization
      const data = await createOrganization({
        organization_name: formData.value.organization_name,
        profile_photo_url: formData.value.profile_photo_url,
        bio: formData.value.bio,
        description: formData.value.description,
      })
      successMessage.value = 'Organization created successfully!'
      console.log(data)
    }
    // Reset the form
    formData.value = {
      organization_name: '',
      profile_photo_url: '',
      bio: '',
      description: '',
    }
  }
  catch (error) {
    console.log(error)
    errorMessage.value = 'Failed to submit form: ' + error.message
  }
}
</script>
