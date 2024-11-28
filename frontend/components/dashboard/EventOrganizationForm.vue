<script setup>
import { ref } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { useToast } from 'vue-toastification/dist/index.mjs'
import { UPDATE_ORGANIZATION } from '~/graphql/mutations/updateOrganization.graphql'
import { CREATE_ORGANIZATION } from '~/graphql/mutations/createOraganization.graphql'
import { apolloClient } from '~/plugins/apollo'
import { GET_ORGANIZATIONS } from '~/graphql/querie/getOrganization.graphql'
import { UPLOAD_ORG_PROFILE_URL } from '~/graphql/mutations/saveOrgProfileImageurl.graphql'
import { uploadImages } from '~/composables/uploadImages'

const { id } = defineProps({
  id: {
    type: String,
    required: false,
  },
})
const toast = useToast()

defineRule('required', (value) => {
  if (!value || !value.length) {
    return 'This field is required'
  }
  return true
})

const base64Thummbnail = ref(null)
const formData = ref({
  organization_name: ref(''),
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

const handleThumbnailChange = async (e) => {
  if (id) {
    try {
      const t = await encodeB64(e.target.files[0])
      const res = await uploadImages(t, [])
      if (res) {
        const { mutate } = useMutation(UPLOAD_ORG_PROFILE_URL)
        const { data } = await mutate({ organization_id: id, profile_photo_url: res.thumbnailImageUrl })
        if (data) {
          toast.success('thumbail changed')
        }
      }
    }
    catch (err) {
      console.log(err)
      toast.error('unbale update thumbnail')
    }
  }
  else {
    try {
      base64Thummbnail.value = await encodeB64(e.target.files[0])
      toast.info('image encoded')
    }
    catch {
      toast.error('some thing went wrong during thumbnail encoding')
    }
  }
}

const handleSubmit = async () => {
  try {
    if (id) {
      const { mutate: createOrganization } = useMutation(UPDATE_ORGANIZATION)
      const { data } = await createOrganization({
        organization_name: formData.value.organization_name,
        bio: formData.value.bio,
        description: formData.value.description,
        organization_id: id,
      })
      if (data) {
        toast.success('success')
      }
      const router = useRouter()
      router.push(`/events/organizations/${data.update_data_organizations_by_pk.organization_id}`)
    }
    else {
      const { mutate: createOrganization } = useMutation(CREATE_ORGANIZATION)
      const { data } = await createOrganization({
        organization_name: formData.value.organization_name,
        bio: formData.value.bio,
        description: formData.value.description,
      })

      const res = await uploadImages(base64Thummbnail.value, [])
      if (res.thumbnailImageUrl) {
        const { mutate } = useMutation(UPLOAD_ORG_PROFILE_URL)
        const { data: d } = await mutate({ organization_id: data.insert_data_organizations_one.organization_id, profile_photo_url: res.thumbnailImageUrl })
        if (d) {
          toast.success('success')
        }
      }

      const router = useRouter()
      router.push(`/events/organizations/${data.insert_data_organizations_one.organization_id}`)
    }
  }
  catch {
    toast.error('Failed to submit form')
  }
}
</script>

<template>
  <div class="p-8 dark:bg-gray-800 dark:text-gray-300">
    <h1 class="text-3xl font-bold mb-6">
      {{ id?'update organization':' Create  Organization' }}
    </h1>

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
            :onchange="handleThumbnailChange"
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
  </div>
</template>
