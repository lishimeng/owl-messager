<template>
  <div class="app-container">
    <el-form
      ref="form"
      status-icon
      :model="state.form"
      :rules="state.rules"
      label-width="180px"
      style="width: 500px"
    >
      <el-form-item label="旧密码" prop="oldPassword">
        <el-input v-model="state.form.oldPassword" type="password" clearable />
      </el-form-item>
      <el-form-item label="新密码" prop="newPassword">
        <el-input v-model="state.form.newPassword" type="password" clearable />
      </el-form-item>
      <el-form-item label="再次确认新密码" prop="newPassword2">
        <el-input v-model="state.form.newPassword2" type="password" clearable />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">修改</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script  setup>
import {
  resetPwdApi
} from "/@/api/user"
import { useRouter } from 'vue-router';
import { ElMessage } from "element-plus"
import { resetRoute } from '/@/router/index';
import { reactive, getCurrentInstance } from "vue"
import { Local, Session } from "/@/utils/storage"
const { proxy } = getCurrentInstance()
const router = useRouter();
const state = reactive({
  form: {
    oldPassword: '',
    newPassword: '',
    newPassword2: ''
  },
  rules: {
    oldPassword: [
      { required: true, message: '请输入旧密码', trigger: 'blur' },
      { min: 6, message: '长度大于6个字符', trigger: 'blur' }
    ],
    newPassword: [
      { required: true, message: '请输入新密码', trigger: 'blur' },
      { min: 6, message: '长度大于6个字符', trigger: 'blur' }
    ],
    newPassword2: [
      { required: true, message: '请输入新密码', trigger: 'blur' },
      { min: 6, message: '长度大于6个字符', trigger: 'blur' }
    ]
  }
})
function onSubmit() {
  proxy.$refs['form'].validate((valid) => {
    if (valid) {
      if (state.form.newPassword !== state.form.newPassword2) {
        ElMessage.success("两次输入密码不一致")
        return
      }
      resetPwdApi({
        uid: parseInt(window.localStorage.getItem('userId') || ""),
        old: state.form.oldPassword,
        new: state.form.newPassword
      }).then((res) => {
        if (res && res.code === 200) {
          ElMessage.success("修改成功")
          Session.clear(); // 清除缓存/token等
          Local.clear();
          resetRoute(); // 删除/重置路由
          router.push('/login');
        } else {
          ElMessage.success('code:' + res.code + ';' + res.message)
        }
      })
    }
  })
}
</script>