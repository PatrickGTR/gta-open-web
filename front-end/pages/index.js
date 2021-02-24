import React from "react";
import fs from "fs";
import path from "path";
import matter from "gray-matter";
import marked from "marked";

import SideBar from "../components/sidebar";
import Layout from "../components/layout";
import Link from "next/link";
import sendRequest from "../utils/sendRequest";

const Home = ({ stats, posts }) => {
  return (
    <Layout title="Home">
      <div className="row">
        <div className="column column-67 home-content">
          {posts.map((post, index) => {
            const mdData = matter(post.content);
            const { date, author } = mdData.data;
            const { content } = mdData;

            return (
              <div key={index}>
                <Link href="blog/[post]" as={`/blog/${post.path}`}>
                  <a>
                    <h1>{mdData.data.title}</h1>
                  </a>
                </Link>
                <div className="author-date">
                  <div className="row">
                    <div className="column">Posted By: {author}</div>
                    <div className="column">Date: {date}</div>
                  </div>
                </div>
                <div
                  dangerouslySetInnerHTML={{
                    // default settings copied from
                    // https://marked.js.org/
                    __html: marked(content, {
                      baseUrl: null,
                      breaks: false,
                      gfm: true,
                      headerIds: true,
                      headerPrefix: "",
                      highlight: null,
                      langPrefix: "language-",
                      mangle: true,
                      pedantic: false,
                      sanitize: false,
                      sanitizer: null,
                      silent: false,
                      smartLists: false,
                      smartypants: false,
                      tokenizer: null,
                      walkTokens: null,
                      xhtml: false,
                    }),
                  }}
                />
              </div>
            );
          })}
        </div>

        <div className="column">
          <SideBar stats={stats} />
        </div>
      </div>
    </Layout>
  );
};

export const getStaticProps = async () => {
  const files = fs.readdirSync("posts");

  let posts = [];

  files.map((file) => {
    const markdownWithMetadata = fs
      .readFileSync(path.join("posts", file))
      .toString();

    posts.push({
      path: file.replace(".md", ""),
      content: markdownWithMetadata,
    });
  });

  let serverStats = {};
  try {
    let response, data;
    // fetch highest kills
    response = await sendRequest("GET", "server/stats?type=1&option=1");
    data = await response.json();
    serverStats = { ...serverStats, highestKill: data.value };

    // fetch highest money
    response = await sendRequest("GET", "server/stats?type=1&option=2");
    data = await response.json();
    serverStats = { ...serverStats, highestMoney: data.value };

    // fetch highest deaths
    response = await sendRequest("GET", "server/stats?type=1&option=3");
    data = await response.json();
    serverStats = {
      ...serverStats,
      highestDeaths: data.value,
    };

    //fetch total accounts
    response = await sendRequest("GET", "server/stats?type=2");
    data = await response.json();
    serverStats = { ...serverStats, playerCount: data.value };
  } catch (e) {
    // write a proper logging
    console.log(
      "an error was caught trying to fetch server data, please check logs",
      e,
    );
  }

  return {
    props: {
      posts,
      stats: serverStats,
    },
  };
};

export default Home;
