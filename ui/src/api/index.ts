
import { get, post, put, del } from '/@/utils/request';
const baseUrl = "/api"

// 获取字典
export const getDictDataApi = (p: object) => get(baseUrl + "/dictData/" + p.groupId + "/zh_CN", p);
// 获取库位列表
export const getLocationListApi = (p: object) => get(baseUrl + "/locationInfo", p);
// 新增库位
export const addLocationInfoApi = (p: object) => post(baseUrl + "/locationInfo", p);
// 修改库位信息
export const updateLocationInfoApi = (p: object) => put(baseUrl + "/locationInfo/" + p.id, p);
// 获取库位信息
export const getLocationInfoApi = (p: object) => get(baseUrl + "/locationInfo/" + p.id, p);
// 获取SPU列表
export const getSpuListApi = (p: object) => get(baseUrl + "/spu", p);
// 修改SPU信息
export const updateSpuInfoApi = (p: object) => put(baseUrl + "/spu/" + p.spuId, p);
// 获取SPU信息
export const getSpuInfoApi = (p: object) => get(baseUrl + "/spu/" + p.id, p);
// 新增SPU
export const addSpuInfoApi = (p: object) => post(baseUrl + "/spu", p);
// 获取SKU列表
export const getSkuListApi = (p: object) => get(baseUrl + "/sku", p);
// 修改SKU信息
export const updateSkuInfoApi = (p: object) => put(baseUrl + "/sku/" + p.skuId, p);
// 获取SKU信息
export const getSkuInfoApi = (p: object) => get(baseUrl + "/sku/" + p.id, p);
// 新增SKU
export const addSkuInfoApi = (p: object) => post(baseUrl + "/sku", p);
// 获取库存列表
export const getStorageListApi = (p: object) => get(baseUrl + "/storage", p);
// 获取收料单列表
export const getReceivingInfoListApi = (p: object) => get(baseUrl + "/receiving/info", p);
// 修改收料单信息
export const updateReceivingInfoApi = (p: object) => put(baseUrl + "/receiving/info/" + p.id, p);
// 获取收料单信息
export const getReceivingInfoApi = (p: object) => get(baseUrl + "/receiving/info/" + p.id, p);
// 新增收料单
export const addReceivingInfoApi = (p: object) => post(baseUrl + "/receiving/info", p);
// 收料单物品列表
export const getReceivingDetailListApi = (p: object) => get(baseUrl + "/receiving/detail/receive/" + p.receiveId, p);
// 新增收料单物品
export const addReceivingDetailApi = (p: object) => post(baseUrl + "/receiving/detail", p);
// 批量新增收料单物品
export const addBatchReceivingDetailApi = (p: object) => post(baseUrl + "/receiving/detail/batch", p);
// 删除收料单物品
export const deleteReceivingDetailApi = (p: object) => del(baseUrl + "/receiving/detail/" + p.id, p);
// 提交收料单
export const submitReceivingApi = (p: object) => put(baseUrl + "/receiving/info/" + p.id + "/submit", p);
// 取消收料单
export const cancelReceivingApi = (p: object) => put(baseUrl + "/receiving/info/" + p.id + "/cancel", p);
// 完成收料
export const accomplishReceivingApi = (p: object) => put(baseUrl + "/receiving/info/" + p.id + "/accomplish", p);
// 获取审核者申请单列表
export const getPickingApplyInfoListByAudiorApi = (p: object) => get(baseUrl + "/apply/info/auditor/" + p.id, p);
// 修改审核者申请单信息
export const updatePickingAppleInfoApi = (p: object) => put(baseUrl + "/apply/info/" + p.id, p);
// 获取申请单信息
export const getPickingAppleInfoApi = (p: object) => get(baseUrl + "/apply/info/" + p.id, p);
// 新增申请单
export const addPickingAppleInfoApi = (p: object) => post(baseUrl + "/apply/info", p);
// 审核申请单
export const PickingAppleInfoAuditorApi = (p: object) => put(baseUrl + "/apply/info/" + p.id + "/audit", p);
// 获取申请单下物品信息
export const getPickingApplyDetailInfoList = (p: object) => get(baseUrl + "/apply/detail/applyId/" + p.id, p);
// 获取申请者申请列表
export const getPickingApplyInfoListByApplicantApi = (p: object) => get(baseUrl + "/apply/info/applicant/" + p.id, p);
// 申请人提交申请单
export const submitPickingApplyInfoApi = (p: object) => put(baseUrl + "/apply/info/" + p.id + "/submit", p);
// 取消申请单
export const cancelPickingApplyInfoApi = (p: object) => get(baseUrl + "/apply/info/" + p.id + "/cancel", p);
// 查询申请单下的物品列表
export const getPickingApplyInfoWithSkuListApi = (p: object) => get(baseUrl + "/apply/detail/applyId/" + p.id, p);
// 新增申请单下的物品
export const addSkuPickingAppleInfoApi = (p: object) => post(baseUrl + "/apply/detail", p);
// 删除申请单下的物品
export const delSkuPickingAppleInfoApi = (p: object) => del(baseUrl + "/apply/detail/" + p.id, p);
// 完成审核单
export const getPickingApplyInfoDoneApi = (p: object) => get(baseUrl + "/apply/info/" + p.id + "/done", p);
// 报废操作单
export const scrapApi = (p: object) => post(baseUrl + "/storage/scrap", p);
// 查询人员列表
export const getUserListApi = (p) => get(baseUrl + '/user/' + p.id, p)
// 查询申请人、审核人报废单列表
export const scrapListApi = (p: object) => get(baseUrl + "/storage/scrap/list", p);
// 修改报废单状态 1取消 2拒绝 3通过
export const scrapSubmitApi = (p: object) => get(baseUrl + "/storage/scrap/submit", p);
// 盘点任务列表
export const takeStockInfoListApi = (p: object) => get(baseUrl + "/takeStockInfo", p);
// id查询任务
export const takeStockInfoApi = (p: object) => get(baseUrl + "/takeStockInfo/" + p.id, p);
// 添加盘点任务
export const addTakeStockInfoApi = (p: object) => post(baseUrl + "/takeStockInfo", p);
// 取消盘点任务
export const cancelTakeStockInfoApi = (p: object) => get(baseUrl + "/takeStockInfo/cancel/" + p.id, p);
// 提交盘点任务
export const submitlTakeStockInfoApi = (p: object) => get(baseUrl + "/takeStockInfo/submit/" + p.id, p);
// 修改盘点任务
export const editTakeStockInfoApi = (p: object) => put(baseUrl + "/takeStockInfo/" + p.id, p);
// 查询盘点任务详情列表
export const takeStockDetailListApi = (p: object) => get(baseUrl + "/takeStockDetail/" + p.id, p);
// 删除盘点任务下的库位
export const delTakeStockDetailApi = (p: object) => del(baseUrl + "/takeStockDetail/" + p.id, p);
// 添加盘点库位
export const addTakeStockDetailApi = (p: object) => post(baseUrl + "/takeStockDetail", p);
// 完成盘点任务
export const finishTakeStockInfoApi = (p: object) => get(baseUrl + "/takeStockInfo/finish/" + p.id, p);
// 查询事件管理列表
export const eventListApi = (p: object) => get(baseUrl + "/operationalEvent", p);
// 查询入库记录
export const receivingDetailLogApi = (p: object) => get(baseUrl + "/itemReceivingDetailLog", p);
// 查询领用记录
export const pickingDetailLogApi = (p: object) => get(baseUrl + "/itemPickingDetailLog", p);
// 查询归还记录
export const pickingDetailRetutnLogApi = (p: object) => get(baseUrl + "/itemPickingDetailReturnLog", p);
// 查询消耗记录
export const expendDetailRetutnLogApi = (p: object) => get(baseUrl + "/itemPickingDetailExpendLog", p);
// 获取供应商列表
export const getSupplierListApi = (p: object) => get(baseUrl + "/supplier", p);
// 新增供应商
export const addSupplierInfoApi = (p: object) => post(baseUrl + "/supplier", p);
// 修改供应商信息
export const updateSupplierInfoApi = (p: object) => put(baseUrl + "/supplier/" + p.id, p);
// 获取库位供应商信息
export const getSupplierInfoApi = (p: object) => get(baseUrl + "/supplier/" + p.id, p);
// 库内调整—修改库存
export const changeStorageInventoryApi = (p: object) => post(baseUrl + "/storage/change_inventory", p);
// 库内调整—销毁库存
export const destroyStorageInventoryApi = (p: object) => post(baseUrl + "/storage/destroy_inventory", p);
// 获取库存汇总列表
export const getStorageViewListApi = (p: object) => get(baseUrl + "/storage/total", p);

// 批量添加盘点库位
export const addBatchTakeStockDetailApi = (p: object) => post(baseUrl + "/takeStockDetail/batch", p);