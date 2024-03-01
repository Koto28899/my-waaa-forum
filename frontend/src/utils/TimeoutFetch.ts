const reason = {
  status: 504,
  message: "请求超时"
}

function timeoutFetch(ms: number, URL: string, init?: RequestInit | undefined) {
  const controller = new AbortController()
  return Promise.race([

    fetch(URL, {
      ...init,
      signal: controller.signal
    }),

    new Promise((resolve) => {
      setTimeout(resolve, ms)
    }).then(() => {
      controller.abort(reason);
      return Promise.reject(reason)
    })

  ])
}

export default timeoutFetch;