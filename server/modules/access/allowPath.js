/**
 * Массив разрешенных путей, без регистрации
 * @param {*} path
 * @returns
 */
export function isAllowPath(path) {
  console.log(" pathhhhhh", path);
  return allowPath.includes(path);
}
export const allowPath = [
  "/",
  "/new",
  "/new/registration2",
  "/unLogin",
  "/isLogin",
  "/accessCheck",
  "/login",
  "/registration",
];
