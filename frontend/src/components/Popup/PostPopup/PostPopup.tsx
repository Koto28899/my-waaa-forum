import React from "react";
import LargePopupWrapper from "@/components/Popup/Wrapper/LargePopupWrapper";

type PropsRequired = {
  setIsPopup: React.Dispatch<React.SetStateAction<boolean>>
}

function PostPopup({ propsRequired }: { propsRequired: PropsRequired }) {
  function closeHandler(e: React.MouseEvent) {
    propsRequired.setIsPopup(false)
  }

  return (
    <LargePopupWrapper>
      <div className="post-popup">
        <h3 className="post-popup__header">发帖</h3>
        <button className="button--close-popup" onClick={closeHandler}>❌</button>
        <div className="post-popup__body">
        </div>
      </div>
    </LargePopupWrapper>
  )
}

export default PostPopup;