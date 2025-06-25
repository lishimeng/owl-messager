<template>
  <div style="width:100%;padding: 10px">
    <el-form style="margin-top: 20px"
             :model="state.formData"
             ref="mailFormFormRef"
             label-width="120px">
      <el-form-item label="配置平台" prop="provider">
        <el-select class="input_width"
                   v-model="state.formData.provider"
                   placeholder="请选择通讯方式">
          <el-option v-for="(item,index) in vendors" :key="index" :label="item" :value="item">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item v-if="state.formData.code.length>0" label="code" prop="code">
        <el-input v-model="state.formData.code" :disabled="state.formData.code.length>0" clearable></el-input>
      </el-form-item>
      <el-form-item label="名称" prop="name">
        <el-input v-model="state.formData.name" clearable></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="state.formData.description" clearable></el-input>
      </el-form-item>
      <el-form-item v-if="state.category=='sms'" label="第三方模板ID" prop="templateId">
        <el-input v-model="state.formData.templateId" clearable></el-input>
      </el-form-item>
      <el-form-item v-if="state.category=='sms'" label="第三方模板签名" prop="signature">
        <el-input type="textarea" v-model="state.formData.signature" clearable></el-input>
      </el-form-item>
      <el-form-item v-if="state.category=='sms'" label="指定发送平台" prop="sender">
        <el-input type="number" v-model="state.formData.sender" clearable></el-input>
      </el-form-item>
      <el-form-item label="模板内容" v-if="htmlTemplate">
        <wngEditor mode="default" height="300px" v-model:getHtml="state.getHtml"
                   v-model:getText="state.getText"></wngEditor>
      </el-form-item>
      <el-form-item v-if="htmlTemplate">
        <el-input
            style="width: 100%"
            type="textarea"
            :autosize="{ minRows: 2, maxRows: 6 }"
            v-model="state.formData.body"
            readonly
        ></el-input>
      </el-form-item>
      <el-form-item label="模板内容" v-else>
        <el-input
            style="width: 100%"
            type="textarea"
            :autosize="{ minRows: 3, maxRows: 20 }"
            v-model="state.formData.body"
        ></el-input>
      </el-form-item>
    </el-form>

  </div>
</template>

<script setup lang="ts">
import {computed, defineAsyncComponent, onMounted, reactive, ref, watch} from "vue";
import {createTemplateApi, getTemplateInfoAPi, updateTemplateApi} from "/@/api/template";
import {ElMessage} from "element-plus";
const mailFormFormRef = ref()
const JsEditor = defineAsyncComponent(() => import('/@/components/js/index.vue'));
const wngEditor = defineAsyncComponent(() => import('/@/components/editor/index.vue'));
const props = defineProps({
  category: {
    type: String,
    required: true,
  },
  templateCode: {
    type: String,
  }
})

// 新增模版: 通讯方式选项
const mailVendors = [
  "smtp",
  "microsoft",
  "tencent"
];

const smsVendors = [
  "ali_yun",
  "tencent_yun",
  "huawei_yun",
];

const imVendors = [
  "fastmsg",
]

onMounted(() => {
  state.category = props.category;
  if (props.templateCode !== undefined) {
    // 编辑：加载配置
    getTemplateInfoAPi({
      code: props.templateCode,
      // category: state.category
    }).then(res => {
      if (res.code && res.code == 200) {
        console.log(res);
        state.formData = res.item
        if (htmlTemplate.value) {
          state.getHtml = res.item.body.replace("<html>", "")
              .replace("</html>", "")
              .replace("<head>", "")
              .replace("</head>", "")
              .replace("<body>", "")
              .replace("</body>", "")
              .replace("<meta charset=\"utf-8\">", "")
        }
      }
    }).catch(err => {
      ElMessage.error("无法加载模板信息")
    })
  }
})

const state = reactive({
  formData: {
    code: "",
    name: "",
    description: "",
    body: "",
    provider: "",
    templateId: "",
    params: "",
    signature: "",
    sender: 0,
    category: "",
  },
  category: "",
  getHtml: "",
  getText: "",
});

watch(() => state.getHtml, (newVal, oldVal) => {
  if (htmlTemplate.value) {
    state.formData.body = "<html>" +
        "<head>" +
        "<meta charset=\"utf-8\">" +
        "</head>" +
        "<body>" + newVal +
        "</body>" +
        "</html>"
  }
  // console.log(state.formData.body)
});

const onSubmit = async () => {
  // todo: 表单验证
  state.formData.category = state.category;
  // console.log(state.getText,state.getHtml);
  if (htmlTemplate.value) {
    if (state.getText) {
      state.formData.body = "<html>" +
          "<head>" +
          "<meta charset=\"utf-8\">" +
          "</head>" +
          "<body>" + state.getHtml +
          "</body>" +
          "</html>";
    } else {
      state.formData.body = state.getText;
    }
  }
  // if (state.formData.sender) {
  //   state.formData.sender = parseInt(state.formData.sender)
  // }
  try {
    let res;
    if (state.formData.code) {
      console.log("更新");
      res = await updateTemplateApi(state.formData);
    } else {
      console.log("创建");
      res = await createTemplateApi(state.formData);
    }

    if (res.code === 200) {
      ElMessage.success("提交成功！");
      resetForm(state.category);
      // state
      return true;
    } else {
      console.log(res);
      ElMessage.error("提交失败");
      return false;
    }
  } catch (err) {
    console.error(err);
    ElMessage.error("网络异常，请重试");
    return false;
  }
};

const resetForm = (category: string) => {
  console.log("resetForm")
  state.category = category;
  state.getHtml = "";
  state.getText = "";
  mailFormFormRef.value.resetFields();
  state.formData.body = "";
}

const vendors = computed(() => {
  switch (state.category) {
    case "sms":
      return smsVendors;
    case "mail":
      return mailVendors;
    case "im":
      return imVendors;
    default:
      return [];
  }
})

const htmlTemplate = computed(() => {
  return state.category !== 'im'
})


defineExpose({
  onSubmit,
  resetForm
})
</script>

<style scoped lang="scss">

</style>