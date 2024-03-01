import BeforeAuthPopupFormEnum from "../Header/RightSideHeader/BeforeAuthPopupForm";

type PropsRequired = {
  setPopup: React.Dispatch<React.SetStateAction<BeforeAuthPopupFormEnum>>
}

function RegisterButton({ propsRequired }: { propsRequired: PropsRequired }) {
  function onClickHandler() {
    propsRequired.setPopup(BeforeAuthPopupFormEnum.register)
  }

  return (
    <button className="button--register" onClick={onClickHandler}>免费注册</button>
  )
}

export default RegisterButton;