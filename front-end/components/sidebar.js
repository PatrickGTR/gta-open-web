import LoginForm from "../components/loginform";
import React, { useEffect, useState } from "react";
import sendRequest from "../utils/sendRequest";

const UNABLE_FETCH_ERR = "error fetching data";

const SideBar = () => {
  const [serverData, setServerData] = useState({
    HighestKill: "",
    HighestDeaths: "",
    HighestMoney: "",
    RegisteredPlayers: 0,
  });
  useEffect(async () => {
    let HighestKill = UNABLE_FETCH_ERR,
      HighestMoney = UNABLE_FETCH_ERR,
      HighestDeaths = UNABLE_FETCH_ERR,
      RegisteredPlayers = -1;
    try {
      let response, data;
      // fetch highest kills
      response = await sendRequest("GET", "server/stats?type=1&option=1");
      data = await response.json();
      HighestKill = data.value;

      // fetch highest money
      response = await sendRequest("GET", "server/stats?type=1&option=2");
      data = await response.json();
      HighestMoney = data.value;

      // fetch highest deaths
      response = await sendRequest("GET", "server/stats?type=1&option=3");
      data = await response.json();
      HighestDeaths = data.value;

      //fetch total accounts
      response = await sendRequest("GET", "server/stats?type=2");
      data = await response.json();
      RegisteredPlayers = data.value;
    } catch {}

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
            <td>{serverData.RegisteredPlayers === -1 && UNABLE_FETCH_ERR}</td>
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
