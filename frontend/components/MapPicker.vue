<script setup>
import { ref, onMounted, watch } from 'vue'

const props = defineProps({
  events: {
    type: Array,
    required: false,
  },
})

const map = ref(null)
const latitude = ref(9.0192)
const longitude = ref(38.7525)

async function start() {
  const L = await import('leaflet') // Ensure Leaflet loads only on the client side

  // Check if the map has already been initialized to avoid re-initialization
  if (!map.value) {
    if (props.events.length === 1 && props.events[0].location) {
      map.value = L.map('map').setView([props.events[0].location.latitude, props.events[0].location.longitude], 10) // Initial coordinates and zoom
    }
    else {
      map.value = L.map('map').setView([latitude.value, longitude.value], 5) // Initial coordinates and zoom
    }

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; <a href="https://osm.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(map.value)
  }

  // Clear existing markers
  map.value.eachLayer((layer) => {
    if (layer instanceof L.Marker) {
      map.value.removeLayer(layer)
    }
  })

  // Add markers for each event
  if (props.events) {
    props.events.forEach((e) => {
      if (e.location && e.location.latitude && e.location.longitude) {
        L.marker([e.location.latitude, e.location.longitude])
          .addTo(map.value)
          .bindPopup(e.title)
      }
    })
  }
}

// Watch for changes in the events prop
watch(() => props.events, () => {
  start()
})

// Call start when the component is mounted
onMounted(() => {
  start()
})
</script>

<template>
  <div
    id="map"
    class="map-container w-full h-96 rounded-lg shadow-lg border border-gray-300 z-0 "
  />
</template>

<style scoped>
.map-container {
  position: relative;
}
</style>
