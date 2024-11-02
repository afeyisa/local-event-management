<template>
  <div class="p-8 dark:bg-gray-800 dark:text-gray-300">
    <h1 class="text-3xl font-bold mb-6">
      {{ id?'update organization':' Create  Organization' }}
    </h1>

    <!-- Form -->
    <Form
      class="space-y-4 "
      @submit="handleSubmit"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          required
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="organization_name"
        />
      </div>

      <div class="flex items-center justify-center w-full">
        <label
          for="thumbnail-file"
          class="flex flex-col items-center justify-center w-full h-32 border-2 border-gray-300 border-dashed rounded-lg cursor-pointer bg-gray-50 dark:hover:bg-gray-800 dark:bg-gray-700 hover:bg-gray-100 dark:border-gray-600 dark:hover:border-gray-500"
        >
          <div class="flex flex-col items-center justify-center pt-5 pb-6">
            <svg
              class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 20 16"
            >
              <path
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"
              />
            </svg>
            <p class="mb-2 text-sm text-gray-500 dark:text-gray-400"><span class="font-semibold">Click to upload</span> </p>
            <p class="text-xs text-gray-500 dark:text-gray-400">images</p>
          </div>
          <Field
            id="thumbnail-file"
            v-model="formData.profile"
            type="file"
            class="hidden"
            accept="image/*"
            name="thumnail"
          />
        </label>
      </div>

      <div>
        <label
          for="bio"
          class="block text-sm font-medium text-gray-700 dark:text-gray-300"
        >Bio</label>
        <Field
          id="bio"
          v-model="formData.bio"
          name="bio"
          type="text"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
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
        >
          Description
        </label>
        <textarea
          id="description"
          v-model="formData.description"
          rows="4"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        />
      </div>

      <button
        type="submit"
        class="px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700"
      >
        {{ id?'Update ':'create' }}
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
import { ref } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { CREATE_ORGANIZATION, UPDATE_ORGANIZATION } from '~/graphql/mutation'
import { apolloClient } from '~/plugins/apollo'
import { GET_ORGANIZATIONS } from '~/graphql/queries'

const { id } = defineProps({
  id: {
    type: String,
    required: false,
  },
})

defineRule('required', (value) => {
  if (!value || !value.length) {
    return 'This field is required'
  }
  return true
})

// Form data and mutation setup
const formData = ref({
  organization_name: ref(''),
  profile: ref(null),
  bio: ref(''),
  description: ref(''),
})

if (id) {
  try {
    const { data } = await apolloClient.query({ query: GET_ORGANIZATIONS, variables: { where: { organization_id: { _eq: id } } } })
    formData.value.organization_name = data.data_organizations[0].organization_name
    formData.value.bio = data.data_organizations[0].bio
    formData.value.description = data.data_organizations[0].description
  }
  catch {
    /** */
  }
}
// Success and error messages
const successMessage = ref('')
const errorMessage = ref('')

// Create or update mutation setup
const handleSubmit = async () => {
  try {
    // upload profile image
    let profile_photo_url
    if (formData.value.profile) {
      const formdata = new FormData()
      formdata.append('thumbnailimage', formData.value.profile)
      const response = await fetch('http://localhost:4000/upload', {
        method: 'POST',
        body: formdata,
      })
      if (response.ok) {
        const { thumbnail_image_url } = await response.json()
        profile_photo_url = thumbnail_image_url
      }
    }

    if (id) {
      const { mutate: createOrganization } = useMutation(UPDATE_ORGANIZATION)
      const { data } = await createOrganization({
        organization_name: formData.value.organization_name,
        profile_photo_url: profile_photo_url,
        bio: formData.value.bio,
        description: formData.value.description,
        organization_id: id,
      })
      const router = useRouter()
      console.log(data)
      router.push(`/events/organizations/${data.update_data_organizations_by_pk.organization_id}`)
    }
    else {
    // create organization
      const { mutate: createOrganization } = useMutation(CREATE_ORGANIZATION)
      const { data } = await createOrganization({
        organization_name: formData.value.organization_name,
        profile_photo_url: profile_photo_url,
        bio: formData.value.bio,
        description: formData.value.description,
      })
      const router = useRouter()
      router.push(`/events/organizations/${data.insert_data_organizations_one.organization_id}`)
    }
  }
  catch (error) {
    errorMessage.value = 'Failed to submit form: ' + error.message
  }
}
</script>
