<template>
  <div class="app-container">
     <el-row>
      <!-- <router-link :to="{ path: 'AddSender'}"> -->
				<el-button type="primary" icon="el-icon-plus" size="small" @click="addsender">新增</el-button>
			<!-- </router-link> -->
    </el-row>
    <el-table :data="state.tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="code" label="发件CODE" />
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

      <!-- <el-table-column fixed="right" label="操作" width="150">
        <template slot-scope="scope">
					<el-button size="mini" type="default" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
					<el-button size="mini" type="danger" icon="el-icon-edit" @click="handleDel(scope.$index, scope.row)">删除</el-button>
				</template>
      </el-table-column> -->
      <el-table-column fixed="right" label="操作" width="150">
      <template v-slot="scope">
              <el-button type="text" size="small" @click="handleEdit(scope.row)">修改</el-button>
              <el-popconfirm :title="'删除'" @confirm="sputDel(scope.row)">
                <template #reference>
              <el-button type="text" size="small">删除</el-button>
      </template>
     </el-popconfirm>
     </template>
     </el-table-column>

    </el-table>
    <div style="text-align: left; margin-top: 15px">
      <el-pagination
        :current-page="state.tableData.pagination.pageNo"
        :pager-count="5"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="state.tableData.pagination.pageSize"
        background
        :total="state.tableData.pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        :page-count="state.tableData.pagination.pageCount"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>


    
    <!-- 新增弹框 -->
    <el-dialog destroy-on-close title="新增" v-model="state.editDialogDetailAdd" center width="50%">
         <el-form ref="ruleForm" :model="state.searchFormAdd" :rules="state.rules" label-width="220px">
                   <el-form-item label="发件CODE" prop="code">
                      <el-input clearable v-model="state.searchFormAdd.code"></el-input>
                   </el-form-item>
                  <el-form-item label="主机地址" prop="host">
                    <el-input clearable v-model="state.searchFormAdd.host"></el-input>
                 </el-form-item>

                <el-form-item label="端口" prop="port">
                  <el-input-number v-model="state.searchFormAdd.port" :min="1"   />
                    <!-- <el-input clearable v-model="state.searchFormAdd.post"></el-input> -->
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input clearable v-model="state.searchFormAdd.email"></el-input>
                </el-form-item>

                <el-form-item label="状态" prop="status">
                  <el-select clearable v-model="state.searchFormAdd.status" placeholder="状态">
                    <el-option v-for="item in state.itemStatus" :label="item.name" :value="item.id" :key="item.id"></el-option>
                  </el-select>
               </el-form-item>

                <el-form-item label="密码" prop="password">
                    <el-input clearable v-model="state.searchFormAdd.password"></el-input>
                </el-form-item>

                <el-form-item prop="createTime" label="创建时间">
                   <el-date-picker v-model="state.searchFormAdd.createTime" placeholder="创建时间"  size="small"
                    :clearable="true">
                  </el-date-picker>
                </el-form-item>
                <el-form-item prop="updateTime" label="修改时间">
                   <el-date-picker v-model="state.searchFormAdd.updateTime" placeholder="修改时间"  size="small"
                    :clearable="true">
                  </el-date-picker>
                </el-form-item>
                <el-form-item>
                    <el-button @click="state.editDialogDetailAdd = false">取消</el-button>
                    <el-button type="primary" @click="determine()">确定</el-button>
                </el-form-item>
            </el-form>
    </el-dialog>


<!-- 修改弹框 -->
<el-dialog destroy-on-close title="修改" v-model="state.editDialogDetailupdate" center width="50%">
         <el-form ref="ruleForm" :model="state.searchFormUpdate" :rules="state.rules" label-width="220px">
                   <el-form-item label="发件CODE" prop="code">
                      <el-input clearable v-model="state.searchFormUpdate.code"></el-input>
                   </el-form-item>
                  <el-form-item label="主机地址" prop="host">
                    <el-input clearable v-model="state.searchFormUpdate.host"></el-input>
                 </el-form-item>

                <el-form-item label="端口" prop="port">
                    <el-input clearable v-model="state.searchFormUpdate.port"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input clearable v-model="state.searchFormUpdate.email"></el-input>
                </el-form-item>

                <el-form-item label="状态" prop="status">
                  <el-select clearable v-model="state.searchFormUpdate.status" placeholder="状态">
                    <el-option v-for="item in state.itemStatus" :label="item.name" :value="item.id" :key="item.id"></el-option>
                  </el-select>
               </el-form-item>

                <el-form-item label="密码" prop="password">
                    <el-input clearable v-model="state.searchFormUpdate.password"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button @click="state.editDialogDetailupdate = false">取消</el-button>
                    <el-button type="primary" @click="updatesender()">确定</el-button>
                </el-form-item>
            </el-form>
    </el-dialog>

  </div>
</template>

<script setup>
import { ElMessage, ElMessageBox } from "element-plus";
import { reactive, onMounted, getCurrentInstance } from "vue"
import { useRouter, useRoute } from "vue-router"
import { formatDate } from "/@/utils/formatTime"
import { getSenderInfoApi,createMailTemplateApi,getsenderlistOneApi,postsenderupdateOneApi,delsenderlistOneApi } from '/@/api/index'
const route = useRoute()
const { proxy } = getCurrentInstance();
const state = reactive({
  editDialogDetailAdd : false,
  editDialogDetailupdate:false,
  userId: parseInt(route.query.userId),
  tableData: {
    pagination: {
      total: 0,
      pageNo: 1,
      pageCount:0,
      pageSize: 10
    },
    rows: []
  },
  itemStatus:[
      {'id':1,'name':'disable'},
      {'id':2,'name':'enable'},
 ],
  searchFormAdd:{
   
  },
  searchFormUpdate:{
    createTime: formatDate(new Date, "YYmmdd"),
    updateTime: formatDate(new Date, "YYmmdd"),
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
    state.tableData.pagination.total=res.more
  }).catch(err => { console.log(err) })
}
function handleDel(index, row) {
  ElMessageBox.confirm("此操作将永久删除该模板, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      deleteSenderInfo(row);
    })
    .catch(() => { });
}


function handleEdit( row){
  getsenderlistOneApi({ id:  row.id}).then((res) => {
        console.log(res);
        state.editDialogDetailupdate = true;
        state.searchFormUpdate = res.data;
    });
}


function sputDel(row){
  delsenderlistOneApi({ id:  row.id}).then((res) => {
      console.log(res);
      getSenderInfoList();
      ElMessage.success("删除成功")
 });

}

// function deleteSenderInfo(row) {
//   deleteSenderInfo({
//     id: row.id
//     .then(res => {
//     ElMessage.success("删除成功")
//   }).catch(err => { console.log(err) })
//   })
// }
function addsender(row){
  state.editDialogDetailAdd=true
}

function determine(row){
  proxy.$refs['ruleForm'].validate().then((value) => {
   if (value) {
   // alert('测试')
   console.log(state.searchFormAdd);

            createMailTemplateApi(state.searchFormAdd).then((res) => {
                console.log(res);
                ElMessage.success('创建成功');
                state.editDialogDetailAdd = false;
                getSenderInfoList()
            });
   }
   });
}
// function uphandleEdit(index, row) {
//   uphandleEdit({
//     id: row.id
//   }).then(res => {
//     ElMessage.success("编辑成功")
//   }).catch(err => { console.log(err) })
// }
function updatesender(row){
  postsenderupdateOneApi(state.searchFormUpdate).then((res) => {
         console.log(res);
         getSenderInfoList();
         state.editDialogDetailupdate = false;
        
    });

}
</script>

<style lang="scss" scoped>
</style>
