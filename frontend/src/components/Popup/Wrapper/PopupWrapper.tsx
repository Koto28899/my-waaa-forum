import React from "react";

const PopupWrapper = (
  { children }:
  { children: React.ReactNode }
) => {
  return (
    <div className="popup">
      <div className="popup__content">
        {children}
      </div>
    </div>
  )
}

export default PopupWrapper;