import request from '@/utils/request'
import log from '@/utils/log'

const apiPrefix = '/demoCrud'

export function GetList(query: any) {
  return request({
    url: `${apiPrefix}${query}`,
    method: 'GET',
  })
}

export function AddObj(obj) {
  return request({
    url: apiPrefix,
    method: 'POST',
    body: obj,
  })
}

export function UpdateObj(obj) {
  return request({
    url: `${apiPrefix}`,
    method: 'PUT',
    body: obj,
  })
}

export function DelObj(obj) {
  log.i(obj, 'obj')
  return request({
    url: `${apiPrefix}/${obj.row.id}`,
    method: 'DELETE',
  })
}

export function GetObj(id) {
  return request({
    url: `${apiPrefix}/${id}`,
    method: 'GET',
    params: { id },
  })
}
