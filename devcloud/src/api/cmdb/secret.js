import request from "@/api/client";

export function LIST_SECRET(params) {
  return request({
    url: "/secrets",
    method: "get",
    params: params,
  });
}

export function ADD_SECRET(data, params) {
  return request({
    url: "/secrets",
    method: "post",
    data,
    params,
  });
}

export function DELETE_SECRET(id, params) {
  return request({
    url: `/secrets/${id}`,
    method: "delete",
    params,
  });
}
