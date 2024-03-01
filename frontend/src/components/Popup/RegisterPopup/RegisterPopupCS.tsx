'use client'

import { useFormState } from "react-dom";
import SubmitButton from "../../Button/SubmitButton";
import BeforeAuthPopupFormEnum from "@/components/Header/RightSideHeader/BeforeAuthPopupForm";
import PopupWrapper from "../Wrapper/PopupWrapper";
import { registerActionCS } from "./RegisterActionCS";
import { useEffect } from "react";
import { ParsedCookie } from "@/components/Header/RightSideHeader/RightSideHeaderCS";
import EmailInput from "@/components/Input/EmailInput";
import NameInput from "@/components/Input/NameInput";
import PasswordInput from "@/components/Input/PasswordInput";
import PasswordAgainInput from "@/components/Input/PasswordAgainInput";

type PropsRequired = {
  setPopup: React.Dispatch<React.SetStateAction<BeforeAuthPopupFormEnum>>
  setParsedCookie: React.Dispatch<React.SetStateAction<ParsedCookie>>
}

const initialState = {
  message: "提 交",
}

function RegisterPopupCS({ propsRequired }: { propsRequired: PropsRequired }) {
  const [state, formAction] = useFormState(registerActionCS, initialState)
  
  function closeHandler() {
    propsRequired.setPopup(BeforeAuthPopupFormEnum.null)
  }

  useEffect(() => {
    if (state.message === "注册成功") {
      propsRequired.setParsedCookie(curr => {return {...curr, isLogin: true}})
      propsRequired.setPopup(BeforeAuthPopupFormEnum.null)
    }
  }, [state])

  return (
    <PopupWrapper>
      <div className="register-popup">
        <div className="register-popup__header">
          <h3>注册论坛账户</h3>
        </div>
        <button className="button--close-popup" onClick={closeHandler}>❌</button>
        <form className="register-popup__form" action={formAction}>
          <EmailInput />
          <NameInput placeholder="昵称"/>
          <PasswordInput />
          <PasswordAgainInput />
          <SubmitButton propsRequired={{
            loadingText: "提交中...",
            text: state.message
          }}/>
        </form>
        <div className="register-popup__portal">
          <span>已经拥有账号？</span><a onClick={e => propsRequired.setPopup(BeforeAuthPopupFormEnum.login)}>登录</a>
        </div>
      </div>
    </PopupWrapper>
  )
}

export default RegisterPopupCS;