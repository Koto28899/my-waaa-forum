import React from "react"
import AdminPopupFormEnum from "./AdminPopupForm"

type PropsRequired = {
  setForm: React.Dispatch<React.SetStateAction<AdminPopupFormEnum>>
}

const selects = [
  {
    form: AdminPopupFormEnum.createSection,
    text: "新增版块"
  },
  {
    form: AdminPopupFormEnum.updateRoleType,
    text: "修改成员类型"
  },
]

function AdminPopupSelect({ propsRequired }: { propsRequired: PropsRequired }) {
  return (
    <div className="admin-popup__body__select">
      {selects.map(item => <button className="button" key={item.text} onClick={e => propsRequired.setForm(item.form)}>{item.text}</button>)}
    </div>
  )
}

export default AdminPopupSelect