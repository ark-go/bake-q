<template>
  <Table-Template
    v-if="tableName"
    flat
    :title="title"
    :tableName="tableName"
    :rows="rows"
    :columns="columns"
    :tableFunc="tableFunc"
    yesBtnEdit
    yesBtnDelete
    @onInfoRow="onInfoRow"
    @onBtnDelete="onDeleteData"
    @onBtnEdit="rowEdit"
    @onRowClick="onRowClick"
    @onAdd="addNew"
    :currentRow="currentRow"
    noExpandPanel
    :noEditTable="false"
    :store="store"
    :rowsPerPage="0"
  >
  </Table-Template>
  <div v-else>не указана таблица</div>
  <Trademark-Dialog
    :sprav="sprav"
    :rowData="rowToDialog"
    v-model:showDialog="showDialog"
    @onSave="dialogSave"
  ></Trademark-Dialog>
</template>
<script>
import {
  defineComponent,
  ref,
  unref,
  defineAsyncComponent,
  onMounted,
  watch,
  watchEffect,
} from "vue";
import { useTableFunc } from "./tableFunc.js";
import { columns } from "./tableColumnList.js";
import { useBakeryStore } from "stores/bakeryStore.js";
import { useSpravStore } from "stores/spravStore";
import { useQuasar } from "quasar";
//import FormDialog from "./FormDialog.vue";
import TrademarkDialog from "./TrademarkDialog.vue";
export default defineComponent({
  name: "TablePanel",
  components: {
    TrademarkDialog,
    TableTemplate: defineAsyncComponent(() => {
      return import("src/components/template/table/TableTemplate.vue");
    }),
  },
  props: {
    modeBody: {
      type: String,
      default: "view",
    },
    tableName: {
      type: String,
      default: "tabTrademark", // для запроса с сервера
    },
    tableInfo: {
      type: Object,
      default: () => ({ tableName: "tabTrademark" }),
    },
    title: {
      type: String,
      default: "Торговые сети",
    },
    panelName: String,
  },
  emits: [""],
  setup(props, { emit }) {
    const $q = useQuasar();
    const showDialog = ref(false);
    const currentRow = ref({});
    const rows = ref([]);
    const tableFunc = useTableFunc(rows, props.tableInfo.tableName);
    const spravStore = useSpravStore();
    const store = useBakeryStore();
    const rowToDialog = ref({});
    const sprav = ref({ brand: [] });

    watch(
      () => spravStore.historyDate,
      async () => {
        await tableFunc.loadTable();
      }
    );
    watch(
      () => spravStore.currentTab,
      async () => {
        // Ловим переключение складки

        if (spravStore.currentTab == "main") {
          await tableFunc.loadTable();
        }
      }
    );
    onMounted(async () => {
      sprav.value.brand = await tableFunc.loadBrand(); // если нужны справочники
      // console.log("brand ", sprav.value.brand);
      await tableFunc.loadTable(
        tableFunc.dateToDateUnix(spravStore.historyDate)
      );
    });
    function rowEdit(row) {
      rowToDialog.value = { ...row };
      console.log("rowEdit", row.name, row);
      showDialog.value = true;
    }
    async function dialogSave(row) {
      console.log("Данные для новой записи", row);
      await tableFunc.saveRowTable(row);
    }
    function addNew() {
      rowToDialog.value = {};
      console.log("addNew");
      showDialog.value = true;
    }
    async function onDeleteData(row) {
      currentRow.value = row;
      $q.dialog({
        title: "Удалить запись?",
        message: row.name,
        cancel: true,
        persistent: true,
        ok: { label: "Удалить", color: "red-3" }, // q-btn
        cancel: { label: "Отменить", color: "blue-5" },
        focus: "cancel",
      }).onOk(() => {
        onDeleteConfirm(row);
      });
    }
    async function onDeleteConfirm(row) {
      console.log("Пришло и готово на удаление ", row.id);
      let mess = "Удаление " + row.name;
      await tableFunc.deleteTable(row);
    }
    return {
      rowToDialog,
      unref,
      showDialog,
      rowEdit,
      addNew,
      onDeleteData,
      dialogSave,
      currentRow,
      store,
      rows,
      columns,
      //  tableBodyMenu,
      tableFunc,
      onInfoRow(row) {
        console.log("info butt", row);
      },
      onRowClick(row) {
        console.log("Нажали по строке");
        currentRow.value = row;
        store.selectedRow = row;
        emit("selectedRow", row);
      },
      sprav,
    };
  },
});
</script>

<style lang="scss" scoped>
.my-sticky-header-table {
  /* height or max-height is important */
  height: 310px;

  .q-table__top,
  .q-table__bottom,
  thead tr:first-child th {
    /* bg color is important for th; just specify one */
    background-color: #c1f4cd;
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
</style>
