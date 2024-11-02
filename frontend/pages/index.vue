<!-- eslint-disable indent -->
<template>
  <div>
    <header class="bg-blue-500  sticky top-0 z-10 shadow-md text-white p-4 dark:bg-gray-800">
      <!-- Navigation Bar -->
      <nav class="container mx-auto flex justify-between items-center">
        <!-- Logo / Branding -->
        <div class="hidden md:block text-lg font-bold">
          <NuxtLink to="/">
            Event Manager
          </NuxtLink>
          <!-- <li class="flex items-center space-x-2 "> -->
        </div>

        <div class="w-full md:w-1/3 flex justify-center items-center">
          <!-- Search Bar -->
          <div class=" flex items-center   space-x-2 md:w-full">
            <button
              v-if="isSearching"
              class="text-white md:hidden"
              @click="toggleFilter"
            >
              <i class="fa fa-filter" />
            </button>

            <i class="fa fa-search   text-gray-400" />
            <input
              v-model="searchInput"
              type="text"
              placeholder="Search events..."
              autoFocus
              class="w-full px-2 py-2 rounded-lg border border-gray-300 text-gray-900 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:text-white"
              @keydown="keyListener"
            >
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
            class="absolute rounded-md right-0 w-64 bg-blue-500 dark:bg-gray-800 shadow-lg p-6"
          >
            <ul class="flex items-center space-x-4">
              <!-- Conditional Rendering based on user login status -->
              <li>
                <NuxtLink
                  to="/events"
                  class="hover:underline"
                >
                  <i class="fa fa-calendar " />
                  My Events
                </NuxtLink>
              </li>
              <li v-show="!isLoggedIn">
                <NuxtLink
                  to="/login"
                  class="hover:underline"
                >
                  <i class="fa fa-sign-in" />
                  Login
                </NuxtLink>
              </li>
              <li v-show="!isLoggedIn">
                <NuxtLink
                  to="/signup"
                  class="hover:underline"
                ><i class="fa fa-user-plus" />
                  Signup
                </NuxtLink>
              </li>
              <ThemeButton />
            </ul>
          </div>
        </div>
        <!-- Navigation Links -->
        <div class="hidden md:block">
          <div class=" relative  flex  items-center space-x-4 ">
            <ul class="flex items-center space-x-4">
              <!-- Conditional Rendering based on user login status -->
              <li>
                <NuxtLink
                  to="/events"
                  class="hover:underline"
                >
                  <i class="fa fa-calendar " />
                  My Events
                </NuxtLink>
              </li>
              <li v-show="!isLoggedIn">
                <NuxtLink
                  to="/login"
                  class="hover:underline"
                >
                  <i class="fa fa-sign-in" />
                  Login
                </NuxtLink>
              </li>
              <li v-show="!isLoggedIn">
                <NuxtLink
                  to="/signup"
                  class="hover:underline"
                ><i class="fa fa-user-plus" />
                  Signup
                </NuxtLink>
              </li>
            </ul>
            <ThemeButton />
          </div>
        </div>
      </nav>
    </header>

    <div class="flex flex-col">
      <!-- Content area: either default content or search results -->
      <div class="flex-grow flex overflow-hidden">
        <!-- Default content displayed if no search results yet -->
        <div
          v-if="!isSearching"
          class="w-full p-4 overflow-y-auto"
        >
          <h2 class="text-2xl font-bold mb-4">
            Featured Events
          </h2>

          <div class="grid grid-cols-2 md:grid-cols-2 lg:grid-cols-2 gap-6">
            <PublicEventCard
              v-for="event in events"
              :key="event.id"
              :event="event"
            />
          </div>
        </div>

        <!-- Search result layout when searching -->
        <div
          v-if="isSearching"
          class="flex w-full text-xl min-h-screen  bg-gray-100 p-4 dark:bg-gray-700 shadow-md"
        >
          <!-- Filters on the left (scrollable) -->
          <template v-if="!isOrgDetail && !isEventDetail">
            <div
              :class="!isFilterOpen? 'hidden md:block flex   w-72 bg-gray-100 p-4 overflow-y-auto text-xl min-h-screen dark:bg-gray-700 shadow-md':'fixed inset-0  bg-opacity-100 transition-opacityabsolute rounded-md right-0 w-64 bg-white dark:bg-gray-800 shadow-lg p-6'"
              @mousedown="initResize"
            >
              <div class=" text-gray-700 dark:text-white font-semibold  ">
                <h2 class="text-xl font-bold mb-4 text-gray-700 dark:text-white">
                  Filters
                </h2>
                <ul>
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
                            v-model="selectedCategories"
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
                              v-model="selecteDays"
                              value="today"
                              type="checkbox"
                              @change="filter"
                            > Today</label>
                        </div>
                        <div class="ml-4 ">
                          <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                            <input
                              v-model="selecteDays"
                              value="tomorrow"
                              type="checkbox"
                              @change="filter"
                            > Tomorrow</label>
                        </div>
                        <div class="ml-4 ">
                          <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                            <input
                              v-model="selecteDays"
                              value="thisWeekend"
                              type="checkbox"
                              @change="filter"
                            > This Weekend</label>
                        </div>
                        <div class="ml-4">
                          <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                            <input
                              v-model="selecteDays"
                              value="nextWeek"
                              type="checkbox"
                              @change="filter"
                            > Next Week</label>
                        </div>
                        <div class="ml-4">
                          <label class=" hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                            <input
                              v-model="selecteDays"
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
                            v-model="isFree"
                            type="radio"
                            value="true"
                            name="isFree"
                            @change="filter"
                          > Free</label>
                        </div>
                        <div class="ml-4 hover:cursor-pointer text-gray-700 dark:text-white font-semibold">
                          <label class=" hover:cursor-pointer text-gray-700 dark:text-white"><input
                            v-model="isFree"
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
                  class="text-xl font-bold py-4 text-gray-700 hover:cursor-pointer dark:text-white"
                  @click="displayMap"
                >
                  map
                </h2>
              </div>

              <!-- <button class="w-1 cursor-ew-resize bg-slate-200" /> -->
            </div>
            <div
              class="flex-grow p-4 overflow-y-auto"
              @click="isFilterOpen=false"
            >
              <h2 class="text-2xl font-bold mb-4 text-gray-700 dark:text-white">
                Search Results
              </h2>
              <div
                v-if="!map"
                class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 gap-6"
              >
                <PublicEventCard
                  v-for="event in events"
                  :key="event.id"
                  :event="event"
                  @eventid="handleEventId"
                  @orgid="handleOrgid"
                />
              </div>
              <div
                v-else
                class="flex justify-center w-full"
              >
                <MapPicker :events="events" />
              </div>
            </div>
          </template>
          <div
            v-else-if="isEventDetail"
            class=" w-full   bg-gray-100 p-4 dark:bg-gray-700   overflow-y-auto flex-shrink-0 text-xl min-h-screen"
          >
            <EventDetail
              :id="eventid"
              @goback="setIsevent"
            />
          </div>
          <div
            v-else-if="isOrgDetail"
            class=" w-full   bg-gray-100 p-4 dark:bg-gray-700   overflow-y-auto flex-shrink-0 text-xl min-h-screen"
          >
            <OrgDetail
              :id="orgid"
              @goback="setIsOrg"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { CHECK_AUTH_QUERY, GET_CATAGORIES, BROWSEEVENTS, PUBLIC_BROWSEEVENTS } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

const isLoggedIn = ref(false)
const searchInput = ref('')
const catagoris = ref([])
const selectedCategories = ref([])
const selecteDays = ref([])
const isCatagoryToggleDown = ref(false)
const isPriceToggleDown = ref(false)
const isDateToggleDown = ref(false)
const isFree = ref('')
const events = ref([])
const isSearching = ref(false)
const map = ref(false)
const isEventDetail = ref(false)
const isOrgDetail = ref(false)
const orgid = ref(null)
const eventid = ref(null)
const isSidebarOpen = ref(false)
const isFilterOpen = ref(false)
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}
const toggleFilter = () => {
  isFilterOpen.value = !isFilterOpen.value
}
// check the state of user
const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
if (data && data.isAuthenticated) {
  isLoggedIn.value = true
}
const { data: catagories_data, loading: catagoris_loading, error: catagoris_error } = await apolloClient.query({ query: GET_CATAGORIES })

if (!catagoris_loading && !catagoris_error) {
  catagoris.value = catagories_data.data_categories
}

const handleEventId = (id) => {
  eventid.value = id
  isEventDetail.value = true
  isOrgDetail.value = false
}

const handleOrgid = (id) => {
  orgid.value = id
  isOrgDetail.value = true
  isEventDetail.value = false
}
const displayMap = () => {
  map.value = !map.value
}

const onSearch = async (whereClause) => {
  isEventDetail.value = false
  isOrgDetail.value = false
  if (isLoggedIn.value) {
    try {
      const { data } = await apolloClient.query({
        query: BROWSEEVENTS,
        variables: { where: whereClause },
      })
      isSearching.value = true
      events.value = data.data_events
    // console.log(events.value)
    }
    catch (err) {
      console.log(err)
    /** empty */
    }
  }
  else {
    try {
      const { data } = await apolloClient.query({
        query: PUBLIC_BROWSEEVENTS,
        variables: { where: whereClause },
      })
      isSearching.value = true
      events.value = data.data_events
    // console.log(events.value)
    }
    catch (err) {
      console.log(err)
    /** empty */
    }
  }
}

const keyListener = (e) => {
  if (e.key === 'Enter' && searchInput.value.length > 0) {
    // console.log(searchInput.value)
    const searchTerm = searchInput.value
    const keywords = searchTerm.split(' ').map(word => `%${word}%`)

    // Construct the `where` clause
    const whereClause = {
      _or: [
        { _or: keywords.map(word => ({ title: { _ilike: word } })) },
        { _or: keywords.map(word => ({ description: { _ilike: word } })) },
        { _or: keywords.map(word => ({ _and: { event_tags: { _and: { tag: { _and: { tag_word: { _ilike: word } } } } } } })) },
        { _or: keywords.map(word => ({ category: { category_name: { _ilike: word } } })) },
      ] }
    events.value = null
    selectedCategories.value = []
    selecteDays.value = []
    isFree.value = ''
    onSearch(whereClause)
  }
}

const filter = () => {
  let filtereddDays = [{ event_date: { } }]
  const today = new Date()

  if (selecteDays.value.length > 0) {
    filtereddDays = selecteDays.value.map((day) => {
      let startDate, endDate
      const currentDay = today.getDay()
      switch (day) {
      /* eslint-disable */
      case 'today':
        startDate = new Date(today)
        endDate = new Date(today)
        break
      case 'tomorrow':
        startDate = new Date(today)
        startDate.setDate(today.getDate() + 1)
        endDate = new Date(startDate)
        break

      case 'thisWeekend':
        startDate = new Date(today)
        startDate.setDate(today.getDate() + (5 - currentDay)) // Friday
        endDate = new Date(startDate)
        endDate.setDate(startDate.getDate() + 2) // Sunday
        break

      case 'nextWeek':
        startDate = new Date(today)
        startDate.setDate(today.getDate() + (8 - currentDay)) // Next Monday
        endDate = new Date(startDate)
        endDate.setDate(startDate.getDate() + 6) // Following Sunday
        break

      case 'nextMonth':
        startDate = new Date(today)
        endDate = new Date(today)
        endDate.setDate(today.getDate() + 30)
        break

      default:
        startDate = new Date(today)
        endDate = new Date(today)
      }

      return {
        _and: [
          { event_date: { _gte: startDate } },
          { event_date: { _lte: endDate } },
        ],
      }
    })
  }
         

  let filteredCatagoryIds = [{ category_id: { } }]
  if (selectedCategories.value.length > 0) {
    filteredCatagoryIds = selectedCategories.value.map(id => ({ category_id: { _eq: id } }))
  }
  let ticketPrice = {}
  if (isFree.value === 'true') {
    ticketPrice = { _eq: 0 }
  }
  if (isFree.value === 'false') {
    ticketPrice = { _gt: 0 }
  }
  const searchTerm = searchInput.value
  const keywords = searchTerm.split(' ').map(word => `%${word}%`)

  // Construct the `where` clause
  const whereClause = {
    _and: [
      {
        _or: [
          { _or: keywords.map(word => ({ title: { _ilike: word } })) },
          { _or: keywords.map(word => ({ description: { _ilike: word } })) },
          { _or: keywords.map(word => ({ _and: { event_tags: { _and: { tag: { _and: { tag_word: { _ilike: word } } } } } } })) },
        ] },
      {
        ticket_price: ticketPrice,
      },
      {
        _or: filteredCatagoryIds,
      },
      {
        _or: filtereddDays,
      },
    ],
  }
  events.value = null
  onSearch(whereClause)
}

const setIsevent = ()=>{
  isEventDetail.value = false
}
const setIsOrg = ()=>{
  isOrgDetail.value = false
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

// resizers
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
</script>
