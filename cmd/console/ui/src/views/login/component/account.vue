<template>
	<el-form size="large" class="login-content-form">
		<el-form-item class="login-animation1">
			<el-input
				type="textarea"
				:rows="3"
				placeholder="请输入管理 Token"
				v-model="state.ruleForm.token"
				clearable
				autocomplete="off"
			>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation2">
			<el-button type="primary" class="login-content-submit" round v-waves @click="onSignIn" :loading="state.loading.signIn">
				<span>登录</span>
			</el-button>
		</el-form-item>
	</el-form>
</template>

<script setup lang="ts" name="loginAccount">
import { reactive, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { useI18n } from 'vue-i18n';
import Cookies from 'js-cookie';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import { initFrontEndControlRoutes } from '/@/router/frontEnd';
import { initBackEndControlRoutes } from '/@/router/backEnd';
import { Session } from '/@/utils/storage';
import { formatAxis } from '/@/utils/formatTime';
import { NextLoading } from '/@/utils/loading';
import { useLoginApi } from '/@/api/login';

const { t } = useI18n();
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const route = useRoute();
const router = useRouter();
const state = reactive({
	ruleForm: {
		token: '',
	},
	loading: {
		signIn: false,
	},
});

const currentTime = computed(() => {
	return formatAxis(new Date());
});

const onSignIn = async () => {
	if (!state.ruleForm.token.trim()) {
		ElMessage.error('请输入授权 Token');
		return;
	}
	state.loading.signIn = true;
	try {
		const res: any = await useLoginApi().signIn({
			token: state.ruleForm.token.trim(),
		});
		if (!res || res.code !== 200 || !res.token) {
			ElMessage.error(res?.message || '登录失败');
			state.loading.signIn = false;
			return;
		}
		Session.set('token', res.token);
		Cookies.set('userName', 'console');
		if (!themeConfig.value.isRequestRoutes) {
			const isNoPower = await initFrontEndControlRoutes();
			signInSuccess(isNoPower);
		} else {
			const isNoPower = await initBackEndControlRoutes();
			signInSuccess(isNoPower);
		}
	} catch {
		ElMessage.error('登录失败');
		state.loading.signIn = false;
	}
};

const signInSuccess = (isNoPower: boolean | undefined) => {
	if (isNoPower) {
		ElMessage.warning('抱歉，您没有登录权限');
		Session.clear();
	} else {
		let currentTimeInfo = currentTime.value;
		if (route.query?.redirect) {
			router.push({
				path: <string>route.query?.redirect,
				query: Object.keys(<string>route.query?.params).length > 0 ? JSON.parse(<string>route.query?.params) : '',
			});
		} else {
			router.push('/');
		}
		const signInText = t('message.signInText');
		ElMessage.success(`${currentTimeInfo}，${signInText}`);
		NextLoading.start();
	}
	state.loading.signIn = false;
};
</script>

<style scoped lang="scss">
.login-content-form {
	margin-top: 20px;
	@for $i from 1 through 2 {
		.login-animation#{$i} {
			opacity: 1;
			animation-name: error-num;
			animation-duration: 0.5s;
			animation-fill-mode: forwards;
			animation-delay: calc($i/10) + s;
		}
	}
	.login-content-submit {
		width: 100%;
		letter-spacing: 2px;
		font-weight: 300;
		margin-top: 15px;
	}
}
</style>
