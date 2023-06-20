<template>
  <vue-html5-editor :content="state.content" :height="500"></vue-html5-editor>
</template>

<script setup lang="ts" name="htmlEditor">
import {reactive} from "vue";

const state=reactive({
  content:''
})
</script>

<style scoped>

</style>