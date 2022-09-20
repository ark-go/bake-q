import { defineStore } from "pinia";
export { storeToRefs } from "pinia";
//import { useQuasar } from "quasar";

export const useProductsStore = defineStore("ProductsStore", {
  state: () => {
    return {
      //все эти свойства будут иметь свой тип автоматически
      /**
       * Выбранная строка продукции
       */
      selectedRow: {},
    };
  },
  // getters: {
  //   isMobile: (state) => {
  // //    let { platform } = useQuasar();
  // //    return platform.is.mobile;
  //   },
  // },
});
