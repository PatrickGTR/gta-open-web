import LoginForm from "../components/loginform";
import React, { useEffect, useState } from "react";

const UNABLE_FETCH_ERR = "error fetching data";

const SideBar = ({ stats }) => {
  const { highestKill, highestDeaths, highestMoney, playerCount } = stats;

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
            <td>{playerCount || UNABLE_FETCH_ERR}</td>
          </tr>
          <tr>
            <td>Banned Users</td>
            <td>6,999</td>
          </tr>
          <tr>
            <td>Most Kills</td>
            <td>{highestKill || UNABLE_FETCH_ERR}</td>
          </tr>
          <tr>
            <td>Most Deaths</td>
            <td>{highestDeaths || UNABLE_FETCH_ERR}</td>
          </tr>
          <tr>
            <td>Most Money</td>
            <td>{highestMoney || UNABLE_FETCH_ERR}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default SideBar;
