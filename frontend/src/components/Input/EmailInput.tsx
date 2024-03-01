import Input from "./Input";

function EmailInput() {
  return (
    <Input 
      id="__email"
      type="text"
      name="__email"
      placeholder="电子邮件地址"
      minLength={3}
      maxLength={80}
      required={true}
    />
  )
}

export default EmailInput;