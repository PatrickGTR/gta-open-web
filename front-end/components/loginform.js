import Router from "next/router";
import { useState } from "react";
import { useToasts } from "react-toast-notifications";
import useStore from "../store/user";
import sendRequest from "../utils/sendRequest";

const LoginForm = () => {
  const setLogin = useStore((state) => state.setLoginStatus);
  const isLoggedIn = useStore((state) => state.loginStatus);
  const { addToast } = useToasts();

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
      const response = await sendRequest("POST", "user", {
        body: JSON.stringify({
          username: accountDetails.username,
          password: accountDetails.password,
        }),
      });

      const { msg } = await response.json();
      if (response.status === 200) {
        setLogin(true);

        addToast(msg, {
          appearance: "success",
        });
        Router.push(`/dashboard`);
      } else {
        addToast(msg, {
          appearance: "error",
        });
        return;
      }
    } catch (e) {
      addToast("Could not connect to the API, please contact a developer.", {
        appearance: "error",
      });
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
