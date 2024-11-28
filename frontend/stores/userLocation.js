import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserLocationStore = defineStore('userlocation', () => {
  const lat = ref(null)
  const lon = ref(null)
  const error = ref(null)
  const loading = ref(false)

  const fetchedAt = ref(null)
  const cacheDuration = 60 * 60 * 1000 // 1 hour

  const fetchLocation = async () => {
    if (lat.value && lon.value && (Date.now() - fetchedAt.value < cacheDuration)) {
      return
    }

    loading.value = true
    error.value = null

    try {
      const response = await fetch('/api/ipinfo')
      const data = await response.json()
      const [latitude, longitude] = data.loc.split(',').map(parseFloat)
      lat.value = latitude
      lon.value = longitude
      fetchedAt.value = Date.now()
    }
    catch (err) {
      error.value = 'Failed to fetch location'
    }
    finally {
      loading.value = false
    }
  }

  const getLocation = () => ({
    lat: lat.value,
    lon: lon.value,
  })

  const isLoading = () => loading.value
  const isError = () => !error.value

  return {
    lat,
    lon,
    error,
    loading,
    fetchLocation,
    getLocation,
    isLoading,
    isError,
  }
})
