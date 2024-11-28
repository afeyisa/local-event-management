import { defineStore } from 'pinia'

export const useSearchType = defineStore('searchtype', () => {
  const searchType = ref('event')

  return { searchType }
})
