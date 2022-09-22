<template>
  <ark-card
    title="Продукция"
    :subTitle="currentLabel"
    :buttonArr="buttonArr"
    @buttonClick="buttonClick"
    :menuObj="{ pdf: 'Получить PDF' }"
    @menuClick="menuClick"
  >
    <!-- <keep-alive include="ProductsTree"> -->
    <div class="row" style="max-height: inherit">
      <div class="col-4" style="max-height: inherit">
        <div v-if="!currentRowText" style="max-height: inherit">
          <div v-if="yesCatalogAll">Выберите продукт</div>
          <products-tree
            v-else
            style="min-width: 100px"
            @on-selected-node="onSelectedNode"
          ></products-tree>
        </div>
        <products-action v-else @on-click-pdf="menuClick('pdf')">
        </products-action>
      </div>

      <div class="col-8 ark-grid-right88" style="max-height: inherit">
        <!-- <transition appear name="comp-fade" mode="out-in"> -->
        <component
          v-if="!!currentTable"
          :is="currentTable"
          v-bind="{
            tableInfo: selectedNode2,
            tabname: currentTabname,
            tablabel: currentLabel,
          }"
        ></component>
        <!-- </transition> -->
      </div>
    </div>
    <!-- </keep-alive> -->
    <Pdf-Dialog
      v-model:showDialog="showPdfDialog"
      :param="pdfDialogParam"
    ></Pdf-Dialog>
  </ark-card>
</template>

<script>
import {
  defineComponent,
  defineAsyncComponent,
  ref,
  computed,
  onActivated,
  nextTick,
  onMounted,
  watch,
} from "vue";
//import NoTable from "components/Products/NoTable.vue";
//import ArkCard from "components/Card/ArkCard.vue";
import ArkCard from "components/Products/ArkCard.vue";
import ProductsTree from "components/Products/ProductsTree.vue";
import { useQuasar } from "quasar";
import { useRoute, useRouter } from "vue-router";
import { arkVuex } from "src/utils/arkVuex.js";
import ProductsAction from "./ProductsAction.vue";
import PdfDialog from "../PDF/PdfDialog.vue";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
import { useProductsStore } from "stores/productsStore.js";
//import { arkVuex } from "src/utils/arkVuex"; // const { pdfWindow } = createArkVuex();
export default defineComponent({
  name: "f-products",
  components: {
    ArkCard,
    ProductsTree,
    ProductsAction,
    PdfDialog,
  },
  props: [],
  emits: ["onToRecept"],
  setup(props, { emit }) {
    const { pdfWindow } = arkVuex();
    const { selectedRow } = storeToRefs(useProductsStore());
    const $q = useQuasar();
    const { cardMain, currentPage } = storeToRefs(usePagesSetupStore());
    // const { selectedRowsVue1x } = arkVuex();
    const selectedNode = ref({});
    const selectedNode2 = ref({});
    const currentTabname = ref("");
    const router = useRouter();
    const route = useRoute();
    const currentTable = ref(null); // подключаемые таблицу
    const currentLabel = ref(null);
    const showPdfDialog = ref(false);
    const pdfDialogParam = ref({});
    const yesCatalogAll = ref(false);
    onMounted(() => {
      currentPage.value = "products";
    });
    const currentRowText = computed(() => {
      if (selectedRow.value?.id)
        return `${selectedRow.value.productvid_name}<br>${selectedRow.value.name} ${selectedRow.value.description}`;
      // let row = selectedRowsVue1x?.products[0];
      // if (row) {
      //   let str =
      //     row.productvid_name + "<br> " + row.name + " " + row.description;
      //   return str;
      // }
      return "";
    });
    onActivated(() => {
      // Возникает при входе при keep-alive
      // если был сброшен выбор дерева то очищаем выбор продукта
      // выбор сбросится при переходе с другой страницы
      // if (!currentLabel.value) selectedRowsVue1x.products.length = 0;
      if (!currentLabel.value) selectedRow.value = {};
    });
    watch(
      () => route.path,
      (val) => {
        // если открыли страницу по пути справочника tbl то симитируем переход по пути дерева, на продукцию
        if (val.includes("/tbl/")) {
          console.log(
            "В продукци имитируем выбор дерева продукция, если вход из справочника"
          );
          yesCatalogAll.value = true;
          onSelectedNode({
            key: "start",
            label: "Продукция",
            table: "products",
          });
        } else {
          onSelectedNode({
            key: "start",
            label: "Продукция",
            table: "products",
          });
        }
      },
      { immediate: true }
    );

    function onSelectedNode(node) {
      //selectedRowsVue1x.products.length = 0;
      //selectedRow.value = {};
      currentLabel.value = node.label;
      switch (node.table) {
        case "producttype":
          currentTabname.value = "producttype";
          currentTable.value = defineAsyncComponent(() =>
            import("./producttype/Table.vue")
          );

          break;
        case "productrawvid":
          currentTabname.value = "productrawvid";
          currentTable.value = defineAsyncComponent(() =>
            import("./productrawvid/Table.vue")
          );

          break;
        case "productassortment":
          currentTabname.value = "productassortment";
          currentTable.value = defineAsyncComponent(() =>
            import("./productassortment/Table.vue")
          );

          break;
        case "productvid":
          currentTabname.value = "productvid";
          currentTable.value = defineAsyncComponent(() =>
            import("./productvid/Table.vue")
          );
          break;
        case "productraw":
          currentTabname.value = "productraw";
          currentTable.value = defineAsyncComponent(() =>
            import("./productraw/Table.vue")
          );
          break;
        case "products":
          currentTabname.value = "products";
          currentTable.value = defineAsyncComponent(() =>
            import("./products/Table.vue")
          );
          break;
        default:
          currentTable.value = null;
          currentLabel.value = null;
          break;
      }
    }

    const buttonArr = ref([
      { key: "backRoute", name: "Назад" },
      //  { key: "Добавить", name: "Второй" },
    ]);
    function buttonClick(val) {
      if (val == "backRoute") {
        console.log(val);
        router.go(-1);
      }
    }
    async function menuClick(val, val2) {
      if (val == "pdf") {
        pdfDialogParam.value = {
          typePdf: "base64", // file/base64
          tgFormat: "pdf", // pdf/jpg
          command: "products",
          fileName: "Продукты",
          // id: selectedRowsVue1x?.products[0]?.id,
          id: selectedRow.value?.id,
        };
        nextTick(() => {
          showPdfDialog.value = true;
        });
      }
    }
    return {
      cardMain,
      pdfDialogParam,
      menuClick,
      showPdfDialog,
      currentTable,
      currentLabel,
      selectedNode,
      selectedNode2,
      onSelectedNode,
      currentTabname,
      buttonArr,
      buttonClick,
      // selectedRowsVue1x,
      currentRowText,
      route,
      yesCatalogAll,
    };
  },
});
</script>

<style lang="scss" scoped></style>

<style lang="scss" scoped>
.ark-grid {
  display: grid;
  //  padding: 6px;
  gap: 10px;
  grid-template-columns: 30% 1fr;
  overflow: auto; // для центра
  max-height: inherit; // размер центра
  .ark-grid-left {
    overflow: auto;
    max-height: inherit;
  }
  .ark-grid-right {
    overflow: auto;
    max-height: inherit;
  }
  @media (max-width: 600px) {
    grid-template-columns: 100%;
    min-width: 94vw;
    .ark-grid-left {
      overflow: unset;
    }
    .ark-grid-right {
      overflow: unset;
    }
  }
}
//
.comp-fade-enter-active,
.comp-fade-leave-active {
  transition: opacity 0.3s ease;
}

.comp-fade-enter-from,
.comp-fade-leave-to {
  opacity: 0;
}
</style>
