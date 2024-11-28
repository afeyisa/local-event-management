<script setup>
import { ref } from 'vue'
import { GET_CATAGORIES } from '~/graphql/querie/getCatagories.graphql'
import { apolloClient } from '~/plugins/apollo'
import { filteredWhereClause, filteredWhereClauseWithLocation } from '~/composables/filterEvents'
import { useEventStore } from '~/stores/events'
import { useIsSearching } from '~/stores/ui/isSearching'
import { useFilters } from '~/stores/filters'
import { useSearchInput } from '~/stores/searchInput'
import { useFilterToggler } from '~/stores/ui/filterToggler'
import { useOnSearch, useOnsearchLocation } from '~/composables/fetchEvents'
import { useisSearchOptionOpen } from '~/stores/ui/isSearchOptionOpen'
import { useSearchType } from '~/stores/ui/searchType'
import { uselocation } from '~/stores/location'
import { usehasMoreEventsOnLocation, usehasMoreEventsOnSearch } from '~/composables/useCheckHasMore'
import { usePagination } from '~/stores/pagination'

const location = uselocation()
const searchType = useSearchType()
const pagination = usePagination()
const searchOptionStore = useisSearchOptionOpen()
const filterToggler = useFilterToggler()
const search = useSearchInput()
const filters = useFilters()
const catagoris = ref([])
const isCatagoryToggleDown = ref(false)
const isPriceToggleDown = ref(false)
const isDateToggleDown = ref(false)
const eventStore = useEventStore()

const isSearchingStore = useIsSearching()
const map = ref(false)
const { data: catagories_data, loading: catagoris_loading, error: catagoris_error } = await apolloClient.query({ query: GET_CATAGORIES })

if (!catagoris_loading && !catagoris_error) {
  catagoris.value = catagories_data.data_categories
}

const displayMap = () => {
  map.value = !map.value
}

const fetchNext = async () => {
  pagination.next()
  if (searchType.searchType === 'location') {
    await usehasMoreEventsOnLocation(filteredWhereClauseWithLocation(filters.selectedDays, filters.selectedCategories, filters.isFree), location.lat, location.lon)
  }
  else {
    await usehasMoreEventsOnSearch(filteredWhereClause(filters.selectedDays, filters.selectedCategories, search.input, filters.isFree))
  }
  if (pagination.has_more) {
    filter()
  }
  else {
    pagination.pev()
  }
}

const fetchPrevious = async () => {
  pagination.pev()
  if (searchType.searchType === 'location') {
    await usehasMoreEventsOnLocation(filteredWhereClauseWithLocation(filters.selectedDays, filters.selectedCategories, filters.isFree), location.lat, location.lon)
  }
  else {
    await usehasMoreEventsOnSearch(filteredWhereClause(filters.selectedDays, filters.selectedCategories, search.input, filters.isFree))
  }
  if (pagination.offset >= 0) {
    filter()
  }
  else {
    pagination.next()
  }
}

const filter = () => {
  eventStore.setEvents(null)
  if (searchType.searchType === 'location' && location.lat && location.lon) {
    useOnsearchLocation(filteredWhereClauseWithLocation(filters.selectedDays, filters.selectedCategories, filters.isFree), location.lat, location.lon)
  }
  else {
    const whereClause = filteredWhereClause(filters.selectedDays, filters.selectedCategories, search.input, filters.isFree)
    useOnSearch(whereClause)
  }
}

// toggler functions
const toggleCatagories = () => {
  isCatagoryToggleDown.value = !isCatagoryToggleDown.value
}
const togglePrice = () => {
  isPriceToggleDown.value = !isPriceToggleDown.value
}
const toggleDate = () => {
  isDateToggleDown.value = !isDateToggleDown.value
}
</script>

<template>
  <div
    @click="searchOptionStore.restDropDown"
  >
    <div class="flex flex-col">
      <div class="flex-grow flex overflow-hidden">
        <div
          v-if="!isSearchingStore.isSearching"
          class="w-full p-4 overflow-y-auto"
        >
          <section
            class="text-center py-12 mb-12 bg-gradient-to-r from-blue-400 to-blue-700 text-white rounded-lg shadow-lg mt-1 w-full"
          >
            <h2 class=" sm:text-5xl text-3xl  font-extrabold mb-4">
              Welcome to events
            </h2>
            <p class="sm:text-xl text-2xl ">
              where you can browse events near by you, create and post events
            </p>
          </section>

          <h2 class="text-2xl text-center text-gray-700 dark:text-white font-bold mb-4">
            {{eventStore.isDefault? 'Recent Events':'events near by you'}}
          </h2>

          <div
            class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 gap-6"
          >
            <PublicEventCard
              v-for="event in eventStore.events"
              :key="event.id"
              :event="event"
            />
          </div>
        </div>
        <div
          v-if="isSearchingStore.isSearching"
          class="flex w-full text-xl min-h-screen  bg-gray-100 p-4 dark:bg-gray-700"
        >
          <div
            :class="!filterToggler.isFilterOpen? 'hidden md:block  rounded-sm w-72 bg-gray-100 z-10 p-4 overflow-y-auto text-xl min-h-screen dark:bg-gray-700 ':'fixed inset-0 z-10 bg-opacity-100 transition-opacityabsolute mt-20 rounded-md right-0 w-64 bg-white dark:bg-gray-800 shadow-lg p-6'"
          >
            <div class=" text-gray-700 dark:text-white font-semibold ">
              <h2 class="text-md sm:text-3xl font-bold mb-4 text-gray-700 dark:text-white">
                Filters
              </h2>
              <ul class=" text-md sm:text-2xl">
                <li class="mb-2">
                  <span
                    class="mb-2  hover:cursor-pointer text-gray-700 dark:text-white font-semibold"
                    @click="toggleCatagories"
                  > Category
                    <i :class="isCatagoryToggleDown?`text-sm fa fa-chevron-down text-gray-700 dark:text-white`:`text-gray-700 dark:text-white text-sm fa fa-chevron-right`" /></span>
                  <div
                    v-if="isCatagoryToggleDown"
                    class="ml-4"
                  >
                    <div
                      v-for="catagory in catagoris"
                      :key="catagory.catagory_id"
                      :value="catagory.catagory_id"

                      class="text-gray-700 dark:text-white font-semibold"
                    >
                      <label class="pl-2 hover:cursor-pointer">
                        <input
                          v-model="filters.selectedCategories"
                          type="checkbox"
                          :name="catagory.catagory_id"
                          :value="catagory.category_id"
                          @change="filter"
                        >
                        {{ catagory.category_name }}</label>
                    </div>
                  </div>
                </li>
                <li class="mb-2 ">
                  <div>
                    <span
                      class="mb-2  hover:cursor-pointer text-gray-700 dark:text-white font-semibold"
                      @click="toggleDate"
                    >Date
                      <i :class="isDateToggleDown?`text-sm fa fa-chevron-down text-gray-700 dark:text-white`:`text-gray-700 dark:text-white text-sm fa fa-chevron-right`" />
                    </span>
                    <div v-if="isDateToggleDown">
                      <div class="ml-4 ">
                        <label class="hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                          <input
                            v-model="filters.selectedDays"
                            value="today"
                            type="checkbox"
                            @change="filter"
                          > Today</label>
                      </div>
                      <div class="ml-4 ">
                        <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                          <input
                            v-model="filters.selectedDays"
                            value="tomorrow"
                            type="checkbox"
                            @change="filter"
                          > Tomorrow</label>
                      </div>
                      <div class="ml-4 ">
                        <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                          <input
                            v-model="filters.selectedDays"
                            value="thisWeekend"
                            type="checkbox"
                            @change="filter"
                          > This Weekend</label>
                      </div>
                      <div class="ml-4">
                        <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                          <input
                            v-model="filters.selectedDays"
                            value="nextWeek"
                            type="checkbox"
                            @change="filter"
                          > Next Week</label>
                      </div>
                      <div class="ml-4">
                        <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                          <input
                            v-model="filters.selectedDays"
                            value="nextMonth"
                            type="checkbox"
                            @change="filter"
                          > Next Month</label>
                      </div>
                    </div>
                  </div>
                </li>
                <li class="mb-2">
                  <div>
                    <span
                      class="mb-2 hover:cursor-pointer text-gray-700 dark:text-white font-semibold"
                      @click="togglePrice"
                    >Price
                      <i :class="isPriceToggleDown?`text-sm fa fa-chevron-down text-gray-700 dark:text-white`:`text-sm fa fa-chevron-right text-gray-700 dark:text-white`" />
                    </span>
                    <fieldset
                      v-if="isPriceToggleDown"
                    >
                      <div class="ml-4 hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                        <label class=" hover:cursor-pointer text-gray-700 dark:text-white"><input
                          v-model="filters.isFree"
                          type="radio"
                          value="true"
                          name="isFree"
                          @change="filter"
                        > Free</label>
                      </div>
                      <div class="ml-4 hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                        <label class=" hover:cursor-pointer text-gray-700 dark:text-white"><input
                          v-model="filters.isFree"
                          type="radio"
                          value="false"
                          name="isFree"
                          @change="filter"
                        > Paid</label>
                      </div>
                    </fieldset>
                  </div>
                </li>
              </ul>
              <h2
                class="text-md sm:text-3xl font-bold py-4 text-gray-700 hover:cursor-pointer dark:text-white"
                @click="displayMap"
              >
                map
              </h2>
            </div>
          </div>
          <div
            class="flex-grow p-4 overflow-y-auto"
            @click="filterToggler.isFilterOpen=false"
          >
            <h2 class="text-2xl font-bold mb-4 text-gray-700 dark:text-white">
              Search Results
            </h2>
            <div
              v-if="!map"
              class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 gap-10"
            >
              <PublicEventCard
                v-for="event in eventStore.events"
                :key="event.id"
                :event="event"
              />
            </div>
            <div
              v-else
              class="flex justify-center w-full "
            >
              <MapPicker
                :events=" eventStore.events"
                class="z-0"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="eventStore.events && eventStore.events.length>0"
      class="flex justify-center gap-20 items-center pb-8 max-w-full bg-gray-100 dark:bg-gray-700"
    >
      <button
        :disabled="!eventStore.events|| !(eventStore.events.length > 0)||!pagination.offset>0"
        class="px-4 py-2 bg-blue-500 text-white font-semibold rounded-md shadow-md hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed"
        @click="fetchPrevious"
      >
        Previous
      </button>

      <button
        :disabled="!eventStore.events|| !(eventStore.events.length > 0)||!pagination.has_more"
        class="px-4 py-2 bg-blue-500 text-white font-semibold rounded-md shadow-md hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed"
        @click="fetchNext"
      >
        Next
      </button>
    </div>
  </div>
</template>
