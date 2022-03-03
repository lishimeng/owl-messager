<template>
	<div class="error">
		<div class="error-flex">
			<div class="left">
				<div class="left-item">
					<div class="left-item-animation left-item-num">Loading</div>
					<div class="left-item-animation left-item-title">跳 转 登 录 中 ......</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import { ElMessage } from "element-plus"
import { reactive, onMounted, getCurrentInstance } from "vue"
import { useRouter, useRoute } from "vue-router"
import { signInCardApi, getInfo } from "/@/api/user"
import { initFrontEndControlRoutes } from '/@/router/frontEnd';
import { Session } from '/@/utils/storage';
const { proxy } = getCurrentInstance();
const route = useRoute()
const state = reactive({
	userId: parseInt(route.query.userId),
})

const router = useRouter()
function signInCard() {
	signInCardApi({
		uid: state.userId
	}).then(res => {
		Session.set('token', res.jwt)
		window.localStorage.setItem("token", res.jwt)
		window.localStorage.setItem('eqRoles', JSON.stringify(res.roles))
		window.localStorage.setItem('oid', res.orgId)
		genUserInfo(state.userId);
	}).catch(() => {
		signInFailed();
	})
}
// 获取用户信息
function genUserInfo(uid) {
	getInfo({
		userId: uid
	}).then(res => {
		window.localStorage.setItem('userName', res.userName)
		window.localStorage.setItem('userId', res.userId)
		initFrontEndControlRoutes();
		signInSuccess();
	}).catch(() => {
		signInFailed();
	})
};
function signInSuccess() {
	router.push('/');
	setTimeout(() => {
		proxy.mittBus.emit('onSignInClick');
	}, 300);
}
function signInFailed() {
	router.push('/login')
	ElMessage.error("未注册用户不能登录")
}
function goSkip() {
	if (state.userId) {
		signInCard()
	} else {
		router.push('/login')
	}
}
onMounted(() => {
	goSkip()
})
</script>

<style scoped lang="scss">
.error {
	height: 100%;
	background-color: var(--el-color-white);
	display: flex;
	.error-flex {
		margin: auto;
		display: flex;
		height: 350px;
		width: 900px;
		.left {
			flex: 1;
			height: 100%;
			align-items: center;
			display: flex;
			.left-item {
				.left-item-animation {
					opacity: 0;
					animation-name: error-num;
					animation-duration: 0.5s;
					animation-fill-mode: forwards;
				}
				.left-item-num {
					color: var(--el-color-info);
					font-size: 55px;
				}
				.left-item-title {
					font-size: 20px;
					color: var(--el-text-color-primary);
					margin: 15px 0 5px 0;
					animation-delay: 0.1s;
				}
				.left-item-msg {
					color: var(--el-text-color-secondary);
					font-size: 12px;
					margin-bottom: 30px;
					animation-delay: 0.2s;
				}
				.left-item-btn {
					animation-delay: 0.2s;
				}
			}
		}
		.right {
			flex: 1;
			opacity: 0;
			animation-name: error-img;
			animation-duration: 2s;
			animation-fill-mode: forwards;
			img {
				width: 100%;
				height: 100%;
			}
		}
	}
}
</style>