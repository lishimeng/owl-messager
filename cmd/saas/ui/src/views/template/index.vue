<template>
  <div class="home-container layout-pd">
    <el-form :inline="true">
      <el-form-item label="通讯方式">
        <el-select v-model="state.category" @change="chooseCategory" placeholder="请选择通讯方式">
          <el-option v-for="(item,index) in state.categoryList" :key="index" :label="item" :value="item">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="showEdit()" icon="ele-CirclePlus">新增</el-button>
      </el-form-item>
    </el-form>
    <div style="margin-top: 10px">
      <el-table :data="state.dataList" border style="width: 100%">
        <el-table-column prop="Id" label="Id" width="60"/>
        <el-table-column prop="Code" label="Code" width="360"/>
        <el-table-column prop="Name" label="模板名称" width="120"/>
        <el-table-column prop="Description" label="模板描述" width="120"/>
        <el-table-column prop="CreateTime" label="创建时间">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.CreateTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column prop="UpdateTime" label="更新时间">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.UpdateTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button icon="ele-Edit" type="primary" @click="showEdit(scope.row)">
              编辑
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin: 10px 10px 10px 10px;">
        <el-pagination background
                       layout="total, sizes, prev, pager, next, jumper"
                       v-model:currentPage="state.queryValue.pageNum"
                       v-model:page-size="state.queryValue.pageSize"
                       :page-sizes="[5,10,15,30,50,100]"
                       :page-count="state.queryValue.totalNum"
                       :total="state.queryValue.totalNum"
                       @size-change="onSizeChange"
                       @current-change="onCurrentChange"
        />
      </div>
    </div>
    <el-drawer
        :title="state.title"
        v-model="state.showDrawer"
        direction="rtl"
        :size="1000"
        :before-close="handleClose">
      <div style="width:100%;padding: 10px">
        <el-form style="margin-top: 20px"
                 :model="state.subForm"
                 ref="mailFormFormRef"
                 label-width="120px">
          <el-form-item label="配置平台" prop="vendor">
            <el-select class="input_width"
                       v-model="state.subForm.vendor"
                       placeholder="请选择通讯方式">
              <el-option v-for="(item,index) in state.vendors" :key="index" :label="item" :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item v-if="state.subForm.code.length>0" label="code" prop="code">
            <el-input v-model="state.subForm.code" :disabled="state.subForm.code.length>0" clearable></el-input>
          </el-form-item>
          <el-form-item label="名称" prop="name">
            <el-input v-model="state.subForm.name" clearable></el-input>
          </el-form-item>
          <el-form-item label="描述" prop="description">
            <el-input v-model="state.subForm.description" clearable></el-input>
          </el-form-item>
          <el-form-item v-if="state.category=='sms'" label="第三方模板ID" prop="templateId">
            <el-input v-model="state.subForm.templateId" clearable></el-input>
          </el-form-item>
          <el-form-item v-if="state.category=='sms'" label="第三方模板签名" prop="signature">
            <el-input type="textarea" v-model="state.subForm.signature" clearable></el-input>
          </el-form-item>
          <el-form-item v-if="state.category=='sms'" label="指定发送平台" prop="sender">
            <el-input type="number" v-model="state.subForm.sender" clearable></el-input>
          </el-form-item>
          <el-form-item label="模板内容">
            <wngEditor mode="default" height="300px" v-model:getHtml="state.getHtml"
                       v-model:getText="state.getText"></wngEditor>
          </el-form-item>
          <el-form-item>
            <el-input disabled style="width: 100%;height: 400px" type="textarea"
                      v-model="state.subForm.body"></el-input>
          </el-form-item>
        </el-form>

      </div>
      <template #footer>
        <div style="margin: 50px">
          <el-button type="primary" @click="onSubmit()">提交</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts" name="Template">
// 引入组件
import {createTemplateApi, getTemplateInfoAPi, getTemplateListAPi, updateTemplateApi} from "/@/api/template";
import {defineAsyncComponent, onMounted, reactive, ref, watch} from 'vue';
import {formatDate} from "../../utils/formatTime";
import {ElMessage} from "element-plus";

const JsEditor = defineAsyncComponent(() => import('/@/components/js/index.vue'))
const wngEditor = defineAsyncComponent(() => import('/@/components/editor/index.vue'));
const mailFormFormRef = ref()
const state = reactive({
  showDrawer: false,
  title: '新增模版',
  category: 'mail',
  categoryList: [
    "mail",
    "sms"
  ],
  dataList: [],
  queryValue: {
    pageNum: 1,
    pageSize: 10,
    totalNum: 0,
    category: ''
  },
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
  vendors: [],
  mailVendors: [
    "smtp",
    "microsoft",
    "tencent"
  ],
  smsVendors: [
    "ali_yun",
    "tencent_yun",
    "huawei_yun",
  ],
  getHtml: "",
  getText: "",
  editorVal: "",
})
onMounted(() => {
  chooseCategory()
})
watch(() => state.getHtml, (newVal, oldVal) => {
  console.log('监听：', newVal)
  state.subForm.body = "<html>" +
      "<head>" +
      "<meta charset=\"utf-8\">" +
      "</head>" +
      "<body>" + newVal +
      "</body>" +
      "</html>"
  console.log(state.subForm.body)
})
const onSubmit = () => {
  state.subForm.category = state.category
  // console.log(state.getText,state.getHtml)
  if (state.getText) {
    state.subForm.body = "<html>" +
        "<head>" +
        "<meta charset=\"utf-8\">" +
        "</head>" +
        "<body>" + state.getHtml +
        "</body>" +
        "</html>"
  } else {
    state.subForm.body = state.getText
  }
  if (state.subForm.sender) {
    state.subForm.sender = parseInt(state.subForm.sender)
  }
  if (state.subForm.code) {
    // console.log("编辑")
    updateTemplateApi(state.subForm).then(res => {
      if (res.code && res.code == 200) {
        ElMessage.success(`提交成功！`);
        state.showDrawer = false
        getTemplateList();
      }
    })
  } else {
    // console.log("新增")
    createTemplateApi(state.subForm).then(res => {
      if (res.code && res.code == 200) {
        ElMessage.success(`提交成功！`);
        state.showDrawer = false
        getTemplateList();
      }
    })
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
  state.showDrawer = false
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
  switch (state.category) {
    case "mail":
      state.vendors = state.mailVendors
      break;
    case "sms":
      state.vendors = state.smsVendors
      break;
  }

}
const showEdit = (data: any) => {
  state.showDrawer = true
  if (data) {
    // console.log("编辑模版")
    state.title = "编辑模版"
    getTemplateInfoAPi({
      code: data.Code,
      category: state.category
    }).then(res => {
      if (res.code && res.code == 200) {
        state.subForm = res.item
        state.getHtml = res.item.body.replace("<html>", "")
            .replace("</html>", "")
            .replace("<head>", "")
            .replace("</head>", "")
            .replace("<body>", "")
            .replace("</body>", "")
            .replace("<meta charset=\"utf-8\">", "")
      }
    })
  } else {
    // console.log("新增模版")
    mailFormFormRef.value.resetFields();
    state.subForm.code = ""
    state.title = "新增模版"
    state.getText = ""
    state.getHtml = ""
    state.subForm.body = ""
  }

}
</script>

<style scoped lang="scss">
.input_width {
  width: 100%;
}
</style>