import request from '@/utils/request'

export function requestFirstCityAll() {
  return request({
    url: '/city/getfirstcity',
    method: 'get'
  })
}
export function requestSecondCity(cityname) {
  return request({
    url: '/city/getsecondcity',
    method: 'get',
    params: { cityname }
  })
}
export function requestThirdCity(cityname) {
  return request({
    url: '/city/getthirdcity',
    method: 'get',
    params: { cityname }
  })
}
