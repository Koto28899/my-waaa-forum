import AdminButton from "@/components/Button/AdminButton";
import AvatarButton from "@/components//Button/AvatarButton";
import FollowsButton from "@/components//Button/FollowsButton";
import HistoryButton from "@/components//Button/HistoryButton";
import NotificationsButton from "@/components//Button/NotificationsButton";
import PostButton from "@/components//Button/PostButton";
import SettingButton from "@/components//Button/SettingButton";
import StarsButton from "@/components//Button/StarsButton";
import { ParsedCookie } from "./RightSideHeaderCS";

type PropsRequired = {
  parsedCookie: ParsedCookie
}

const AfterAuthRightSideHeader = ({ propsRequired }: { propsRequired: PropsRequired }) => {
  return (
    <div className="after-auth-buttons">
      <AvatarButton propsRequired={{
        roleID: propsRequired.parsedCookie.roleID
      }}/>
      <NotificationsButton />
      <FollowsButton />
      <StarsButton />
      <HistoryButton />
      <PostButton />
      <SettingButton />
      {(propsRequired.parsedCookie.roleType == "admin") && <AdminButton />}
    </div>
  )
}

export default AfterAuthRightSideHeader;