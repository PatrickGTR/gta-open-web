export const parseCookie = (cookieContent, lookingFor) => {
  let cookie = {};

  cookieContent.split(";").forEach((element) => {
    let [k, v] = element.split("=");
    cookie[k.trim()] = v;
  });
  return cookie[lookingFor];
};
