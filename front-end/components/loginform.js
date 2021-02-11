import Router from "next/router";
import { useState } from "react";
import useStore from "../store/user";

const LoginForm = () => {
  const setLogin = useStore((state) => state.setLoginStatus);
  const isLoggedIn = useStore((state) => state.loginStatus);

  const [errorMessage, setErrorMesssage] = useState(null);

  const [accountDetails, setAccountDetails] = useState({
    username: "",
    password: "",
  });

  const onChange = (e) => {
    e.preventDefault();
    setAccountDetails({ ...accountDetails, [e.target.name]: e.target.value });
  };

  const doLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8000/user/", {
        method: "POST",
        body: JSON.stringify({
          username: accountDetails.username,
          password: accountDetails.password,
        }),
        credentials: "include",
      });

      if (response.status === 200) {
        setLogin(true);

        Router.push(`/dashboard`);
        return;
      }

      setErrorMesssage(
        "Invalid credentials, try again with valid credentials.",
      );
    } catch {
      setErrorMesssage(
        "Could not connect to the API, please contact a developer.",
      );
    }
  };

  return (
    <>
      {!isLoggedIn && (
        <form>
          <label>Username</label>
          <input
            name="username"
            type="text"
            placeholder="Username"
            onChange={onChange}
          />

          <label>Password</label>
          <h6 className="error-msg">{errorMessage !== null && errorMessage}</h6>
          <input
            name="password"
            type="password"
            placeholder="Passowrd"
            onChange={onChange}
          />

          <input
            className="button-primary"
            type="submit"
            value="Login"
            onClick={doLogin}
          />
        </form>
      )}
    </>
  );
};

export default LoginForm;
