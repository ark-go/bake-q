import { defineStore } from "pinia";
export { storeToRefs } from "pinia";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
//import merge from "merge";
import { dom, Notify, useQuasar, Platform, Screen } from "quasar";
//let $q = useQuasar();
//import { nextTick, unref, ref, toRefs } from "vue";
/**
 * Настройки страницы для пользователя
 */
const { style, height } = dom;
export const useArkCardStore = defineStore("ArkCardStore", {
  state: () => {
    return {
      /**
       * Размеры arkCard
       */
      pagePaddingY: 10,
      topSectionSize: {},
      infoSectionSize: {},
      bodySectionSize: {},
      bottomSectionSize: {},
    };
  },
  actions: {
    resetSectionSize() {},
    // onResizeDiv(val) {
    //   console.log("resize action", val);
    // },
  },

  getters: {
    /**Рабочий размер Height для ArkCard с padding  ( pagePaddingY )*/
    // arkCardHeight: (state) => {
    //   const pageSetup = usePagesSetupStore();
    //   // вычичслим зону
    //   return (
    //     pageSetup.pageHeight - pageSetup.pageOffset - pageSetup.pagePaddingY
    //   );
    // },

    maxBodyHeight(state) {
      const pageSetup = usePagesSetupStore();
      // bodySectionSize - не использовать в расчете !!
      let topBottom = 0;
      topBottom += state.topSectionSize.height
        ? state.topSectionSize.height
        : 0;
      topBottom += state.infoSectionSize.height
        ? state.infoSectionSize.height
        : 0;
      topBottom += state.bottomSectionSize.height
        ? state.bottomSectionSize.height
        : 0;
      // console.log(
      //   "++++++++++++ size",
      //   topBottom,
      //   state.topSectionSize?.height,
      //   state.infoSectionSize?.height,
      //   state.bottomSectionSize?.height,
      //   pageSetup.pageBodyHeight
      // );
      return `${pageSetup.pageBodyHeight - topBottom - state.pagePaddingY}px `;
    },
    maxBodyHeightCss(state) {
      return {
        maxHeight: this.maxBodyHeight,
        minHeight: this.maxBodyHeight,
        overflow: "auto",
      };
    },
  },
});
// function copyObject(o) {
//   // глубокое копиование ?
//   if (o === null) return null;
//   var output, v, key;
//   output = Array.isArray(o) ? [] : {};
//   for (key in o) {
//     v = o[key];
//     output[key] = typeof v === "object" ? copyObject(v) : v;
//   }
//   return output;
// }
// const pageSetup = usePagesSetupStore();

// // прочитаем при загрузке данные из компа
// const lockStat = localStorage.getItem("pageSetup");
// // внимание надо объединить данные из кода и данные прочитанные из LocalStorage компа
// // иначе измененые данные никогда не задействуются, потому что будут переписаны из компьютера
// if (lockStat) {
//   // для глубокого слияния используем модуль merge
//   let mergeLockComp = merge.recursive(
//     true,
//     pageSetup.$state,
//     JSON.parse(localStorage.getItem("pageSetup"))
//   );
//   // заносим данные из кода и и данные из LocalStorage компьютера в наш магазин
//   pageSetup.$state = mergeLockComp;
// }
// // подписываемся на события и сохраняем при изменение в локальный комп
// pageSetup.$subscribe((mutation, state) => {
//   if (state.page[state.currentPage]?.resetStoreUser) {
//     let buf = state.currentPage; // запомним текущюю страницу
//     Notify.create({
//       color: "silver",
//       textColor: "white",
//       icon: "reset_tv",
//       message: "Сбросить настройки экрана?",
//       caption: "Будут установлены настройки по умолчанию.",
//       position: "center",
//       // avatar,
//       multiLine: true,
//       timeout: 0,
//       actions: [
//         {
//           label: "Сброс",
//           color: "orange-9",
//           handler: () => {
//             localStorage.removeItem("pageSetup");
//             pageSetup.$reset();
//             state.currentPage = buf; // восстановитьпосле сброса текущую страницу
//             // в этом же тике запишем нашу текущюю страницу обратно в LocalStorage
//             // иначе после сброса, не получается :)
//             localStorage.setItem("pageSetup", JSON.stringify(state));
//           },
//         },
//         {
//           label: "Не хочу",
//           color: "green",
//           handler: () => {
//             // раз отказались то вернем ключ сброса на место;
//             state.page[state.currentPage].resetStoreUser = false; //
//             /* console.log('wooow') */
//           },
//         },
//       ],
//     });
//     // pageSetup.$reset();
//     // state.currentPage = buf;
//   }
//   localStorage.setItem("pageSetup", JSON.stringify(state));
//});
