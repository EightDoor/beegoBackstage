import { RequestQueryBuilder } from '@nestjsx/crud-request'
import type { CreateQueryParams } from '@nestjsx/crud-request/lib/interfaces'
import log from '@/utils/log'

interface ParamsType {
  page: {
    currentPage: number
    pageSize: number
  }
  sort: any
  form: any
}
/**
 * 接口请求 查询条件
 * @param params
 */
const searchParam = (params: any) => {
  let param = '?'
  for (const item in params)
    param += `${item}=${params[item]}&`

  param = param.substring(0, param.length - 1)
  return param
}

/**
 * crud 请求
 * @param params
 * @returns
 */
const crudSearchParam = (params: ParamsType) => {
  log.d(params, 'crud请求参数')
  const result: CreateQueryParams = {
    limit: params.page.pageSize ?? 10,
    page: params.page.currentPage ?? 1,
  }
  return `?${RequestQueryBuilder.create(result).query()}`
}
export { searchParam, crudSearchParam }
