import request from '@/api/client'

export function LIST_HOST(params) {
    return request({
        url: '/hosts',
        method: 'get',
        params: params
    })
}

export function GET_HOST(id, params) {
    return request({
      url: `/hosts/${id}`,
      method: "get",
      params: params,
    });
  }
  