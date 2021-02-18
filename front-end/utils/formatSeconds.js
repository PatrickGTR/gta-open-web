export function formatSeconds(seconds) {
  seconds || 0;
  seconds = Number(seconds);
  seconds = Math.abs(seconds);
  seconds === 0 ? (seconds = 1) : seconds;

  const d = Math.floor(seconds / (3600 * 24));
  const y = Math.floor(seconds / (3600 * 24) / 365);
  const h = Math.floor((seconds % (3600 * 24)) / 3600);
  const m = Math.floor((seconds % 3600) / 60);
  const s = Math.floor(seconds % 60);

  let format;
  if (y > 0) {
    format = d > 0 ? d + " " + (d == 1 ? "year" : "years") : "";
  } else if (d > 0) {
    format = d > 0 ? d + " " + (d == 1 ? "day" : "days") : "";
  } else if (h > 0) {
    format = h > 0 ? h + " " + (h == 1 ? "hour" : "hours") : "";
  } else if (m > 0) {
    format = m > 0 ? m + " " + (m == 1 ? "minute" : "minutes") : "";
  } else if (s > 0) {
    format = s > 0 ? s + " " + (s == 1 ? "second" : "seconds") : "";
  }

  return format;
}
