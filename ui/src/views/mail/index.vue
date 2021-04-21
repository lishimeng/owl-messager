<template>
  <div class="app-container">
    <el-table :data="tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="发送时间" />
      <el-table-column prop="messageId" label="消息ID" />
      <el-table-column prop="templateId" label="模板ID" />
      <el-table-column prop="params" label="内容" />
      <el-table-column prop="status" label="状态" width="80px">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="success" size="mini" disable-transitions>已发送</el-tag>
          <el-tag v-else-if="scope.row.status === 2" type="danger" size="mini" disable-transitions>未知</el-tag>
        </template>
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
import { getMailApi, getMailByMessageIdApi } from '../../api/mail.js'

export default {
  name: 'Message',
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
    this.getMailList()
  },
  methods: {
    handleSizeChange(val) {
      this.tableData.pagination.pageSize = val
      this.getMaterialOperationList()
    },
    handleCurrentChange(val) {
      this.tableData.pagination.pageNo = val
      this.getMaterialOperationList()
    },
    getMailList() {
      if (this.messageId) {
        this.showPage = false
        getMailByMessageIdApi({
          messageId: this.messageId,
          pageNo: this.tableData.pageNo,
          pageSize: this.tableData.pageSize
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
        getMailApi({
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

<style lang="scss" scoped>
  .chart-container{
  width: 100%;
  height: 300px;
}
</style>
