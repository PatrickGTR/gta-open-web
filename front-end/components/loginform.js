import Router from "next/router";
import { useState } from "react";
import useStore from "../store/user";
import sendRequest from "../utils/sendRequest";
import { useMessage } from "../utils/message";

const LoginForm = () => {
  const setLogin = useStore((state) => state.setLoginStatus);
  const isLoggedIn = useStore((state) => state.loginStatus);

  const [accountDetails, setAccountDetails] = useState({
    username: "",
    password: "",
  });

  const { notifyError, notifySuccess } = useMessage();

  const onChange = (e) => {
    e.preventDefault();
    setAccountDetails({ ...accountDetails, [e.target.name]: e.target.value });
  };

  const doLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await sendRequest("POST", "user", {
        body: JSON.stringify({
          username: accountDetails.username,
          password: accountDetails.password,
        }),
      });

      const { msg } = await response.json();
      if (response.status === 200) {
        setLogin(true);

        notifySuccess(msg);
        Router.push(`/dashboard`);
      } else {
        notifyError(msg);
        return;
      }
    } catch (e) {
      notifyError("Could not connect to the API, please contact a developer.");
    }
  };

  return (
    <>
      {!isLoggedIn && (
        <>
          <h3>User Control Panel</h3>
          <form>
            <input
              placeholder="Username"
              name="username"
              type="text"
              onChange={onChange}
            />
            <input
              placeholder="Password"
              name="password"
              type="password"
              onChange={onChange}
            />
            <input
              style={{ width: "100%" }}
              className="button-primary"
              type="submit"
              value="Sign in"
              onClick={doLogin}
            />
          </form>
        </>
      )}
    </>
  );
};

export default LoginForm;
