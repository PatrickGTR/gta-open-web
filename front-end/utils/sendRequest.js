const sendRequest = async (method, endpoint) => {
  const url = `http://localhost:8000/${endpoint}`;
  const response = await fetch(url, {
    method: method,
    credentials: "include",
  });

  return response;
};

export default sendRequest;
