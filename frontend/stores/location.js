import { defineStore } from 'pinia'

export const uselocation = defineStore('location', () => {
  const lat = ref(null)
  const lon = ref(null)

  return { lat, lon }
})
