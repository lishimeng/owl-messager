<template>
  <div class="app-container">
    <div>
      <el-form :inline="true" class="demo-form-inline" size="small">
        <el-form-item label>
          <el-input v-model.trim="state.userNo" clearable placeholder="用户编号" />
        </el-form-item>
        <el-form-item label>
          <el-select v-model="state.status" clearable placeholder="状态">
            <el-option
              v-for="item in state.statusList"
              :key="item.itemId"
              :label="item.itemName"
              :value="item.itemId"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" @click="queryList">查询</el-button>
        </el-form-item>
        <el-form-item>
          <!-- <router-link :to="{ path: '/user/add' }"> -->
          <el-button type="primary" icon="el-icon-plus" size="small" @click="state.drawer = true">新增</el-button>
          <!-- </router-link> -->
        </el-form-item>
      </el-form>
      <el-table
        v-loading="state.loading"
        border
        :data="state.tableData.rows"
        max-height="400px"
        :cell-style="{ padding: '3px' }"
        stripe
        style="width: 100%"
      >
        <el-table-column prop="userNo" label="用户编号" width="100px" />
        <el-table-column prop="userName" label="用户名称" />
        <el-table-column prop="phone" label="手机号" width="120px" />
        <el-table-column prop="email" label="邮箱" />
        <el-table-column label="状态" width="80px">
          <template #default="scope">
            <el-tag
              v-if="scope.row.status === 20"
              type="success"
              size="mini"
              disable-transitions
            >{{ getDictDataItemName(scope.row.status, 12) }}</el-tag>
            <el-tag
              v-else-if="scope.row.status === 50"
              type="danger"
              size="mini"
              disable-transitions
            >{{ getDictDataItemName(scope.row.status, 12) }}</el-tag>
            <el-tag v-else type="info" size="mini" disable-transitions>
              {{
                getDictDataItemName(scope.row.status, 12)
              }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="400">
          <template #default="scope">
            <el-button
              size="mini"
              type="default"
              icon="el-icon-edit"
              :disabled="scope.row.userId === 1 ? true : false"
              @click="handleEdit(scope.$index, scope.row)"
            >编辑</el-button>
            <el-button
              size="mini"
              type="primary"
              icon="el-icon-setting"
              :disabled="scope.row.userId === 1 ? true : false"
              @click="handleRoleConfig(scope.$index, scope.row)"
            >配置角色</el-button>
            <el-button
              size="mini"
              type="primary"
              icon="el-icon-RefreshRight"
              :disabled="scope.row.userId === 1 ? true : false"
              @click="handleResetPwd(scope.$index, scope.row)"
            >重置密码</el-button>
            <el-button
              v-if="scope.row.status != 50"
              size="mini"
              type="danger"
              icon="el-icon-circle-close"
              :disabled="scope.row.userId === 1 ? true : false"
              @click="setUserStatus(scope.row, 50)"
            >停用</el-button>
            <el-button
              v-if="scope.row.status === 50"
              size="mini"
              type="warning"
              icon="el-icon-circle-check"
              :disabled="scope.row.userId === 1 ? true : false"
              @click="setUserStatus(scope.row, 20)"
            >启用</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="text-align: right; right:15px; bottom: 15px; position:absolute;">
        <el-pagination
          :current-page="state.tableData.pagination.pageNum"
          :page-sizes="[10, 20, 50, 100]"
          :page-size="state.tableData.pagination.pageSize"
          background
          layout="total, sizes, prev, pager, next, jumper"
          :page-count="state.tableData.pagination.totalPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
      <!--修改用户信息-->
      <el-dialog
        title="修改用户信息"
        v-model="state.showUserInfoDialog"
        :close-on-click-modal="false"
        @close="colseEdit"
      >
        <el-form
          ref="userInfoForm"
          status-icon
          size="small"
          :model="state.userInfoForm"
          :rules="state.rules"
          label-width="100px"
        >
          <el-form-item label="用户编号" prop="userNo">
            <el-input v-model="state.userInfoForm.userNo" clearable />
          </el-form-item>
          <el-form-item label="用户姓名" prop="userName">
            <el-input v-model="state.userInfoForm.userName" clearable />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="state.userInfoForm.email" clearable />
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="state.userInfoForm.phone" clearable maxlength="11" />
          </el-form-item>
          <el-form-item label="状态">
            <el-radio-group
              v-model="state.userInfoForm.status"
              @change="setUserStatus(state.userInfoForm, state.userInfoForm.status)"
            >
              <el-radio :label="10">未激活</el-radio>
              <el-radio :label="20">已激活</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-form>
        <span class="dialog-footer">
          <el-button size="small" @click="state.showUserInfoDialog = false">取 消</el-button>
          <el-button size="small" type="primary" @click="editUserInfo()">确 定</el-button>
        </span>
      </el-dialog>
      <!--配置用户角色-->
      <el-dialog
        title="配置用户角色"
        v-model="state.showRoleDialog"
        :close-on-click-modal="false"
        :close="colseRole"
      >
        <div class="select-tree">
          <el-scrollbar
            tag="div"
            class="is-empty"
            wrap-class="el-select-dropdown__wrap"
            view-class="el-select-dropdown__list"
          >
            <el-tree
              ref="roles"
              :data="state.roleTree"
              show-checkbox
              check-on-click-node
              node-key="id"
              @check="handleCheckChange"
            />
          </el-scrollbar>
        </div>
        <span class="dialog-footer">
          <el-button size="small" type="primary" @click="setRole()">确 定</el-button>
        </span>
      </el-dialog>
    </div>
    <el-drawer
      title="新增用户"
      v-model="state.drawer"
      :destroy-on-close="true"
      custom-class="drawer"
      ref="drawer"
      :close-on-press-escape="true"
      :before-close="handleClose"
    >
      <AddUser />
    </el-drawer>
  </div>
</template>
<script  setup>
import {
  getUserListApi,
  getInfo,
  setUserStatusApi,
  addUserRoleApi,
  deleteUserRoleApi,
  editUserApi,
  resetPwdByAdminApi
} from "/@/api/user"
import AddUser from "/@/views/user/userAdd.vue"
import {
  getDictDataApi,
} from "/@/api/index"
import { ElAlert, ElMessage, ElMessageBox } from "element-plus"
import { reactive, onMounted, getCurrentInstance } from "vue"
import { done } from "nprogress"
import { copyText } from 'vue3-clipboard'
const { proxy } = getCurrentInstance()
const state = reactive({
  drawer: false,
  firstOpen: false,
  userNo: '',
  status: '',
  statusList: [],
  showUserInfoDialog: false,
  showRoleDialog: false,
  formLabelWidth: '120px',
  selectUserId: '',
  userId: parseInt(window.localStorage.getItem('userId') || ""),
  roleTree: [
    { id: 1, label: 'Administrator' },
    // { id: 2, label: '收料员' },
    { id: 3, label: '仓库管理员' },
    { id: 4, label: '设备管理员' },
    { id: 5, label: '设备使用者' },
    { id: 6, label: '系统管理员' }
  ],
  permission: true,
  tableData: {
    pagination: {
      total: 0,
      pageNum: 1,
      pageSize: 10,
      totalPage: 0
    },
    rows: []
  },
  loading: false,
  userInfoForm: {
    userId: '',
    userName: '',
    userNo: '',
    email: '',
    phone: '',
    status: ''
  },
  rules: {
    userNo: [{ required: true, message: '请输入用户编号', trigger: 'blur' }],
    userName: [
      { required: true, message: '请输入用户名称', trigger: 'blur' }
    ],
    email: [
      { required: true, message: '请输入邮箱地址', trigger: 'blur' },
      { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur'] }
    ],
    phone: [
      { min: 11, max: 11, message: '长度为11个字符', trigger: 'blur' }
    ],
    password: [
      { required: true, message: '请输入密码', trigger: 'blur' },
      { min: 6, message: '长度大于6个字符', trigger: 'blur' }
    ]
  },
})

// onComputed(() => {
//   function isRole() {
//     if (window.localStorage.getItem('roles').indexOf(10) > -1) {
//       return true
//     } else {
//       return false
//     }
//   }
// })

onMounted(() => {
  getUserList()
  getDictData(12)
})

function getDictDataItemName(itemId, groupId) {
  var list = []
  switch (groupId) {
    // 状态
    case 12:
      list = state.statusList
      break;
  }
  for (let i = 0; i < list.length; i++) {
    if (itemId == list[i].itemId) {
      return list[i].itemName
    }
  }
  return ""
}

// 获取库位状态字典
function getDictData(groupId) {
  getDictDataApi({
    groupId: groupId,
  }).then(res => {
    if (res.code == 200 && res.items.length > 0) {
      switch (groupId) {
        // 状态
        case 12:
          state.statusList = res.items
          break;
      }
    }
  }).catch((err) => { console.log(err) })
}

function handleSizeChange(val) {
  state.tableData.pagination.pageSize = val
  getUserList()
}
function handleCurrentChange(val) {
  state.tableData.pagination.pageNum = val
  getUserList()
}
function queryList() {
  state.tableData.pagination.pageNum = 1
  getUserList()
}


// 获取用户列表
function getUserList() {
  getUserListApi({
    pageNo: state.tableData.pagination.pageNum,
    pageSize: state.tableData.pagination.pageSize,
    userNo: state.userNo,
    status: state.status
  }).then((res) => {
    if (res && res.code === 200) {
      state.tableData.pagination.totalPage = res.totalPage
      state.tableData.rows = res.items ? res.items : []
    } else {
      // state.getException(res.code)
    }
  })
}

// 打开修改页面
function handleEdit(index, row) {
  getUserInfo(row)
}
// 关闭修改页面
function colseEdit() {
  proxy.$refs['userInfoForm'].resetFields()
}
// 修改用户
function editUserInfo() {
  proxy.$refs['userInfoForm'].validate((valid) => {
    if (valid) {
      editUserApi({
        userId: state.userInfoForm.userId,
        userName: state.userInfoForm.userName,
        userNo: state.userInfoForm.userNo,
        email: state.userInfoForm.email,
        phone: state.userInfoForm.phone
      }).then((res) => {
        if (res && res.code === 200) {
          state.showUserInfoDialog = false
          ElMessage.success("修改成功")
          getUserList()
        } else {
          // state.getException(res.code)
        }
      })
    }
  })
}

// 打开当前用户权限
function handleRoleConfig(index, row) {
  state.showRoleDialog = true
  state.selectUserId = row.userId
  getInfo({
    userId: row.userId
  }).then((res) => {
    if (res && res.code === 200) {
      if (res.roles && res.roles.length > 0) {
        proxy.$refs.roles.setCheckedKeys(res.roles)
      } else {
        proxy.$refs.roles.setCheckedKeys([])
      }
    } else {
      // getException(res.code)
    }
  })
}
// 关闭权限页面
function colseRole() {
  proxy.$refs.roles.setCheckedKeys([])
  state.selectUserId = ''
}

// 配置用户权限
function handleCheckChange(data, node) {
  for (var i = 0; i < node.checkedKeys.length; i++) {
    if (data.id === node.checkedKeys[i]) {
      addUserRole(data)
      return
    }
  }
  deleteUserRole(data)
}
// 添加用户权限
function addUserRole(data) {
  addUserRoleApi({
    uid: state.selectUserId,
    rid: data.id
  }).then(res => {
    if (res && res.code === 200) {
      ElMessage.success("添加权限成功：" + data.label)
    } else {
      ElMessage.success("操作失败")
    }
  }).catch(err => {
    console.error(err)
  })
}
// 删除用户权限
function deleteUserRole(data) {
  if (!(state.selectUserId && data.id)) return
  deleteUserRoleApi({
    userId: state.selectUserId,
    roleId: data.id
  }).then(res => {
    if (res && res.code === 200) {
      ElMessage.success("取消权限成功：" + data.label)
    } else {
      ElMessage.success("操作失败")
    }
  }).catch(err => {
    console.error(err)
  })
}
// 管理员重置密码
function handleResetPwd(index,row) {
  var userName = row.userName
  ElMessageBox({
      title: "重置密码提示",
      message: '此操作将重置用户 ' + userName + ' 的登录密码, 是否继续?',
      showCancelButton: true,
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      showClose: false,
      beforeClose: (action, instance, done) => {
        if (action === 'confirm') {
          instance.confirmButtonLoading = true;
          instance.showCancelButton = false;
          instance.confirmButtonText = '重置中...';
          setTimeout(() => {
            resetPwd(row)
            done();
            setTimeout(() => {
              instance.confirmButtonLoading = false;
            }, 300);
          }, 3000);
        } else {
          done();
        }
      },
    }).then(action => {
      ElMessage({
        type: 'success',
        message: '操作成功'
      })
    }).catch(action =>{});
}
function resetPwd(row) {
  resetPwdByAdminApi({
    uid: row.userId
  }).then(res =>{
    ElMessageBox({
      message: '新密码(请复制保存): ' + res.password ,
      confirmButtonText: '确定',
      center: true,
      beforeClose: (action, instance, done) => {
        if (action === 'confirm') {
          copyText(res.password, undefined, (error) => {
            if (error) {
              ElMessage.warning(`复制失败: ${error} ！`);
            } else {
              ElMessage.success(`复制: ${res.password} 成功！`);
            }
          });
          done();
        } else {
          done();
        }
      }
    })
  }).catch(err =>{})
}
function setRole() {
  proxy.$refs.roles.setCheckedKeys([])
  state.selectUserId = ''
  state.showRoleDialog = false
}

// 获取用户信息
function getUserInfo(row) {
  getInfo({
    userId: row.userId
  }).then((res) => {
    if (res && res.code === 200) {
      state.showUserInfoDialog = true
      state.userInfoForm.userId = res.userId
      state.userInfoForm.userName = res.userName
      state.userInfoForm.userNo = res.userNo
      state.userInfoForm.email = res.email
      state.userInfoForm.phone = res.phone
      state.userInfoForm.status = res.status
    } else {
      // getException(res.code)
    }
  })
}

// 修改用户状态
function setUserStatus(row, status) {
  setUserStatusApi({
    userId: row.userId,
    status: status
  }).then((res) => {
    if (res && res.code === 200) {
      ElMessage.success("修改成功")
      getUserList()
    } else {
      // state.getException(res.code)
    }
  })
}

function handleClose(done){
  getUserList()
  done()
}
</script>
<style scoped>
</style>
