import request from "@/api/client";

export function SEARCH(params) {
  return request({
    url: "/search",
    method: "get",
    params: params,
  });
}
