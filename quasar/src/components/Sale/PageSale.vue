<template>
  <q-page class="flex flex-center" style="min-width: 360px">
    <Sale :pageMaxHeight="pageMaxHeight"></Sale>
  </q-page>
</template>

<script>
import { defineComponent, ref, toRefs, computed } from "vue";
import Sale from "./Sale.vue";
import { useQuasar } from "quasar";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
export default defineComponent({
  name: "PageSale",
  components: {
    Sale,
  },
  setup() {
    const $q = useQuasar();
    const pageSetup = usePagesSetupStore();
    pageSetup.currentPage = "sale";
    const { cardMain } = storeToRefs(usePagesSetupStore());
    // const {fontSize} = toRefs(state)
    const pageMaxHeight = computed(() => {
      return {
        minHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
        maxHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
      };
    });

    // function panelFnHeight(offset, height2) {
    //   if (!$q.fullscreen.isActive) {
    //     //  console.log("Обычный Экран!");
    //     let height = `calc(100vh - ${offset}px)`;
    //     let heightChild = `calc(100vh - ${offset}px)`;
    //     pageMaxHeight.value = {
    //       minHeight: heightChild,
    //       maxHeight: heightChild,
    //     };
    //     return { minHeight: height, maxHeight: height };
    //   } else {
    //     console.log("Полный Экран!");
    //     pageMaxHeight.value = {
    //       minHeight: "100vh",
    //       maxHeight: "100vh",
    //     };
    //     return {};
    //   }
    // }
    function clickHelp() {}
    return {
      pageMaxHeight,
      cardMain,
      clickHelp,
    };
  },
});
</script>
<style lang="scss" scoped>
:deep(.q-table tbody td) {
  font-size: v-bind("cardMain.fontSize.curr + 'px'");
}
:deep(.q-tree) {
  font-size: v-bind("cardMain.fontSize.curr + 'px'");
}
</style>
