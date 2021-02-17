import Link from "next/link";
import Router from "next/router";
import React, { useState } from "react";
import { useToasts } from "react-toast-notifications";
import Layout from "../../components/layout";
import sendRequest from "../../utils/sendRequest";

function AddMedia() {
  const [inputData, setInputData] = useState({
    link: "",
    title: "",
  });

  const { addToast } = useToasts();

  const onSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await sendRequest("POST", "media", {
        body: JSON.stringify({
          youtubeLink: inputData.link,
          title: inputData.title,
        }),
      });

      const data = await response.json();

      console.log(data);

      // Redirect user back to '/media' page.
      // on success
      if (response.status === 200) {
        addToast(data.msg, { appearance: "success" });
        Router.push("/media");
      }
    } catch (e) {
      console.log(e);
    }
  };

  const onInputChange = (e) => {
    e.preventDefault();

    setInputData({
      ...inputData,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <Layout title="Adding Media...">
      <h2>How to submit</h2>
      <strong>
        Follow the guidelines below when submitting, failing to do so will get
        your post deleted
      </strong>
      <ul>
        <li>Your video should not contain pronography</li>
        <li>Advertisments of other SA-MP servers, discord, and such.</li>
        <li>Religion hate, discrimination, toxic reaction to players.</li>
      </ul>
      <p>
        To submit a video, simply insert a youtube link e.g
        <code>https://www.youtube.com/watch?v=Nl1PDCFTxFY</code> and provide a
        title with less tha 50 characters, click submit and you are good to go!
      </p>
      <hr />
      <h1>Adding new media</h1>
      <form>
        <div className="row">
          <div className="column">
            <label>Youtube Link</label>
            <input
              name="link"
              type="text"
              placeholder="Insert link"
              onChange={onInputChange}
            />
          </div>
          <div className="column">
            <label>Title</label>
            <input
              name="title"
              type="text"
              placeholder="Insert title"
              onChange={onInputChange}
            />
          </div>
        </div>
        <a className="button" type="button" onClick={onSubmit}>
          Submit
        </a>
        <Link href="/media">
          <a className="button button-outline" type="button">
            Go back
          </a>
        </Link>
      </form>
    </Layout>
  );
}

export default AddMedia;