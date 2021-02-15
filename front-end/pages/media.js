import React from "react";
import Layout from "../components/layout";

function Media() {
  const dummy = [
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
    12,
    13,
    14,
    15,
    16,
    17,
    18,
    19,
    20,
    21,
    22,
    23,
    24,
  ];

  return (
    <Layout title="Media">
      <div className="row">
        <div className="column" style={{ textAlign: "center" }}>
          {dummy.map((index) => (
            <>
              <figure className="media">
                <img
                  className="media-thumbnail"
                  height="128"
                  src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.ytimg.com%2Fvi%2FQ7nmsLEeMIs%2Fmaxresdefault.jpg&f=1&nofb=1"
                />
                <figcaption className="media-caption">
                  <strong>
                    This is a very long title for this post and I should
                    probably stop but I'm not going to.
                  </strong>
                  <br />
                  PatrickGTR
                  <br />
                  60 views | 1 hour ago
                </figcaption>
              </figure>
            </>
          ))}
        </div>
      </div>
    </Layout>
  );
}

export default Media;
