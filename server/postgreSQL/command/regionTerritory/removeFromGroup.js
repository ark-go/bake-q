import { pool } from "../../initPostgreSQL.js";
import { botSendMessage } from "../../../tg/startTgBot.js";
export async function removeFromGroup(req, res, tabname, timezone) {
  let sqlP = {
    text: /*sql*/ `
 
    SELECT * from region_x_territory_delete($1,$2 at time zone $3,$4); -- перенесли
    
          `,
    values: [
      req.body?.childId,
      req.body?.dateStart,
      timezone,
      null,
      //req.session?.user?.id,
    ],
  };
  try {
    //   console.log("region_territory_delete", sqlP);
    let result = await pool.query(sqlP);
    result = result.rowCount > 0 ? result.rows : null;
    let mess = `Deell ${tabname}:` + "\n" + req?.session?.user?.email;
    //console.log(mess, result);
    mess += "\n" + JSON.stringify(result);
    botSendMessage(mess, req);
    return {
      result: result,
    };
  } catch (err) {
    console.log("Ошибка чтения ", tabname, err.toString());
    return {
      error: err.toString(),
    };
  }
}
