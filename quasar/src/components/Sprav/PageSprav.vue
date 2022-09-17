<template>
  <!-- :style-fn="panelFnHeight" -->
  <q-page class="flex flex-center" style="min-width: 360px">
    <sprav :pageMaxHeight="pageMaxHeight"></sprav>
  </q-page>
</template>

<script>
import { defineComponent, ref, toRefs, computed } from "vue";
import { useRoute } from "vue-router";
import Sprav from "components/Sprav/Sprav.vue";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
export default defineComponent({
  name: "PageSprav",
  components: {
    Sprav,
  },
  setup() {
    const pageSetup = usePagesSetupStore();
    const route = useRoute();
    pageSetup.currentPage = "sprav";
    const { cardMain } = storeToRefs(usePagesSetupStore());
    // const {fontSize} = toRefs(state)
    // const pageMaxHeight = ref();
    const pageMaxHeight = computed(() => {
      let HH = {
        minHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
        maxHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
      };
      return HH;
    });
    // function panelFnHeight(offset, height2) {
    //   let height = `calc(100vh - ${offset}px)`;

    //   pageMaxHeight.value = {
    //     minHeight: height,
    //     maxHeight: height,
    //   };
    //   console.log("nmmmmmmmmmmmmmmmmmmmmmmmm", pageMaxHeight.value);
    //   return { minHeight: height, maxHeight: height };
    // }
    function clickHelp() {}
    return {
      //  panelFnHeight,
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
