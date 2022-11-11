
import { get, post, put, del } from '/@/utils/request';
const baseUrl = "/api"

export const getMessageInfoApi = (p: object) => get(baseUrl + '/message', p)
export const getSenderInfoApi = (p: object) => get(baseUrl + '/mail_sender', p)
export const getMailTemplateApi = (p: object) => get(baseUrl + '/mail_template', p)
export const getTaskApi = (p: object) => get(baseUrl + '/task', p)
export const getTaskInfoByMessageIdApi = (p: object) => get(baseUrl + '/task/message/' + p.messageId, p)
export const getMailApi = (p: object) => get(baseUrl + '/mail', p)
export const getMailByMessageIdApi = (p: object) => get(baseUrl + '/mail/message/' + p.messageId, p)

export const createMailTemplateApi = (p: object) => post(baseUrl + '/mail_sender/add/', p)

export const getsenderlistOneApi = (p: object) => get(baseUrl + '/mail_sender/one/'+ p.id, p)
export const postsenderupdateOneApi = (p: object) => post(baseUrl + '/mail_sender/update/', p)

export const delsenderlistOneApi = (p: object) => get(baseUrl + '/mail_sender/del/'+ p.id, p)


export const createMaiTemplateMailApi = (p: object) => post(baseUrl + '/mail_template/add/', p)

export const gettemplatelistOneApi = (p: object) => get(baseUrl + '/mail_template/one/'+ p.id, p)

export const posttemplateupdateOneApi = (p: object) => post(baseUrl + '/mail_template/update/', p)

export const deltemplatelistOneApi = (p: object) => get(baseUrl + '/mail_template/del/'+ p.id, p)





