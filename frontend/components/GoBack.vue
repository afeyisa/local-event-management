<template>
  <div>
    <button
      class="text-yellow-500 hover:text-yellow-700"
      @click="goBack"
    >
      <i class="fa fa-chevron-left" /> Go Back
    </button>
  </div>
</template>

<script setup>
const router = useRouter()
function goBack() {
  router.go(-1)
}
</script>
