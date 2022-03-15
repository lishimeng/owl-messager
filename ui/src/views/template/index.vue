<template>
  <div class="app-container">
    <el-table :data="state.tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="templateCode" label="模板" width="200px" />
      <el-table-column prop="templateBody" label="正文" />
      <!-- <el-table-column prop="category" label="分类" width="50px"></el-table-column> -->
      <!-- <el-table-column prop="description" label="描述"></el-table-column> -->
      <el-table-column prop="createTime" label="创建时间" width="200px" />
      <el-table-column prop="updateTime" label="更新时间" width="200px" />
      <el-table-column prop="status" label="状态" width="80px">
        <template #default="scope">
          <el-tag v-if="scope.row.status === 1" type="success" size="mini" disable-transitions>已启用</el-tag>
          <el-tag v-else-if="scope.row.status === 2" type="danger" size="mini" disable-transitions>已停用</el-tag>
        </template>
      </el-table-column>
      <!-- <el-table-column prop="name" label="模板名称"></el-table-column> -->
      <!-- <el-table-column prop="" label="操作" width="180px">
        <template slot-scope="">
					<el-button size="mini" type="default" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
					<el-button size="mini" type="danger" icon="el-icon-edit" @click="handleDel(scope.$index, scope.row)">删除</el-button>
				</template>
      </el-table-column> -->
    </el-table>
    <div style="text-align: left; margin-top: 15px">
      <el-pagination
        :current-page="state.tableData.pagination.pageNo"
        :pager-count="5"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="state.tableData.pagination.pageSize"
        background
        layout="total, sizes, prev, pager, next, jumper"
        :page-count="state.tableData.pagination.pageCount"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { reactive, onMounted } from "vue"
import { getMailTemplateApi } from '/@/api/index'
const state = reactive({
  tableData: {
    pagination: {
      total: 0,
      pageNo: 1,
      pageSize: 10,
      pageCount: 0
    },
    rows: []
  }
})
onMounted(() => {
})

function handleSizeChange(val) {
  state.tableData.pagination.pageSize = val
  getMailTemplate()
}
function handleCurrentChange(val) {
  state.tableData.pagination.pageNo = val
  getMailTemplate()
}
function getMailTemplate() {
  getMailTemplateApi({
    pageNo: state.tableData.pagination.pageNo,
    pageSize: state.tableData.pagination.pageSize
  }).then(res => {
    state.tableData.rows = res.items
    state.tableData.pagination.pageCount = res.totalPage
  }).catch(()=>{})
}
</script>

<style lang="scss" scoped>
</style>