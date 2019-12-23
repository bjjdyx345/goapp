import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/village/list',
    method: 'get',
    params: query
  })
}
export function requestAll() {
  return request({
    url: '/village/getall',
    method: 'get'
  })
}

export function requestUpdate(data) {
  return request({
    url: '/village/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/village/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/village/delete',
    method: 'post',
    data
  })
}

export function requestDetail(id) {
  return request({
    url: '/village/detail',
    method: 'get',
    params: { id }
  })
}
export function requestMenuButton(menucode) {
  return request({
    url: '/menu/menubuttonlist',
    method: 'get',
    params: { menucode }
  })
}

