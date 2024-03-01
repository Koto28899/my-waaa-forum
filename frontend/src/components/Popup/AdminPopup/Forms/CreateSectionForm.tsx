'use client'

import SubmitButton from "@/components/Button/SubmitButton";
import { useFormState } from "react-dom";
import { createSectionAction } from "../Actions/CreateSectionAction";
import StatementTextarea from "@/components/Textarea/StatementTextarea";
import { ManagerIDInput, NewSectionNameInput } from "@/components/Input";

const initialState = {
  message: "提 交"
}

function CreateSecitonForm() {
  const [state, formAction] = useFormState(createSectionAction, initialState)

  return (
    <div className="form--create-section">
      <h3>新增版块</h3>
      <form action={formAction}>
        <ManagerIDInput/>
        <NewSectionNameInput/>
        <StatementTextarea placeholder="论坛简介"/>
        <SubmitButton propsRequired={{
          loadingText: "提交中...",
          text: state.message
        }}/>
      </form>
    </div>
  )
}

export default CreateSecitonForm;