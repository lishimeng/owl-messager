<template>
  <imp-panel title>
    <el-form status-icon size="small" :model="state.form" label-width="180px">
      <el-form-item label="邮件开关">
        <el-switch
          v-model="state.form.debug"
          width="50"
          active-color="#13ce66"
          inactive-color="#ff4949"
          active-text="开启"
          inactive-text="关闭"
          :active-value="false"
          :inactive-value="true"
        />
      </el-form-item>
    </el-form>
  </imp-panel>
</template>

<script setup>
import { reactive, onMounted } from "vue"
import {  } from 'element-plus';
const state = reactive({
  form: {
    debug: ''
  }
})
onMounted(() => {
})
onMounted(()=> {
})

</script>
