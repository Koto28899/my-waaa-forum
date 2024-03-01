import Link from "next/link";
import Image from "next/image";
import commonAvatar from "@/app/commonAvatar.jpg"

type PropsRequired = {
  roleID: number
}

function AvatarButton({ propsRequired }: { propsRequired: PropsRequired }) {
  return (
    <Link className="button--avatar" href={`/profile/${propsRequired.roleID}`}>
      <Image src={commonAvatar} width={48} alt="avatar"/>
    </Link>
  )
}

export default AvatarButton;