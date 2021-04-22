import {
  get,
  post,
  postForm
} from '../api/request'

export const getMessageInfoApi = (p) => get('api/message', p)
export const getSenderInfoApi = (p) => get('api/mail_sender', p)
export const getMailTemplateApi = (p) => get('api/mail_template', p)
export const getTaskApi = (p) => get('api/task', p)
export const getTaskInfoByMessageIdApi = (p) => get('api/task/message/' + p.messageId, p)
export const getMailApi = (p) => get('api/mail', p)
export const getMailByMessageIdApi = (p) => get('api/mail/message/' + p.messageId, p)

export const createMailTemplateApi = (p) => post('api/mail_template', p)
export const getCurrentSwitchStatusApi = (p) => get('coreapi/908/001/1', p)
export const changeSwitchStatusApi = (p) => postForm('coreapi/908/001/2', p)
// export const addSenderInfoApi = (p) => post('api/mail_sender', p)
// export const deleteSenderInfoApi = (p) => get('api/mail_sender/delete', p)
// export const deleteMailTemplateApi = (p) => get('api/mail_template/delete', p)

// import request from '@/utils/request'
// import { baseURL } from './request'

// export function getSenderInfoApi(query) {
//   return request({
//     url: '/mail_sender',
//     method: 'get',
//     params: query,
//     baseURL: 'http://192.168.1.76/messager'
//   })
// }

// export function getMessageInfoApi(query) {
//   return request({
//     url: '/message',
//     method: 'get',
//     params: query,
//   })
// }
