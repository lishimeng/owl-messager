<template>
  <codemirror
    v-model="state.code"
    placeholder="Code goes here..."
    :style="{ height: '400px' }"
    :autofocus="true"
    :indent-with-tab="true"
    :tab-size="2"
    :extensions="extensions"
    @ready="handleReady"
    @change="onChange"
  />
</template>

<script setup lang="ts" name="jsEditor">
import { reactive, shallowRef } from 'vue'
import { Codemirror } from 'vue-codemirror'
import { javascript } from '@codemirror/lang-javascript'
import { oneDark } from '@codemirror/theme-one-dark'

const extensions = [javascript(), oneDark]

// 参数列表
const prop = defineProps({
  code: String
})

const state = reactive({
  code: prop.code
})

// Codemirror EditorView instance ref
const view = shallowRef()
const handleReady = (payload:any) => {
view.value = payload.view
}

const emit = defineEmits(['onChange'])

const onChange = () => {
  emit('onChange', state.code)
}

</script>