<template>
  <div class="app-container">
    <!-- <el-row>
      <router-link :to="{ path: 'AddMailTemplate'}">
				<el-button type="primary" icon="el-icon-plus" size="small">新增</el-button>
			</router-link>
    </el-row> -->
    <el-table :data="tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="50px" />
      <el-table-column prop="templateCode" label="模板" />
      <el-table-column prop="templateBody" label="正文" />
      <!-- <el-table-column prop="category" label="分类" width="50px"></el-table-column> -->
      <!-- <el-table-column prop="description" label="描述"></el-table-column> -->
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="更新时间" />
      <el-table-column prop="status" label="状态" width="80px">
        <template slot-scope="scope">
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
import { getMailTemplateApi } from '../../api/mail.js'
export default {
  name: 'MailTemplate',
  props: {},
  data() {
    return {
      tableData: {
        pagination: {
          total: 0,
          pageNum: 1,
          pageSize: 20,
          totalSize: 0
        },
        rows: []
      }
    }
  },
  created() {},
  mounted() {
    this.getMailTemplate()
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
    getMailTemplate() {
      getMailTemplateApi({
        pageNum: this.tableData.pagination.pageNum,
        pageSize: this.tableData.pagination.pageSize
      }).then(res => {
        console.log(res)
        if (res) {
          this.tableData.rows = res.items
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
