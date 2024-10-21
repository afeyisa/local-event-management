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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
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
      <Field
        v-slot="{ field }"
        name="thumbnail_image_url"
        rules="required|url"
      >
        <label
          for="thumbnail_image_url"
          class="block text-gray-700 dark:text-gray-300"
        >Thumbnail Image URL</label>
        <input
          type="text"
          v-bind="field"
          class="mt-1 block w-full border border-gray-300 dark:border-gray-600 p-2 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
          placeholder="Enter image URL"
        >
      </Field>

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
import { Field, Form } from 'vee-validate'

// Event form logic (submit handler, validation, etc.) goes here
const handleSubmit = (values) => {
  console.log('Form Submitted', values)
}
</script>
