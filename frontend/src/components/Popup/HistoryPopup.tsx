import React from "react"
import LargePopupWrapper from "./Wrapper/LargePopupWrapper"

type PropsRequired = {
  setIsPopup: React.Dispatch<React.SetStateAction<boolean>>
}

function HistoryPopup({ propsRequired }: { propsRequired: PropsRequired }) {
  function closeHandler(e: React.MouseEvent) {
    propsRequired.setIsPopup(false)
  }

  return (
    <LargePopupWrapper>
      <div className="history-popup">
        <h3 className="history-popup__header">历史记录</h3>
          <button className="button--close-popup" onClick={closeHandler}>❌</button>
          <div className="history-popup__body">
          
        </div>
      </div>
    </LargePopupWrapper>
  );
}

export default HistoryPopup