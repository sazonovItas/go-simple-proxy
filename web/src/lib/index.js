export function validateEmail(email) {
  if (/^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/.test(email)) {
    return "";
  }

  return "invalid email form";
}

export function validateLogin(login) {
  if (/^[a-zA-Z0-9]{4,40}$/.test(login)) {
    return "";
  }

  return "invalid login form";
}

export function validatePassword(password) {
  if (/^[a-zA-Z0-9]{4,60}$/.test(password)) {
    return "";
  }

  return "invalid password form";
}

export function prettyBytes(num, precision = 3, addSpace = true) {
  // Define an array of units in increasing order of size
  const UNITS = ["B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
  // If the absolute value of the number is less than 1, return it with the smallest unit 'B'
  if (Math.abs(num) < 1) return num + (addSpace ? " " : "") + UNITS[0];
  // Determine the appropriate unit based on the magnitude of the number
  const exponent = Math.min(
    Math.floor(Math.log10(num < 0 ? -num : num) / 3),
    UNITS.length - 1,
  );
  // Calculate the value in the chosen unit with the specified precision
  const n = Number(
    ((num < 0 ? -num : num) / 1000 ** exponent).toPrecision(precision),
  );
  // Construct and return the human-readable string representation
  return (num < 0 ? "-" : "") + n + (addSpace ? " " : "") + UNITS[exponent];
}
