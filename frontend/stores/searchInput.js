import { defineStore } from 'pinia'

export const useSearchInput = defineStore('searchInput', () => {
  const input = ref('')
  return { input }
})
