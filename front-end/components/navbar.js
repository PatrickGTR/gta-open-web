import Link from "next/link";
import useStore from "../store/user";

const NavBar = () => {
  const isLoggedIn = useStore((state) => state.loginStatus);

  const GlobalNav = () => {
    return (
      <>
        <li>
          <Link href="/">
            <a>Home</a>
          </Link>
        </li>
        <li>
          <Link href="/">
            <a>Forums</a>
          </Link>
        </li>
      </>
    );
  };

  const LoggedUserNav = () => {
    const clear = useStore((state) => state.clearState);
    const doLogout = (e) => {
      //clear();
    };

    return (
      <>
        <li>
          <Link href="/dashboard">
            <a>Dashboard</a>
          </Link>
        </li>
        <li>
          <Link href="/">
            <a onClick={doLogout}>Logout</a>
          </Link>
        </li>
      </>
    );
  };

  return (
    <header className="header">
      <a href="" className="logo">
        <img
          width="64"
          src="https://camo.githubusercontent.com/11857964d64562f7c921ba7ce05fd363ae4f0ed0654ecb24ac95ffa51aa4d241/68747470733a2f2f696d616765732d6578742d312e646973636f72646170702e6e65742f65787465726e616c2f39626e714d4a523842454c45674942503870795a7a58527432574a304e6d495770734e6a77637674644d732f68747470732f692e6962622e636f2f53524c7a7a636e2f6774616f70656e2d72656464616464792e706e67"
        />
      </a>
      <input className="menu-btn" type="checkbox" id="menu-btn" />
      <label className="menu-icon" htmlFor="menu-btn">
        <span className="navicon"></span>
      </label>
      <ul className="menu">
        <GlobalNav />
        {isLoggedIn && <LoggedUserNav />}
      </ul>
    </header>
  );
};

export default NavBar;
