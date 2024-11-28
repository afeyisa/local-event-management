import { defineStore } from 'pinia'

export const useFilters = defineStore('filters', () => {
  const selectedCategories = ref([])
  const selectedDays = ref([])
  const isFree = ref('')
  return { selectedCategories, selectedDays, isFree }
})
