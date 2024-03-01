import Image from "next/image"
import commonAvatar from "@/app/commonAvatar.jpg"

type PropsRequired = {
  createdAt: string,
  roleName: string,
  type: string,
  statement: string
}

const info = [
  {
    key: "id",
    value: "🆔用户 ID"
  },
  {
    key: "createdAt",
    value: "📅加入时间",
  },
  {
    key: "roleName",
    value: "🙂论坛昵称",
  },
  {
    key: "type",
    value: "🏷️账号类型",
  },
  {
    key: "statement",
    value: "📜个性签名",
  },
]

function RoleInfo({ propsRequired }: { propsRequired: PropsRequired }) {
  return (
    <div className="role-info">
      <Image src={commonAvatar} alt="avatar"/>
      <span className="role-info__header">个人简介</span>
      <ol className="role-info__lists">
        {
          info.map(item => {return(
            <li className="role-info__list" key={item.key}>
              <span className="role-info__list__key">{item.value}:</span>
              <p className="role-info__list__value">{propsRequired[item.key as keyof PropsRequired]}</p>
            </li>
          )})
        }
      </ol>
    </div>
  )
}

export default RoleInfo;