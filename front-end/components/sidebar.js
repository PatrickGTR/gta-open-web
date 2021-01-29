import LoginForm from "../components/loginform";

const SideBar = () => {
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
