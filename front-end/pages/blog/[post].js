import React from "react";
import fs from "fs";
import path from "path";
import matter from "gray-matter";
import Head from "next/head";
import marked from "marked";
import Layout from "../../components/layout";

const Post = ({ htmlString, data }) => {
  return (
    <>
      <Head>
        <meta title={data.title} content={data.description} />
      </Head>
      <Layout>
        <h1>{data.title}</h1>
        <div className="author-date">
          <div className="row">
            <div className="column">Posted By: {data.author}</div>
            <div className="column column-offset-40">Date: {data.date}</div>
          </div>
        </div>
        <div dangerouslySetInnerHTML={{ __html: htmlString }} />
      </Layout>
    </>
  );
};

export const getStaticPaths = async () => {
  const files = fs.readdirSync("posts");

  const paths = files.map((filename) => ({
    params: {
      post: filename.replace(".md", ""),
    },
  }));

  return {
    paths,
    fallback: false,
  };
};

export const getStaticProps = async ({ params: { post } }) => {
  const markdownWithMetadata = fs
    .readFileSync(path.join("posts", post + ".md"))
    .toString();

  const parsedMarkdown = matter(markdownWithMetadata);

  const htmlString = marked(parsedMarkdown.content);

  return {
    props: {
      htmlString,
      data: parsedMarkdown.data,
    },
  };
};

export default Post;
