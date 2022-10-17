import Redis from "ioredis";
import dotenv from "dotenv";
dotenv.config();
console.log("IP for Redis", process.env.Redis_host);
import connectRedis from "connect-redis";
import session from "express-session";
import { createRequire } from "module";
const require = createRequire(import.meta.url);
const { name } = require("../package.json");

export const redisStore = connectRedis(session);
export const redisClient = new Redis({
  //showFriendlyErrorStack: true, // только для отладки.. для вывода нормальных ошибок
  keyPrefix: process.env.Redis_prefix
    ? process.env.Redis_prefix + ":"
    : name + ":",
  port: process.env.Redis_port, // Redis port
  host: process.env.Redis_host, // Redis host
  family: 4, // 4 (IPv4) or 6 (IPv6)
  password: process.env.Redis_password,
  db: process.env.Redis_db,
  reconnectOnError(err) {
    // https://github.com/luin/ioredis#reconnect-on-error
    console.log("Redis reconnectOnError", err.message);
    const targetError = "ETIMEDOUT"; // "READONLY";
    if (err.message.includes(targetError)) {
      // Only reconnect when the error contains "READONLY"
      return true; // or `return 1;`
    }
  },
});

redisClient.on("error", function (err) {
  if (err?.code != "EHOSTUNREACH") {
    console.log(
      "redis error::",
      process.env.Redis_host,
      err?.address + "!",
      err.toString()
    );
  }
});
redisClient.on("connect", function () {
  console.info("Redis", process.env.Redis_host, "Подключился");
});

redisClient.on("ready", function () {
  console.info("Redis", process.env.Redis_host, "Готов");
});

//! redisClient.connect().catch(console.error)
//export { RedisStore };
//export const redisStore = new RedisStore({ client: redisClient });
