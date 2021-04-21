<template>
  <div class="app-container">
    <el-table :data="tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" />
      <el-table-column prop="subject" label="邮件主题" />
      <el-table-column prop="priority" label="优先级" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="发件时间" />
      <el-table-column prop="nextSendTime" label="下次发件时间" />
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
    <!-- <el-row :gutter="32" style="background:#fff;padding:16px 16px 0;margin-top:32px;">
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <raddar-chart />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <pie-chart />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <bar-chart />
        </div>
      </el-col>
    </el-row> -->
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
          pageNum: 1,
          pageSize: 10,
          totalSize: 0
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
      this.getMaterialOperationList()
    },
    handleCurrentChange(val) {
      this.tableData.pagination.pageNum = val
      this.getMaterialOperationList()
    },
    getMessageInfo() {
      getMessageInfoApi({
        pageNum: this.tableData.pageNum,
        pageSize: this.tableData.pageSize
      }).then(res => {
        if (res) {
          this.tableData.rows = res.items
        }
      })
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
