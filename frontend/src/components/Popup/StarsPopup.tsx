import LargePopupWrapper from "./Wrapper/LargePopupWrapper";
import PopupWrapper from "./Wrapper/PopupWrapper";

type PropsRequired = {
  setIsPopup: React.Dispatch<React.SetStateAction<boolean>>
}

function StarsPopup({ propsRequired }: { propsRequired: PropsRequired }) {
  function closeHandler(e: React.MouseEvent) {
    propsRequired.setIsPopup(false)
  }

  return (
    <LargePopupWrapper>
      <div className="stars-popup">
        <h3 className="stars-popup__header">收藏列表</h3>
          <button className="button--close-popup" onClick={closeHandler}>❌</button>
          <div className="stars-popup__body">
          
        </div>
      </div>
    </LargePopupWrapper>
  );
}

export default StarsPopup;