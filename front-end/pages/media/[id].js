import React, { useEffect, useState } from "react";
import Layout from "../../components/layout";
import sendRequest from "../../utils/sendRequest";
import { formatSeconds } from "../../utils/formatSeconds";
import useStore from "../../store/user";
import Link from "next/link";

import { useMessage } from "../utils/message";

function SpecificMedia({ postid, post, commentData }) {
  const isLoggedIn = useStore((state) => state.loginStatus);
  const [didClickComment, setClickComment] = useState(false); // button toggle when user clicks the text area
  const [comment, setComment] = useState(""); // comment textbox
  const [comments, setComments] = useState(commentData); // comment data coming from the API
  const [isUpToDate, upToDate] = useState(true); // to update the section when this is called.

  const { notifySuccess, notifyError } = useMessage();

  // increase views everytime this resets
  useEffect(() => {
    sendRequest("POST", "media/add_views", {
      body: JSON.stringify({ mediaid: postid }),
    });
  }, []);

  useEffect(async () => {
    const response = await sendRequest("GET", `media/comment/${postid}`);
    const data = await response.json();

    setComments(data);
    upToDate(true);
  }, [isUpToDate]);

  return (
    <Layout title={post.title}>
      <div className="row">
        <div className="column column-67">
          <div className="iframe-container">
            <iframe
              src={
                `https://www.youtube.com/embed/` +
                post.youtubeLink.split("=")[1]
              }
              className="iframe-responsive"
              frameBorder="0"
              allowFullScreen
            />
          </div>
          <div style={{ marginTop: "5rem" }}>
            <h4>
              <strong>{post.title}</strong>
            </h4>
            <p>
              {post.views +
                ` ${post.views > 1 ? `views` : `view`}` +
                " | " +
                post.datePosted}
            </p>
            <a>{post.author}</a>
          </div>
        </div>
        <div className="column column-33">
          <h3>Comments</h3>
          <hr style={{ marginTop: "1rem", marginBottom: "1rem" }} />
          <div style={{ fontSize: "1.3rem" }}>
            <form>
              {isLoggedIn ? (
                <>
                  <textarea
                    value={comment || ""}
                    style={{ resize: "none" }}
                    placeholder="Add a public comment..."
                    onClick={(e) => {
                      e.preventDefault();
                      setClickComment(true);
                    }}
                    onChange={(e) => {
                      e.preventDefault();
                      setComment(e.target.value);
                    }}
                  />

                  {didClickComment && (
                    <div style={{ textAlign: "right" }}>
                      <a
                        className="button button-clear"
                        onClick={(e) => setClickComment(false)}
                      >
                        Cancel
                      </a>
                      <a
                        className="button"
                        onClick={(e) => {
                          e.preventDefault();

                          if (comment == "") {
                            notifyError("You can't send an empty comment");
                            return;
                          }

                          setComment("");
                          upToDate(false);
                          sendRequest("POST", "media/comment", {
                            body: JSON.stringify({
                              mediaid: postid,
                              comment: comment,
                            }),
                          });

                          notifySuccess("You have posted your comment");
                        }}
                      >
                        Comment
                      </a>
                    </div>
                  )}
                </>
              ) : (
                <p>
                  <Link href="/">Login</Link> to add a comment
                </p>
              )}
            </form>
            {comments.length ? (
              comments.map((comment, index) => (
                <p key={index}>
                  <strong>{comment.author}</strong>{" "}
                  <i>{formatSeconds(comment.datePosted) + " ago"}</i>
                  <br />
                  {comment.comment}
                </p>
              ))
            ) : (
              <p>This post has 0 comments</p>
            )}
          </div>
        </div>
      </div>
    </Layout>
  );
}

export const getServerSideProps = async (ctx) => {
  const postid = ctx.query.id;

  const postResponse = await sendRequest("GET", `media/${postid}`);
  const post = await postResponse.json();

  const commentResponse = await sendRequest("GET", `media/comment/${postid}`);
  const comments = await commentResponse.json();

  return {
    props: {
      postid,
      post,
      commentData: comments,
    },
  };
};

export default SpecificMedia;
