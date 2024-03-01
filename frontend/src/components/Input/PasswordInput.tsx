import Input from "./Input";

function PasswordInput() {
  return (
    <Input
      id="__password"
      type="password"
      name="__password"
      placeholder="密码"
      minLength={1}
      maxLength={9}
      required={true}
    />
  )
}

export default PasswordInput;