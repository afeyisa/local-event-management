<template>
  <div>
    <div class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-6">
      <!-- Loop through the events array and pass each event as a prop to EventCard -->
      <EventCard
        v-for="event in events"
        :key="event.id"
        :event="event"
      />
    </div>
  </div>
</template>

<script setup>
import EventCard from '~/components/dashboard/EventCard.vue'

const events = [
  {
    id: 1,
    title: 'Vue.js Conference 2024',
    description: 'A global Vue.js conference bringing together developers from around the world.',
    date: '2024-11-15',
    location: 'San Francisco, CA',
    image: 'https://via.placeholder.com/400x300.png?text=Vue+Conference',
  },
  {
    id: 2,
    title: 'Nuxt Meetup 2024',
    description: 'Join the Nuxt community for a day of talks, workshops, and networking.',
    date: '2024-12-05',
    location: 'New York, NY',
    image: 'https://via.placeholder.com/400x300.png?text=Nuxt+Meetup',
  },
  {
    id: 3,
    title: 'TailwindCSS Workshop',
    description: 'A hands-on workshop to master TailwindCSS and build beautiful UIs.',
    date: '2024-10-25',
    location: 'Online',
    image: 'https://via.placeholder.com/400x300.png?text=Tailwind+Workshop',
  },
]

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
</script>
