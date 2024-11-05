<script setup>
import { Field, Form, defineRule, ErrorMessage } from 'vee-validate'
import { useMutation } from '@vue/apollo-composable'
import { ref, onMounted } from 'vue'
import { CREATE_EVENT, DELETE_IMAGE, UPDATE_EVENT } from '~/graphql/mutation'
import { GET_ORGANIZATIONS_MY_ORG, GET_TAGS, GET_CATAGORIES, GET_EVENT_DETAILS, GET_MY_ID } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

const organizations = ref([])
const tags = ref([])
const selectedTag = ref([])
const catagoris = ref([])
const isInPerson = ref(true)
const currentImageIndex = ref(0)
const pevSelectedImages = ref(null)
const map = ref(null)

const latitude = ref(null)
const longitude = ref(null)
const markerPosition = ref(null)

function updateCoordinates(event) {
  const { lat, lng } = event.latlng

  // Update latitude and longitude
  latitude.value = lat
  longitude.value = lng
  console.log(lat, lng)

  // Remove the existing marker from the map
  if (markerPosition.value) {
    markerPosition.value.remove()
  }

  // Create a new marker at the updated coordinates
  markerPosition.value = L.marker([latitude.value, longitude.value])
    .addTo(map.value)
    .bindPopup(formData.value.title)
    .openPopup()
}

const { id } = defineProps({
  id: {
    type: String,
    required: false,
  },
})

const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
if (!loading && !error) {
  const myId = userData.data_users[0].user_id

  const { data, loading, error } = await apolloClient.query({ query: GET_ORGANIZATIONS_MY_ORG, variables: { where: { organizes: { organizer_id: { _eq: myId } } } } })
  if (!loading && !error) {
    organizations.value = data.data_organizations
  }
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
  // isOnline: ref(null),
  thumbnailImage: ref(null),
  images: ref([]),
  ticketPrice: ref(0),
  totalAvailableTickets: ref(0),
  region: ref(null),
  city: ref(null),
  street: ref(null),
  tags: ref([]),
  description: ref(null),

})
let thumbnail_url
onMounted(async () => {
  start()
  if (id) {
    try {
      const { data } = await apolloClient.query({
        query: GET_EVENT_DETAILS,
        variables: { id: id },
      })

      const eventDetails = data.data_events_by_pk
      if (eventDetails) {
        formData.value = {
          byOrganizationId: ref(eventDetails.by_organization_id),
          title: ref(eventDetails.title),
          eventDate: ref(eventDetails.event_date),
          categoryId: ref(eventDetails.category_id),
          venue: ref(eventDetails.venue),
          // isOnline: ref(eventDetails.is_online),
          ticketPrice: ref(eventDetails.ticket_price),
          totalAvailableTickets: ref(eventDetails.total_available_tickets),
          region: ref(eventDetails.address.region_name),
          city: ref(eventDetails.address.city_name),
          street: ref(eventDetails.address.region_name),
          description: ref(eventDetails.description),
        }
        if (eventDetails.thumbnail_image_url) {
          thumbnail_url = eventDetails.thumbnail_image_url
        }
        if (eventDetails.location && eventDetails.location.latitude && eventDetails.location.longitude) {
          latitude.value = eventDetails.location.latitude
          longitude.value = eventDetails.location.longitude
        }

        if (eventDetails.images && eventDetails.images.length > 0) {
          pevSelectedImages.value = [...eventDetails.images]
        }

        const tagWordIds = eventDetails.event_tags.map(tag => tag.tag_word_id)

        selectedTag.value = tag.data_tags.filter((_, i) => tagWordIds.includes(tag.data_tags[i].tag_id))
        tags.value = tag.data_tags.filter((_, i) => !tagWordIds.includes(tag.data_tags[i].tag_id))
        formData.value.tags = selectedTag.value.map(tag => tag.tag_id)
      }
    }
    catch (error) {
      console.error('Error fetching event details:', error)
    }
  }
})

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
const previousImage = () => {
  if (currentImageIndex.value > 0) {
    currentImageIndex.value--
  }
}

// Navigate Next Image
const nextImage = () => {
  if (currentImageIndex.value < pevSelectedImages.value.length - 1) {
    currentImageIndex.value++
  }
}

const handleDeleteImage = async () => {
  const { mutate } = useMutation(DELETE_IMAGE)
  await mutate({ event_id: id, image_url: pevSelectedImages.value[currentImageIndex.value].image_url })
  pevSelectedImages.value = pevSelectedImages.value.filter((_, i) => i !== currentImageIndex.value)
  currentImageIndex.value = 0
}

const handleSubmit = async () => {
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
  let imageObjects = []
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
    if (thumbnail_image_url) {
      thumbnail_url = thumbnail_image_url
    }
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

  if (id) {
    const variables = {
      id: id,
      category_id: formData.value.categoryId,
      description: formData.value.description,
      event_date: formData.value.eventDate,
      thumbnail_image_url: thumbnail_url,
      ticket_price: formData.value.ticketPrice,
      title: formData.value.title,
      total_available_tickets: formData.value.totalAvailableTickets,
      venue: formData.value.venue,
      tagged_event_id: id,

      address: {
        street_name: formData.value.street,
        city_name: formData.value.city,
        region_name: formData.value.region,
      },

      by_organization_id: formData.value.byOrganizationId,
      location: {
        longitude: longitude.value,
        latitude: latitude.value,
      },
      tags: formData.value.tags.map(tag_id => ({
        tag_word_id: tag_id,
        tagged_event_id: id,
      })),
      images: imageObjects
        ? imageObjects.map(image => ({
          event_id: id,
          image_url: image.image_url,
        }))
        : [],
    }
    try {
      const { mutate: updateEvent } = useMutation(UPDATE_EVENT)
      await updateEvent(variables)
      const router = useRouter()
      router.push(`/events/${id}`)
    }
    catch (error) {
      console.error('Error updating event:', error)
    }
  }
  else {
    try {
      const { mutate: inserEvent } = useMutation(CREATE_EVENT)
      const { data, loading } = await inserEvent(
        {
          by_organization_id: formData.value.byOrganizationId,
          title: formData.value.title,
          ticket_price: formData.value.ticketPrice,
          total_available_tickets: formData.value.totalAvailableTickets,
          event_date: formData.value.eventDate,
          category_id: formData.value.categoryId,
          description: formData.value.description,
          venue: formData.value.venue,
          thumbnail_image_url: thumbnail_url,
          images: imageObjects,
          location: {
            longitude: longitude.value,
            latitude: latitude.value,
          },
          address: {
            street_name: formData.value.street,
            city_name: formData.value.city,
            region_name: formData.value.region,
          },
          tags: eventtags,
        },

      )
      if (!loading) {
        const router = useRouter()
        router.push(`/events/${data.insert_data_events_one.event_id}`)
      }
    }
    catch (err) {
      console.log('Upload error:', err)
    }
  }
}

defineRule('required', (value) => {
  if (!value || !value.length) {
    return 'This field is required'
  }
  return true
})

defineRule('futuredate', (value) => {
  const today = new Date()
  const selectedDate = new Date(value)
  if (selectedDate < today) {
    return 'The event date and time must be in the future.'
  }

  return true
})

defineRule('greaterthanzero', (value) => {
  if (Number(value) < 0) {
    return 'this field must not be less than zero'
  }
  return true
})

const hasSelectedImages = computed(() => pevSelectedImages.value && pevSelectedImages.value.length > 0)

async function start() {
  const L = await import('leaflet')

  map.value = L.map('map').setView([9.0192, 38.7525], 13)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
  }).addTo(map.value)

  markerPosition.value = L.marker([9.0192, 38.7525])
    .addTo(map.value)
  map.value.on('click', updateCoordinates)
}
</script>

<template>
  <div class="event-form pb-20 max-w-xl mx-auto bg-white dark:bg-gray-800">
    <h1 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
      {{ id?"Update Event":"Create Event" }}
    </h1>
    <Form
      class="space-y-4"
      @submit="handleSubmit"
    >
      <h2 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
        About Event
      </h2>
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
          :value="org.organization_id"
          class="bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 p-2"
        >
          {{ org.organization_name }}
        </option>
      </Field>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="by_organization_id"
      />
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
      <label
        for="event_date"
        class="block text-gray-700 dark:text-gray-300"
      >Event Date</label>
      <Field
        v-model="formData.eventDate"
        name="event_date"
        rules="futuredate"
        type="date"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="event_date"
      />

      <label
        for="category_id"
        class="block text-gray-700 dark:text-gray-300"
      >Category</label>
      <Field
        v-model="formData.categoryId"
        name="category_id"
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

      <h2 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
        About Event Tickets
      </h2>
      <label
        for="ticket_price"
        class="block text-gray-700 dark:text-gray-300"
      >Ticket Price</label>
      <Field
        v-model="formData.ticketPrice"
        name="ticket_price"
        type="number"
        rules="greaterthanzero"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Enter ticket price"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="ticket_price"
      />

      <label
        for="total_available_tickets"
        class="block text-gray-700 dark:text-gray-300"
      >Total Available Tickets</label>
      <Field
        v-model="formData.totalAvailableTickets"
        name="total_available_tickets"
        type="number"
        rules="greaterthanzero"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Total tickets"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="total_available_tickets"
      />
      <div v-if="isInPerson">
        <h2 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
          About Event Event Location
        </h2>
        <div
          id="map"
          :lat-lng="markerPosition"
          class="map-container  w-full h-96 rounded-lg shadow-lg border border-gray-300 z-0"
        />

        <label
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
        <label
          for="city"
          class="block text-gray-700 dark:text-gray-300"
        >City</label>
        <Field
          v-model="formData.city"
          name="city"
          type="text"
          class="mt-1 block w-full mb-4 border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="City"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="city"
        />
        <label
          for="street"
          class="block text-gray-700 dark:text-gray-300"
        >Street</label>
        <Field
          v-model="formData.street"
          name="street"
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
          type="text"
          class="mt-1 block w-full border mb-4 border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter venue"
        />
        <ErrorMessage
          class="text-red-500 mb-4"
          name="venue"
        />
      </div>
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
      <label
        for="description"
        class="block text-gray-700 dark:text-gray-300"
      >Description</label>
      <Field
        v-model="formData.description"
        name="description"
        as="textarea"
        class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline focus:outline-1 focus:outline-blue-500 focus:outline-none"
        placeholder="Enter description"
      />
      <ErrorMessage
        class="text-red-500 mb-4"
        name="description"
      />
      <div
        v-if="id && hasSelectedImages"
        class="relative"
      >
        <button
          v-if="Array.isArray(pevSelectedImages)&&currentImageIndex > 0 && pevSelectedImages.length > 0"
          class="absolute left-0 top-1/2 transform -translate-y-1/2"
          @click="previousImage"
        >
          <i :class="(currentImageIndex > 0 && Array.isArray(pevSelectedImages) && pevSelectedImages.length > 0)?'fa fa-chevron-left dark:text-gray-500 text-gray-800 p-2 text-xl':''" />
        </button>

        <img
          :src="pevSelectedImages[currentImageIndex].image_url"
          alt="Event Image"
          class="w-full h-64 object-cover rounded-md mb-6"
        >

        <button
          v-if="currentImageIndex < pevSelectedImages.length - 1 && Array.isArray(pevSelectedImages) && pevSelectedImages.length > 0"
          class="absolute right-0 top-1/2 transform -translate-y-1/2"
          @click="nextImage"
        >
          <i :class="(currentImageIndex < pevSelectedImages.length - 1 && Array.isArray(pevSelectedImages) && pevSelectedImages.length > 0) ? 'fa fa-chevron-right dark:text-gray-500 text-gray-800 p-2 text-xl' : ''" />
        </button>
        <button

          type="button"
          class="text-red-500 ml-2 text-ellipsis"
          @click="handleDeleteImage"
        >
          Remove
        </button>
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
            <p class="text-xs text-gray-500 dark:text-gray-400">thumnail image</p>
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
        class="text-red-500 mb-4 w-fullrounded-md"

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

      <button
        type="submit"
        class="w-full bg-blue-500 rounded-md text-white py-2 hover:bg-blue-600 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700"
      >
        {{ id?"Udate Event":"Create Event" }}
      </button>
    </Form>
  </div>
</template>

<style scoped>
.map-container {
  position: relative;
}
</style>
