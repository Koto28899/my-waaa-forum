'use client'

import { useFormState } from "react-dom";
import { updateRoleTypeAction } from "../Actions/UpdateRoleTypeAction";

const initialState = {
  message: "提交"
}

function UpdateRoleTypeForm() {
  const [state, formAction] = useFormState(updateRoleTypeAction, initialState)

  return (
    <div className="form--update-role-type">
      <h3>修改成员类型</h3>
      
    </div>
  )
}

export default UpdateRoleTypeForm;