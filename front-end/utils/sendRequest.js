const sendRequest = async (method, endpoint, custom) => {
  const isProd = !process.env.DEV
    ? "http://vps-bd1b8740.vps.ovh.net:8000/"
    : "http://localhost:8000/";

  const url = isProd + endpoint;
  const response = await fetch(url, {
    method: method,
    credentials: "include",
    ...custom,
  });

  return response;
};

export default sendRequest;
