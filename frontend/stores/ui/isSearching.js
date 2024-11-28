import { defineStore } from 'pinia'

export const useIsSearching = defineStore('isSearching', () => {
  const isSearching = ref(false)
  const setIsSearching = (s) => {
    isSearching.value = s
  }
  const resetIsearching = () => {
    isSearching.value = false
  }
  return { isSearching, setIsSearching, resetIsearching }
})
