import { boot } from "quasar/wrappers";
import { useUserStore, storeToRefs } from "stores/userStore.js";
export default boot(({ app }) => {
  function userInfo() {
    const userStore = useUserStore();
    return userStore;
  }
  app.provide("userInfo", userInfo);
  //   app.config.globalProperties.$user = user;
});
