import service from '@/utils/request'

// @Tags CsTenant
// @Summary 创建CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenant true "创建CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CsTenant/createCsTenant [post]
export const createCsTenant = (data) => {
  return service({
    url: '/csTenant/createCsTenant',
    method: 'post',
    data
  })
}

// @Tags CsTenant
// @Summary 删除CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenant true "删除CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CsTenant/deleteCsTenant [delete]
export const deleteCsTenant = (data) => {
  return service({
    url: '/csTenant/deleteCsTenant',
    method: 'delete',
    data
  })
}

// @Tags CsTenant
// @Summary 删除CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CsTenant/deleteCsTenant [delete]
export const deleteCsTenantByIds = (data) => {
  return service({
    url: '/csTenant/deleteCsTenantByIds',
    method: 'delete',
    data
  })
}

// @Tags CsTenant
// @Summary 更新CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenant true "更新CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /CsTenant/updateCsTenant [put]
export const updateCsTenant = (data) => {
  return service({
    url: '/csTenant/updateCsTenant',
    method: 'put',
    data
  })
}

// @Tags CsTenant
// @Summary 用id查询CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CsTenant true "用id查询CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /CsTenant/findCsTenant [get]
export const findCsTenant = (params) => {
  return service({
    url: '/csTenant/findCsTenant',
    method: 'get',
    params
  })
}

// @Tags CsTenant
// @Summary 分页获取CsTenant列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CsTenant列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CsTenant/getCsTenantList [get]
export const getCsTenantList = (params) => {
  return service({
    url: '/csTenant/getCsTenantList',
    method: 'get',
    params
  })
}



export const setTenantApis = (data) =>{
  return service({
    url: '/csTenant/setTenantApis',
    method: 'put',
    data
  })
}

export const setTenantMenus = (data) =>{
  return service({
    url: '/csTenant/setTenantMenus',
    method: 'put',
    data
  })
}


export const getApisByTenantID = (params) =>{
  return service({
    url: '/csTenant/getApisByTenantID',
    method: 'get',
    params
  })
}

export const getMenusByTenantID = (params) =>{
  return service({
    url: '/csTenant/getMenusByTenantID',
    method: 'get',
    params
  })
}

export const createCsTenantAdmin = (tenantID) =>{
  return service({
    url: '/csTenant/createCsTenantAdmin?tenantID='+tenantID,
    method: 'post'
  })
}
