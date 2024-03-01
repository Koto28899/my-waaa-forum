import React from "react";
import PopupWrapper from "./Wrapper/PopupWrapper";
import LargePopupWrapper from "./Wrapper/LargePopupWrapper";

type PropsRequired = {
  setIsPopup: React.Dispatch<React.SetStateAction<boolean>>
}

function SettingPopup({ propsRequired }: { propsRequired: PropsRequired }) {
  function closeHandler(e: React.MouseEvent) {
    propsRequired.setIsPopup(false)
  }

  return (
    <LargePopupWrapper>
      <div className="setting-popup">
        <h3 className="setting-popup__header">设置</h3>
          <button className="button--close-popup" onClick={closeHandler}>❌</button>
          <div className="setting-popup__body">
          
        </div>
      </div>
    </LargePopupWrapper>
  );
}

export default SettingPopup;