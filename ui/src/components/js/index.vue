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
import { ref, reactive, shallowRef, emit } from 'vue'
import { Codemirror } from 'vue-codemirror'
import { javascript } from '@codemirror/lang-javascript'
import { oneDark } from '@codemirror/theme-one-dark'
import { stat } from 'fs'

const code = ref(`console.log('Hello, world!')`)
const extensions = [javascript(), oneDark]

const prop = defineProps({
  code: String
})

const state = reactive({
  code: prop.code
})

const log = console.log

// Codemirror EditorView instance ref
const view = shallowRef()
const handleReady = (payload:any) => {
view.value = payload.view
}

const emit = defineEmits(['code'])

const onChange = () => {
  emit('code', state.code)
}

// Status is available at all times via Codemirror EditorView
const getCodemirrorStates = () => {
const state = view.value.state
const ranges = state.selection.ranges
const selected = ranges.reduce((r:any, range:any) => r + range.to - range.from, 0)
const cursor = ranges[0].anchor
const length = state.doc.length
const lines = state.doc.lines
// more state info ...
// return ...
}
</script>