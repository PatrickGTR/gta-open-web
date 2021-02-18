import { useEffect, useState } from "react";
import useStore from "../store/user";
import Router from "next/router";
import Layout from "../components/layout";

import { parseCookie } from "../utils/cookie";
import sendRequest from "../utils/sendRequest";

const DashBoard = ({ data }) => {
  const isLogged = useStore((state) => state.loginStatus);

  // redirect user if there's  no localstorage
  // or isLogged is not set to true.
  useEffect(() => {
    if (!isLogged) {
      Router.push("/");
      return;
    }
  }, []);

  const DisplayAccountData = () => {
    return (
      <>
        <table>
          <tbody>
            <tr>
              <td>Account ID:</td>
              <td>{data.account.uid}</td>
            </tr>
            <tr>
              <td>Register Date</td>
              <td>{data.account.register_date}</td>
            </tr>
            <tr>
              <td>Last Login</td>
              <td>{data.account.last_login}</td>
            </tr>
            <tr>
              <td>Skin</td>
              <td>
                <img
                  width="121"
                  src={`https://open.mp/images/skins/${data.stats.skin}.png`}
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
          {Object.keys(data.stats).map((key, value) => (
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
          {Object.keys(data.items).map((key, value) => (
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
        <h2>{data.account.username}</h2>
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
      {Object.keys(data) == 0 ? (
        "There was an issue loading your data at the moment"
      ) : (
        <>
          <DisplayInfo />
        </>
      )}
    </Layout>
  );
};

export const getServerSideProps = async (ctx) => {
  const cookie = ctx.req.headers.cookie;

  let data = {};
  try {
    const response = await sendRequest("GET", `user`, {
      headers: ctx.req ? { cookie: cookie } : undefined,
    });

    data = await response.json();
  } catch {}

  return {
    props: {
      data,
    },
  };
};

export default DashBoard;
