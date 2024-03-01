'use client'

import React, { useState } from "react";
import LargePopupWrapper from "../Wrapper/LargePopupWrapper";
import AdminPopupSelect from "./Select";
import AdminPopupForm from "./Forms/Form";
import AdminPopupFormEnum from "./AdminPopupForm";

type PropRequired = {
  setIsPopup: React.Dispatch<React.SetStateAction<boolean>>
}

function AdminPopup({ propsRequired }: { propsRequired: PropRequired }) {
  const [form, setForm] = useState<AdminPopupFormEnum>(AdminPopupFormEnum.null)

  function closeHandler() {
    propsRequired.setIsPopup(false)
  }

  return (
    <LargePopupWrapper>
      <div className="admin-popup">
        <h3 className="admin-popup__header">管理员控制台</h3>
        <button className="button--close-popup" onClick={closeHandler}>❌</button>
        <div className="admin-popup__body">
          <AdminPopupSelect propsRequired={{
            setForm: setForm
          }}/>
          <AdminPopupForm propsRequired={{
            form: form
          }}/>
        </div>
      </div>
    </LargePopupWrapper>
  )
}

export default AdminPopup;