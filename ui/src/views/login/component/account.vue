<template>
	<el-form class="login-content-form">
		<el-form-item>
			<el-input
				type="text"
				:placeholder="$t('message.account.accountPlaceholder1')"
				prefix-icon="el-icon-user"
				v-model="ruleForm.userName"
				clearable
				autocomplete="off"
			>
			</el-input>
		</el-form-item>
		<el-form-item>
			<el-input
				:type="isShowPassword ? 'text' : 'password'"
				:placeholder="$t('message.account.accountPlaceholder2')"
				prefix-icon="el-icon-lock"
				v-model="ruleForm.password"
				autocomplete="off"
			>
				<template #suffix>
					<i
						class="iconfont el-input__icon login-content-password"
						:class="isShowPassword ? 'icon-yincangmima' : 'icon-xianshimima'"
						@click="isShowPassword = !isShowPassword"
					>
					</i>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item>
			<el-button type="primary" class="login-content-submit" round @click="onSignIn" :loading="loading.signIn">
				<span>{{ $t('message.account.accountBtnText') }}</span>
			</el-button>
		</el-form-item>
	</el-form>
</template>

<script lang="ts">
import { toRefs, reactive, defineComponent, computed, getCurrentInstance } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { useI18n } from 'vue-i18n';
import { initFrontEndControlRoutes } from '/@/router/frontEnd';
import { useStore } from '/@/store/index';
import { Session } from '/@/utils/storage';
import { formatAxis } from '/@/utils/formatTime';
import { signIn, getInfo } from '/@/api/user'
export default defineComponent({
	name: 'loginAccount',
	setup() {
		const { t } = useI18n();
		const { proxy } = getCurrentInstance() as any;
		const store = useStore();
		const route = useRoute();
		const router = useRouter();
		const state = reactive({
			isShowPassword: false,
			ruleForm: {
				userName: '',
				password: '',
			},
			loading: {
				signIn: false,
			},
		});
		// 时间获取
		const currentTime = computed(() => {
			return formatAxis(new Date());
		});
		// 点击登录
		const onSignIn = async () => {
			state.loading.signIn = true;
			signIn({
				loginName: state.ruleForm.userName,
				password: state.ruleForm.password,
			}).then(res =>{
				Session.set('token', res.jwt)
				window.localStorage.setItem('token', res.jwt)
				window.localStorage.setItem('eqRoles', JSON.stringify(res.roles))
				window.localStorage.setItem('oid', res.orgId)
				genUserInfo(res.uid)
			}).catch(err =>{
				state.loading.signIn = false;
				console.log(err)
			})
		};
		// 获取用户信息
		const genUserInfo = async (uid) => {
			getInfo({
				userId: uid
			}).then(res =>{
				window.localStorage.setItem('userName', res.userName)
				window.localStorage.setItem('userId', res.userId)
				if (!store.state.themeConfig.themeConfig.isRequestRoutes) {
					// 前端控制路由，2、请注意执行顺序
					initFrontEndControlRoutes();
					signInSuccess();
				}
			}).catch(err =>{console.log(err)})
		};
		// 登录成功后的跳转
		const signInSuccess = () => {
			// 初始化登录成功时间问候语
			let currentTimeInfo = currentTime.value;
			// 登录成功，跳到转首页
			// 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
			// 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
			if (route.query?.redirect) {
				router.push({
					path: route.query?.redirect,
					query: Object.keys(route.query?.params).length > 0 ? JSON.parse(route.query?.params) : '',
				});
			} else {
				router.push('/');
			}
			// 登录成功提示
			setTimeout(() => {
				// 关闭 loading
				state.loading.signIn = false;
				const signInText = t('message.signInText');
				ElMessage.success(`${currentTimeInfo}，${signInText}`);
				// 修复防止退出登录再进入界面时，需要刷新样式才生效的问题，初始化布局样式等(登录的时候触发，目前方案)
				proxy.mittBus.emit('onSignInClick');
			}, 300);
		};
		return {
			currentTime,
			onSignIn,
			...toRefs(state),
		};
	},
});
</script>

<style scoped lang="scss">
.login-content-form {
	margin-top: 20px;
	.login-content-password {
		display: inline-block;
		width: 25px;
		cursor: pointer;
		&:hover {
			color: #909399;
		}
	}
	.login-content-code {
		display: flex;
		align-items: center;
		justify-content: space-around;
		.login-content-code-img {
			width: 100%;
			height: 40px;
			line-height: 40px;
			background-color: #ffffff;
			border: 1px solid rgb(220, 223, 230);
			color: #333;
			font-size: 16px;
			font-weight: 700;
			letter-spacing: 5px;
			text-indent: 5px;
			text-align: center;
			cursor: pointer;
			transition: all ease 0.2s;
			border-radius: 4px;
			user-select: none;
			&:hover {
				border-color: #c0c4cc;
				transition: all ease 0.2s;
			}
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
