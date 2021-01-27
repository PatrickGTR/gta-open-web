import Router from "next/router";
import { useState } from "react";

import useStore from "../store/user";

const SideBar = () => {
  const [accountDetails, setAccountDetails] = useState({
    username: "",
    password: "",
  });

  const [errorMessage, setErrorMesssage] = useState(null);

  const setLogin = useStore((state) => state.setLoginStatus);

  const doLogin = async (e) => {
    e.preventDefault();

    let formData = new FormData();
    formData.append("username", accountDetails.username);
    formData.append("password", accountDetails.password);

    const response = await fetch("http://localhost:8000/user/", {
      method: "POST",
      body: formData,
    });

    if (response.status === 200) {
      const { token } = await response.json();

      localStorage.setItem("jwt-token", token);

      setLogin(true);

      Router.push(`/dashboard`);
      return;
    }
  };

  const onChange = (e) => {
    e.preventDefault();
    setAccountDetails({ ...accountDetails, [e.target.name]: e.target.value });
  };

  return (
    <div>
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

      <h3>Server Statistics</h3>
      <table className="table-statistics">
        <tbody>
          <tr>
            <td>Online Players</td>
            <td>20 / 100</td>
          </tr>
          <tr>
            <td>Registered Users</td>
            <td>123,000</td>
          </tr>
          <tr>
            <td>Banned Users</td>
            <td>6,999</td>
          </tr>
          <tr>
            <td>Most Kills</td>
            <td>IAlwaysKill</td>
          </tr>
          <tr>
            <td>Most Deaths</td>
            <td>IAlwaysDie</td>
          </tr>
          <tr>
            <td>Most Money</td>
            <td>IAmMadeOfMoney</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default SideBar;
