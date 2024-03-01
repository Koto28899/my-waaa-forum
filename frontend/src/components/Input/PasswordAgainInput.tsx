import Input from "./Input";

function PasswordAgainInput() {
  return (
    <Input
      id="__passwordAgain"
      type="password"
      name="__passwordAgain"
      placeholder="确认密码"
      minLength={1}
      maxLength={9}
      required={true}
    />
  )
}

export default PasswordAgainInput;