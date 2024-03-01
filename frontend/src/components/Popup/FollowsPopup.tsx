import React from "react";
import LargePopupWrapper from "./Wrapper/LargePopupWrapper";

type PropRequired = {
  setIsPopup: React.Dispatch<React.SetStateAction<boolean>>
}

function FollowsPopup({ propsRequired }: { propsRequired: PropRequired }) {
  function closeHandler(e: React.MouseEvent) {
    propsRequired.setIsPopup(false)
  }

  return (
    <LargePopupWrapper>
      <div className="follows-popup">
        <h3 className="follows-popup__header">关注列表</h3>
          <button className="button--close-popup" onClick={closeHandler}>❌</button>
          <div className="follows-popup__body">
          
        </div>
      </div>
    </LargePopupWrapper>
  );
}

export default FollowsPopup;