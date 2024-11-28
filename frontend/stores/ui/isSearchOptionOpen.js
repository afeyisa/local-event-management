import { defineStore } from 'pinia'

export const useisSearchOptionOpen = defineStore('issearchoptionopen', () => {
  const dropdownOpen = ref(false)

  function toggleDropdown() {
    dropdownOpen.value = !dropdownOpen.value
  }
  function restDropDown() {
    dropdownOpen.value = false
  }
  return { dropdownOpen, toggleDropdown, restDropDown }
})
