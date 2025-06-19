<template>
  <div class="home-container layout-pd" style="margin-top: 10px">
    <el-form :inline="true">
      <el-form-item>
        <el-button type="primary" @click="showCreate()" icon="ele-CirclePlus">新增</el-button>
      </el-form-item>
    </el-form>
    <div style="margin-top: 10px">
      <el-table :data="state.dataList" border style="width: 100%">
        <el-table-column prop="id" label="Id" width="120"/>
        <el-table-column prop="name" label="名称" width="200" show-overflow-tooltip/>
        <el-table-column prop="appId" label="AppId" show-overflow-tooltip/>
        <el-table-column prop="createTime" label="创建时间" width="220">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.createTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="180">
          <template #default="scope">
            <el-button size="small" type="text" @click="showSecret(scope.row)">
              查看密钥
            </el-button>
            <el-popconfirm
                title="是否删除?"
                @confirm="deleteClient(scope.row)">
              <template #reference>
                <el-button size="small" type="text">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin: 10px 10px 10px 10px;">
        <el-pagination background
                       v-if="state.queryValue.totalNum>1"
                       layout="sizes, prev, pager, next, jumper"
                       v-model:currentPage="state.queryValue.pageNum"
                       v-model:page-size="state.queryValue.pageSize"
                       :page-sizes="[5,10,15,30,50,100]"
                       :page-count="state.queryValue.totalNum"
                       @size-change="onSizeChange"
                       @current-change="onCurrentChange"
        />
      </div>
    </div>
    <el-dialog
        v-model="state.showSecret"
        width="70%"
        center>
      <el-input v-model="state.secret" readonly/>
    </el-dialog>
    <el-dialog
        v-model="state.showCreate"
        title="创建客户端"
        width="50%"
        center>
      <el-form :model="state.formData" label-width="120px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="state.formData.name" :readonly="state.createSuccess"></el-input>
        </el-form-item>
        <el-form-item label="AppId" prop="appId" v-if="state.createSuccess">
          <el-input v-model="state.formData.appId" readonly></el-input>
        </el-form-item>
        <el-form-item label="密钥" prop="secret" v-if="state.createSuccess">
          <el-input v-model="state.formData.secret" readonly></el-input>
        </el-form-item>
      </el-form>
      <span class="dialog-footer mt30" style="display: flex; justify-content: center;">
        <el-button type="primary" v-if="state.createSuccess" @click="closeDialog()">完成</el-button>
        <el-button type="primary" v-else @click="createClient()">提交</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">

import {onMounted, reactive} from "vue";
import {formatDate} from "/@/utils/formatTime";
import {createClientApi, delClientApi, getClientListAPi, querySecretApi} from "/@/api/client";
import {ElMessage} from "element-plus";

const state = reactive({
  showSecret: false,
  showCreate: false,
  createSuccess: false,
  formData: {
    name: "",
    appId: "",
    secret: "",
  },
  secret: "",
  queryValue: {
    pageNum: 1,
    pageSize: 10,
    totalNum: 0,
    org: 1,
  },
  dataList: []
})

onMounted(() => {
  getClientList()
})
const onSizeChange = (val: object) => {
  state.queryValue.pageSize = val
  getClientList()
}
const onCurrentChange = (val: object) => {
  state.queryValue.pageNum = val
  getClientList()
}

const getClientList = () => {
  getClientListAPi(state.queryValue).then(res => {
    if (res && res.code == 200 && res.items) {
      state.dataList = res.items
      state.queryValue.totalNum = res.totalPage
    } else {
      state.dataList = []
      state.queryValue.totalNum = 0
    }
  })
}

const showCreate = () => {
  state.formData.name = ""
  state.formData.appId = ""
  state.formData.secret = ""
  state.createSuccess = false
  state.showCreate = true
}

const createClient = () => {
  createClientApi({
    org: 1,
    name: state.formData.name,
  }).then(res => {
    if (res.code == 200) {
      ElMessage.success(`创建成功`)
      getClientList()
      state.formData.appId = res.appId
      state.formData.secret = res.secret
      state.createSuccess = true
    } else {
      ElMessage.error(`失败:` + res.message)
    }
  }).catch(error => {
    ElMessage.error(`创建失败！`)
  })
}

const closeDialog = () => {
  state.showCreate = false
}

const showSecret = (row: object) => {
  querySecretApi({
    appId: row.appId
  }).then(res => {
    if (res.code == 200) {
      state.secret = res.secret
      state.showSecret = true
    } else {
      ElMessage.error(`失败:` + res.message)
    }
  }).catch(error => {
    ElMessage.error(`查询失败！`)
  })
}

const deleteClient = (row: object) => {
  delClientApi(row).then(res => {
    if (res.code == 200) {
      ElMessage.success(`删除成功`)
      getClientList()
    } else {
      ElMessage.error(`失败:` + res.message)
    }
  }).catch(error => {
    ElMessage.error(`删除失败！`)
  })
}
</script>

<style scoped lang="scss">

</style>