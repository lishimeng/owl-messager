<template>
  <div class="app-container">
  </div>
</template>

<script setup>
import { reactive, onMounted } from "vue"
const state = reactive({
  tableData: {
    pagination: {
      total: 0,
      pageNo: 1,
      pageSize: 10,
      pageCount: 0
    },
    rows: []
  }
})
onMounted(() => {
})

</script>

<style lang="scss" scoped>
</style>
