<template>
  <div class="app-container">
    <el-table :data="tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" />
      <el-table-column prop="messageId" label="邮件ID" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="修改时间" />
      <el-table-column prop="status" label="状态" width="80px">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.Status === 1" type="success" size="mini" disable-transitions>已启用</el-tag>
          <el-tag v-else-if="scope.row.Status === 2" type="danger" size="mini" disable-transitions>已停用</el-tag>
        </template>
      </el-table-column>
    </el-table>
    <div style="text-align: left; margin-top: 15px">
      <el-pagination
        :current-page="tableData.pagination.pageNum"
        :pager-count="5"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="tableData.pagination.pageSize"
        background
        layout="total, sizes, prev, pager, next, jumper"
        :total="tableData.pagination.totalSize"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import { getTaskApi } from '../../api/mail.js'
export default {
  name: 'Task',
  props: {},
  data() {
    return {
      tableData: {
        pagination: {
          total: 0,
          pageNum: 1,
          pageSize: 10,
          totalSize: 0
        },
        rows: []
      },
      status: 2
    }
  },
  created() {},
  mounted() {
    this.getTaskList()
  },
  methods: {
    handleSizeChange(val) {
      this.tableData.pagination.pageSize = val
      this.getMaterialOperationList()
    },
    handleCurrentChange(val) {
      this.tableData.pagination.pageNum = val
      this.getMaterialOperationList()
    },
    getTaskList() {
      getTaskApi({
        status: this.status,
        pageNum: this.tableData.pagination.pageNum,
        pageSize: this.tableData.pagination.pageSize
      }).then(res => {
        if (res) {
          this.tableData.rows = res.items
        }
      })
    }
  }
}
</script>
