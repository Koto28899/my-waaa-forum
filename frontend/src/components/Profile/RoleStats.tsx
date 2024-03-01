type PropsRequired = {
  postsCount: number,
  commentsCount: number,
  repliesCount: number,
  likesCount: number,
  helpsCount: number,
  fansCount: number,
}

const stats = [
  {
    key: "postsCount",
    value: "📃发帖次数",
  },
  {
    key: "commentsCount",
    value: "📤评论次数",
  },
  {
    key: "repliesCount",
    value: "💬回复次数",
  },
  {
    key: "likesCount",
    value: "👍被赞次数",
  },
  {
    key: "helpsCount",
    value: "🤝帮助次数",
  },
  {
    key: "fansCount",
    value: "🍝粉丝总数",
  },
]

function RoleStats({ propsRequired }: { propsRequired: PropsRequired }) {
  return (
    <div className="role-stats">
      <span className="role-stats__header">数据统计</span>
      <ol className="role-stats__lists">
        {
          stats.map(item => {return(
            <li className="role-stats__list" key={item.key}>
              <span className="role-stats__list__key">{item.value}: </span>
              <span className="role-stats__list__value">{propsRequired[item.key as keyof PropsRequired]}</span>
            </li>
          )})
        }
      </ol>
    </div>
  )
}

export default RoleStats;