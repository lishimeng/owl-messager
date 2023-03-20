import {
  get,
  post,
  put,
  del
} from '/@/utils/request';

const baseURL = 'api'
// 登录
export const signIn = (p) => post(baseURL + '/user/sign_in', p)
// 获取用户信息
export const getInfo = (p) => get(baseURL + '/authUser/' + p.userId, p)
// 退出登录
export const logOut = (p) => post(baseURL + '/user/logout', p)
// 获取用户列表
export const getUserListApi = (p) => get(baseURL + '/authUser', p)
// 修改用户状态
export const setUserStatusApi = (p) => put(baseURL + '/authUser/' + p.userId + "/status", p)
// 添加用户权限
export const addUserRoleApi = (p) => post(baseURL + '/authRoles', p)
// 删除用户权限
export const deleteUserRoleApi = (p) => del(baseURL + '/authRoles/' + p.userId + '/' + p.roleId, p)
// 修改用户信息
export const editUserApi = (p) => put(baseURL + '/authUser/' + p.userId, p)
// 新增用户
export const addUserApi = (p) => post(baseURL + '/authUser/add', p)
// 修改密码
export const resetPwdApi = (p) => post(baseURL + '/user/password/change', p)
// 管理员重置密码
export const resetPwdByAdminApi = (p) => post(baseURL + '/user/password/reset', p)
// 登录
export const signInCardApi = (p) => post(baseURL + '/user/sign_in_card', p)
// 获取验证码
export const getCaptchaApi = (p) => post(baseURL + '/user/getcaptcha', p)