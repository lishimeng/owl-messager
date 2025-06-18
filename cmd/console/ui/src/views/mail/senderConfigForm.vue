<template>
  <div>
    <!--    smtp      -->
    <div v-show="props.vendor=='smtp'">
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
    <div v-show="props.vendor=='microsoft'">
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
    <div v-show="props.vendor=='tencent'">
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
    <div v-show="props.vendor=='ali_yun'">
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
    <div v-show="props.vendor=='tencent_yun'">
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
    <div v-show="props.vendor=='huawei_yun'">
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
</template>

<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";

const smtpFormRef = ref();
const microsoftFormRef = ref();
const tencentFormRef = ref();
const aliYunFormRef = ref();
const tencentYunFormRef = ref();
const huaweiYunFormRef = ref();

const props = defineProps({
  vendor: {
    type: String,
    required: true,
  },
  config: {
    type: String,
    required: true,
  },
});

onMounted(() => {
  if (props.config) {
    loadJson(props.vendor, props.config);
  }
})

const state = reactive({
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
});

const exportJson = () => {
  switch (props.vendor) {
    case 'smtp':
      return JSON.stringify(state.smtp);
    case 'microsoft':
      return JSON.stringify(state.microsoft);
    case 'tencent':
      return JSON.stringify(state.tencent);
    case 'ali_yun':
      return JSON.stringify(state.aliYun);
    case 'tencent_yun':
      return JSON.stringify(state.tencentYun);
    case 'huawei_yun':
      return JSON.stringify(state.huaweiYun);
  }
};

const loadJson = (vendor: string, config: string) => {
  switch (vendor) {
    case 'smtp':
      state.smtp = JSON.parse(config)
      break;
    case 'microsoft':
      state.microsoft = JSON.parse(config)
      break;
    case 'tencent':
      state.tencent = JSON.parse(config)
      break;
    case 'ali_yun':
      state.aliYun = JSON.parse(config)
      break;
    case 'tencent_yun':
      state.tencentYun = JSON.parse(config)
      break;
    case 'huawei_yun':
      state.huaweiYun = JSON.parse(config)
      break;
  }
}

const clearForm = () => {
  smtpFormRef.value?.resetFields();
  microsoftFormRef.value?.resetFields();
  tencentFormRef.value?.resetFields();
  aliYunFormRef.value?.resetFields();
  tencentYunFormRef.value?.resetFields();
  huaweiYunFormRef.value?.resetFields();
};

defineExpose({
  exportJson,
  loadJson,
  clearForm,
})

</script>

<style scoped lang="scss">

</style>