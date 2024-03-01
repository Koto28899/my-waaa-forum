'use client'

type ReturnType = {
  message: string
}

async function loginActionCS(prevState: any, formData: FormData): Promise<ReturnType> {
  let resultMessage: string = "未知错误"
  const requestBody = JSON.stringify({
    email: formData.get("__email"),
    pwd: formData.get("__password")
  })
  await fetch("http://localhost:8080/api/login", {
    method: "POST",
    body: requestBody,
    credentials: "include"
  }).then(response => {
    if (!response.ok) {
      return Promise.reject({
        status: response.status,
        statusText: response.statusText
      })
    } else {
      resultMessage = "登录成功"
    }
  }).catch(err => {
    switch (err.status) {
      case 400: {
        resultMessage = `请求体错误`
        break
      }
      case 401: {
        resultMessage = "账号不存在或密码错误"
        break
      }
      case 500: {
        resultMessage = "服务器错误"
        break
      }
      default: resultMessage = "请求被拒绝"
    }
  })
  return {
    message: resultMessage
  }
}

export {
  loginActionCS
}