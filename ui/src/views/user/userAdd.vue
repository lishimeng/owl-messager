<template>
  <div class="app-container">
    <el-form
      ref="ruleForm"
      status-icon
      :model="state.form"
      :rules="state.rules"
      size="small"
      label-width="100px"
    >
      <el-form-item label="用户编号" prop="userNo">
        <el-input v-model="state.form.userNo" />
      </el-form-item>
      <el-form-item label="用户姓名" prop="userName">
        <el-input v-model="state.form.userName" />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="state.form.password" type="password" />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="state.form.email" />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="state.form.phone" maxlength="11" />
      </el-form-item>
      <el-form-item label="状态">
        <el-radio-group v-model="state.form.status">
          <el-radio :label="10">未激活</el-radio>
          <el-radio :label="20">已激活</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script setup>
import {
  addUserApi
} from "/@/api/user"
import { reactive, getCurrentInstance } from "vue"
import { ElMessage } from "element-plus"
import { useRouter } from "vue-router"
const router = useRouter()
const { proxy } = getCurrentInstance()

const state = reactive({
  roles: [],
  form: {
    userNo: '',
    userName: '',
    password: '',
    email: '',
    phone: '',
    status: 20
  },
  rules: {
    userNo: [{ required: true, message: '请输入用户编号', trigger: 'blur' }],
    userName: [
      { required: true, message: '请输入用户名称', trigger: 'blur' }
    ],
    password: [
      { required: true, message: '请输入密码', trigger: 'blur' },
      { min: 6, message: '长度大于6个字符', trigger: 'blur' }
    ],
    email: [
      { required: true, message: '请输入邮箱地址', trigger: 'blur' },
      {
        type: 'email',
        message: '请输入正确的邮箱地址',
        trigger: ['blur', 'change']
      }
    ],
    phone: [
      { min: 11, max: 11, message: '长度为11个字符', trigger: 'blur' }
    ]
  }
})

function onSubmit() {
  proxy.$refs['ruleForm'].validate().then((value) => {
    if (value) {
      addUserApi({
        userNo: state.form.userNo,
        userName: state.form.userName,
        password: state.form.password,
        email: state.form.email,
        phone: state.form.phone,
        status: state.form.status
      }).then((res) => {
        if (res && res.code === 200) {
          ElMessage.success("创建成功")
          router.push({
            path: '/user',
          })
        } else {
          ElMessage.success(res.code + ':' + res.message)
        }
      })
    }
  })
}
</script>
