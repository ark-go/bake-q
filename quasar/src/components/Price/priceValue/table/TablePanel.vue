<!--
    :tableName="tableName" - имя таблицы для обслуживания
    :rows="rows" - [] строки таблицы
    :columns="columns" [] - колонки 
     :tableBodyMenu="tableBodyMenu" - null - компонент, обработки меню правой мыши
    :tableFunc - Function  функция, ее подключаем по месту
    noExpandPanel  - скрыть расширение заголовка для управления
    noTitlePanel  -  скрыть заголовок таблицы
    noTopBtn  - не показывать кнопку с плюсом
    noTopFind - не показывать поле поиска
    noTopColumnSelect - не показывать выбор колонок
    noInfoBtn - убрать кнопку Info из строки
    yesBtnEdit - показывать кнопку редактирования
    yesBtnDelete - показывать кнопку удаления
    noEditTable - удаляет кнопки Edit и Delete и кнопку плюс в Top
    @onInfoRow - по кнопке Инфо
    @onBtnDelete - кнопка удалить в строке
    @onBtnEdit - кнопка едит в строке
    @onRowClick - по строке
    @onAdd - кнопка плюс в заголовке
    :iconBtnEdit="" - иконка для редактирования
    :iconBtnDelete="" - иконка для удаления
    :rowsPerPage  - кол-во строк таблице - странице
-->
<template>
  <Table-Template
    v-if="tableName"
    :title="title"
    :tableName="tableName"
    :rows="RowsPriceValue"
    :columns="columns"
    :tableBodyMenu="tableBodyMenu"
    @onInfoRow="$emit('onRowDblClick')"
    @onBtnDelete="onInfoRow"
    @onBtnEdit="onInfoRow"
    @onRowClick="onRowClick"
    @onAdd="onRowClick"
    @onRowDblClick="$emit('onRowDblClick')"
    :currentRow="selectedRowPrice"
    noExpandPanel
    noEditTable
    :rowsPerPage="0"
  >
  </Table-Template>
  <div v-else>не указана таблица</div>
</template>
<script>
import {
  defineComponent,
  ref,
  defineAsyncComponent,
  onMounted,
  watch,
} from "vue";
import { columns } from "./tableColumnList.js";
import { usePriceStore, storeToRefs } from "stores/priceStore.js";
import TableTemplate from "src/components/template/table/TableTemplate.vue";
//import { useLoadPriceValue } from "../../loadPriceValue.js";
export default defineComponent({
  name: "TablePanel",
  components: { TableTemplate },
  props: {
    modeBody: {
      type: String,
      default: "view",
    },
    tableName: {
      type: String,
      default: "tabPrice", // для запроса с сервера
    },
    title: {
      type: String,
      default: "Прайс",
    },
    checkSave: Boolean,
    panelName: String,
  },
  emits: [""],
  setup(props, { emit }) {
    //   const loadPriceValue = useLoadPriceValue();
    const tableBodyMenu = defineAsyncComponent(() => {
      return import("./TableBodyMenu.vue");
    });
    const priceStore = usePriceStore();
    const {
      RowsPriceValue,
      selectedRowDoc,
      selectedRowPrice,
      priceValueCount,
    } = storeToRefs(usePriceStore());
    //const rows = ref([]);
    // const currentRow = ref({});
    const pagination = ref({
      rowsPerPage: 10,
    });
    // onMounted(async () => {
    //   await tableFunc.loadTable();
    // });
    // priceStore.watchStore(() => {
    //   return watch(
    //     // сигнал на перезагрузку таблицы
    //     [() => props.checkSave], // () => selectedRowDoc.value.id], строку документа ловим в priceDoc
    //     async () => {
    //       // await loadTable();
    //       await loadPriceValue.loadTable();
    //     }
    //   );
    // });
    function onRowClick(row) {
      console.log("Нажали по строке");
      selectedRowPrice.value = row;
      //    priceStore.selectedRowPrice = row;
      //  emit("selectedRow", row);
    }
    return {
      //    currentRow,
      selectedRowPrice,
      priceStore,
      pagination,
      RowsPriceValue,
      columns,
      tableBodyMenu,
      onInfoRow(row) {
        console.log("info butt", row);
      },
      onRowClick,
    };
  },
});
</script>
