<template>
  <ark-card
    flat
    class="ark-card-panel"
    style="overflow: auto"
    :title="'Составление рецепта' + ttkNumberStr"
    :subTitle="curentRowName"
    :buttonArr="buttonArr"
    @buttonClick="buttonClick"
    :menuObj="{ pdf: 'Получить PDF' }"
    @menuClick="menuClick"
  >
    <template v-slot:leftTop>
      <button-down
        @on-change-razdel="onChangeRazdel"
        :label="labelButtonDown"
      ></button-down>
    </template>
    <template v-slot:leftCenter>
      <table-raw
        :tabname="currentTableName"
        :loadTableName="loadTableName"
        @add-ingredient="addIngredient"
        :ref="(el) => (refTable = el)"
      ></table-raw>
    </template>
    <template v-slot:rightTop>
      <q-field dense class="text-center text-black">
        <div
          style="
            justify-content: center;
            align-items: center;
            width: 100%;
            display: flex;
            color: black;
          "
        >
          ИНГРЕДИЕНТЫ
        </div>
      </q-field>
    </template>
    <template v-slot:rightCenter>
      <table-ingred
        :tabname="currentTableName"
        :isLeft="true"
        :ref="(el) => (refTableIngred = el)"
      ></table-ingred>
    </template>
  </ark-card>
</template>

<script>
import {
  defineComponent,
  onMounted,
  onUnmounted,
  defineAsyncComponent,
  ref,
  watchEffect,
} from "vue";
import ArkCard from "./ArkCard.vue";
//import { arkVuex } from "src/utils/arkVuex.js";
import { dom } from "quasar";
//import FormSelectIngr from "./FormSelectIngr.vue";
import ButtonDown from "./ButtonDown.vue";
import TableRaw from "./TableRaw.vue";
import TableIngred from "./TableIngred.vue";
import { useProductsStore, storeToRefs } from "stores/productsStore.js";
import { useRoute, useRouter } from "vue-router";

export default defineComponent({
  name: "p-recept",
  props: [],
  components: { ArkCard, TableRaw, TableIngred, ButtonDown },
  // emits: ["onToMain"],
  setup(props, { emit }) {
    const { height } = dom;
    //const { selectedRowsVue1x } = arkVuex();
    const { selectedRow } = storeToRefs(useProductsStore());
    const route = useRoute();
    const router = useRouter();
    const currentRow = ref({});
    const curentRowName = ref("");
    const ttkNumberStr = ref("");
    const currentTableName = ref("productraw"); // products
    const labelButtonDown = ref("");
    const topElement = ref({});
    const refTable = ref(null);
    const refTableIngred = ref(null);
    const loadTableName = ref("loadRaw");
    function onChangeRazdel(val) {
      currentTableName.value = val;

      switch (val) {
        case "products":
          labelButtonDown.value = "Раздел продукции";
          loadTableName.value = "loadProducts";
          break;
        case "productraw":
          labelButtonDown.value = "Раздел сырья";
          loadTableName.value = "loadRaw";
          break;
        default:
          break;
      }
    }
    async function addIngredient() {
      await refTableIngred.value.loadTable(); // перечитаем что там по умолчанию
      // при обновлении ингредиентов.. обновляем продукты, если у  нас таблица продуктов активна
      // там могла поменятся сумма
      if (currentTableName.value == "products") {
        await refTable.value.loadTable("loadProducts", "recept");
      }
    }
    onMounted(() => {
      if (selectedRow.value?.id) {
        onChangeRazdel("productraw");
      } else {
        //Данных о выбранном продукте не поступило
        router.push({ name: "tbl", params: { tblRouteParam: "products" } });
      }
    });

    watchEffect(() => {
      if (selectedRow.value?.id) {
        currentRow.value = selectedRow.value;
        curentRowName.value =
          currentRow.value.productvid_name + ", " + currentRow.value.name;
        ttkNumberStr.value = " ТТК№ " + currentRow.value.document_num;
      } else {
        currentRow.value = null;
        ttkNumberStr.value = "";
      }
    });

    function onBtnToogle(val) {
      console.log("Выбор: ", val);
    }
    const buttonArr = ref([
      { key: "back", name: "Обратно" },
      //  { key: "Добавить", name: "Второй" },
    ]);
    function buttonClick(val) {
      if (val == "back") {
        if (route.path.includes("/recept/")) {
          router.go(-1);
        }
        // else {
        //   emit("onToMain");
        // }
      }
      ////////////////////////////////
    }
    return {
      buttonArr,
      buttonClick,
      addIngredient,
      currentRow,
      curentRowName,
      ttkNumberStr,
      onBtnToogle,
      labelButtonDown,
      onChangeRazdel,
      currentTableName,
      topElement,
      refTable,
      refTableIngred,
      loadTableName,
    };
  },
});
</script>
<style lang="scss" scoped>
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
