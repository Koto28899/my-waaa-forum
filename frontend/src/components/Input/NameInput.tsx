import Input from "./Input";

function NameInput({ placeholder }: { placeholder: string }) {
  return (
    <Input
      id="__name"
      type="text"
      name="__name"
      placeholder={placeholder}
      minLength={2}
      maxLength={8}
      required={true}
    />
  )
}

export default NameInput;