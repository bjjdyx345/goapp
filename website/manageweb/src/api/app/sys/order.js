import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/order/list',
    method: 'get',
    params: query
  })
}
export function requestAll() {
  return request({
    url: '/order/getall',
    method: 'get'
  })
}

export function requestUpdate(data) {
  return request({
    url: '/order/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/order/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/order/delete',
    method: 'post',
    data
  })
}

export function requestDetail(id) {
  return request({
    url: '/order/detail',
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

/*
export function requestApartID(){
return request({
  url: '/card/detail',
  method: 'get',
  params: { id }
})
}
*/
