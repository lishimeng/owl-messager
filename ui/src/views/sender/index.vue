<template>
  <div class="app-container">
    <!-- <el-row>
      <router-link :to="{ path: 'AddSender'}">
				<el-button type="primary" icon="el-icon-plus" size="small">新增</el-button>
			</router-link>
    </el-row>-->
    <el-table :data="state.tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="senderCode" label="发件CODE" />
      <el-table-column prop="host" label="主机地址" />
      <el-table-column prop="port" label="端口" width="80px" />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="status" label="状态" width="80px">
        <template #default="scope">
          <el-tag v-if="scope.row.status === 2" type="success" size="mini" disable-transitions>已启用</el-tag>
          <el-tag
            v-else-if="scope.row.status === 1"
            type="danger"
            size="mini"
            disable-transitions
          >已停用</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="password" label="密码" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="修改时间" />
      <!-- <el-table-column prop="" label="操作">
        <template slot-scope="scope">
					<el-button size="mini" type="default" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
					<el-button size="mini" type="danger" icon="el-icon-edit" @click="handleDel(scope.$index, scope.row)">删除</el-button>
				</template>
      </el-table-column>-->
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
import { ElMessage, ElMessageBox } from "element-plus";
import { reactive, onMounted, getCurrentInstance } from "vue"
import { useRouter, useRoute } from "vue-router"
import { getSenderInfoApi } from '/@/api/index'
const route = useRoute()
const state = reactive({
  userId: parseInt(route.query.userId),
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
  getSenderInfoList()
})
function handleSizeChange(val) {
  state.tableData.pagination.pageSize = val
  getSenderInfoList()
}
function handleCurrentChange(val) {
  state.tableData.pagination.pageNo = val
  getSenderInfoList()
}
function getSenderInfoList() {
  getSenderInfoApi({
    pageNo: state.tableData.pagination.pageNo,
    pageSize: state.tableData.pagination.pageSize
  }).then(res => {
    state.tableData.rows = res.items
    state.tableData.pagination.pageCount = res.totalPage
  }).catch(err => { console.log(err) })
}
// function handleDel(index, row) {
//   ElMessageBox.confirm("此操作将永久删除该模板, 是否继续?", "提示", {
//     confirmButtonText: "确定",
//     cancelButtonText: "取消",
//     type: "warning"
//   })
//     .then(() => {
//       deleteSenderInfo(row);
//     })
//     .catch(() => { });
// }
// function deleteSenderInfo(row) {
//   deleteSenderInfoApi({
//     id: row.id,
//   }).then(res => {
//     ElMessage.success("删除成功")
//     getSenderInfoList()
//   }).catch(err => { console.log(err) })
// }
// function uphandleEdit(index, row) {
//   updateSenderInfoApi({
//     id: row.id
//   }).then(res => {
//     ElMessage.success("编辑成功")
//   }).catch(err => { console.log(err) })
// }
</script>

<style lang="scss" scoped>
</style>
