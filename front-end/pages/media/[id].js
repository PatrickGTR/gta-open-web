import { useRouter } from "next/router";
import React from "react";
import Layout from "../../components/layout";
import sendRequest from "../../utils/sendRequest";

function SpecificMedia({ data }) {
  return (
    <Layout title="Viewing specific media">
      <div className="row">
        <div className="column column-67">
          <div className="iframe-container">
            <iframe
              src={
                `https://www.youtube.com/embed/` +
                data.youtubeLink.split("=")[1]
              }
              className="iframe-responsive"
              frameBorder="0"
              allowFullScreen
            />
          </div>
          <div style={{ marginTop: "5rem" }}>
            <h4>
              <strong>{data.title}</strong>
            </h4>
            <p>
              {data.views} {data.views > 1 ? `views` : `view`} |{" "}
              {data.datePosted}
            </p>
            <a>{data.author}</a>
          </div>
        </div>
        <div className="column column-33">
          <h3>Comments</h3>
          <hr />
          <div style={{ fontSize: "1.3rem" }}>
            <p>
              <strong>Patrick</strong> <i>1 day ago</i> <br />
              Cool Video Bro!
            </p>
            <p>
              <strong>VeryLongUser</strong> <i>1 month ago</i> <br />
              Song name?
            </p>
            <p>
              <strong>Syntax</strong> <i>1 month ago</i> <br />
              This is a test comment Very long ass comment lol
            </p>
            <p>
              <strong>Syntax</strong> <i>1 month ago</i> <br />
              You're so trash, get a better gameplay kiddo.
            </p>
            <p>
              <strong>Syntax</strong> <i>1 month ago</i> <br />
              SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM
              SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM SPAM
              SPAM SPAM SPAM SPAM SPAM SPAM
            </p>
            <p>
              <strong>Syntax</strong> <i>1 month ago</i> <br />
              This is a test comment
            </p>
            <p>
              <strong>Syntax</strong> <i>1 month ago</i> <br />
              This is a test comment
            </p>
            <p>
              <strong>Syntax</strong> <i>1 month ago</i> <br />
              This is a test comment
            </p>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export const getServerSideProps = async (ctx) => {
  const postid = ctx.query.id;

  const response = await sendRequest("GET", `media/${postid}`);
  const data = await response.json();

  return {
    props: {
      data,
    },
  };
};

export default SpecificMedia;
