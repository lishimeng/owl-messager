import { Module } from 'vuex';
// 此处加上 `.ts` 后缀报错，具体原因不详
import { UserInfosState, RootStateTypes } from '/@/store/interface/index';

const userInfosModule: Module<UserInfosState, RootStateTypes> = {
	namespaced: true,
	state: {
		userInfos: {},
	},
	mutations: {
		// 设置用户信息
		getUserInfos(state: any, data: object) {
			state.userInfos = data;
		},
	},
	actions: {
		// 设置用户信息
		async setUserInfos({ commit }, data: object) {
			if (data) {
				commit('getUserInfos', data);
			} else {
				if (window.localStorage.getItem('eqRoles')) {
					let data = {}
					data.userName = window.localStorage.getItem('userName') || ""
					data.time = new Date().getTime()

					let auth: Array<string> = ['base']
					let roles = window.localStorage.getItem("eqRoles")
					// admin
					if (roles && roles.indexOf('1') > 0) {data.authPageList = auth.push('admin');}
					data.authPageList = auth
					commit('getUserInfos', data)
				}
			}
		},
	},
};

export default userInfosModule;
