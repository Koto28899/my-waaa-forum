import RoleInfo from "@/components/Profile/RoleInfo"
import RolePosts from "@/components/Profile/RolePosts"
import RoleStats from "@/components/Profile/RoleStats"
import timeoutFetch from "@/utils/TimeoutFetch"
import { revalidateTag } from "next/cache"
import { notFound } from "next/navigation"
import { Suspense } from "react"

type RoleBasis = {
  id: number,
  createdAt: string,
  roleName: string,
  type: string,
  statement: string,
  postsCount: number,
  commentsCount: number,
  repliesCount: number,
  likesCount: number,
  helpsCount: number,
  fansCount: number,
}

async function getRoleBasis(roleID: string) {
  const URL = `http://localhost:8080/api/role/basis?id=${roleID}`
  const response = await timeoutFetch(10000, URL, {
    method: "GET",
    next: {
      tags: ["roleBasis"]
    }
  })
  if (!response.ok) {
    switch (response.status) {
      case 400: {
        throw new Error('ID 无效')
      }
      case 404: {
        throw new Error('用户不存在')
      }
      case 500: {
        throw new Error('服务端错误')
      }
      case 504: {
        throw new Error('请求已超时')
      }
      default: {
        throw new Error('客户端错误')
      }
    }
  }
  return response.json()
}

async function RoleProfilePage({ params }: { params: { id: string } }) {
  const roleBasis: RoleBasis = await getRoleBasis(params.id)
  revalidateTag("roleBasis")
  
  return(
    <div className="profile">
      <RoleInfo  propsRequired={{
        ...roleBasis
      }}/>
      <RoleStats propsRequired={{
        ...roleBasis
      }}/>
      <RolePosts propsRequired={{
        roleID: params.id
      }}/>
    </div>
  )
}

export default RoleProfilePage