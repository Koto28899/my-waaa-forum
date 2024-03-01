import BeforeAuthPopupFormEnum from "../Header/RightSideHeader/BeforeAuthPopupForm";

type PropsRequired = {
  setPopup: React.Dispatch<React.SetStateAction<BeforeAuthPopupFormEnum>>
}

function LoginButton({ propsRequired }: { propsRequired: PropsRequired }) {
  function onClickHandler() {
    propsRequired.setPopup(BeforeAuthPopupFormEnum.login)
  }

  return (
    <button className="button--login" onClick={onClickHandler}>登录</button>
  )
}

export default LoginButton;