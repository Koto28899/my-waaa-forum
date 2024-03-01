'use client'

import { useEffect, useState } from "react";
import BeforeAuthRightSideHeaderCS from "./BeforeAuthRightSideHeaderCS";
import AfterAuthRightSideHeader from "./AfterAuthRightSideHeader";

type ParsedCookie = {
  isLogin: boolean;
  roleID: number,
  roleType: string,
}

function RightSideHeaderCS() {
  const [parsedCookie, setParsedCookie] = useState<ParsedCookie>({
    isLogin: false,
    roleID: 0,
    roleType: ""
  })

  useEffect(() => {
    fetch("http://localhost:8080/api/parseCookie", {
      method: "GET",
      credentials: "include",
    }).then(response => {
      if (!response.ok) {
        return Promise.reject({
          status: response.status,
          statusText: response.statusText
        })
      } else {
        response.json().then(json => {
          setParsedCookie({...json, isLogin:true})
        })
      }
    }).catch(err => {

    })
  }, [])

  return (
    <div className="header--right">
      {
        parsedCookie.isLogin
        ? <AfterAuthRightSideHeader propsRequired={{
          parsedCookie: parsedCookie
        }}/>
        : <BeforeAuthRightSideHeaderCS propsRequired={{
          setParsedCookie: setParsedCookie
        }}/>
      }
    </div>
  )
}

export default RightSideHeaderCS;
export type {
  ParsedCookie
}