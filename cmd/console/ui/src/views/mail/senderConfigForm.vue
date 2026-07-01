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
    <!--    tencent mail (SES)      -->
    <div v-show="props.vendor=='tencent' && props.category=='mail'">
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
        <el-form-item label="发件邮箱" prop="senderEmail">
          <el-input v-model="state.tencent.senderEmail" clearable placeholder="noreply@mail.example.com"></el-input>
        </el-form-item>
        <el-form-item label="发件人别名" prop="senderAlias">
          <el-input v-model="state.tencent.senderAlias" clearable placeholder="可选，不能含冒号"></el-input>
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
    <!--    tencent sms      -->
    <div v-show="props.vendor=='tencent' && props.category=='sms'">
      <el-form style="margin-top: 20px"
               :model="state.tencentSms"
               ref="tencentSmsFormRef"
               label-width="120px">
        <el-form-item label="appId" prop="appId">
          <el-input v-model="state.tencentSms.appId" clearable></el-input>
        </el-form-item>
        <el-form-item label="appKey" prop="appKey">
          <el-input v-model="state.tencentSms.appKey" clearable></el-input>
        </el-form-item>
        <el-form-item label="smsAppId" prop="smsAppId">
          <el-input v-model="state.tencentSms.smsAppId" clearable></el-input>
        </el-form-item>
        <el-form-item label="region" prop="region">
          <el-input v-model="state.tencentSms.region" clearable></el-input>
        </el-form-item>
        <el-form-item label="signName" prop="signName">
          <el-input v-model="state.tencentSms.signName" clearable></el-input>
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
    <!--        fastmsg -->
    <div v-show="props.vendor=='fastmsg'">
      <el-form style="margin-top: 20px"
               :model="state.fastmsg"
               ref="fastmsgFormRef"
               label-width="120px">
        <el-form-item label="host" prop="host">
          <el-input v-model="state.fastmsg.host" clearable></el-input>
        </el-form-item>
        <el-form-item label="token" prop="token">
          <el-input v-model="state.fastmsg.token" clearable></el-input>
        </el-form-item>
        <el-form-item label="username" prop="username">
          <el-input v-model="state.fastmsg.username" clearable></el-input>
        </el-form-item>
        <el-form-item label="password" prop="password">
          <el-input v-model="state.fastmsg.password" clearable></el-input>
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
const tencentSmsFormRef = ref();
const huaweiYunFormRef = ref();
const fastmsgFormRef = ref();

const props = defineProps({
  vendor: {
    type: String,
    required: true,
  },
  category: {
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
    senderEmail: "",
    senderAlias: "",
    sender: ""
  },
  aliYun: {
    appKey: "",
    appSecret: "",
    region: "",
    signName: ""
  },
  tencentSms: {
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
  fastmsg: {
    host: "",
    token: "",
    username: "",
    password: "",
  },
});

const resolveVendorKey = (vendor: string, category: string) => {
  if (vendor === 'tencent_yun' || (vendor === 'tencent' && category === 'sms')) {
    return 'tencentSms';
  }
  return vendor;
};

const normalizeTencentMailConfig = (cfg: Record<string, string>) => {
  const next = {
    appId: cfg.appId || "",
    secret: cfg.secret || "",
    region: cfg.region || "",
    senderEmail: cfg.senderEmail || "",
    senderAlias: cfg.senderAlias || "",
    sender: cfg.sender || "",
  };
  if (!next.senderEmail && next.sender) {
    const m = next.sender.match(/^(.+?)\s+<([^>]+)>$/);
    if (m) {
      next.senderAlias = next.senderAlias || m[1].trim();
      next.senderEmail = m[2].trim();
    } else {
      next.senderEmail = next.sender.trim();
    }
  }
  return next;
};

const validateConfig = (): string | null => {
  const key = resolveVendorKey(props.vendor, props.category);
  if (key === 'tencent') {
    const c = state.tencent;
    if (!c.senderEmail?.trim() && !c.sender?.trim()) {
      return '请填写腾讯云发件邮箱（senderEmail）';
    }
    if (c.senderAlias?.includes(':')) {
      return '发件人别名不能包含冒号';
    }
  }
  if (key === 'smtp' && !state.smtp.senderEmail?.trim()) {
    return '请填写 SMTP 发件邮箱';
  }
  return null;
};

const exportJson = () => {
  const err = validateConfig();
  if (err) {
    return null;
  }
  const key = resolveVendorKey(props.vendor, props.category);
  switch (key) {
    case 'smtp':
      return JSON.stringify(state.smtp);
    case 'microsoft':
      return JSON.stringify(state.microsoft);
    case 'tencent':
      return JSON.stringify(state.tencent);
    case 'tencentSms':
      return JSON.stringify(state.tencentSms);
    case 'ali_yun':
      return JSON.stringify(state.aliYun);
    case 'huawei_yun':
      return JSON.stringify(state.huaweiYun);
    case 'fastmsg':
      return JSON.stringify(state.fastmsg);
  }
};

const loadJson = (vendor: string, config: string) => {
  const key = resolveVendorKey(vendor, props.category);
  switch (key) {
    case 'smtp':
      state.smtp = JSON.parse(config)
      break;
    case 'microsoft':
      state.microsoft = JSON.parse(config)
      break;
    case 'tencent':
      state.tencent = normalizeTencentMailConfig(JSON.parse(config))
      break;
    case 'tencentSms':
      state.tencentSms = JSON.parse(config)
      break;
    case 'ali_yun':
      state.aliYun = JSON.parse(config)
      break;
    case 'huawei_yun':
      state.huaweiYun = JSON.parse(config)
      break;
    case 'fastmsg':
      state.fastmsg = JSON.parse(config)
      break;
  }
}

const clearForm = () => {
  smtpFormRef.value?.resetFields();
  microsoftFormRef.value?.resetFields();
  tencentFormRef.value?.resetFields();
  aliYunFormRef.value?.resetFields();
  tencentSmsFormRef.value?.resetFields();
  huaweiYunFormRef.value?.resetFields();
  fastmsgFormRef.value?.resetFields();
};

defineExpose({
  exportJson,
  loadJson,
  clearForm,
})

</script>

<style scoped lang="scss">

</style>