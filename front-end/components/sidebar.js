export default () => {
    return (
    <div>
        <form>
            <label>Username</label>

            <input type="text" placeholder="Username"/>

            <label>Password</label>

            <input type="text" placeholder="Username"/>
            <input class="button-primary" type="submit" value="Login"/>
        </form>

        <h3>Server Statistics</h3>
        <table class="table-statistics">
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

    )
}