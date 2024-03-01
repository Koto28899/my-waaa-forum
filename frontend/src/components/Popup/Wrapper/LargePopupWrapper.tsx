import React from "react";

const LargePopupWrapper = (
  { children }:
  { children: React.ReactNode }
) => {
  return (
    <div className="popup">
      <div className="popup__content--large">
        {children}
      </div>
    </div>
  )
}

export default LargePopupWrapper;