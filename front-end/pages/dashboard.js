import { useEffect, useState } from "react";
import Layout from "../components/layout";

import useStore from "../store/user";

const DashBoard = () => {
  const [userData, setUserData] = useState({});
  const [isLoading, setLoading] = useState(true);

  const isLogged = useStore((state) => state.loginStatus);

  useEffect(async () => {
    const response = await fetch("http://localhost:8000/user/1", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + localStorage.getItem("jwt-token"),
      },
    });

    if (response.status === 200) {
      const data = await response.json();

      setUserData(data);
      setLoading(false);

      console.log(isLogged);
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

  const DisplayInfo = () =>
    Object.keys(userData).length != 0 ? (
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
    ) : (
      <p>Could not redeem data, empty objet passed</p>
    );

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
