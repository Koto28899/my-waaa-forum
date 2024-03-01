import React from "react";
import PopupWrapper from "./Wrapper/PopupWrapper";
import BeforeAuthPopupFormEnum from "../Header/RightSideHeader/BeforeAuthPopupForm";
import EmailInput from "../Input/EmailInput";

type PropsRequired = {
  setPopup: React.Dispatch<React.SetStateAction<BeforeAuthPopupFormEnum>>
}

const ForgotPasswordPopup = ({ propsRequired }: { propsRequired: PropsRequired }) => {

  function closeHandler() {
    propsRequired.setPopup(BeforeAuthPopupFormEnum.null)
  }

  async function submitHandler(formData: FormData) {
    try {
      await Promise.race([
        // fetch here
        new Promise((_, reject) => setTimeout(reject, 1000)),
      ])
    } catch (err) {
      
    }
  }

  return (
    <PopupWrapper>
      <div className="forgot-password-popup">
        <div className="forgot-password-popup__header">
          <h2>找回论坛账户</h2>
          <p>如果您提供的邮箱绑定了论坛账户，我们将向您的邮箱发送一封邮件以重置密码</p>
        </div>
        <button className="button--close-popup" onClick={closeHandler}>❌</button>
        <form className="forgot-password-popup__form" action={formData => submitHandler(formData)}>
          <EmailInput />
          <button className="button" type="submit">发送邮件</button>
        </form>
        <div className="forgot-password-popup__portal">
          <span>还未注册？</span><a onClick={e => propsRequired.setPopup(BeforeAuthPopupFormEnum.register)}>注册</a>
        </div>
        <div className="forgot-password-popup__portal">
          <span>回忆起密码？</span><a onClick={e => propsRequired.setPopup(BeforeAuthPopupFormEnum.login)}>登录</a>
        </div>
      </div>
    </PopupWrapper>
  )
}

export default ForgotPasswordPopup;