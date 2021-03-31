import LoginForm from "../components/loginform";
import { noAvailableServer } from "../utils/message";

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
            <td>{playerCount || noAvailableServer}</td>
          </tr>
          <tr>
            <td>Banned Users</td>
            <td>6,999</td>
          </tr>
          <tr>
            <td>Most Kills</td>
            <td>{highestKill || noAvailableServer}</td>
          </tr>
          <tr>
            <td>Most Deaths</td>
            <td>{highestDeaths || noAvailableServer}</td>
          </tr>
          <tr>
            <td>Most Money</td>
            <td>{highestMoney || noAvailableServer}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default SideBar;
