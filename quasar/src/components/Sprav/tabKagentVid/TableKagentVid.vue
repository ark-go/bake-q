<template>
  <div class="column no-wrap" style="display: grid; max-height: inherit">
    <Table-Template
      flat
      :title="tableInfo.label"
      :rows="rows"
      :columns="columns"
      yesBtnEdit
      yesBtnDelete
      @onBtnDelete="onDeleteData"
      @onBtnEdit="onEditRow"
      @onRowClick="onRowClick"
      @onAdd="addNew"
      @onRowDblClick="onEditRow"
      :currentRow="spravStore?.selectedRow"
      noExpandPanel
      :noEditTable="false"
      :store="store"
      :rowsPerPage="0"
      noInfoBtn
    >
    </Table-Template>
  </div>
  <One-Value-Dialog
    :inputArr="inputArr"
    v-model:showDialog="showDialog"
    @inputForSave="onSaveRow"
  ></One-Value-Dialog>
</template>

<script>
import {
  defineComponent,
  defineAsyncComponent,
  ref,
  onMounted,
  computed,
  watch,
  watchEffect,
  unref,
} from "vue";
import { useArkUtils } from "src/utils/arkUtils"; // const arkUtils = useArkUtils();
import { useQuasar } from "quasar";
import OneValueDialog from "./OneValueDialog.vue";
import { useSpravStore } from "stores/spravStore";
export default defineComponent({
  name: "TableKagentVid",
  components: {
    OneValueDialog,
    TableTemplate: defineAsyncComponent(() => {
      return import("src/components/template/table/TableTemplate.vue");
    }),
  },
  props: {
    //tableName: String,
    tableInfo: {
      type: Object,
      default: () => {
        return {
          label: "Вид контрагента",
          tableName: "kagentvid",
        };
      },
    },
  },
  setup(props) {
    const $q = useQuasar();
    const arkUtils = useArkUtils();
    const spravStore = useSpravStore();
    const rows = ref([]);
    const inputArr = ref({});
    const showDialog = ref(false);
    async function restartComponent() {
      // необходимо при создании, или вставке :is компонета срабатывает
      if (props.tableInfo.tableName) {
        await loadTable();
      } else {
        rows.value = [];
      }
    }
    function onRowClick(row) {
      spravStore.selectedRow = row;
    }
    onMounted(async () => {
      await restartComponent();
    });
    watch(
      () => props.tableInfo.tableName,
      async () => {
        await restartComponent();
      }
    );

    async function loadTable() {
      let dat = {
        tableNameLoad: props.tableInfo.tableName,
      };
      let mess = props.tableInfo.label;
      let res = await arkUtils.dataLoad("/api/spravLoad", dat, mess);
      if (res.result) {
        rows.value = res.result;
      } else {
        rows.value = [];
      }
    }
    async function onEditRow(row) {
      inputArr.value = { ...row };
      showDialog.value = true;
    }
    function addNew() {
      inputArr.value = {};
      showDialog.value = true;
    }
    async function onSaveRow(row) {
      if (row.id) {
        await onSaveData(row);
      } else {
        await inputForSave(row);
      }
    }
    async function onSaveData(row) {
      let dataToBase = {};
      dataToBase.name = row.name;
      dataToBase.id = row.id;
      dataToBase["tableName"] = props.tableInfo.tableName;
      console.log("состав:", dataToBase);
      let mess = "Обновление " + props.tableInfo.label;
      let res = await arkUtils.dataLoad("/api/spravUpdate", dataToBase, mess);
      if (res.error) {
        return;
      }
      await loadTable();
    }
    // ----------------------------
    async function onDeleteConfirm(row) {
      let dataToBase = {
        id: row.id,
        tableName: props.tableInfo.tableName,
      };
      let mess = "Удаление " + props.tableInfo.label + " / " + row.name;
      let res = await arkUtils.dataLoad("/api/spravDelete", dataToBase, mess);
      if (res.error) {
        return;
      }
      await loadTable();
    }
    async function onDeleteData(row) {
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
      //--------------------
    }

    //--------------------
    async function inputForSave(row) {
      let dataToBase = {};
      dataToBase.name = row.name;
      dataToBase.id = row.id;
      dataToBase["tableName"] = props.tableInfo.tableName;
      console.log("состав:", dataToBase);
      let mess = "Добавление " + props.tableInfo.label + "/" + row.name;
      let res = await arkUtils.dataLoad("/api/spravAdd", dataToBase, mess);
      if (res.error) {
        return;
      }
      await loadTable();
    }

    const columns = ref([
      {
        name: "name",
        label: "Наименование",
        align: "left",
        field: "name",
        required: true,
      },
      {
        name: "user_email",
        label: "Е-Mail",
        align: "left",
        field: "user_email",
        hidden: true,
      },
      {
        name: "user_date",
        label: "Дата",
        align: "left",
        field: "user_date",
        hidden: true,
      },
    ]);
    return {
      spravStore,
      onRowClick,
      rows,
      columns,
      onSaveData,
      addNew,
      inputArr,
      showDialog,
      inputForSave,
      onDeleteData,
      onEditRow,
      onSaveRow,
    };
  },
});
</script>
<style lang="scss">
.table-sprav-column-table {
  /* указав максимальную ширину, чтобы пример мог
выделить липкий столбец в любом окне браузера */
  //max-width: 40%;

  thead tr:first-child th:first-child {
    /* bg цвет важен для th; просто укажите один */
    background-color: #fff;
  }
  td:first-child {
    background-color: #fdfdfc;
  }
  th:first-child,
  td:first-child {
    position: sticky;
    left: 0;
    z-index: 1;
    max-width: 30%;
  }
}
</style>
