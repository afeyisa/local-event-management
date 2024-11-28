<script setup>
import { ref } from 'vue'
import { constructWhereClause } from '~/composables/filterEvents'
import { useEventStore } from '~/stores/events'
import { useIsSearching } from '~/stores/ui/isSearching'
import { useSearchInput } from '~/stores/searchInput'
import { useFilters } from '~/stores/filters'
import { useOnSearch, useOnsearchLocation, defaultEventHomeEvents } from '~/composables/fetchEvents'
import { useFilterToggler } from '~/stores/ui/filterToggler'
import { useUserLocationStore } from '~/stores/userLocation'
import { useisSearchOptionOpen } from '~/stores/ui/isSearchOptionOpen'
import { useSearchType } from '~/stores/ui/searchType'
import { uselocation } from '~/stores/location'
import { usePagination } from '~/stores/pagination'

const pagination = usePagination()
const searchType = useSearchType()
const location = uselocation()
const searchOptionStore = useisSearchOptionOpen()
const eventStore = useEventStore()
const locationStore = useUserLocationStore()
const filterToggler = useFilterToggler()
const filters = useFilters()
const search = useSearchInput()
const eveventStore = useEventStore()
const isSearchSore = useIsSearching()
const isSidebarOpen = ref(false)
const searchOptions = [
  { label: 'Title', value: 'event' },
  { label: 'Location', value: 'location' },
]

const reset = () => {
  isSearchSore.resetIsearching()
  eveventStore.events = []
  filterToggler.resetToClosedFilter()
  fetchintialContent()
}

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

const router = useRouter()
const onSearch = async (whereClause) => {
  if (router.currentRoute.value.path !== '/') {
    router.push('/')
  }
  isSearchSore.setIsSearching(true)
  pagination.reset()
  if (searchType.searchType === 'location') {
    const { lat, lon } = await getCoordinatesWithGeoapify(search.input)
    location.lat = lat
    location.lon = lon
    await useOnsearchLocation(whereClause, lat, lon)
  }
  else {
    useOnSearch(whereClause)
  }// it better if we make this function return error for notification
}

const keyListener = async (e) => {
  if (e.key === 'Enter' && search.input.length > 0) {
    eveventStore.setEvents(null)
    filters.selectedCategories = []
    filters.selectedDays = []
    filters.isFree = ''
    const searchTerm = search.input
    const keywords = searchTerm.split(' ').map(word => `%${word}%`)
    const whereClause = searchType.searchType === 'location' ? {} : constructWhereClause(keywords)
    onSearch(whereClause)
  }
}

async function getCoordinatesWithGeoapify(address) {
  try {
    const data = await $fetch('/api/geocode', { params: { address } })
    if (data.features && data.features.length > 0) {
      const { lat, lon } = data.features[0].properties
      return { lat, lon }
    }
    else {
      console.log('No results found.')
      return null
    }
  }
  catch (error) {
    console.error('Failed to fetch coordinates:', error)
    return null
  }
}

const fetchintialContent = async () => {
  if (eventStore.events === null || eventStore.events.length === 0) {
    if (locationStore.lat && locationStore.lon) {
      useOnsearchLocation({}, locationStore.lat, locationStore.lon)
    }
    else {
      defaultEventHomeEvents()
      eveventStore.isDefault = true
    }
  }
}

onMounted(async () => {
  await locationStore.fetchLocation()
  await fetchintialContent()
})

function selectOption(value) {
  searchType.searchType = value
  searchOptionStore.dropdownOpen = false
  console.log('Selected Search Type:', searchType.searchType)
}
</script>

<template>
  <header class="bg-blue-500  sticky top-0 z-20 shadow-md text-white p-4 dark:bg-gray-800">
    <nav class="container mx-auto flex justify-between items-center">
      <div class="hidden md:block text-xl font-bold">
        <NuxtLink
          to="/"
          @click="reset"
        >
          Event Manager
        </NuxtLink>
      </div>

      <div class="w-full md:w-1/3 flex justify-center items-center">
        <div class=" flex items-center   space-x-2 md:w-full">
          <button
            v-if="isSearchSore.isSearching && router.currentRoute.value.path === '/'"
            class="text-white md:hidden"
            @click="filterToggler.toggleFilter"
          >
            <i class="fa fa-filter" />
          </button>

          <i class="fa fa-search   text-gray-400" />
          <input
            v-model="search.input"
            type="text"
            placeholder="Search events..."
            autoFocus
            class="w-full px-2 py-2 text-sm sm:text-lg rounded-lg border border-gray-300 text-gray-900 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:text-white"
            @keydown="keyListener"
          >
          <div class="relative">
            <i
              :class="searchOptionStore.dropdownOpen
                ? 'text-sm fa fa-chevron-down text-gray-700 hover:cursor-pointer dark:text-white'
                : 'text-gray-700 hover:cursor-pointer dark:text-white text-sm fa fa-chevron-right'"
              @click="searchOptionStore.toggleDropdown"
            />

            <div
              v-if="searchOptionStore.dropdownOpen"
              class="absolute left-0 mt-2 m-w-48 bg-blue-500 border border-gray-200 rounded-lg shadow-lg dark:bg-gray-800 dark:border-gray-700 z-50"
            >
              <ul class="py-1">
                <li
                  v-for="option in searchOptions"
                  :key="option.value"
                  class="px-4 py-2 hover:bg-gray-100  hover:text-gray-900 dark:hover:bg-gray-600 dark:hover:text-gray-100 cursor-pointer"
                  @click="selectOption(option.value)"
                >
                  {{ option.label }}
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <button
        class="ml-2 text-white focus:outline-none md:hidden"
        @click="toggleSidebar"
      >
        <i class="fa fa-bars" />
      </button>

      <div
        :class="['fixed inset-0 rounded-sm bg-gray-800 bg-opacity-75 transition-opacity', { hidden: !isSidebarOpen }]"
        @click="toggleSidebar"
      >
        <div
          class="absolute rounded-md right-0 w-fit bg-blue-500 flex justify-end items-center dark:bg-gray-800 shadow-lg p-6"
        >
          <PublicHeader />
        </div>
      </div>
      <div class="hidden md:block justify-end">
        <div class=" relative  flex justify-end sm:items-center space-x-4 ">
          <PublicHeader />
        </div>
      </div>
    </nav>
  </header>
</template>
