import React from "react";
import Layout from "../components/layout";
import sendRequest from "../utils/sendRequest";

function SecToMHS(seconds) {
  seconds = seconds || 0;
  seconds = Number(seconds);
  seconds = Math.abs(seconds);

  const y = Math.floor(seconds / (3600 * 24) / 365);
  const d = Math.floor(seconds / (3600 * 24));
  const h = Math.floor((seconds % (3600 * 24)) / 3600);
  const m = Math.floor((seconds % 3600) / 60);
  const s = Math.floor(seconds % 60);

  let format;
  if (y > 0) {
    format = d > 0 ? d + " " + (d == 1 ? "year" : "years") : "";
  } else if (d > 0) {
    format = d > 0 ? d + " " + (d == 1 ? "day" : "days") : "";
  } else if (h > 0) {
    format = h > 0 ? h + " " + (h == 1 ? "hour" : "hours") : "";
  } else if (m > 0) {
    format = m > 0 ? m + " " + (m == 1 ? "minute" : "minutes") : "";
  } else if (s > 0) {
    format = s > 0 ? s + " " + (s == 1 ? "second" : "seconds") : "";
  }

  return format;
}

function Media({ data }) {
  return (
    <Layout title="Media">
      <div className="row">
        <div className="column" style={{ textAlign: "center" }}>
          {data.map((data, index) => (
            <figure key={index} className="media">
              <img
                className="media-thumbnail"
                height="128"
                src={
                  `http://img.youtube.com/vi/` +
                  data.youtubeLink.split("=")[1] +
                  `/mqdefault.jpg`
                }
              />
              <figcaption className="media-caption">
                <strong>
                  This is a very long title for this post and I should probably
                  stop but I'm not going to.
                </strong>
                <br />
                {data.author}
                <br />
                60 views | {SecToMHS(data.datePosted) + ` ` + `ago`}
              </figcaption>
            </figure>
          ))}
        </div>
      </div>
    </Layout>
  );
}

export const getServerSideProps = async (context) => {
  const response = await sendRequest("GET", "media");
  const data = await response.json();

  return {
    props: {
      data,
    },
  };
};

export default Media;
