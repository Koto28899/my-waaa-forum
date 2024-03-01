import Textarea from "./Textarea";

function StatementTextarea({ placeholder }: { placeholder: string }) {
  return (
    <Textarea propsRequired={{
      id: "__statement",
      placeholder: placeholder,
      name: "__statement",
      rows: 5,
      cols: 28,
      minLength: 0,
      maxLength: 42,
      required: false
    }}/>
  )
}

export default StatementTextarea;