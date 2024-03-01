import LeftSideHeader from "./LeftSideHeader";
import RightSideHeaderCS from "./RightSideHeader/RightSideHeaderCS";

const Header = () => {
  return (
    <header className="header">
      <LeftSideHeader />
      <RightSideHeaderCS />
    </header>
  );
}

export default Header;