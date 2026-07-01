<template>
  <div class="home-container layout-pd" style="margin-top: 10px">
    <el-form :inline="true">
      <el-form-item label="租户">
        <el-input v-model="state.queryValue.tenantCode" clearable placeholder="可选" style="width: 140px" @change="getMailSenders"/>
      </el-form-item>
      <el-form-item label="通讯方式">
        <el-select v-model="state.category" @change="getMailSenders" placeholder="请选择通讯方式" style="width: 120px">
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
        <el-table-column prop="id" label="id" width="60"/>
        <el-table-column prop="code" label="code" width="380"/>
        <el-table-column prop="defaultSender" label="默认发送人" width="120">
          <template v-slot="scope">
            <span v-if="scope.row.defaultSender">
              是
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="vendor" label="Vendor" width="120"/>
        <el-table-column prop="createTime" label="创建时间" show-overflow-tooltip>
          <template #default="scope">
            {{ formatDate(new Date(scope.row.createTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column prop="updateTime" label="更新时间" show-overflow-tooltip>
          <template #default="scope">
            {{ formatDate(new Date(scope.row.updateTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="240">
          <template #default="scope">
            <el-button size="small" link @click="showEdit(scope.row)">
              编辑
            </el-button>
            <el-button size="small" link v-if="scope.row.defaultSender" @click="showTest(scope.row)">
              测试发送
            </el-button>
            <el-button size="small" link v-else @click="setDefaultSender(scope.row)">
              设为默认
            </el-button>
            <el-popconfirm
                title="是否删除?"
                @confirm="deleteSender(scope.row)">
              <template #reference>
                <el-button size="small" link>删除</el-button>
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
    <el-dialog v-model="state.showDialog" :title="state.title" width="50%" center>
      <div style="margin: 10px">
        <el-form :model="state.form" label-width="120px">
          <el-form-item v-if="!state.isDisabled" label="租户" prop="tenantCode">
            <el-input v-model="state.form.tenantCode" placeholder="tenant.code"></el-input>
          </el-form-item>
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
        <SenderConfigForm
            :vendor="state.form.vendor"
            :category="state.category"
            :config="state.form.config"
            ref="senderForm"/>
      </div>
      <span class="dialog-footer">
        <el-button @click="state.showDialog = false">取消</el-button>
        <el-button type="primary" @click="onSubmit()">提交</el-button>
      </span>
    </el-dialog>
    <el-dialog
        v-model="state.showTest"
        :title="state.testTitle"
        width="50%"
        center>
      <el-form style="margin-top: 20px"
               :model="state.testForm"
               label-width="120px">
        <el-form-item label="Messager 地址" prop="host">
          <el-input v-model="state.testForm.host" :placeholder="state.messagerHostHint"></el-input>
          <div class="params-hint">{{ state.messagerHostTip }}</div>
        </el-form-item>
        <el-form-item label="AppId" prop="appId">
          <el-input v-model="state.testForm.appId" placeholder="open_client.app_id"></el-input>
        </el-form-item>
        <el-form-item label="Secret" prop="secret">
          <el-input v-model="state.testForm.secret" type="password" show-password placeholder="open_client.secret"></el-input>
        </el-form-item>
        <el-form-item label="平台" prop="vendor">
          <el-input v-model="state.testForm.vendor" disabled></el-input>
        </el-form-item>
        <el-form-item label="发送人" prop="code">
          <el-input v-model="state.testForm.code" disabled></el-input>
        </el-form-item>
        <el-form-item :label="state.testReceiver" prop="receiver">
          <el-input v-model="state.testForm.receiver"></el-input>
        </el-form-item>
        <el-form-item label="使用模板">
          <el-select v-model="state.testForm.tpl" placeholder="已配置的模板..." style="width: 100%" @change="onTestTemplateChange">
            <el-option v-for="v in state.tplList" :key="v.code" :label="v.name" :value="v.code">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="模板参数" prop="tplParam">
          <el-input
              v-model="state.testForm.tplParam"
              type="textarea"
              :autosize="{ minRows: 4, maxRows: 10 }"
              placeholder='JSON，key 为模板映射对外参数名，例如：&#10;{&#10;  "userName": "张三",&#10;  "code": "839201"&#10;}'
          ></el-input>
          <div class="params-hint">对应发信 API 的 params 字段；留空时短信默认 {"code":"123456"}</div>
        </el-form-item>
      </el-form>
      <span class="dialog-footer mt30" style="display: flex; justify-content: center;">
        <el-button type="primary" @click="sendTest()">发送</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script setup lang="ts" name="Mail">
import {onMounted, reactive, ref} from 'vue';
import {
  createMailSenderConfigApi, delSenderApi,
  getMailSendersApi,
  getSenderInfoByCategoryAPi, senderTestApi, setDefaultSenderApi,
  getSenderTestConfigApi,
  updateMailSenderConfigApi
} from "/@/api/mail";
import {ElMessage} from "element-plus";
import {formatDate} from "/@/utils/formatTime";
import SenderConfigForm from "/@/views/mail/senderConfigForm.vue";
import {getTemplateListAPi} from "/@/api/template";
import {getProvidersApi} from "/@/api/dict";

const senderForm = ref();

const state = reactive({
  isShowText: true,
  isDisabled: false,
  title: "新增",
  showDialog: false,
  showTest: false,
  queryValue: {
    pageSize: 10,
    pageNum: 1,
    totalNum: 0,
    category: "",
    tenantCode: "",
  },
  dataList: [],
  category: "mail",
  categoryList: [
    "mail",
    "sms",
    "im"
  ],
  vendors: [] as string[],
  form: {
    tenantCode: "",
    code: "",
    defaultSender: 1,
    vendor: "",
    config: "",
    category: "",
  },
  testTitle: "",
  testReceiver: "",
  testForm: {
    host: "",
    appId: "",
    secret: "",
    code: "",
    vendor: "",
    receiver: "",
    tpl: "",
    tplParam: "{}",
  },
  tplList: [],
  messagerHostHint: "留空=本地直发；或 http://localhost:81",
  messagerHostTip: "远程测试填 owl-messager 根地址（非 Console :80）。发信路径 POST /messages/mail",
})
const MESSAGER_PROXY_KEY = 'messagerTestProxy';

const loadMessagerProxy = () => {
  try {
    const raw = localStorage.getItem(MESSAGER_PROXY_KEY);
    if (!raw) return;
    const saved = JSON.parse(raw);
    state.testForm.host = saved.host || '';
    state.testForm.appId = saved.appId || '';
    state.testForm.secret = saved.secret || '';
  } catch {
    // ignore
  }
};

const saveMessagerProxy = () => {
  localStorage.setItem(MESSAGER_PROXY_KEY, JSON.stringify({
    host: state.testForm.host,
    appId: state.testForm.appId,
    secret: state.testForm.secret,
  }));
};

onMounted(() => {
  loadMessagerProxy();
  loadSenderTestConfig();
  getMailSenders();
})
const loadSenderTestConfig = () => {
  getSenderTestConfigApi().then(res => {
    if (res?.defaultHost) {
      state.messagerHostHint = `留空=本地直发；默认 ${res.defaultHost}`;
      if (!state.testForm.host) {
        state.testForm.host = res.defaultHost;
      }
    }
    if (res?.message) {
      state.messagerHostTip = res.message + (res.sendPath ? `（${res.sendPath}）` : '');
    }
  }).catch(() => {
    // ignore
  });
};
const showEdit = (row: object) => {
  // console.log(row)
  if (row) {
    state.title = '编辑'
    state.isDisabled = true
    state.isShowText = true
    getSenderInfoByCategory(row.code)
  } else {
    state.showDialog = true
    state.isDisabled = false
    state.isShowText = false
    state.title = '新增'
    state.form.vendor = ''
    state.form.tenantCode = ''
  }
}
const showTest = async (row: object) => {
  switch (state.category) {
    case 'mail':
      state.testTitle = "邮件发送测试"
      state.testReceiver = "收件邮箱"
      break
    case 'sms':
      state.testTitle = "SMS发送测试"
      state.testReceiver = "收信手机"
      break
    case 'im':
      state.testTitle = "即时通讯发送测试"
      state.testReceiver = "收信人"
      break
  }
  state.testForm.vendor = row.vendor
  state.testForm.code = row.code
  state.testForm.tpl = ""
  state.testForm.tplParam = "{}"
  await loadTemplate(state.category, state.testForm.vendor)
  state.showTest = true
}
const buildParamsExample = (paramsStr: string) => {
  try {
    const mapping = JSON.parse(paramsStr || '{}');
    const example: Record<string, string> = {};
    for (const key of Object.keys(mapping)) {
      example[key] = key === 'code' ? '123456' : '示例值';
    }
    return JSON.stringify(example, null, 2);
  } catch {
    return '{}';
  }
};
const onTestTemplateChange = (code: string) => {
  const tpl = state.tplList.find((item: { code: string }) => item.code === code);
  if (tpl?.params) {
    state.testForm.tplParam = buildParamsExample(tpl.params);
  } else {
    state.testForm.tplParam = '{}';
  }
};
const parseTestParams = () => {
  const raw = state.testForm.tplParam?.trim();
  if (!raw || raw === '{}') {
    return undefined;
  }
  try {
    return JSON.parse(raw);
  } catch {
    ElMessage.error('模板参数不是合法 JSON');
    return null;
  }
};
const loadTemplate = async (category: string, vendor: string) => {
  const res = await getTemplateListAPi({
    pageNum: 1,
    pageSize: 100,
    category: category,
    provider: vendor,
  })
  if (res.items && res.items.length > 0) {
    state.tplList = res.items
  } else {
    state.tplList = []
  }
}
const sendTest = () => {
  if (!state.testForm.appId || !state.testForm.secret) {
    ElMessage.error('请填写 AppId 和 Secret');
    return;
  }
  if (!state.testForm.tpl) {
    ElMessage.error('请选择模板');
    return;
  }
  if (!state.testForm.receiver?.trim()) {
    ElMessage.error(`请填写${state.testReceiver}`);
    return;
  }
  const params = parseTestParams();
  if (params === null) {
    return;
  }
  saveMessagerProxy();
  const payload: Record<string, unknown> = {
    category: state.category,
    host: state.testForm.host,
    appId: state.testForm.appId,
    secret: state.testForm.secret,
    receiver: state.testForm.receiver,
    template: state.testForm.tpl,
  };
  if (params !== undefined) {
    payload.params = params;
  }
  if (state.category === 'mail') {
    payload.subject = 'Owl-messager: 测试邮件';
  }
  senderTestApi(payload).then(res => {
    if (res && res.code == 200) {
      state.showTest = false;
      ElMessage.success('发送成功，请检查发送历史与收件方');
      getMailSenders();
    } else {
      ElMessage.error(res?.message || '发送失败');
    }
  }).catch(err => {
    console.log(err)
    ElMessage.error(err?.message || '发送失败')
  })
}
const setDefaultSender = (row: object) => {
  setDefaultSenderApi({
    code: row.code,
    category: state.category,
    provider: row.vendor,
  }).then(res => {
    if (res && res.code == 200) {
      getMailSenders();
    }
  })
}
const deleteSender = (row: object) => {
  delSenderApi({
    code: row.code
  }).then(res => {
    if (res && res.code == 200) {
      ElMessage.success(`删除成功！`);
      getMailSenders();
    }
  }).catch(err => {
    console.log(err)
    ElMessage.error(`删除失败`)
  })
}
const onSubmit = () => {
  const config = senderForm.value.exportJson();
  if (!config) {
    ElMessage.error('请完善发件人配置');
    return;
  }
  state.form.config = config;
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
    if (!state.form.tenantCode?.trim()) {
      ElMessage.error('请填写租户 code');
      return;
    }
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
  senderForm.value?.clearForm();
}
const getSenderInfoByCategory = (code: string) => {
  getSenderInfoByCategoryAPi({
    category: state.category,
    code: code
  }).then(res => {
    if (res.item) {
      state.form.vendor = res.item.vendor
      state.form.code = res.item.code
      state.form.defaultSender = res.item.defaultSender === 1 ? 1 : 0
      state.form.config = res.item.config
      senderForm.value?.loadJson(res.item.vendor, res.item.config);
      state.showDialog = true
    }
  })
}
const loadVendors = async () => {
  try {
    const res = await getProvidersApi(state.category);
    if (res?.code === 200 && res.items) {
      state.vendors = res.items.map((item: { name: string }) => item.name);
    } else {
      state.vendors = [];
    }
  } catch {
    state.vendors = [];
  }
};
const getMailSenders = async () => {
  await loadVendors();
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
.params-hint {
  margin-top: 6px;
  font-size: 12px;
  color: var(--el-text-color-secondary);
  line-height: 1.5;
}
</style>