import fs from "fs";
import moment from "moment-timezone";

import path from "path";
import { fileURLToPath } from "url";
const __dirname = path.dirname(fileURLToPath(import.meta.url));
//console.log("+++++++++++++++++++++++++++++", __dirname);
export function startCheckVersionSite(socket) {
  let timerInt = setInterval(() => {
    sendVersionSite(socket, timerInt);
  }, 60000);
}
function sendVersionSite(socket) {
  // BUG Жесткая привязка!
  let pathPackage = path.join(
    __dirname,
    "..",
    "..",
    "..",
    "quasar",
    "package.json"
  );
  //  console.log("+++++++++++++++++++++++++++++", pathPackage);
  fs.readFile(pathPackage, "utf8", (err, data) => {
    // process.env.PackagePath
    if (err) {
      console.log(`Не прочитать Package файл от Quasar: ${err}`);
    } else {
      const pack = JSON.parse(data);
      // console.log("Версия сайта на сервере:", pack.version);
      let mom = moment.tz(Date.now(), socket.request?.session?.timezone);
      socket.emit("updateSite", {
        date:
          socket.request?.session?.timezone +
          ": " +
          mom.format("DD-MM-YYYY HH:mm"),
        versionSiteServer: pack.version,
      });
    }
  });
}
