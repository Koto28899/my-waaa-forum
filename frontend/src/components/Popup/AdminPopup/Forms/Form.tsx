import AdminPopupFormEnum from "../AdminPopupForm";
import CreateSecitonForm from "./CreateSectionForm";
import UpdateRoleTypeForm from "./UpdateRoleTypeForm";
import { Suspense } from "react";

type PropsRequired = {
  form: AdminPopupFormEnum
}

function AdminPopupForm({ propsRequired }: { propsRequired: PropsRequired }) {
  return (
    <Suspense fallback={ <p>Loading...</p> }>
      <div className="admin-popup__body__form">
        {(() => {
          switch (propsRequired.form) {
            case AdminPopupFormEnum.createSection: {
              return <CreateSecitonForm />
            }
            case AdminPopupFormEnum.updateRoleType: {
              return <UpdateRoleTypeForm />
            }
            default: {
              return (
                <h3>你可以在左侧选择你要进行的行为</h3>
              )
            }
          }
        })()}
      </div>
    </Suspense>
  )
}

export default AdminPopupForm;