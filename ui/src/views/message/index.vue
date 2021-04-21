<template>
  <div class="app-container">
    <el-table :data="tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="subject" label="邮件主题" />
      <el-table-column prop="priority" label="优先级" width="80px" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="发件时间" />
      <el-table-column prop="nextSendTime" label="下次发件时间" />
      <el-table-column prop="" label="邮件详情" width="100px">
        <template slot-scope="scope">
          <el-button size="mini" type="primary" icon="" @click="showInfo(1,scope.row)">查看邮件</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="" label="任务详情" width="100px">
        <template slot-scope="scope">
          <el-button size="mini" type="primary" icon="" @click="showInfo(2,scope.row)">查看任务</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div style="text-align: left; margin-top: 15px">
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
import { getMessageInfoApi } from '../../api/mail.js'

export default {
  name: 'Message',
  props: {},
  data() {
    return {
      tableData: {
        pagination: {
          total: 0,
          pageNo: 1,
          pageSize: 10,
          pageCount: 0
        },
        rows: []
      }
    }
  },
  created() {},
  mounted() {
    this.getMessageInfo()
  },
  methods: {
    handleSizeChange(val) {
      this.tableData.pagination.pageSize = val
      this.getMessageInfo()
    },
    handleCurrentChange(val) {
      this.tableData.pagination.pageNo = val
      this.getMessageInfo()
    },
    getMessageInfo() {
      getMessageInfoApi({
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
    },
    showInfo(type, row) {
      if (type === 1) {
        this.$router.push({
          path: '/mail',
          query: {
            messageId: row.id
          }
        })
      } else if (type === 2) {
        this.$router.push({
          path: '/task',
          query: {
            messageId: row.id
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
