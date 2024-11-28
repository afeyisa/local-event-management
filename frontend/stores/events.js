import { defineStore } from 'pinia'

export const useEventStore = defineStore('events', () => {
  const events = ref([])
  const isDefault = ref(false)

  const setEvents = (e) => {
    events.value = e
  }
  return { isDefault, events, setEvents }
})
