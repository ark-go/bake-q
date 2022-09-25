<template>
  <component
    v-if="formFactor == 'page'"
    :is="currentComponent"
    :key="route.path"
  ></component>
  <q-page
    v-else-if="formFactor == 'card'"
    class="flex flex-center"
    style="min-width: 360px"
  >
    <Ark-Card @click.right.prevent>
      <transition
        appear
        enter-active-class="animated bounceOutRight"
        leave-active-class="animated fadeOut"
      >
        <component :is="currentComponent" :key="route.path"></component>
      </transition>
    </Ark-Card>
  </q-page>
  <q-page v-else class="flex flex-center" style="min-width: 360px">
    <component :is="currentComponent" :key="route.path"></component>
  </q-page>
</template>

<script>
import { defineComponent, defineAsyncComponent, ref, watch } from "vue";
import { useRoute } from "vue-router";
//import Sprav from "components/Sprav/Sprav.vue";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
import ArkCard from "./ArkCard.vue";
export default defineComponent({
  name: "SpravPageTables",
  components: { ArkCard },
  props: ["tblRouteParam", "formFactor"],
  setup(props) {
    const pageSetup = usePagesSetupStore();
    const route = useRoute();
    const currentComponent = ref(null);
    pageSetup.currentPage = "SpravPageTables";
    const { cardMain } = storeToRefs(usePagesSetupStore());
    function clickHelp() {}
    watch(
      () => props.tblRouteParam,
      (val) => {
        selectTable(val);
        console.log("параметры справочника из пути", route.path, props);
      },
      { immediate: true } // первый проход тоже срабатывать
    );
    function selectTable(val) {
      switch (val) {
        case "title":
          currentComponent.value = defineAsyncComponent(
            () => import("src/components/CatalogAllTables/PageTbl.vue") //"src/components/Sprav/tabBakery/TablePanel.vue")
          );
          setTitle("Справочники");
          break;
        case "bakery":
          currentComponent.value = defineAsyncComponent(
            () => import("components/Bakery/BakeryTable.vue") //"src/components/Sprav/tabBakery/TablePanel.vue")
          );
          setTitle("Пекарни");
          break;
        case "kagent":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Kagent/KagentTable.vue")
          );
          setTitle("Контрагенты");
          break;
        case "brand":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Sprav/tabBrand/TablePanel.vue")
          );
          setTitle("Бренды");
          break;
        //---------------------
        case "assortament":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productassortment/Table.vue")
          );
          setTitle("Ассортимент");
          break;
        case "raw":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productraw/Table.vue")
          );
          setTitle("Сырьё");
          break;
        case "typeraw":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productrawvid/Table.vue")
          );
          setTitle("Тип сырья");
          break;
        case "products":
          currentComponent.value = defineAsyncComponent(
            //() => import("src/components/Products/products/Table.vue")
            () => import("src/components/Products/PageProductsExt.vue")
          );
          setTitle("Продукция");
          break;
        case "typeproduct":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/producttype/Table.vue")
          );
          setTitle("Тип продукции");
          break;
        case "tovar":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/productvid/Table.vue")
          );
          setTitle("Товар");
          break;
        case "city":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/City/Table.vue")
          );
          setTitle("Города");
          break;
        case "department":
          currentComponent.value = defineAsyncComponent(() =>
            import("components/Users/PageUsersTree.vue")
          );
          setTitle("Пользователи");
          break;
        case "bakeryconfig":
          currentComponent.value = defineAsyncComponent(() =>
            import("components/Sprav/PageSprav.vue")
          );
          setTitle("Конфиг. пекарен");
          break;
        case "kagentvid":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Sprav/tabKagentVid/TableKagentVid.vue")
          );
          setTitle("Виды контрагента");
          break;
        case "kagentreg":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Sprav/tabKagentReg/TableKagentVidReg.vue")
          );
          setTitle("Виды рег. контрагента");
          break;
        case "kagentgroups":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Sprav/tabKagentGroups/TableKagentGroups.vue")
          );
          setTitle("Группы контрагента");
          break;
        case "trademark":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Sprav/tabTrademark/TablePanel.vue")
          );
          setTitle("Торговые сети");
          break;
        case "recept":
          currentComponent.value = defineAsyncComponent(() =>
            import("src/components/Products/Recepts/Recepts.vue")
          );
          setTitle("x Рецепт");
          break;
        case "sale":
          currentComponent.value = defineAsyncComponent(() =>
            import("components/Sale/PageSale.vue")
          );
          setTitle("Продажи");
          break;
        case "products":
        case "products2":
          currentComponent.value = defineAsyncComponent(() =>
            import("components/Products/PageProducts.vue")
          );
          setTitle("Продукция");
          break;
        case "price":
          currentComponent.value = defineAsyncComponent(() =>
            import("components/Price/PagePrice.vue")
          );
          setTitle("Прайс ТТК");
          break;
        default:
          currentComponent.value = defineAsyncComponent(
            () => import("src/components/CatalogAllTables/PageTbl.vue") //"src/components/Sprav/tabBakery/TablePanel.vue")
          );
          setTitle("Справочники");
          break;
        // currentComponent.value = defineAsyncComponent(() =>
        //   import("src/components/Sprav/tabBrand/TablePanel.vue")
        // );
      }
    }
    function setTitle(val) {
      let title = val;
      document.title = "ХиТ";
      document.title += title ? " | " + title : "";
    }
    return {
      cardMain,
      clickHelp,
      currentComponent,
      setTitle,
      route,
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

<style lang="scss" scoped>
// .slide-fade-enter-active {
//   transition: all 0.3s ease-out;
// }

// .slide-fade-leave-active {
//   transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
// }

// .slide-fade-enter-from,
// .slide-fade-leave-to {
//   transform: translateX(200px);
//   opacity: 0;
// }
</style>
<style lang="scss" scoped>
:deep(.arkadii-sticky-header-table) {
  /* height or max-height is important */
  // height: 310px;

  // .q-table__top,
  // .q-table__bottom,
  thead tr:first-child th {
    /* bg color is important for th; just specify one */
    background-color: #f2f8fd;
  }
  thead tr th {
    position: sticky;
    z-index: 1;
  }
  thead tr:first-child th {
    top: 0;
  }
  /* this is when the loading indicator appears */
  &.q-table--loading thead tr:last-child th {
    /* height of all previous header rows */
    top: 48px;
  }
}
:deep(.arkadii-table-header) {
  text-align: center;
}
</style>
