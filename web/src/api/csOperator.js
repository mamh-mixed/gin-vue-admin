import service from '@/utils/request'

// @Tags CsOperator
// @Summary 创建CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperator true "创建CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperator/createCsOperator [post]
export const createCsOperator = (data) => {
  return service({
    url: '/csOperator/createCsOperator',
    method: 'post',
    data
  })
}

// @Tags CsOperator
// @Summary 删除CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperator true "删除CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperator/deleteCsOperator [delete]
export const deleteCsOperator = (data) => {
  return service({
    url: '/csOperator/deleteCsOperator',
    method: 'delete',
    data
  })
}

// @Tags CsOperator
// @Summary 删除CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperator/deleteCsOperator [delete]
export const deleteCsOperatorByIds = (data) => {
  return service({
    url: '/csOperator/deleteCsOperatorByIds',
    method: 'delete',
    data
  })
}

// @Tags CsOperator
// @Summary 更新CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperator true "更新CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csOperator/updateCsOperator [put]
export const updateCsOperator = (data) => {
  return service({
    url: '/csOperator/updateCsOperator',
    method: 'put',
    data
  })
}

// @Tags CsOperator
// @Summary 用id查询CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CsOperator true "用id查询CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csOperator/findCsOperator [get]
export const findCsOperator = (params) => {
  return service({
    url: '/csOperator/findCsOperator',
    method: 'get',
    params
  })
}

// @Tags CsOperator
// @Summary 分页获取CsOperator列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CsOperator列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperator/getCsOperatorList [get]
export const getCsOperatorList = (params) => {
  return service({
    url: '/csOperator/getCsOperatorList',
    method: 'get',
    params
  })
}

export const createCsOperatorAdmin = (operatorID) =>{
  return service({
    url: '/csOperator/createCsOperatorAdmin?operatorID='+operatorID,
    method: 'post'
  })
}



export const setOperatorApis = (data) =>{
  return service({
    url: '/csOperator/setOperatorApis',
    method: 'put',
    data
  })
}

export const setOperatorMenus = (data) =>{
  return service({
    url: '/csOperator/setOperatorMenus',
    method: 'put',
    data
  })
}




export const getApisByOperatorID = (params) =>{
  return service({
    url: '/csOperator/getApisByOperatorID',
    method: 'get',
    params
  })
}

export const getMenusByOperatorID = (params) =>{
  return service({
    url: '/csOperator/getMenusByOperatorID',
    method: 'get',
    params
  })
}
