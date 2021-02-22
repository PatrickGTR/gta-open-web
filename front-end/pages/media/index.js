import React from "react";
import Link from "next/link";
import Layout from "../../components/layout";
import sendRequest from "../../utils/sendRequest";
import useStore from "../../store/user";
import { formatSeconds } from "../../utils/formatSeconds";
import Head from "next/head";

function Media({ data }) {
  const isLoggedIn = useStore((state) => state.loginStatus);

  return (
    <Layout title="Media">
      <Head>
        <meta name="title" content="GTA Open" />
        <meta
          name="keywords"
          content="Grand Theft Auto, San Andreas, SA-MP, GTA, GTA Open, open."
        />
        <meta
          name="description"
          content="GTA Open Media, contents posted by GTA Open players"
        />
        <meta
          name="image"
          content="https://camo.githubusercontent.com/11857964d64562f7c921ba7ce05fd363ae4f0ed0654ecb24ac95ffa51aa4d241/68747470733a2f2f696d616765732d6578742d312e646973636f72646170702e6e65742f65787465726e616c2f39626e714d4a523842454c45674942503870795a7a58527432574a304e6d495770734e6a77637674644d732f68747470732f692e6962622e636f2f53524c7a7a636e2f6774616f70656e2d72656464616464792e706e67"
        />
        <meta name="url" content="/media/" />
      </Head>
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
