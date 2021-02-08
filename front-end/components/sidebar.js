import LoginForm from "../components/loginform";
import React, { useEffect, useState } from "react";
import sendRequest from "../utils/sendRequest";

const SideBar = ({ response, data }) => {
  const [serverData, setServerData] = useState({
    HighestKill: "",
    HighestDeaths: "",
    HighestMoney: "",
    RegisteredPlayers: 0,
  });
  useEffect(async () => {
    let response, data;

    // fetch highest kills
    response = await sendRequest("GET", "server/stats?type=1&option=1");
    data = await response.json();
    const HighestKill = data.value;

    // fetch highest money
    response = await sendRequest("GET", "server/stats?type=1&option=2");
    data = await response.json();
    const HighestMoney = data.value;

    // fetch highest deaths
    response = await sendRequest("GET", "server/stats?type=1&option=3");
    data = await response.json();
    const HighestDeaths = data.value;

    //fetch total accounts
    response = await sendRequest("GET", "server/stats?type=2");
    data = await response.json();
    const RegisteredPlayers = data.value;

    setServerData({
      HighestKill,
      HighestDeaths,
      HighestMoney,
      RegisteredPlayers,
    });
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
            <td>{serverData.RegisteredPlayers}</td>
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
