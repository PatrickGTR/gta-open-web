import React from "react";
import Link from "next/link";
import Layout from "../../components/layout";
import sendRequest from "../../utils/sendRequest";
import useStore from "../../store/user";

function SecToMHS(seconds) {
  seconds = seconds || 0;
  seconds = Number(seconds);
  seconds = Math.abs(seconds);

  const d = Math.floor(seconds / (3600 * 24));
  const y = Math.floor(seconds / (3600 * 24) / 365);
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
  const isLoggedIn = useStore((state) => state.loginStatus);

  return (
    <Layout title="Media">
      <div className="row">
        <h2 className="column">Most recent media</h2>
        {isLoggedIn && (
          <Link href="/media/add">
            <a className="button">Add Media</a>
          </Link>
        )}
      </div>
      <hr style={{ marginBottom: "1rem", marginTop: "1rem" }} />
      <div className="row">
        <div className="column" style={{ textAlign: "center" }}>
          {data.map((data, index) => (
            <Link
              key={index}
              href="/media/[id]"
              as={`/media/${encodeURIComponent(data.id)}`}
            >
              <figure className="media">
                <img
                  className="media-thumbnail"
                  src={
                    `http://img.youtube.com/vi/` +
                    data.youtubeLink.split("=")[1] +
                    `/mqdefault.jpg`
                  }
                />
                <figcaption className="media-caption">
                  <div className="row">
                    <div className="column">
                      <strong>{data.title}</strong>
                      <br />
                      {data.author}
                      <br />
                      60 views | {SecToMHS(data.datePosted) + ` ` + `ago`}
                    </div>
                  </div>
                </figcaption>
              </figure>
            </Link>
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
