<template>
  <div class="event-form pb-20 max-w-xl mx-auto bg-white dark:bg-gray-800">
    <h1 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
      Create Event
    </h1>
    <Form
      class="space-y-4"
      @submit="handleSubmit"
    >
      <h2 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
        About Event
      </h2>
      <!-- Organization ID -->
      <label
        for="by_organization_id"
        class="block text-gray-700 dark:text-gray-300"
      >For Organization</label>
      <Field
        v-model="formData.byOrganizationId"
        name="by_organization_id"
        rules="required"
        as="select"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Choose Your Organization"
      >
        <option
          v-for="org in organizations"
          :key="org.organization_id"
          :value="org. organization_id"
          class="bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 p-2"
        >
          {{ org.organization_name }}
        </option>
      </Field>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="by_organization_id"
      />
      <!-- Title -->
      <label
        for="title"
        class="block text-gray-700 dark:text-gray-300"
      >Title</label>
      <Field
        v-model="formData.title"
        name="title"
        rules="required"
        type="text"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Enter event title"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="title"
      />
      <!-- Event Date -->
      <label
        for="event_date"
        class="block text-gray-700 dark:text-gray-300"
      >Event Date</label>
      <Field
        v-model="formData.eventDate"
        name="event_date"
        rules="required"
        type="date"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="event_date"
      />

      <!-- Category ID -->
      <label
        for="category_id"
        class="block text-gray-700 dark:text-gray-300"
      >Category</label>
      <Field
        v-model="formData.categoryId"
        name="category_id"
        rules="required"
        as="select"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
      >
        <option
          v-for="ctg in catagoris"
          :key="ctg.category_id"
          :value="ctg.category_id"
          class="bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 p-2"
        >
          {{ ctg.category_name }}
        </option>
      </Field>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="category_id"
      />
      <!-- Venue -->
      <!-- Event Type (Online or In-Person) -->
      <label
        for="is_online"
        class="block text-gray-700 dark:text-gray-300"
      >Event Type</label>
      <Field
        v-model="formData.isOnline"
        name="is_online"
        rules="required"
        as="select"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        @change="setIsOnline"
      >
        <option value="true">
          Online
        </option>
        <option value="false">
          In-Person
        </option>
      </Field>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="is_online"
      />

      <h2 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
        About Event Tickets
      </h2>
      <!-- Ticket Price -->
      <label
        for="ticket_price"
        class="block text-gray-700 dark:text-gray-300"
      >Ticket Price</label>
      <Field
        v-model="formData.ticketPrice"
        name="ticket_price"
        type="number"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Enter ticket price"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="ticket_price"
      />

      <!-- Total Available Tickets -->
      <label
        for="total_available_tickets"
        class="block text-gray-700 dark:text-gray-300"
      >Total Available Tickets</label>
      <Field
        v-model="formData.totalAvailableTickets"
        name="total_available_tickets"
        type="number"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Total tickets"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="total_available_tickets"
      />
      <div v-if="isInPerson">
        <!-- Region  -->
        <h2 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
          About Event Event Location
        </h2>      <label
          for="region"
          class="block text-gray-700 dark:text-gray-300"
        >Region</label>
        <Field
          v-model="formData.region"
          name="region"
          type="text"
          class="mt-1 block w-full border mb-4 border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Region"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="region"
        />
        <!-- city name -->
        <label
          for="city"
          class="block text-gray-700 dark:text-gray-300"
        >City</label>
        <Field
          v-model="formData.city"
          name="city"
          rules="required"
          type="text"
          class="mt-1 block w-full mb-4 border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="City"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="city"
        />
        <!-- street name -->
        <label
          for="street"
          class="block text-gray-700 dark:text-gray-300"
        >Street</label>
        <Field
          v-model="formData.street"
          name="street"
          rules="required"
          type="text"
          class="mt-1 block w-full  mb-4 border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Street"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="Street"
        />
        <label
          for="venue"
          class="block text-gray-700 dark:text-gray-300"
        >Venue</label>
        <Field
          v-model="formData.venue"
          name="venue"
          rules="required"
          type="text"
          class="mt-1 block w-full border mb-4 border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter venue"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="venue"
        />
      </div>
      <!-- Tags -->
      <div
        class="mt-2"
      >
        <h3 class="block text-gray-700 dark:text-gray-300">
          Select tags
        </h3>
        <div class="mt-1 min-h-10 block border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100">
          <ul class="flex flex-wrap gap-2">
            <li
              v-for="(t, index) in selectedTag"
              :key="t"
              class="inline-block p-1 bg-gray-200 dark:bg-gray-600 rounded-md"
            >
              <button
                type="button"
                class="ml-2  text-gray-700 dark:text-gray-300"
                @click="removeTag(index, t)"
              >
                {{ t.tag_word }}
              </button>
            </li>
          </ul>
        </div>
      </div>

      <div
        v-if="tags && tags.length > 0"
        class="mt-2"
      >
        <div class="mt-1 block border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100">
          <ul class="flex flex-wrap gap-2">
            <li
              v-for="(t, index) in tags"
              :key="t"
              class="inline-block p-1 bg-gray-200 dark:bg-gray-600 rounded-md"
            >
              <!-- <span class="inline-block truncate">{{ t.tag_word }}</span> -->
              <button
                type="button"
                class="ml-2  text-gray-700 dark:text-gray-300"
                @click="selectTag(index, t)"
              >
                {{ t.tag_word }}
              </button>
            </li>
          </ul>
        </div>
      </div>
      <!-- Description -->
      <label
        for="description"
        class="block text-gray-700 dark:text-gray-300"
      >Description</label>
      <Field
        v-model="formData.description"
        name="description"
        rules="required"
        as="textarea"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Enter description"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="description"
      />
      <!-- Thumbnail Image URL -->
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
            v-model="formData.thumbnailImage"
            type="file"
            class="hidden"
            accept="image/*"
            name="thumnail"
          />
        </label>
      </div>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="thumnail"
      />
      <!-- for additional images -->

      <div class="flex items-center justify-center w-full">
        <label
          for="dropzone-file"
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
            <p class="mb-2 text-sm text-gray-500 dark:text-gray-400"><span class="font-semibold">Click to upload</span></p>
            <p class="text-xs text-gray-500 dark:text-gray-400">images</p>
          </div>
          <Field
            id="dropzone-file"
            v-model="formData.images"
            type="file"
            class="hidden"
            accept="image/*"
            multiple
            name="images"
          />
        </label>
      </div>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="images"
      />
      <div
        v-if="formData.images && formData.images.length > 0"
        class="mt-2"
      >
        <h3 class="block text-gray-700 dark:text-gray-300 text-ellipsis">
          Selected images:
        </h3>
        <ul class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none">
          <li
            v-for="(image, index) in formData.images"
            :key="index"
            class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          >
            <span class="inline-block w-full truncate">{{ image.name }}</span>
            <button
              type="button"
              class="text-red-500 ml-2 text-ellipsis"
              @click="removeSelectedImage(index)"
            >
              X
            </button>
          </li>
        </ul>
      </div>

      <!-- Submit Button -->
      <button
        type="submit"
        class="w-full bg-blue-500 rounded-md text-white py-2 hover:bg-blue-600 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700"
      >
        Create Event
      </button>
    </Form>
  </div>
</template>

<script setup>
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { useMutation } from '@vue/apollo-composable'
import { CREATE_EVENT } from '~/graphql/mutation'
import { GET_ORGANIZATIONS_FOR_EVENT, GET_TAGS, GET_CATAGORIES } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

const organizations = ref([])
const tags = ref([])
const selectedTag = ref([])
const catagoris = ref([])
const isInPerson = ref(true)

const { data, loading, error } = await apolloClient.query({ query: GET_ORGANIZATIONS_FOR_EVENT })
// const {data,loading,error} = useQuery(query)
if (!loading && !error) {
  organizations.value = data.data_organizations
  // console.log(data.data_organizations)
}
const { data: tag, loading: tag_loading, error: tag_error } = await apolloClient.query({ query: GET_TAGS })

if (!tag_loading && !tag_error) {
  tags.value = tag.data_tags
}

const { data: catagories_data, loading: catagoris_loading, error: catagoris_error } = await apolloClient.query({ query: GET_CATAGORIES })

if (!catagoris_loading && !catagoris_error) {
  catagoris.value = catagories_data.data_categories
}

const formData = ref({
  byOrganizationId: ref(null),
  title: ref(null),
  eventDate: ref(null),
  categoryId: ref(null),
  venue: ref(null),
  isOnline: ref(null),
  thumbnailImage: ref(null),
  images: ref([]),
  ticketPrice: ref(null),
  totalAvailableTickets: ref(null),
  region: ref(null),
  city: ref(null),
  street: ref(null),
  tags: ref([]),
  description: ref(null),

})

const setIsOnline = () => {
  isInPerson.value = formData.value.isOnline === 'false'
}
// Function to remove a selected image from the preview and formData
const removeSelectedImage = (index) => {
  formData.value.images.splice(index, 1)
}
const removeTag = (index, tag) => {
  formData.value.tags = formData.value.tags.filter((_, i) => i !== index)
  selectedTag.value = selectedTag.value.filter((_, i) => i !== index)
  tags.value = [...tags.value, tag]
}
const selectTag = (index, tag) => {
  formData.value.tags = [...formData.value.tags, tag.tag_id]
  selectedTag.value = [...selectedTag.value, tag]
  tags.value = tags.value.filter((_, i) => i !== index)
}

const handleSubmit = async () => {
  try {
    // Append thumbnail image to formdata
    const formdata = new FormData()
    if (formData.value.thumbnailImage) {
      formdata.append('thumbnailimage', formData.value.thumbnailImage)
    }

    // Append other images to formdata
    if (formData.value.images && formData.value.images.length > 0) {
      formData.value.images.forEach((file) => {
        formdata.append('other_images', file)
      })
    }
    // Upload files via fetch
    let response
    let imageObjects = {}
    let thumbnail_url
    if (formData.value.thumbnailImage || (formData.value.images && formData.value.images.length > 0)) {
      response = await fetch('http://localhost:4000/upload', {
        method: 'POST',
        body: formdata,
      })
      // Check response status and handle error
      if (!response.ok) {
        const errorMessage = await response.text()
        throw new Error(`image upload failed: ${errorMessage}`)
      }
      const { thumbnail_image_url, other_images_urls } = await response.json()
      thumbnail_url = thumbnail_image_url
      if (other_images_urls) {
        imageObjects = other_images_urls.map(url => ({
          image_url: url,
        }))
      }
    }
    // Handle success
    const eventtags = formData.value.tags.map(id => ({
      tag_word_id: id,
    }))

    const { mutate: inserEvent } = useMutation(CREATE_EVENT)
    const { data } = await inserEvent(
      {
        by_organization_id: formData.value.byOrganizationId,
        title: formData.value.title,
        ticket_price: formData.value.ticketPrice,
        total_available_tickets: formData.value.totalAvailableTickets,
        event_date: formData.value.eventDate,
        category_id: formData.value.categoryId,
        description: formData.value.description,
        venue: formData.value.venue,
        is_online: formData.value.isOnline,
        thumbnail_image_url: thumbnail_url,
        images: imageObjects,
        location: {
          longitude: -122.4194,
          latitude: 37.7749,
        },
        address: {
          street_name: formData.value.street,
          city_name: formData.value.city,
          region_name: formData.value.region,
        },
        tags: eventtags,
      },

    )
    const router = useRouter()
    router.push('/events/myevents')
    // console.log(data)
    // formData.value = {
    //   byOrganizationId: ref(null),
    //   title: ref(null),
    //   eventDate: ref(null),
    //   categoryId: ref(null),
    //   venue: ref(null),
    //   isOnline: ref(null),
    //   thumbnailImage: ref(null),
    //   images: ref([]),
    //   ticketPrice: ref(null),
    //   totalAvailableTickets: ref(null),
    //   region: ref(null),
    //   city: ref(null),
    //   street: ref(null),
    //   tags: ref([]),
    //   description: ref(null),

    // }
  }
  catch (err) {
    console.log('Upload error:', err)
  }
}

defineRule('required', (value) => {
  return true
  if (!value || !value.length) {
    return 'This field is required'
  }
  return true
})
</script>
