<template>
  <component v-if="formFactor == 'page'" :is="currentComponent"></component>
  <q-page
    v-else-if="formFactor == 'card'"
    class="flex flex-center"
    style="min-width: 360px"
  >
    <Ark-Card>
      <component :is="currentComponent"></component>
    </Ark-Card>
  </q-page>
  <q-page v-else class="flex flex-center" style="min-width: 360px">
    <component :is="currentComponent"></component>
  </q-page>
</template>

<script>
import {
  defineComponent,
  defineAsyncComponent,
  ref,
  watch,
  computed,
} from "vue";
import { useRoute } from "vue-router";
//import Sprav from "components/Sprav/Sprav.vue";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
import ArkCard from "./ArkCard.vue";
export default defineComponent({
  name: "SpravPageTables",
  components: { ArkCard },
  props: ["paramTabName", "formFactor"],
  setup(props) {
    const pageSetup = usePagesSetupStore();
    const route = useRoute();
    const currentComponent = ref(null);
    pageSetup.currentPage = "SpravPageTables";
    const { cardMain } = storeToRefs(usePagesSetupStore());
    function clickHelp() {}
    console.log("propssssssssss", props.paramTabName);
    watch(
      () => props.paramTabName,
      (val) => {
        selectTable(val);
        console.log("параметры справочника из пути", props);
      },
      { immediate: true } // первый проход тоже срабатывать
    );
    function selectTable(val) {
      switch (val) {
        case "bakery":
          currentComponent.value = defineAsyncComponent(
            () => import("components/Bakery/BakeryTable.vue") //"src/components/Sprav/tabBakery/TablePanel.vue")
          );
          break;
        case "kagent":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Kagent/KagentTable.vue")
          );
          break;
        case "brand":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Sprav/tabBrand/TablePanel.vue")
          );
          break;
        //---------------------
        case "assortament":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productassortment/Table.vue")
          );
          break;
        case "raw":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productraw/Table.vue")
          );
          break;
        case "typeraw":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productrawvid/Table.vue")
          );
          break;
        case "products":
          currentComponent.value = defineAsyncComponent(
            //() => import("src/components/Products/products/Table.vue")
            () => import("src/components/Products/PageProductsExt.vue")
          );
          break;
        case "typeproduct":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/producttype/Table.vue")
          );
          break;
        case "tovar":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productvid/Table.vue")
          );
          break;
        case "city":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/City/Table.vue")
          );
          break;
        default:
        // currentComponent.value = defineAsyncComponent(() =>
        //   import("src/components/Sprav/tabBrand/TablePanel.vue")
        // );
      }
    }
    return {
      cardMain,
      clickHelp,
      currentComponent,
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
