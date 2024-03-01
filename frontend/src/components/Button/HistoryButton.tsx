'use client'

import { useState } from "react";
import HistoryPopup from "../Popup/HistoryPopup";

function HistoryButton() {
  const [isPopup, setIsPopup] = useState(false)

  function onClickHandler() {
    setIsPopup(true)
  }

  return (
    <>
      {isPopup && <HistoryPopup propsRequired={{
        setIsPopup: setIsPopup
      }}/>}
      <button className="button--after-auth" onClick={onClickHandler}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
          <path strokeLinecap="round" strokeLinejoin="round" d="M12 6v6h4.5m4.5 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
        </svg>
        <span>历史</span>
      </button>
    </>
  )
}

export default HistoryButton;