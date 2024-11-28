import { defineStore } from 'pinia'

export const usePagination = defineStore('pagination', () => {
  const limit = ref(2)
  const offset = ref(0)
  const has_more = ref(true)
  const has_pev = ref(false)
  const next = () => {
    if (has_more.value) {
      offset.value = offset.value + limit.value
    }
  }
  const pev = () => {
    offset.value = offset.value - limit.value
  }

  const reset = () => {
    has_more.value = true
    has_pev.value = false
    offset.value = 0
  }
  return { limit, offset, next, pev, has_more, has_pev, reset }
})
