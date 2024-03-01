'use client'

import React, { useEffect } from "react";
import PopupWrapper from "../Wrapper/PopupWrapper";
import BeforeAuthPopupFormEnum from "@/components/Header/RightSideHeader/BeforeAuthPopupForm";
import { useFormState } from "react-dom";
import { loginActionCS } from "./LoginActionCS";
import SubmitButton from "@/components/Button/SubmitButton";
import { ParsedCookie } from "@/components/Header/RightSideHeader/RightSideHeaderCS";
import EmailInput from "@/components/Input/EmailInput";
import PasswordInput from "@/components/Input/PasswordInput";

type PropsRequired = {
  setPopup: React.Dispatch<React.SetStateAction<BeforeAuthPopupFormEnum>>
  setParsedCookie: React.Dispatch<React.SetStateAction<ParsedCookie>>
}

const initialState = {
  message: "登 录",
}

function LoginPopupCS({ propsRequired }: { propsRequired: PropsRequired }) {

  const [state, formAction] = useFormState(loginActionCS, initialState)

  useEffect(() => {
    if (state.message === "登录成功") {
      propsRequired.setParsedCookie(curr => {return {...curr, isLogin: true}})
      propsRequired.setPopup(BeforeAuthPopupFormEnum.null)
    }
  }, [state])

  function closeHandler() {
    propsRequired.setPopup(BeforeAuthPopupFormEnum.null)
  }

  return (
    <PopupWrapper>
      <div className="login-popup">
        <div className="login-popup__header">
          <h3>登录论坛账户</h3>
        </div>
        <button className="button--close-popup" onClick={closeHandler}>❌</button>
        <form className="login-popup__form" action={formAction}>
          <EmailInput />
          <PasswordInput />
          <SubmitButton propsRequired={{
            loadingText: "登录中",
            text: state.message
          }}/>
        </form>
        <div className="login-popup__portal">
          <span>还未注册？</span><a onClick={e => propsRequired.setPopup(BeforeAuthPopupFormEnum.register)}>注册</a>
        </div>
        <div className="login-popup__portal">
          <span>忘记密码？</span><a onClick={e => propsRequired.setPopup(BeforeAuthPopupFormEnum.forgotPassword)}>找回</a>
        </div>
      </div>
    </PopupWrapper>
  )
}

export default LoginPopupCS;