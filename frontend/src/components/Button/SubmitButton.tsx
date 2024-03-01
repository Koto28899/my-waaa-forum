import { useFormStatus } from "react-dom";

type PropsRequired = {
  loadingText: string
  text: string
}

function SubmitButton({ propsRequired }: { propsRequired: PropsRequired }) {
  const { pending } = useFormStatus();
  return (
    <button className="button" type="submit" disabled={pending}>
      {pending ? propsRequired.loadingText : propsRequired.text}
    </button>
  );
}

export default SubmitButton;