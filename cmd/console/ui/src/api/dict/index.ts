import {get} from '/src/utils/request';

const baseUrl = "/api/dict"

export const getProvidersApi = (category: string) => get(`${baseUrl}/providers/${category}`, {});
