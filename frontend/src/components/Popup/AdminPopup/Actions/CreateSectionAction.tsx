'use client'

import timeoutFetch from "@/utils/TimeoutFetch"

async function createSectionAction(prevState: any, formData: FormData) {
  var resultMessage = "未知错误"

  const requestBody = JSON.stringify({
    sectionName: formData.get("__name"),
    statement: formData.get("__statement"),
    managerID: parseInt(formData.get("__ManagerID") as string)
  })

  await timeoutFetch(2000, "http://localhost:8080/api/section", {
    method: "POST",
    body: requestBody,
    credentials: "include"
  }).then(response => {
    if (!response.ok) {
      return response.json().then(json => {
        return Promise.reject({
          status: response.status,
          error: json.error
        })
      })
    } else {
      resultMessage = "创建成功"
    }
  }).catch(err => {
    switch (err.status) {
      case 400: {
        resultMessage = err.error == undefined ? "请求体错误" : err.error
        break
      }
      case 500: {
        resultMessage = "服务端错误"
        break
      }
      case 504: {
        resultMessage = "请求超时"
        break
      }
      default: {
        resultMessage = "请求被拒绝"
      }
    }
  })
  return {
    message: resultMessage
  }
}

export {
  createSectionAction
}