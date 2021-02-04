import LoginForm from "../components/loginform";
import fetch from "isomorphic-fetch";
import React, { useEffect, useState } from "react";

const SideBar = ({ response, data }) => {
  const [serverData, setServerData] = useState({
    HighestKill: "",
    HighestDeaths: "",
    HighestMoney: "",
  });
  useEffect(async () => {
    let response, data;

    // fetch highest kills
    response = await fetch("http://localhost:8000/server/stats?type=1", {
      method: "GET",
      credential: "include",
    });

    data = await response.json();
    const HighestKill = data.username;

    // fetch highest money
    response = await fetch("http://localhost:8000/server/stats?type=2", {
      method: "GET",
      credential: "include",
    });

    data = await response.json();
    const HighestMoney = data.username;

    // fetch highest deaths
    response = await fetch("http://localhost:8000/server/stats?type=3", {
      method: "GET",
      credential: "include",
    });

    data = await response.json();
    const HighestDeaths = data.username;

    setServerData({ HighestKill, HighestDeaths, HighestMoney });
  }, []);

  return (
    <div>
      <LoginForm />

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
            <td>{serverData.HighestKill}</td>
          </tr>
          <tr>
            <td>Most Deaths</td>
            <td>{serverData.HighestDeaths}</td>
          </tr>
          <tr>
            <td>Most Money</td>
            <td>{serverData.HighestMoney}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default SideBar;
