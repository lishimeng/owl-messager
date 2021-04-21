<template>
  <div class="app-container">
    <!-- <el-row>
      <router-link :to="{ path: 'AddSender'}">
				<el-button type="primary" icon="el-icon-plus" size="small">新增</el-button>
			</router-link>
    </el-row> -->
    <el-table :data="tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" />
      <el-table-column prop="senderCode" label="发件CODE" />
      <el-table-column prop="host" label="主机地址" />
      <el-table-column prop="port" label="端口" />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="status" label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 2" type="success" size="mini" disable-transitions>已启用</el-tag>
          <el-tag v-else-if="scope.row.status === 1" type="danger" size="mini" disable-transitions>已停用</el-tag>
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
import { getSenderInfoApi } from '../../api/mail.js'
export default {
  name: 'Sender',
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
    this.getSenderInfoList()
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
    getSenderInfoList() {
      getSenderInfoApi({
        pageNum: this.tableData.pageNum,
        pageSize: this.tableData.pageSize
      }).then(res => {
        console.log(res)
        if (res) {
          this.tableData.rows = res.items
        }
      })
    }
    // handleDel(index, row) {
    // 	this.$confirm("此操作将永久删除该模板, 是否继续?", "提示", {
    // 	  confirmButtonText: "确定",
    // 	  cancelButtonText: "取消",
    // 	  type: "warning"
    // 	})
    // 	 .then(() => {
    // 	    this.deleteSenderInfo(row);
    // 	  })
    // 	.catch(() => {});
    // },
    // deleteSenderInfo(row){
    // 	deleteSenderInfoApi({
    // 	  id: row.id,
    // 	}).then(res => {
    // 	 if (res && res.code == 0) {
    // 		 this.$message.success("删除成功");
    // 		  this.getSenderInfoList();
    // 		} else {
    // 		  this.getException(res.code);
    // 		}
    // 	});
    // },
    // uphandleEdit(index, row){
    //   // updateSenderInfoApi({
    //   //   id: row.id
    //   // }).then(res =>{
    //   //   if (res && res.code == 0){
    //   //     this.$message("编辑成功")
    //   //   }
    //   // })
    // },
  }
}
</script>

<style lang="scss" scoped>

</style>
