'use client'

import LoginButton from "@/components/Button/LoginButton";
import RegisterButton from "@/components/Button/RegisterButton";
import ForgotPasswordPopup from "@/components/Popup/ForgotPasswordPopup";
import LoginPopupCS from "@/components/Popup/LoginPopup/LoginPopupCS";
import RegisterPopupCS from "@/components/Popup/RegisterPopup/RegisterPopupCS";

import { ParsedCookie } from "./RightSideHeaderCS";
import React, { SetStateAction, useState } from "react";
import BeforeAuthPopupFormEnum from "./BeforeAuthPopupForm";

type PropsRequired = {
  setParsedCookie: React.Dispatch<SetStateAction<ParsedCookie>>
}

function BeforeAuthRightSideHeaderCS({ propsRequired }: { propsRequired: PropsRequired }) {
  const [popup, setPopup] = useState(BeforeAuthPopupFormEnum.null)

  return (
    <>
      {(() => {
        switch (popup) {
        case BeforeAuthPopupFormEnum.login:
          return <LoginPopupCS propsRequired={{
            setPopup: setPopup,
            setParsedCookie: propsRequired.setParsedCookie
          }}/>
        case BeforeAuthPopupFormEnum.register:
          return <RegisterPopupCS propsRequired={{
            setPopup: setPopup,
            setParsedCookie: propsRequired.setParsedCookie
          }}/>
        case BeforeAuthPopupFormEnum.forgotPassword:
          return <ForgotPasswordPopup propsRequired={{
            setPopup: setPopup
          }}/>
        default:
          return <></>
        }
      })()}
      <div className="before-auth-buttons">
        <LoginButton propsRequired={{
            setPopup: setPopup
          }}/>
        <RegisterButton propsRequired={{
            setPopup: setPopup
          }}/>
      </div>
    </>
  )
}

export default BeforeAuthRightSideHeaderCS;