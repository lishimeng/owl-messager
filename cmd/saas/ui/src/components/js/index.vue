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
      :disabled="state.disabled"
  />
</template>

<script setup lang="ts" name="jsEditor">
import {reactive, shallowRef, watch} from 'vue'
import {Codemirror} from 'vue-codemirror'
import {javascript} from '@codemirror/lang-javascript'
import {oneDark} from '@codemirror/theme-one-dark'

const extensions = [javascript(), oneDark]

// 参数列表
const prop = defineProps({
  code: String,
  disabled: {
    type: Boolean,
    default: () => false,
  }
})

const state = reactive({
  code: prop.code,
  disabled: prop.disabled,
  options: {
   /* mode: "text/html",
    htmlMode: true,
    lineNumBers: true,*/
  }
})

// Codemirror EditorView instance ref
const view = shallowRef()
const handleReady = (payload: any) => {
  view.value = payload.view
}

const emit = defineEmits(['onChange','update:code'])

const onChange = (val) => {
  console.log(val)
  emit('onChange', state.code)
  emit('update:code', state.code)
}
watch(
    () => state.code,
    (val) => {
      state.code = val
      console.log("code:",state.code)
    },
    {deep: true,}
)
</script>
<style scoped>
.CodeMirror {
  width: 800px !important;
}

</style>