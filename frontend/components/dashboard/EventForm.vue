<template>
  <div class="event-form pb-20 max-w-xl mx-auto bg-white dark:bg-gray-800">
    <h1 class="text-2xl font-semibold mb-6 text-gray-900 dark:text-gray-100">
      Create Event
    </h1>
    <Form
      class="space-y-4"
      @submit="handleSubmit"
    >
      <!-- Title -->
      <Field
        v-slot="{ field }"
        name="title"
        rules="required"
      >
        <label
          for="title"
          class="block text-gray-700 dark:text-gray-300"
        >Title</label>
        <input
          type="text"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter event title"
        >
      </Field>

      <!-- Ticket Price -->
      <Field
        v-slot="{ field }"
        name="ticket_price"
        rules="required|decimal"
      >
        <label
          for="ticket_price"
          class="block text-gray-700 dark:text-gray-300"
        >Ticket Price</label>
        <input
          type="number"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter ticket price"
        >
      </Field>

      <!-- Total Available Tickets -->
      <Field
        v-slot="{ field }"
        name="total_available_tickets"
        rules="required|integer"
      >
        <label
          for="total_available_tickets"
          class="block text-gray-700 dark:text-gray-300"
        >Total Available Tickets</label>
        <input
          type="number"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Total tickets"
        >
      </Field>

      <!-- Event Date -->
      <Field
        v-slot="{ field }"
        name="event_date"
        rules="required|date"
      >
        <label
          for="event_date"
          class="block text-gray-700 dark:text-gray-300"
        >Event Date</label>
        <input
          type="date"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
        >
      </Field>

      <!-- Category ID -->
      <Field
        v-slot="{ field }"
        name="category_id"
        rules="required|integer"
      >
        <label
          for="category_id"
          class="block text-gray-700 dark:text-gray-300"
        >Category</label>
        <input
          type="number"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter category ID"
        >
      </Field>

      <!-- Description -->
      <Field
        v-slot="{ field }"
        name="description"
        rules="required|min:10"
      >
        <label
          for="description"
          class="block text-gray-700 dark:text-gray-300"
        >Description</label>
        <textarea
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter description"
        />
      </Field>

      <!-- Venue -->
      <Field
        v-slot="{ field }"
        name="venue"
        rules="required"
      >
        <label
          for="venue"
          class="block text-gray-700 dark:text-gray-300"
        >Venue</label>
        <input
          type="text"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter venue"
        >
      </Field>

      <!-- Organization ID -->
      <Field
        v-slot="{ field }"
        name="by_organization_id"
        rules="required|integer"
      >
        <label
          for="by_organization_id"
          class="block text-gray-700 dark:text-gray-300"
        >Organization ID</label>
        <input
          type="number"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter organization ID"
        >
      </Field>

      <!-- Event Type (Online or In-Person) -->
      <Field
        v-slot="{ field }"
        name="is_online"
        rules="required"
      >
        <label
          for="is_online"
          class="block text-gray-700 dark:text-gray-300"
        >Event Type</label>
        <select
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
        >
          <option value="true">
            Online
          </option>
          <option value="false">
            In-Person
          </option>
        </select>
      </Field>

      <!-- Thumbnail Image URL -->
      <div class="mb-4">
        <label
          for="thumbnail_image_url"
          class="block text-gray-700 dark:text-gray-300"
        >Thumbnail Image URL</label>
        <Field
          type="text"
          name="thumbnail_image_url"
          rules="required|url"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:outline-blue-500 focus:outline-none"
          placeholder="Enter image URL"
        />
      </div>
      <ErrorMessage
        class="text-red-500 mb-4"
        name="thumbnail_image_url"
      />

      <!-- Submit Button -->
      <button
        type="submit"
        class="w-full py-2 px-4 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-500 dark:bg-blue-700 dark:hover:bg-blue-600"
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

// Event form logic (submit handler, validation, etc.) goes here
const handleSubmit = (values) => {
  console.log('Form Submitted', values)
}
defineRule('required', (value) => {
  if (!value || !value.length) {
    return 'This field is required'
  }
  return true
})

const { mutate: inserEvent } = useMutation(CREATE_EVENT)

const k = true

if (k) {
  const { data } = await inserEvent(
    {
      by_organization_id: '6bde1a8b-146d-4a6c-bc1d-f86f879fc4c7',
      title: 'Tech Conference 2024',
      ticket_price: 100.00,
      total_available_tickets: 200,
      event_date: '2024-11-10',
      category_id: '45ce3080-9f3b-4c9a-88e3-00b464aa5de0',
      description: 'A conference for tech enthusiasts.',
      venue: '123 Event Venue, San Francisco',
      is_online: false,
      thumbnail_image_url: 'https://example.com/event-thumbnail.jpg',
      images: [
        {
          image_url: 'https://example.com/event-image-1.jpg',
        },
        {
          image_url: 'https://example.com/event-image-2.jpg',
        },
      ],
      location: {
        longitude: -122.4194,
        latitude: 37.7749,
      },
      address: {
        street_name: '123 Main St',
        city_name: 'San Francisco',
        region_name: 'CA',
      },
      // tags: [
      //   {
      //     tag_word_id: '1',
      //   },
      //   {
      //     tag_word_id: '2',
      //   },
      // ],
    },

  )
  console.log(data)
}

// const k = {
//   by_organization_id: '6bc1e4d3-95d8-4e1f-89bb-84349e8f10de',
//   category_id: '8a5c1b2c-098c-47b8-8fb6-dc2f4c76b5a2',
//   created_at: '2024-10-22T10:00:00Z',
//   description: 'This is a sample event description.',
//   event_date: '2024-11-05T15:00:00Z',
//   is_online: false,
//   thumbnail_image_url: 'https://example.com/event-thumbnail.jpg',
//   ticket_price: 50.00,
//   title: 'Tech Conference 2024',
//   total_available_tickets: 100,
//   venue: '123 Event Venue, San Francisco',
// }
</script>
