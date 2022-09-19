<template>
  <!-- :style-fn="panelFnHeight" -->
  <q-page class="flex flex-center" style="min-width: 360px">
    <sprav></sprav>
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
    function clickHelp() {}
    return {
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
