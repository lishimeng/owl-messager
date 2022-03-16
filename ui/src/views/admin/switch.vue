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
          @change="changeSwitchStatus()"
        />
      </el-form-item>
    </el-form>
  </imp-panel>
</template>

<script setup>
import { getCurrentSwitchStatusApi, changeSwitchStatusApi } from '/@/api/mail'
import { reactive, onMounted } from "vue"
import { ElMessage } from 'element-plus';
const state = reactive({
  form: {
    debug: ''
  }
})
onMounted(() => {
})
onMounted(()=> {
  getCurrentSwitchStatus()
})
function getCurrentSwitchStatus() {
  getCurrentSwitchStatusApi({
    operaterId: window.localStorage.getItem('userId')
  }).then(res => {
    state.form.debug = res
  })
}
function changeSwitchStatus() {
  changeSwitchStatusApi({
    debug: state.form.debug,
    operaterId: window.localStorage.getItem('userId')
  }).then(res => {
    if (res && res.code === 0) {
      ElMessage.success(res.message)
    } else {
      ElMessage.error(res.message)
    }
    getCurrentSwitchStatus()
  })
}
</script>
