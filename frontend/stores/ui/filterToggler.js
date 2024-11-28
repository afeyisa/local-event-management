import { defineStore } from 'pinia'

export const useFilterToggler = defineStore('filtertoggler', () => {
  const isFilterOpen = ref(false)
  const toggleFilter = () => {
    isFilterOpen.value = !isFilterOpen.value
  }
  const resetToClosedFilter = () => {
    isFilterOpen.value = false
  }
  return { isFilterOpen, toggleFilter, resetToClosedFilter }
})
