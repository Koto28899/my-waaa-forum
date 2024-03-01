'use client'

type ReturnType = {
  message: string
}

async function registerActionCS(prevState: any, formData: FormData): Promise<ReturnType> {
  var resultMessage = "未知错误"
  const registerPopupPassword = formData.get("__password");
  const registerPopupPasswordAgin = formData.get("__passwordAgain")
  if (registerPopupPassword != registerPopupPasswordAgin ) {
    return {
      message: (
        "前后输入密码不一致"
      )
    }
  }
  const requestBody = JSON.stringify({
    email: formData.get("__email"),
    name: formData.get("__name"),
    pwd: registerPopupPassword
  })
  await fetch("http://localhost:8080/api/register", {
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
      resultMessage = "注册成功"
    }
  }).catch(err => {
    switch (err.status) {
      case 400: {
        resultMessage = `请求体错误`
        break
      }
      case 500: {
        resultMessage = "服务器错误"
        break
      }
      case 403: {
        resultMessage = "邮箱已使用 "
        break
      }
      default: resultMessage = "请求被拒绝"
    }
  })
  return {
    message: resultMessage
  }
}

export { registerActionCS }