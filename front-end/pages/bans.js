import React from "react";
import Layout from "../components/layout";
import sendRequest from "../utils/sendRequest";

function Bans({ datas }) {
  return (
    <Layout title="Banned Players">
      {!datas.length ? (
        <h3>An error occured while fetching data, try again later.</h3>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Username</th>
              <th>Admin</th>
              <th>Reason</th>
              <th>Date</th>
              <th>Unban date</th>
            </tr>
          </thead>
          <tbody>
            {datas.map((data, idx) => {
              return (
                <tr key={idx}>
                  <td>{data.username}</td>
                  <td>{data.by}</td>
                  <td>{data.reason}</td>
                  <td>{data.banDate}</td>
                  <td>{data.unbanDate}</td>
                </tr>
              );
            })}
          </tbody>
        </table>
      )}
    </Layout>
  );
}

export const getServerSideProps = async () => {
  let datas = [];
  try {
    const response = await sendRequest("GET", "server/banlist");
    datas = await response.json();
  } catch {
    console.log("could not connect to the server");
  }

  return {
    props: { datas },
  };
};

export default Bans;
