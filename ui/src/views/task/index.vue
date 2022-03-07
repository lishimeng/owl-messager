<template>
  <div>
    <el-table :data="state.tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="messageId" label="邮件ID" width="80px" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="修改时间" />
      <el-table-column prop="status" label="状态" width="80px">
      </el-table-column>
    </el-table>
    <div v-if="state.showPage" style="text-align: left; margin-top: 15px">
      <el-pagination
        :current-page="state.tableData.pagination.pageNo"
        :pager-count="5"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="state.tableData.pagination.pageSize"
        background
        layout="sizes, prev, pager, next, jumper"
        :page-count="state.tableData.pagination.pageCount"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { reactive, onMounted } from "vue"
import { getTaskApi, getTaskInfoByMessageIdApi } from '/@/api/index'
import { useRouter, useRoute } from "vue-router"
const route = useRoute()
const state = reactive({
  messageId: route.query.messageId,
  tableData: {
    pagination: {
      total: 0,
      pageNo: 1,
      pageSize: 10,
      pageCount: 0
    },
    rows: []
  },
  showPage: true
})
onMounted(() => {
  getTaskList()
})
function handleSizeChange(val) {
  state.tableData.pagination.pageSize = val
  getTaskList()
}
function handleCurrentChange(val) {
state.tableData.pagination.pageNo = val
getTaskList()
}
function getTaskList() {
  if (state.messageId) {
    state.showPage = false
    getTaskInfoByMessageIdApi({
      messageId: state.messageId,
      pageNo: state.tableData.pagination.pageNo,
      pageSize: state.tableData.pagination.pageSize
    }).then(res => {
      state.tableData.rows = []
      state.tableData.rows.push(res)
    })
  } else {
    getTaskApi({
      pageNo: state.tableData.pagination.pageNo,
      pageSize: state.tableData.pagination.pageSize
    }).then(res => {
      state.tableData.rows = res.items ? res.items : []
      state.tableData.pagination.pageCount = res.totalPage
    })
  }
}
</script>
