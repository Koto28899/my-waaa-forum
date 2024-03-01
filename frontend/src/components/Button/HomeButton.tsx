import Link from "next/link";
import Image from 'next/image';
import logoImage from "@/app/favicon.ico";

const HomeButton = () => {
  return (
    <Link className="button--logo" href="/">
      <Image className="button--logo__img" height={48} src={logoImage} alt="Logo"/>
      <span className="button--logo__label">W.A.A.A</span>
    </Link>
  )
}

export default HomeButton;