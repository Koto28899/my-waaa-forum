'use client'

import { useState } from "react"
import AdminPopup from "../Popup/AdminPopup/AdminPopup"

export default function AdminButton() {
  const [isPopup, setIsPopup] = useState(false)

  function onClickHandler() {
    setIsPopup(true)
  }

  return (
    <>
      {isPopup && <AdminPopup propsRequired={{
          setIsPopup: setIsPopup
      }}/>}
      <button className="button--after-auth" onClick={onClickHandler}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
          <path strokeLinecap="round" strokeLinejoin="round" d="m6.75 7.5 3 2.25-3 2.25m4.5 0h3m-9 8.25h13.5A2.25 2.25 0 0 0 21 18V6a2.25 2.25 0 0 0-2.25-2.25H5.25A2.25 2.25 0 0 0 3 6v12a2.25 2.25 0 0 0 2.25 2.25Z" />
        </svg>
        <span>管理</span>
      </button>
    </>
  )
}