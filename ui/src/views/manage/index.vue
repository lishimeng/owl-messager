<template>
  <imp-panel title>
    <el-form status-icon size="small" :model="form" :rules="rules" label-width="180px">
      <el-form-item label="邮件开关">
        <el-switch
          v-model="form.debug"
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

<script>
import { getCurrentSwitchStatusApi, changeSwitchStatusApi } from '../../api/mail.js'
export default {
  data() {
    return {
      form: {
        debug: ''
      }
    }
  },
  mounted() {
    this.getCurrentSwitchStatus()
  },
  methods: {
    getCurrentSwitchStatus() {
      getCurrentSwitchStatusApi({
        operaterId: window.localStorage.getItem('userId')
      }).then(res => {
        this.form.debug = res
      })
    },
    changeSwitchStatus() {
      changeSwitchStatusApi({
        debug: this.form.debug,
        operaterId: window.localStorage.getItem('userId')
      }).then(res => {
        if (res && res.code === 0) {
          this.$message.success(res.message)
          this.getCurrentSwitchStatus()
        } else {
          this.$message.error(res.message)
          this.getCurrentSwitchStatus()
        }
      })
    }
  }
}
</script>
