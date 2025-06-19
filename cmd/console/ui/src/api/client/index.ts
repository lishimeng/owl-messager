import {get, post} from '/src/utils/request';

const baseUrl = "/api/client"
export const createClientApi = (p: object) => post(baseUrl + "/create", p);
export const getClientListAPi = (p: object) => get(baseUrl + "/list", p);
export const delClientApi = (p: object) => post(baseUrl + "/del", p);
export const querySecretApi = (p: object) => get(baseUrl + "/secret", p);