type PropsRequired = {
  roleID: string
}

function RolePosts({ propsRequired }: { propsRequired: PropsRequired }) {
  return (
    <div className="role-posts">
      你可以在这里查找该用户发过的帖子
    </div>
  )
}

export default RolePosts;