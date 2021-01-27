import Navbar from "../components/navbar";
import Header from "../components/header";

const Layout = ({ title, children }) => {
  return (
    <>
      <Header title={title} />
      <Navbar />
      <div className="container">{children}</div>
      <footer className="footer">
        Made with 💖 by Patrick Subang | Copyright 2021 GTA Open
      </footer>
    </>
  );
};

export default Layout;
