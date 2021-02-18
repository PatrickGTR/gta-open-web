import React from "react";
import Link from "next/link";
import Layout from "../../components/layout";
import sendRequest from "../../utils/sendRequest";
import useStore from "../../store/user";
import { formatSeconds } from "../../utils/formatSeconds";

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
                      {data.views +
                        ` ${data.views > 1 ? "views" : "views"}` +
                        " | " +
                        formatSeconds(data.datePosted) +
                        ` ago`}
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
