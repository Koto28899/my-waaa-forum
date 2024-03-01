import React, { SetStateAction } from "react";
import LargePopupWrapper from "./Wrapper/LargePopupWrapper";

type PropRequired = {
  setIsPopup: React.Dispatch<SetStateAction<boolean>>
}

function NotificationsPopup({ propsRequired }: { propsRequired: PropRequired }) {
  function closeHandler(e: React.MouseEvent) {
    propsRequired.setIsPopup(false)
  }

  return (
    <LargePopupWrapper>
      <div className="notifications-popup">
        <h3 className="notifications-popup__header">消息列表</h3>
          <button className="button--close-popup" onClick={closeHandler}>❌</button>
          <div className="notifications-popup__body">
          
        </div>
      </div>
    </LargePopupWrapper>
  );
}

export default NotificationsPopup;