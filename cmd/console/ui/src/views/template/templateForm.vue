<template>
  <div style="width:100%;padding: 10px">
    <el-form style="margin-top: 20px"
             :model="state.formData"
             ref="mailFormFormRef"
             label-width="120px">
      <el-form-item label="租户" prop="tenantCode" v-if="!state.formData.code">
        <el-input v-model="state.formData.tenantCode" clearable placeholder="tenant.code"></el-input>
      </el-form-item>
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
      <el-form-item label="第三方模板 ID" prop="cloudTemplate">
        <el-input v-model="state.formData.cloudTemplate" clearable placeholder="云平台模板 ID，本地模板可留空"></el-input>
      </el-form-item>
			<el-form-item label="模板参数" prop="params">
				<el-input
					style="width: 100%"
					type="textarea"
					:autosize="{ minRows: 6, maxRows: 16 }"
					v-model="state.formData.params"
					placeholder='JSON 映射，例如：&#10;{&#10;  "userName": {&#10;    "description": "用户名",&#10;    "attr": ["UserName"]&#10;  },&#10;  "code": {&#10;    "attr": ["code"]&#10;  }&#10;}'
				></el-input>
				<div class="params-hint">
					key 为发信 API 的 params 字段名；attr 为模板/云平台变量名（可不同大小写）。也支持简写：<code>userName, code</code>
				</div>
			</el-form-item>
<!--      <el-form-item v-if="state.category=='sms'" label="第三方模板签名" prop="signature">-->
<!--        <el-input type="textarea" v-model="state.formData.signature" clearable></el-input>-->
<!--      </el-form-item>-->
<!--      <el-form-item v-if="state.category=='sms'" label="指定发送平台" prop="sender">-->
<!--        <el-input type="number" v-model="state.formData.sender" clearable></el-input>-->
<!--      </el-form-item>-->
			<div>
				<el-form-item label="模板内容" v-if="htmlTemplate">
					<wngEditor
              :key="props.templateCode || 'create'"
              mode="default"
              height="320px"
              v-model:getHtml="state.getHtml"
              v-model:getText="state.getText"
          ></wngEditor>
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
			</div>
    </el-form>

  </div>
</template>

<script setup lang="ts">
import {computed, defineAsyncComponent, onMounted, reactive, ref, watch} from "vue";
import {createTemplateApi, getTemplateInfoAPi, updateTemplateApi} from "/@/api/template";
import {getProvidersApi} from "/@/api/dict";
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

// 通讯平台选项来自 /api/dict/providers/{category}

const extractEditorHtml = (body: string) => {
  if (!body) return '';
  const match = body.match(/<body[^>]*>([\s\S]*?)<\/body>/i);
  return match ? match[1].trim() : body;
};

const formatLoadedParams = (params: string) => {
  if (!params?.trim()) {
    return '{}';
  }
  if (params.trim().startsWith('{')) {
    try {
      return JSON.stringify(JSON.parse(params), null, 2);
    } catch {
      return params;
    }
  }
  return params;
};

const vendorOptions = ref<string[]>([]);

const loadVendors = async (category: string) => {
  try {
    const res = await getProvidersApi(category);
    if (res?.code === 200 && res.items) {
      vendorOptions.value = res.items.map((item: { name: string }) => item.name);
    } else {
      vendorOptions.value = [];
    }
  } catch {
    vendorOptions.value = [];
  }
};

onMounted(() => {
  state.category = props.category;
  loadVendors(props.category);
  if (props.templateCode !== undefined) {
    // 编辑：加载配置
    getTemplateInfoAPi({
      code: props.templateCode,
      // category: state.category
    }).then(res => {
      if (res.code && res.code == 200) {
        state.formData = res.item
        state.formData.params = formatLoadedParams(res.item.params || '{}')
        if (htmlTemplate.value && res.item.body) {
          state.getHtml = extractEditorHtml(res.item.body);
        }
      }
    }).catch(err => {
      ElMessage.error("无法加载模板信息")
    })
  }
})

const state = reactive({
  formData: {
    tenantCode: "",
    code: "",
    name: "",
    description: "",
    body: "",
    provider: "",
		cloudTemplate: "",
    params: "{}",
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

const validateParams = (): boolean => {
  const raw = state.formData.params?.trim();
  if (!raw) {
    state.formData.params = '{}';
    return true;
  }
  if (raw.startsWith('{')) {
    try {
      JSON.parse(raw);
      return true;
    } catch {
      ElMessage.error('模板参数不是合法 JSON');
      return false;
    }
  }
  return true;
};

const onSubmit = async () => {
  if (!validateParams()) {
    return false;
  }
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
      res = await updateTemplateApi(state.formData);
    } else {
      if (!state.formData.tenantCode?.trim()) {
        ElMessage.error("请填写租户 code");
        return false;
      }
      res = await createTemplateApi(state.formData);
    }

    if (res.code === 200) {
      ElMessage.success("提交成功！");
      resetForm(state.category);
      // state
      return true;
    } else {
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
  state.category = category;
  loadVendors(category);
  state.getHtml = "";
  state.getText = "";
  mailFormFormRef.value.resetFields();
  state.formData.body = "";
  state.formData.params = "{}";
}

const vendors = computed(() => vendorOptions.value)

const htmlTemplate = computed(() => {
  return state.category !== 'im'
})

defineExpose({
  onSubmit,
  resetForm
})
</script>

<style scoped lang="scss">
.params-hint {
  margin-top: 6px;
  font-size: 12px;
  color: var(--el-text-color-secondary);
  line-height: 1.5;

  code {
    padding: 0 4px;
    background: var(--el-fill-color-light);
    border-radius: 3px;
  }
}
</style>