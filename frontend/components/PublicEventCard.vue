<script setup>
import { formatDate } from '~/composables/formatDate'
import { useBookMark } from '~/composables/useBookMark'
import { fetchBase64Image } from '~/composables/fetchImage'
import { useTicket } from '~/composables/useTicket'

const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
})
const e = ref(props.event)
const base64Image = ref(null)
const prefetchImages = async () => {
  base64Image.value = await fetchBase64Image(e.value.thumbnail_image_url)
}
// :to="`/${event.event_id}/ticket`"

const handleTicket = async () => {
  if (!(e.value.ticket_price > 0)) {
    const { buyFreeTicket } = useTicket()
    await buyFreeTicket(e.value.event_id)
  }
  else {
    const router = useRouter()
    router.push(`/${e.value.event_id}/ticket`)
  }
}

const { booker, bookmarks } = await useBookMark(e.value.bookmarks ? e.value.bookmarks : [], e.value.event_id)
onMounted(() => {
  prefetchImages()
})
</script>

<template>
  <div
    class="e-card min-w-72 max-h-fit rounded overflow-hidden shadow-lg bg-white dark:bg-gray-900"
  >
    <div class="grid  sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-1 mb-2">
      <img
        :src="base64Image || 'https://via.placeholder.com/400x300.png?text=Vue+Conference'"
        alt="Event Image"
        class="w-full h-52 rounded-sm  object-cover"
      >
      <div class="px-6 pt-4  pb-2 h-52">
        <div class=" text-xl sm:text-3xl font-bold text-gray-600 dark:text-gray-300 ">
          {{ e.title }}
        </div>
        <div class="text-md sm:text-xl text-gray-600 dark:text-gray-400">
          <i class="fa fa-tags p-2" /><strong>{{ event.category?.category_name }}</strong>
        </div>

        <div class="flex items-center text-md sm:text-xl text-gray-600 dark:text-gray-400">
          <i class=" p-2 fa fa-calendar" />
          {{ formatDate(e.event_date) }}
        </div>

        <div class="flex items-center text-md text-md sm:text-xl text-gray-600 dark:text-gray-400 mt-2">
          <i class="p-2 fa fa-map-marker" />
          {{ e.address?.street_name || '' }} {{ e.address?.city_name || '' }} {{ e.address?.region_name || '' }}
        </div>
        <div class=" text-gray-600  text-md sm:text-xl dark:text-gray-400 ">
          <i class="p-2 fa fa-ticket" />
          Price
          {{ e.ticket_price===0?'free':e.ticket_price +' ETB' }}
        </div>
        <div class=" text-gray-600 text-md sm:text-xl dark:text-gray-400 ">
          <i class="p-2 fa fa-ticket" />
          Available Tickets
          {{ e.total_available_tickets }}
        </div>
      </div>
    </div>
    <div class="px-6" />
    <div class=" items-baseline text-gray-600 dark:text-gray-400 ">
      <div
        v-if="e.organization&&e.organization.organization_name"
        class="px- py w-full"
      >
        <NuxtLink
          class="w-full  dark:text-gray-200 px-1 text-gray-700"
          :to="`/org/${event.organization.organization_id}`"
        >
          <PublicOrganizationCard
            class="w-full"
            :organization="event.organization"
          />
        </NuxtLink>
      </div>
      <p
        v-else
        class="ml-2 text-sm"
      >
        {{ e.organization?e.organization.organization_name:'unknown' }}
      </p>
    </div>

    <div class="px-6 py-1 flex justify-between space-x-4">
      <div class="px-6 py-1 flex space-x-4">
        <button>
          <NuxtLink
            :key="event.event_id"
            class="inline-block pb-2 text-blue-500 hover:text-blue-700 dark:hover:text-blue-500"
            :to="`/${event.event_id}`"
          >
            Event Details
          </NuxtLink>
        </button>

        <button
          class="text-green-500 hover:text-green-700 pb-2"
          @click="handleTicket"
        >
          <i class="fa fa-ticket" />
          <!-- <NuxtLink
            :to="`/${event.event_id}/ticket`"
          > -->
          Buy Tickets
          <!-- </NuxtLink> -->
        </button>
      </div>
      <button
        v-if="bookmarks"
        class="text-yellow-500 hover:text-yellow-700 pb-2"
        @click="booker"
      >
        <i
          :class="bookmarks.length>0?'fa fa-bookmark':'fa fa-bookmark-o '"
        />
      </button>
    </div>
  </div>
</template>
