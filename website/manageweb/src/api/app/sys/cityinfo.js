import request from '@/utils/request'

export function requestAllCity1() {
  return request({
    url: '/cityinfo/getallcity1',
    method: 'get'
  })
}
export function requestCity2(city1) {
  return request({
    url: '/cityinfo/getcity2',
    method: 'get',
    params: { city1 }
  })
}
export function requestCity3(city2) {
  return request({
    url: '/cityinfo/getcity3',
    method: 'get',
    params: { city2 }
  })
}
