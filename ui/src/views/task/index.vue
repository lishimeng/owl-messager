<template>
  <div class="app-container">
    <el-table :data="tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="messageId" label="邮件ID" width="80px" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="修改时间" />
      <el-table-column prop="status" label="状态" width="80px">
        <!-- <template slot-scope="scope">
          <el-tag v-if="scope.row.Status === 1" type="success" size="mini" disable-transitions>已启用</el-tag>
          <el-tag v-else-if="scope.row.Status === 2" type="danger" size="mini" disable-transitions>已停用</el-tag>
        </template> -->
      </el-table-column>
    </el-table>
    <div v-if="showPage" style="text-align: left; margin-top: 15px">
      <el-pagination
        :current-page="tableData.pagination.pageNo"
        :pager-count="5"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="tableData.pagination.pageSize"
        background
        layout="total, sizes, prev, pager, next, jumper"
        :page-count="tableData.pagination.pageCount"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import { getTaskApi, getTaskInfoByMessageIdApi } from '../../api/mail.js'
export default {
  name: 'Task',
  props: {},
  data() {
    return {
      messageId: this.$route.query.messageId,
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
    }
  },
  created() {},
  mounted() {
    this.getTaskList()
  },
  methods: {
    handleSizeChange(val) {
      this.tableData.pagination.pageSize = val
      this.getTaskList()
    },
    handleCurrentChange(val) {
      this.tableData.pagination.pageNo = val
      this.getTaskList()
    },
    getTaskList() {
      if (this.messageId) {
        this.showPage = false
        getTaskInfoByMessageIdApi({
          messageId: this.messageId,
          pageNo: this.tableData.pagination.pageNo,
          pageSize: this.tableData.pagination.pageSize
        }).then(res => {
          if (res && res.code === 200) {
            var list = []
            list.push(res)
            this.tableData.rows = list
          } else {
            this.$message.error(res.message)
          }
        })
      } else {
        getTaskApi({
          pageNo: this.tableData.pagination.pageNo,
          pageSize: this.tableData.pagination.pageSize
        }).then(res => {
          if (res && res.code === 200) {
            this.tableData.rows = res.items
            this.tableData.pagination.pageCount = res.totalPage
          } else {
            this.$message.error(res.message)
          }
        })
      }
    }
  }
}
</script>
