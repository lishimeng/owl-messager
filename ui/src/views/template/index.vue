<template>
  <div class="app-container">
    <el-button type="primary" icon="el-icon-plus" size="small" @click="addtemplate">新增</el-button>
    <el-table :data="state.tableData.rows" stripe border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80px" />
      <el-table-column prop="name" label="模板名称"></el-table-column> 
      <el-table-column prop="code" label="模板" width="150px" />
      
      <!-- <el-table-column prop="templateBody" label="正文" /> -->

      <el-table-column align='center' prop="body" label="正文" min-width="180px" 
      :show-overflow-tooltip="true"/>


       <el-table-column prop="category" label="分类" width="80px"></el-table-column> 
       <el-table-column prop="description" label="描述"></el-table-column> 
     
      <el-table-column prop="status" label="状态" width="80px">
        <template #default="scope">
          <el-tag v-if="scope.row.status === 1" type="success" size="mini" disable-transitions>已启用</el-tag>
          <el-tag v-else-if="scope.row.status === 2" type="danger" size="mini" disable-transitions>已停用</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="createTime" label="创建时间" width="150px" />
      <el-table-column prop="updateTime" label="更新时间" width="150px" />

      <!-- <el-table-column prop="" label="操作" width="180px">
        <template slot-scope="">
					<el-button size="mini" type="default" icon="el-icon-edit" @click="edittemplate(scope.row)">编辑</el-button>
					<el-button size="mini" type="danger" icon="el-icon-edit" @click="deltemplate(scope.$index, scope.row)">删除</el-button>
				</template>
      </el-table-column>  -->

      <el-table-column fixed="right" label="操作" width="150">
      <template v-slot="scope">
              <el-button type="text" size="small" @click="edittemplate(scope.row)">修改</el-button>
              <el-popconfirm :title="'删除'" @confirm="deltemplate(scope.row)">
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
  </div>
  <!-- 新增弹框 -->
  <el-dialog destroy-on-close :title="'添加'" v-model="state.editDialogDetailAdd" center width="50%">
       <el-form ref="ruleForm" :model="state.searchFormAdd" :rules="state.rules" label-width="220px">
                <!-- <el-form-item label="Key" prop="value">
                 <el-input clearable v-model="state.searchFormAdd.value"></el-input>
               </el-form-item> -->
               <el-form-item label="模板名称" prop="name">
                <el-input clearable v-model="state.searchFormAdd.name" ></el-input>
            </el-form-item>

                <el-form-item label="模板" prop="code">
                <el-input clearable v-model="state.searchFormAdd.code" ></el-input>
              </el-form-item>

                <el-form-item label="正文" prop="body">
                <el-input 
                type="textarea"  :maxlength="9999" :autosize="{ minRows: 5, maxRows: 5}"
                clearable v-model="state.searchFormAdd.body" ></el-input>
              </el-form-item>

              <!-- <el-form-item label="分类" prop="category">
                <el-input clearable v-model="state.searchFormAdd.category" ></el-input>
              </el-form-item> -->
              <el-form-item label="分类" prop="category">
                  <el-select clearable v-model="state.searchFormAdd.status" placeholder="分类">
                    <el-option v-for="item in state.itemCategory" :label="item.name" :value="item.id" :key="item.id"></el-option>
                  </el-select>
               </el-form-item>

                <el-form-item label="描述" prop="description">
                <el-input clearable v-model="state.searchFormAdd.description" ></el-input>
            </el-form-item>
            <!-- <el-form-item label="状态" prop="status">
                <el-input clearable v-model="state.searchFormAdd.status" ></el-input>
            </el-form-item> -->
              <el-form-item label="状态" prop="status">
                  <el-select clearable v-model="state.searchFormAdd.status" placeholder="状态">
                    <el-option v-for="item in state.itemStatus" :label="item.name" :value="item.id" :key="item.id"></el-option>
                  </el-select>
               </el-form-item>
               <el-form-item prop="createTime" label="创建时间">
                   <el-date-picker v-model="state.searchFormAdd.createTime" placeholder="创建时间"  size="small"
                    :clearable="true">
                  </el-date-picker>
                </el-form-item>
                <el-form-item prop="updateTime" label="更新时间">
                   <el-date-picker v-model="state.searchFormAdd.updateTime" placeholder="修改时间"  size="small"
                    :clearable="true">
                  </el-date-picker>
                </el-form-item>

               
                <!-- <el-form-item label="创建时间" prop="createTime">
                <el-input clearable v-model="state.searchFormAdd.createTime" ></el-input>
            </el-form-item>
                <el-form-item label="更新时间" prop="updateTime">
                <el-input clearable v-model="state.searchFormAdd.updateTime" ></el-input>
            </el-form-item> -->
               
           
            <!-- <el-form-item label="操作" prop="operation">
                <el-input clearable v-model="state.searchFormAdd.operation" ></el-input>
            </el-form-item> -->
               <el-form-item>
              <el-button @click="state.editDialogDetailAdd = false">取消</el-button>
                <el-button type="primary" @click="deters()">确定</el-button>
              </el-form-item>
        </el-form>
        </el-dialog>
  <!-- 修改弹框 -->
  <el-dialog destroy-on-close :title="'修改'" v-model="state.editDialogDetailupdate" center width="50%">
       <el-form ref="ruleForm" :model="state.searchFormupdate" :rules="state.rules" label-width="220px">
                <!-- <el-form-item label="Key" prop="value">
                 <el-input clearable v-model="state.searchFormAdd.value"></el-input>
               </el-form-item> -->
               <el-form-item label="模板名称" prop="name">
                <el-input clearable v-model="state.searchFormupdate.name" ></el-input>
            </el-form-item>

                <el-form-item label="模板" prop="code">
                <el-input clearable v-model="state.searchFormupdate.code" ></el-input>
              </el-form-item>

                <el-form-item label="正文" prop="body">
                <el-input 
                type="textarea"  :maxlength="9999" :autosize="{ minRows: 5, maxRows: 5}"
                clearable v-model="state.searchFormupdate.body" ></el-input>
              </el-form-item>

              <!-- <el-form-item label="分类" prop="category">
                <el-input clearable v-model="state.searchFormAdd.category" ></el-input>
              </el-form-item> -->
              <el-form-item label="分类" prop="category">
                  <el-select clearable v-model="state.searchFormupdate.category" placeholder="分类">
                    <el-option v-for="item in state.itemCategory" :label="item.name" :value="item.id" :key="item.id"></el-option>
                  </el-select>
               </el-form-item>

                <el-form-item label="描述" prop="description">
                <el-input clearable v-model="state.searchFormupdate.description" ></el-input>
            </el-form-item>
            <!-- <el-form-item label="状态" prop="status">
                <el-input clearable v-model="state.searchFormAdd.status" ></el-input>
            </el-form-item> -->
              <el-form-item label="状态" prop="status">
                  <el-select clearable v-model="state.searchFormupdate.status" placeholder="状态">
                    <el-option v-for="item in state.itemStatus" :label="item.name" :value="item.id" :key="item.id"></el-option>
                  </el-select>
               </el-form-item>
               <el-form-item prop="createTime" label="创建时间">
                   <el-date-picker v-model="state.searchFormupdate.createTime" placeholder="创建时间"  size="small"
                    :clearable="true">
                  </el-date-picker>
                </el-form-item>
                <el-form-item prop="updateTime" label="更新时间">
                   <el-date-picker v-model="state.searchFormupdate.updateTime" placeholder="修改时间"  size="small"
                    :clearable="true">
                  </el-date-picker>
                </el-form-item>

               
                <!-- <el-form-item label="创建时间" prop="createTime">
                <el-input clearable v-model="state.searchFormAdd.createTime" ></el-input>
            </el-form-item>
                <el-form-item label="更新时间" prop="updateTime">
                <el-input clearable v-model="state.searchFormAdd.updateTime" ></el-input>
            </el-form-item> -->
               
           
            <!-- <el-form-item label="操作" prop="operation">
                <el-input clearable v-model="state.searchFormAdd.operation" ></el-input>
            </el-form-item> -->
               <el-form-item>
              <el-button @click="state.editDialogDetailAdd = false">取消</el-button>
                <el-button type="primary" @click="detertemplate()">确定</el-button>
              </el-form-item>
        </el-form>
        </el-dialog>



</template>
  

<script setup>
import { reactive, onMounted,getCurrentInstance } from "vue"
import { getMailTemplateApi,createMaiTemplateMailApi,gettemplatelistOneApi,posttemplateupdateOneApi,deltemplatelistOneApi} from '/@/api/index'
import { ElMessage, ElMessageBox } from "element-plus";
const { proxy } = getCurrentInstance();

const state = reactive({
  editDialogDetailAdd:false,
  editDialogDetailupdate:false,
  searchFormAdd:{},
  searchFormupdate:{},
  itemStatus:[
    {'id':1,'name':'enable'},
      {'id':0,'name':'disable'},
],
itemCategory:[
     {'id':2,'name':'2'},
      {'id':1,'name':'1'},
],
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
  getMailTemplate()
})
function addtemplate(row){
  state.editDialogDetailAdd=true
}

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
    state.tableData.pagination.total=res.more
  }).catch(()=>{})
}
function edittemplate(row){
  gettemplatelistOneApi({ id:  row.id}).then((res) => {
   console.log(res);
   state.editDialogDetailupdate = true;
   state.searchFormupdate = res.data;
   });
}

function deters(row){
  proxy.$refs['ruleForm'].validate().then((value) => {
   if (value) {
   // alert('测试')
   console.log(state.searchFormAdd);
         gettemplatelistOneApi(state.searchFormAdd).then((res) => {
             console.log(res);
             ElMessage.success('创建成功');
             state.editDialogDetailAdd = false;
            getMailTemplate()
          });
   }
   });
}
function detertemplate(row){
  posttemplateupdateOneApi(state.searchFormupdate).then((res) => {
         console.log(res);
         getMailTemplate()
         state.editDialogDetailupdate = false;
        
    });

}
function deltemplate(row){
  deltemplatelistOneApi({ id:  row.id}).then((res) => {
      console.log(res);
      getMailTemplate()
      ElMessage.success("删除成功")
 });

}
</script>

<style lang="scss" scoped>
</style>