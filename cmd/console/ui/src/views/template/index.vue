<template>
  <div class="home-container layout-pd" style="margin-top: 10px">
    <el-form :inline="true">
      <el-form-item label="通讯方式">
        <el-select v-model="state.category" @change="chooseCategory" placeholder="请选择通讯方式" style="width: 120px">
          <el-option v-for="(item,index) in state.categoryList" :key="index" :label="item" :value="item">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="openCreate" icon="ele-CirclePlus">新增</el-button>
      </el-form-item>
    </el-form>
    <div style="margin-top: 10px">
      <el-table :data="state.dataList" border style="width: 100%">
        <el-table-column prop="id" label="Id" width="60"/>
        <el-table-column prop="code" label="Code" width="180" show-overflow-tooltip/>
        <el-table-column prop="name" label="模板名称" width="120"/>
        <el-table-column prop="description" label="模板描述" width="180" show-overflow-tooltip/>
        <el-table-column prop="createTime" label="创建时间">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.createTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column prop="UpdateTime" label="更新时间">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.updateTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="140" fixed="right">
          <template #default="scope">
            <el-button size="small" type="text" @click="openEdit(scope.row.code)">
              编辑
            </el-button>
            <el-popconfirm
                title="是否删除?"
                @confirm="deleteDeviceRow(scope.row.code)">
              <template #reference>
                <el-button size="small" type="text">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin: 10px 10px 10px 10px;">
        <el-pagination background
                       layout="sizes, prev, pager, next, jumper"
                       v-model:currentPage="state.queryValue.pageNum"
                       v-model:page-size="state.queryValue.pageSize"
                       :page-sizes="[5,10,15,30,50,100]"
                       :page-count="state.queryValue.totalNum"
                       @size-change="onSizeChange"
                       @current-change="onCurrentChange"
        />
      </div>
    </div>
    <el-drawer
        :title="'创建模版'"
        v-model="state.showCreateDrawer"
        direction="rtl"
        :size="1000"
        :before-close="handleClose">
      <TemplateForm
          :category="state.category"
          ref="createRef"
      ></TemplateForm>
      <template #footer>
        <div style="margin: 50px">
          <el-button type="primary" @click="onSubmit">提交</el-button>
        </div>
      </template>
    </el-drawer>
    <el-drawer
        :title="'编辑模版'"
        v-model="state.showEditDrawer"
        direction="rtl"
        :destroy-on-close="true"
        :size="1000"
        :before-close="handleClose">
      <TemplateForm
          :category="state.category"
          :template-code="state.selected"
          ref="editRef"
      ></TemplateForm>
      <template #footer>
        <div style="margin: 50px">
          <el-button type="primary" @click="onEditSubmit">保存</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts" name="Template">
// 引入组件
import {delTemplateApi, getTemplateListAPi} from "/@/api/template";
import {onMounted, reactive, ref} from 'vue';
import {formatDate} from "../../utils/formatTime";
import TemplateForm from "/@/views/template/templateForm.vue";
import {ElMessage} from "element-plus";

const createRef = ref();
const editRef = ref();
const state = reactive({
  showCreateDrawer: false,
  showEditDrawer: false,
  title: '新增模版',
  category: 'mail',
  categoryList: [
    "mail",
    "sms",
    "im"
  ],
  dataList: [],
  queryValue: {
    pageNum: 1,
    pageSize: 10,
    totalNum: 0,
    category: ''
  },
  // 新增模版
  subForm: {
    code: "",
    name: "",
    description: "",
    body: "",
    vendor: "",
    templateId: "",
    params: "",
    signature: "",
    sender: 0,
    category: "",
  },
  getHtml: "",
  getText: "",
  editorVal: "",
  selected:"",
})
onMounted(() => {
  chooseCategory()
})

const onSubmit = async () => {
  const success = await createRef.value.onSubmit();
  if (success) {
    state.showCreateDrawer = false;
    getTemplateList();
  }
}

const onEditSubmit = async () => {
  const success = await editRef.value.onSubmit();
  console.log(success);
  if (success) {
    state.showEditDrawer = false;
    getTemplateList();
  }
}

const getTemplateList = () => {
  state.queryValue.category = state.category
  getTemplateListAPi(state.queryValue).then(res => {
    if (res.items && res.items.length > 0) {
      state.dataList = res.items
      state.queryValue.totalNum = res.totalPage
    } else {
      state.dataList = []
      state.queryValue.totalNum = 0
    }
  })
}
const handleClose = (done: any) => {
  // state.showCreateDrawer = false
  done();
}
const onSizeChange = (val: any) => {
  state.queryValue.pageSize = val
  getTemplateList();
}
const onCurrentChange = (val: any) => {
  state.queryValue.pageNum = val
  getTemplateList();
}
const chooseCategory = () => {
  getTemplateList();
  createRef.value?.resetForm(state.category);
}
const openCreate = () => {
  state.showCreateDrawer = true
};

const openEdit = (code:string) => {
  state.selected = code;
  state.showEditDrawer = true
};

const deleteDeviceRow = (code:string) => {
  delTemplateApi({
    code:code,
  })
      .then((res) => {
        console.log(res);
        ElMessage.success('删除成功');
        getTemplateList();
      })
      .catch((err) => {
        console.log(err);
      });
}

</script>

<style scoped lang="scss">
.input_width {
  width: 100%;
}
</style>