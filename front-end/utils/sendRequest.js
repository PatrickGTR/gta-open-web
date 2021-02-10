const sendRequest = async (method, endpoint, custom) => {
  const url = `http://localhost:8000/${endpoint}`;
  const response = await fetch(url, {
    method: method,
    credentials: "include",
    ...custom,
  });

  return response;
};

export default sendRequest;
