import React from "react";
import fs from "fs";
import path from "path";
import matter from "gray-matter";
import marked from "marked";

import SideBar from "../components/sidebar";
import Layout from "../components/layout";
import Link from "next/link";

const Home = ({ posts }) => {
  return (
    <Layout title="Home">
      <div className="row">
        <div className="home-content">
          {posts.map((post) => {
            const mdData = matter(post.content);
            const { date, author } = mdData.data;
            const { content } = mdData;

            return (
              <div key={mdData.data.title}>
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

        <div className="column column-offset-5">
          <SideBar />
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

  return {
    props: {
      posts,
    },
  };
};

export default Home;
