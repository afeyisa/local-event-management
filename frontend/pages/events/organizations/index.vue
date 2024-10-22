<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols gap-6">
    <p v-if="error">Something went wrong...</p>
    <p v-if="loading">Loading...</p>
    <OrganizationCard
      v-else
      v-for="organization in organizations"
      :key="organization.organization_name"
      :organization="organization"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
// import { useQuery } from '@vue/apollo-composable'
import OrganizationCard from '~/components/dashboard/OrganizationCard.vue'
import { GET_ORGANIZATIONS } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'
import { useQuery } from '@vue/apollo-composable'


definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
const organizations = ref([
  {
    organization_name: 'Tech Innovators Inc.',
    bio: 'Leading the world in cutting-edge technology solutions.',
    created_at: '2023-04-10T10:15:30Z',
    description: 'We specialize in developing innovative tech solutions for businesses around the globe.',
    profile_photo: 'https://via.placeholder.com/150?text=Tech+Innovators',
  },
  {
    organization_name: 'Green Earth Initiative',
    bio: 'Working towards a sustainable future.',
    created_at: '2022-06-25T08:45:10Z',
    description: 'Our organization focuses on environmental sustainability projects worldwide.',
    profile_photo: 'https://via.placeholder.com/150?text=Green+Earth',
  },
  {
    organization_name: 'Healthcare Heroes',
    bio: '',
    created_at: '2021-09-12T14:30:00Z',
    description: 'A non-profit organization aimed at providing healthcare services to underserved communities.',
    profile_photo: 'https://via.placeholder.com/150?text=Healthcare+Heroes',
  },
  {
    organization_name: 'Art & Culture Foundation',
    bio: 'Promoting art and culture globally.',
    created_at: '2020-12-20T12:00:00Z',
    description: 'We support and sponsor events, exhibitions, and workshops to promote cultural diversity and appreciation of the arts.',
    profile_photo: 'https://via.placeholder.com/150?text=Art+Culture',
  },
  {
    organization_name: 'Tech Innovators Inc.',
    bio: 'Leading the world in cutting-edge technology solutions.',
    created_at: '2023-04-10T10:15:30Z',
    description: 'We specialize in developing innovative tech solutions for businesses around the globe.',
    profile_photo: 'https://via.placeholder.com/150?text=Tech+Innovators',
  },
  {
    organization_name: 'Green Earth Initiative',
    bio: 'Working towards a sustainable future.',
    created_at: '2022-06-25T08:45:10Z',
    description: 'Our organization focuses on environmental sustainability projects worldwide.',
    profile_photo: 'https://via.placeholder.com/150?text=Green+Earth',
  },
  {
    organization_name: 'Healthcare Heroes',
    bio: '',
    created_at: '2021-09-12T14:30:00Z',
    description: 'A non-profit organization aimed at providing healthcare services to underserved communities.',
    profile_photo: 'https://via.placeholder.com/150?text=Healthcare+Heroes',
  },
  {
    organization_name: 'Art & Culture Foundation',
    bio: 'Promoting art and culture globally.',
    created_at: '2020-12-20T12:00:00Z',
    description: 'We support and sponsor events, exhibitions, and workshops to promote cultural diversity and appreciation of the arts.',
    profile_photo: 'https://via.placeholder.com/150?text=Art+Culture',
  },
])

const { data, loading, error } = await apolloClient.query({ query: GET_ORGANIZATIONS })
// const {data,loading,error} = useQuery(query)
if (!loading && !error) {
  organizations.value = data.data_organizations
  console.log(data.data_organizations)
}
</script>

<!--
<script setup>
import { ref } from 'vue'
import EventCard from '~/components/dashboard/EventCard.vue'

definePageMeta({
  layout: 'mydashboard',
  middleware: 'auth',
})
// Example event data
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

const defaultEvents = [
  { id: 2, title: 'Rock Concert 2024', description: 'The best rock bands performing live this summer.' },
]

// Dummy search results to display after search
const searchResults = ref([
  { id: 1, title: 'Jazz Night in New York', description: 'An evening of live jazz in a cozy venue.' },
])

// Resizing behavior for the filters panel
const initResize = (event) => {
  const div = event.target.parentElement
  const startX = event.clientX
  const startWidth = div.offsetWidth

  const doDrag = (e) => {
    div.style.width = `${startWidth + e.clientX - startX}px`
  }

  const stopDrag = () => {
    document.removeEventListener('mousemove', doDrag)
    document.removeEventListener('mouseup', stopDrag)
  }

  document.addEventListener('mousemove', doDrag)
  document.addEventListener('mouseup', stopDrag)
}

// Search state
const isSearching = ref(false)

// Simulate a search action
function onSearch() {
  isSearching.value = true // Display search results after search
}

// Simulate initial search call
// onSearch()
</script> -->
