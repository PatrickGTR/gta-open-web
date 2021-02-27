const sendRequest = async (method, endpoint, custom) => {
  const isProd =
    process.env.NODE_ENV === "development"
      ? "http://localhost:8000/"
      : "https://api.gta-open.ga/";

  const url = isProd + endpoint;
  const response = await fetch(url, {
    method: method,
    mode: "cors",
    credentials: "include",
    ...custom,
  });

  return response;
};

export default sendRequest;
