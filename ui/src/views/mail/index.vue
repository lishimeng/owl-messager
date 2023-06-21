<template>
  <div class="home-container layout-pd">
    <el-form :inline="true">
      <el-form-item label="通讯方式">
        <el-select v-model="state.category" @change="getMailSenders" placeholder="请选择通讯方式">
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
        <el-table-column prop="Id" label="id" width="60"/>
        <el-table-column prop="Code" label="code" width="360"/>
        <el-table-column prop="Default" label="Default" width="120"/>
        <el-table-column prop="Vendor" label="Vendor" width="120"/>
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
    <el-dialog v-model="state.showDialog" :title="state.title" width="50%" center>
      <div style="margin: 10px">
        <el-form :model="state.form" label-width="120px">
          <el-form-item label="配置平台" prop="vendor">
            <el-select class="input_width" :disabled="state.isDisabled"
                       v-model="state.form.vendor"
                       @change="getMailSenderInfo"
                       placeholder="请选择通讯方式">
              <el-option v-for="(item,index) in state.vendors" :key="index" :label="item" :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item v-if="state.isShowText" label="CODE" prop="code">
            <el-input :disabled="state.isDisabled" v-model="state.form.code"></el-input>
          </el-form-item>
          <el-form-item label="默认发送人" prop="defaultSender">
            <el-radio-group v-model="state.form.defaultSender" class="ml-4">
              <el-radio :label="1">是</el-radio>
              <el-radio :label="0">否</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-form>
        <!--    smtp      -->
        <div v-show="state.form.vendor=='smtp'">
          <el-form ref="smtpFormRef"
                   style="margin-top: 20px"
                   :model="state.smtp"
                   label-width="120px">
            <el-form-item label="host" prop="host">
              <el-input v-model="state.smtp.host" clearable></el-input>
            </el-form-item>
            <el-form-item label="port" prop="port">
              <el-input v-model="state.smtp.port" type="number" clearable></el-input>
            </el-form-item>
            <el-form-item label="senderEmail" prop="senderEmail">
              <el-input v-model="state.smtp.senderEmail" clearable></el-input>
            </el-form-item>
            <el-form-item label="senderAlias" prop="senderAlias">
              <el-input v-model="state.smtp.senderAlias" clearable></el-input>
            </el-form-item>
            <el-form-item label="authUser" prop="authUser">
              <el-input v-model="state.smtp.authUser" clearable></el-input>
            </el-form-item>
            <el-form-item label="authPass" prop="authPass">
              <el-input v-model="state.smtp.authPass" clearable></el-input>
            </el-form-item>
          </el-form>
        </div>
        <!--    microsoft      -->
        <div v-show="state.form.vendor=='microsoft'">
          <el-form style="margin-top: 20px"
                   :model="state.microsoft"
                   ref="microsoftFormRef"
                   label-width="120px">
            <el-form-item label="clientId" prop="clientId">
              <el-input v-model="state.microsoft.clientId" clearable></el-input>
            </el-form-item>
            <el-form-item label="tenant" prop="tenant">
              <el-input v-model="state.microsoft.tenant" clearable></el-input>
            </el-form-item>
            <el-form-item label="scope" prop="scope">
              <el-input v-model="state.microsoft.scope" clearable></el-input>
            </el-form-item>
            <el-form-item label="sender" prop="sender">
              <el-input v-model="state.microsoft.sender" clearable></el-input>
            </el-form-item>
            <el-form-item label="certificate" prop="certificate">
              <el-input v-model="state.microsoft.certificate" :rows="4"
                        type="textarea" clearable></el-input>
            </el-form-item>
            <el-form-item label="certificateKey" prop="certificateKey">
              <el-input v-model="state.microsoft.certificateKey" :rows="4"
                        type="textarea" clearable></el-input>
            </el-form-item>
          </el-form>
        </div>
        <!--    tencent      -->
        <div v-show="state.form.vendor=='tencent'">
          <el-form style="margin-top: 20px"
                   :model="state.tencent"
                   ref="tencentFormRef"
                   label-width="120px">
            <el-form-item label="appId" prop="appId">
              <el-input v-model="state.tencent.appId" clearable></el-input>
            </el-form-item>
            <el-form-item label="secret" prop="secret">
              <el-input v-model="state.tencent.secret" clearable></el-input>
            </el-form-item>
            <el-form-item label="region" prop="region">
              <el-input v-model="state.tencent.region" clearable></el-input>
            </el-form-item>
            <el-form-item label="sender" prop="sender">
              <el-input v-model="state.tencent.sender" clearable></el-input>
            </el-form-item>
          </el-form>
        </div>
        <!--        阿里云sms ali_yun-->
        <div v-show="state.form.vendor=='ali_yun'">
          <el-form style="margin-top: 20px"
                   :model="state.aliYun"
                   ref="aliYunFormRef"
                   label-width="120px">
            <el-form-item label="appKey" prop="appKey">
              <el-input v-model="state.aliYun.appKey" clearable></el-input>
            </el-form-item>
            <el-form-item label="appSecret" prop="appSecret">
              <el-input v-model="state.aliYun.appSecret" clearable></el-input>
            </el-form-item>
            <el-form-item label="region" prop="region">
              <el-input v-model="state.aliYun.region" clearable></el-input>
            </el-form-item>
            <el-form-item label="signName" prop="signName">
              <el-input v-model="state.aliYun.signName" clearable></el-input>
            </el-form-item>
          </el-form>
        </div>
        <!--        腾讯云sms tencent_yun-->
        <div v-show="state.form.vendor=='tencent_yun'">
          <el-form style="margin-top: 20px"
                   :model="state.tencentYun"
                   ref="tencentYunFormRef"
                   label-width="120px">
            <el-form-item label="appId" prop="appId">
              <el-input v-model="state.tencentYun.appId" clearable></el-input>
            </el-form-item>
            <el-form-item label="appKey" prop="appKey">
              <el-input v-model="state.tencentYun.appKey" clearable></el-input>
            </el-form-item>
            <el-form-item label="smsAppId" prop="smsAppId">
              <el-input v-model="state.tencentYun.smsAppId" clearable></el-input>
            </el-form-item>
            <el-form-item label="region" prop="region">
              <el-input v-model="state.tencentYun.region" clearable></el-input>
            </el-form-item>
            <el-form-item label="signName" prop="signName">
              <el-input v-model="state.tencentYun.signName" clearable></el-input>
            </el-form-item>
          </el-form>
        </div>
        <!--        华为云sms huawei_yun-->
        <div v-show="state.form.vendor=='huawei_yun'">
          <el-form style="margin-top: 20px"
                   :model="state.huaweiYun"
                   ref="huaweiYunFormRef"
                   label-width="120px">
            <el-form-item label="host" prop="host">
              <el-input v-model="state.huaweiYun.host" clearable></el-input>
            </el-form-item>
            <el-form-item label="appId" prop="appId">
              <el-input v-model="state.huaweiYun.appId" clearable></el-input>
            </el-form-item>
            <el-form-item label="appKey" prop="appKey">
              <el-input v-model="state.huaweiYun.appKey" clearable></el-input>
            </el-form-item>
            <el-form-item label="sender" prop="sender">
              <el-input v-model="state.huaweiYun.sender" clearable></el-input>
            </el-form-item>
            <el-form-item label="signName" prop="signName">
              <el-input v-model="state.huaweiYun.signName" clearable></el-input>
            </el-form-item>
          </el-form>
        </div>
      </div>
      <span class="dialog-footer">
        <el-button @click="state.showDialog = false">取消</el-button>
        <el-button type="primary" @click="onSubmit()">提交</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script setup lang="ts" name="Mail">
import {onMounted, reactive, ref} from 'vue';
import {
  createMailSenderConfigApi,
  getMailSendersApi,
  getSenderInfoByCategoryAPi,
  updateMailSenderConfigApi
} from "/@/api/mail";
import {ElMessage} from "element-plus";
import {formatDate} from "/@/utils/formatTime";

const formRef = ref()
const smtpFormRef = ref()
const microsoftFormRef = ref()
const tencentFormRef = ref()
const aliYunFormRef = ref()
const tencentYunFormRef = ref()
const huaweiYunFormRef = ref()

const state = reactive({
  isShowText: true,
  isDisabled: false,
  title: "新增",
  showDialog: false,
  queryValue: {
    pageSize: 10,
    pageNum: 1,
    totalNum: 0,
    category: ""
  },
  dataList: [],
  category: "mail",
  categoryList: [
    "mail",
    "sms"
  ],
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
  form: {
    code: "",
    defaultSender: 1,
    vendor: "",
    config: "",
    category: "",
  },
  smtp: {
    host: "",
    port: 0,
    senderEmail: "",
    senderAlias: "",
    authUser: "",
    authPass: ""
  },
  microsoft: {
    clientId: "",
    tenant: "",
    scope: "",
    sender: "",
    certificate: "",
    certificateKey: ""
  },
  tencent: {
    appId: "",
    secret: "",
    region: "",
    sender: ""
  },
  aliYun: {
    appKey: "",
    appSecret: "",
    region: "",
    signName: ""
  },
  tencentYun: {
    appId: "",
    appKey: "",
    smsAppId: "",
    region: "",
    signName: ""
  },
  huaweiYun: {
    host: "",
    appId: "",
    appKey: "",
    sender: "",
    signName: ""
  },
})
onMounted(() => {
  getMailSenders();
})
const showEdit = (row: object) => {
  state.showDialog = true
  // console.log(row)
  if (row) {
    state.title = '编辑'
    state.isDisabled = true
    state.isShowText = true
    getSenderInfoByCategory(row.Code)
  } else {
    state.isDisabled = false
    state.isShowText = false
    state.title = '新增'
    state.form.vendor = ''
  }
}
const onSubmit = () => {
  switch (state.form.vendor) {
    case 'smtp':
      state.smtp.port = parseInt(state.smtp.port)
      state.form.config = JSON.stringify(state.smtp)
      break;
    case 'microsoft':
      state.form.config = JSON.stringify(state.microsoft)
      break;
    case 'tencent':
      state.form.config = JSON.stringify(state.tencent)
      break;
    case 'ali_yun':
      state.form.config = JSON.stringify(state.aliYun)
      break;
    case 'tencent_yun':
      state.form.config = JSON.stringify(state.tencentYun)
      break;
    case 'huawei_yun':
      state.form.config = JSON.stringify(state.huaweiYun)
      break;
  }
  state.form.category = state.category
  createConfig();
}
const createConfig = () => {
  if (state.form.code) {
    console.log("编辑")
    updateMailSenderConfigApi(state.form).then(res => {
      if (res && res.code == 200) {
        ElMessage.success(`提交成功！`);
        state.showDialog = false
        getMailSenders()
      }
    })
  } else {
    console.log("新增")
    createMailSenderConfigApi(state.form).then(res => {
      if (res && res.code == 200) {
        ElMessage.success(`提交成功！`);
        state.showDialog = false
        getMailSenders()
      }
    })
  }

}
const getMailSenderInfo = () => {
  state.form.code = ''
  state.form.defaultSender = 1
  state.form.config = ''
  switch (state.form.vendor) {
    case 'smtp':
      smtpFormRef.value.resetFields();
      break;
    case 'microsoft':
      microsoftFormRef.value?.resetFields();
      break;
    case 'tencent':
      tencentFormRef.value?.resetFields();
      break;
    case 'ali_yun':
      aliYunFormRef.value?.resetFields();
      break;
    case 'tencent_yun':
      tencentYunFormRef.value?.resetFields();
      break;
    case 'huawei_yun':
      huaweiYunFormRef.value?.resetFields();
      break;
  }
}
const getSenderInfoByCategory = (code: string) => {
  getSenderInfoByCategoryAPi({
    category: state.category,
    code: code
  }).then(res => {
    if (res.item) {
      state.form.vendor = res.item.Vendor
      state.form.code = res.item.Code
      state.form.defaultSender = res.item.Default
      state.form.config = res.item.Config
      switch (res.item.Vendor) {
        case 'smtp':
          state.smtp = JSON.parse(res.item.Config)
          break;
        case 'microsoft':
          state.microsoft = JSON.parse(res.item.Config)
          break;
        case 'tencent':
          state.tencent = JSON.parse(res.item.Config)
          break;
        case 'ali_yun':
          state.aliYun = JSON.parse(res.item.Config)
          break;
        case 'tencent_yun':
          state.tencentYun = JSON.parse(res.item.Config)
          break;
        case 'huawei_yun':
          state.huaweiYun = JSON.parse(res.item.Config)
          break;
      }
    }
  })
}
const getMailSenders = () => {
  switch (state.category) {
    case 'mail':
      state.vendors = state.mailVendors
      break
    case 'sms':
      state.vendors = state.smsVendors
      break
  }
  state.queryValue.category = state.category
  getMailSendersApi(state.queryValue).then(res => {
    if (res && res.code == 200 && res.items) {
      state.dataList = res.items
      state.queryValue.totalNum = res.totalPage
    } else {
      state.dataList = []
      state.queryValue.totalNum = 0
    }
    // console.log(state.queryValue.totalNum)
  })
}
const onSizeChange = (val: object) => {
  state.queryValue.pageSize = val
  getMailSenders();
}
const onCurrentChange = (val: object) => {
  state.queryValue.pageNum = val
  getMailSenders();
}
</script>

<style scoped lang="scss">
.input_width {
  width: 100%;
}
</style>