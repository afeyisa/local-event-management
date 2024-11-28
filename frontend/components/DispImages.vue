<script setup>
import { fetchBase64Image } from '~/composables/fetchImage'

const props = defineProps({
  images: {
    type: Array,
    required: true,
  },
})
const base64Images = ref([])
const currentImageIndex = ref(0)
const previousImage = () => {
  if (currentImageIndex.value > 0) {
    currentImageIndex.value--
  }
}

// Navigate Next Image
const nextImage = () => {
  if (currentImageIndex.value < props.images.length - 1) {
    currentImageIndex.value++
  }
}

const prefetchImages = async () => {
  base64Images.value = await Promise.all(
    props.images.map(img => fetchBase64Image(img.image_url)),
  )
}

watch(props.images, prefetchImages, { immediate: true })

onMounted(() => {
  prefetchImages()
})
</script>

<template>
  <div
    class="relative"
  >
    <button
      v-if="currentImageIndex > 0 && Array.isArray(images) && images.length > 0"
      class="absolute left-0 top-1/2 transform -translate-y-1/2"
      @click="previousImage"
    >
      <i :class="(currentImageIndex > 0 && Array.isArray(images) && images.length > 0)?'fa fa-chevron-left dark:text-gray-500 text-gray-800 p-2 text-xl':''" />
    </button>

    <img
      v-if="Array.isArray(images) && images.length > 0 && images[currentImageIndex]?.image_url"
      :src="base64Images[currentImageIndex]"
      alt="Event Image"
      class="w-full h-64 object-cover rounded-md mb-6"
    >

    <button
      v-if="currentImageIndex < images.length - 1 && Array.isArray(images) && images.length > 0"
      class="absolute right-0 top-1/2 transform -translate-y-1/2"
      @click="nextImage"
    >
      <i :class="(currentImageIndex < images.length - 1 && Array.isArray(images) && images.length > 0) ? 'fa fa-chevron-right dark:text-gray-500 text-gray-800 p-2 text-xl' : ''" />
    </button>
  </div>
</template>
