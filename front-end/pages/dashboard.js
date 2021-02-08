import { useEffect, useState } from "react";
import useStore from "../store/user";
import Router from "next/router";
import Layout from "../components/layout";

import getCookie from "../utils/getcookie";
import sendRequest from "../utils/sendRequest";

const DashBoard = () => {
  const [userData, setUserData] = useState({});
  const [isLoading, setLoading] = useState(true);

  const isLogged = useStore((state) => state.loginStatus);

  useEffect(async () => {
    if (!isLogged) {
      Router.push("/");
      return;
    }

    const response = await sendRequest(
      "GET",
      `user/${getCookie("db_user_id")}`,
    );
    if (response.status === 200) {
      const data = await response.json();

      setUserData(data);
      setLoading(false);
    }
  }, []);

  const DisplayAccountData = () => {
    return (
      <>
        <table>
          <tbody>
            <tr>
              <td>Account ID:</td>
              <td>{userData.account.uid}</td>
            </tr>
            <tr>
              <td>Register Date</td>
              <td>{userData.account.register_date}</td>
            </tr>
            <tr>
              <td>Last Login</td>
              <td>{userData.account.last_login}</td>
            </tr>
            <tr>
              <td>Skin</td>
              <td>
                <img
                  width="121"
                  src={`https://open.mp/images/skins/${userData.stats.skin}.png`}
                />
              </td>
            </tr>
          </tbody>
        </table>
      </>
    );
  };

  const DisplayStats = () => (
    <>
      <table>
        <thead>
          <tr>
            <th>Statistic</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
          {Object.keys(userData.stats).map((key, value) => (
            <tr key={key}>
              <td>{key.charAt(0).toUpperCase() + key.slice(1)}</td>
              <td>{value}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );

  const DisplayItems = () => (
    <>
      <table>
        <thead>
          <tr>
            <th>Item</th>
            <th>Quantity</th>
          </tr>
        </thead>
        <tbody>
          {Object.keys(userData.items).map((key, value) => (
            <tr key={key}>
              <td>{key.charAt(0).toUpperCase() + key.slice(1)}</td>
              <td>{value}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );

  const DisplayInfo = () => {
    return (
      <>
        <h2>{userData.account.username}</h2>
        <div className="row">
          <div className="column">
            <DisplayAccountData />
          </div>
          <div className="column">
            <DisplayStats />
          </div>
          <div className="column">
            <DisplayItems />
          </div>
        </div>
      </>
    );
  };

  return (
    <Layout title="Dashboard">
      {isLoading ? (
        "Please wait... Loading user data"
      ) : (
        <>
          <DisplayInfo />
        </>
      )}
    </Layout>
  );
};

export default DashBoard;
