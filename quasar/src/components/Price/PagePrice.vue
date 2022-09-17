<template>
  <q-page class="flex flex-center" style="min-width: 360px">
    <Price :pageMaxHeight="pageMaxHeight"></Price>
  </q-page>
</template>

<script>
import { defineComponent, ref, toRefs, computed } from "vue";
import Price from "./Price.vue";
import { useQuasar } from "quasar";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
export default defineComponent({
  name: "PagePrice",
  components: {
    Price,
  },
  setup() {
    const $q = useQuasar();
    const pageSetup = usePagesSetupStore();
    pageSetup.currentPage = "price";
    const { cardMain } = storeToRefs(usePagesSetupStore());
    const pageMaxHeight = computed(() => {
      return {
        minHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
        maxHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
      };
    });

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
