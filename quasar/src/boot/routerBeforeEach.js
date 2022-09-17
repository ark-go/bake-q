import { boot } from "quasar/wrappers";
import { axios } from "boot/axios";
// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(async ({ router } /* { app, router, ... } */) => {
  let yesReplace = true;
  router.beforeEach(async (to, from) => {
    // console.log("router boot: router BeforeEach", to.meta.checkAccess);
    if (to.meta?.checkAccess) {
      let check = null;
      try {
        check = await checkAccess(to.path, to.meta?.title);
        if (!check) {
          return "/";
          // return false; // никуда не переходим, запрещено
        }
      } catch (error) {
        console.error("Ошибка проверки доступа", error);
        // ошибка при запросе, остаемся на месте
        return false; // никуда не переходим
      }
    }
    // BUG: не работает, сам replace в return

    // в роутере прописана meta.replace. при первом проходе
    // надо войти в блок и сказать, раз мы уходим со страницы которую надо переписать
    // делаем replace,
    // yesReplace - для предотвращения цикла, при повторном проходе, по return
    if (from.meta?.replace && yesReplace) {
      // если стояли на том который не надо запоминать
      yesReplace = false; // флаг, мы делаем замену
      // console.log("Замена истории");
      return { path: to.path, replace: true };
    }
    yesReplace = true; // флаг ожидания, на разрешение замены, устанавливаем при любом проходе
  });

  async function checkAccess(path, title) {
    let resp = await axios.post("/api/accessCheck", { path: path });
    let respData = resp.data;
    let keyRes = Object.keys(respData);
    if (!keyRes.includes("result") && !keyRes.includes("error")) {
      throw "При проверке доступа не пришел ожидаемый ответ";
    }
    if (respData.error) {
      //
      console.error("Доступ запрещен:.", path);
      return false; //router.push("/");
    }
    if (respData.result) {
      // запрос прошел, и без ошибки
      console.warn("Доступ разрешен :", path);
      return true;
    }
  }
  // Меняем заголовок сайта
  router.afterEach((to, from) => {
    let title = to.meta?.title;
    document.title = "ХиТ";
    document.title += title ? " | " + title : "";
  });
});
