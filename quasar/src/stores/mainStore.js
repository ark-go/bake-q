import { defineStore } from "pinia";
export { storeToRefs } from "pinia";
//import { useQuasar } from "quasar";

export const useMainStore = defineStore("MainStore", {
  state: () => {
    return {
      //все эти свойства будут иметь свой тип автоматически
      rightDrawerOpen: false,
      leftDrawerOpen: false,
      isLeftDrawer: true,
      modalLoginOpen: false,
    };
  },
  // getters: {
  //   isMobile: (state) => {
  // //    let { platform } = useQuasar();
  // //    return platform.is.mobile;
  //   },
  // },
});
