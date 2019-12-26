import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/card/list',
    method: 'get',
    params: query
  })
}

export function requestquery(query) {
  return request({
    url: '/card/query',
    method: 'get',
    params: query
  })
}

export function requestAll() {
  return request({
    url: '/card/getall',
    method: 'get'
  })
}

export function requestUpdate(data) {
  return request({
    url: '/card/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/card/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/card/delete',
    method: 'post',
    data
  })
}

export function requestDetail(id) {
  return request({
    url: '/card/detail',
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

export function requestBindVillage(card_id, card_type, village_id) {
  return request({
    url: '/card/bindvillage',
    method: 'post',
    params: { card_id, card_type, village_id }
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
