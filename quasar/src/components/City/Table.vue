<template>
  <div class="column no-wrap" style="display: grid; max-height: inherit">
    <Table-Template
      flat
      title="Города"
      :rows="rows"
      :columns="columns"
      :tableFunc="tableFunc"
      yesBtnEdit
      yesBtnDelete
      @onInfoRow="onInfoRow"
      @onBtnDelete="onDeleteData"
      @onBtnEdit="onEditData"
      @onRowClick="onRowClick"
      @onAdd="addNew"
      :currentRow="currentRow"
      noExpandPanel
      :noEditTable="false"
      :store="store"
      :rowsPerPage="0"
    >
    </Table-Template>
  </div>
  <trademark-dialog
    v-model:showDialog="showDialog"
    :sprav="sprav"
    :row-data="rowToDialog"
    @onSave="onSaveData"
  ></trademark-dialog>
</template>

<script>
import {
  defineComponent,
  ref,
  isRef,
  onMounted,
  computed,
  watch,
  watchEffect,
  unref,
} from "vue";
import { useArkUtils } from "src/utils/arkUtils"; // const arkUtils = useArkUtils();
import { getSprav } from "./getSprav";
import { useQuasar } from "quasar";
import TableTemplate from "../template/table/TableTemplate.vue";
import TrademarkDialog from "./Dialog.vue";
export default defineComponent({
  name: "CityTable",
  components: {
    TableTemplate,
    TrademarkDialog,
  },
  props: {
    subTitle: String,
    buttonArr: Object,
    menuObj: Object,
    maxWidth: String,
  },
  setup(props) {
    const $q = useQuasar();
    const arkUtils = useArkUtils();
    const tableNameSting = ref("");
    const rows = ref([]);
    const visibleColumns = ref([]);
    const filter = ref("");
    const addNewEnabled = ref(false); //включаем кнопку
    const showDialog = ref(false);
    const confirmDelete = ref({});
    const nameElement = ref("");
    const url = ref("/api/");
    const sprav = ref({}); // Справочники сюда
    const rowToDialog = ref({});

    const paginationСatalog = ref({
      rowsPerPage: 10,
    });
    async function restartComponent() {
      // необходимо при создании, или вставке :is компонета срабатывает
      sprav.value.region = await getSprav("region_kl", "Регионы"); // если нужны справочники
      console.log("region", sprav.value.region);
      url.value = url.value + "city";
      await loadTable(); // читаем, и там заполняем rows
    }
    onMounted(async () => {
      // срабатывает при назначении таблицы  :is=
      await restartComponent(); // подготовка переменных, загрузка
    });

    // ------------- //!не применяем  -------------------------------
    function dblClickRow(val) {
      console.log("sprav dblClick", val);
    }
    // ------------- закрываем/сворачиваем лишние колонки  -------------------------------
    // function columnFilter() {
    //   visibleColumns.value = [];
    //   columns.value.forEach((item, index, array) => {
    //     if (item.name == "user_email" || item.name == "user_date") return;
    //     visibleColumns.value.push(item.name);
    //   });
    //   console.log("Колонки показать", visibleColumns.value);
    // }

    // -------------  Запись строки Update-------------------------------
    async function onSaveData(val) {
      let valUn = isRef(val) ? val.value : val;
      let cmd = "add";
      if (valUn.id) {
        cmd = "update";
      }
      console.log("Пришло и готово на запись ", cmd, valUn);
      let mess = "Обновление города"; // сообщение в загрузчик для
      let res = await arkUtils.dataLoad(
        url.value,
        { ...valUn, cmd: cmd }, // если нет cmd то update
        mess
      );
      if (res.error) {
        return console.log("Ошибка записи:");
      }
      await loadTable(); // перечитывам таблицу
    }
    //------------ Добавление новой строки
    // async function onAddData(val) {
    //   console.log("Данные для новой записи", val);
    //   await onSaveData(val, "add");
    // }
    // ------------ Удаление ----------------
    async function deleteData(val) {
      // приходит строка таблицы в row
      console.log("Пришло и готово на удаление ", val.id);
      let mess = "Удаление " + tableNameSting.value + " / " + val.name;
      let res = await arkUtils.dataLoad(
        url.value,
        { ...val, cmd: "delete" },
        mess
      );
      if (res.error) {
        return console.log("Ошибка удаления:", dataToBase);
      }
      await loadTable();
    }
    // ------------ Запрос подтверждения удаления
    async function onDeleteData(val) {
      console.log("FFFFFFFFFFFFFFFFFFFFFFFFFFFFF", val);
      $q.dialog({
        title: "Удалить запись?",
        message: val.name,
        cancel: true,
        persistent: true,
        //! цвета вынести
        ok: { label: "Удалить", color: "red-3" }, // q-btn
        cancel: { label: "Отменить", color: "blue-5" },
        focus: "cancel",
      }).onOk(() => {
        deleteData(val);
      });
    }

    // ------------- открываем диалог
    function addNew() {
      rowToDialog.value = {};
      showDialog.value = true;
    }
    function onEditData(val) {
      console.log("ЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖЖ", val);
      rowToDialog.value = { ...val };
      showDialog.value = true;
    }
    // ------------------- читаем все данные
    async function loadTable() {
      tableNameSting.value = "city";
      let dat = {
        cmd: "load",
      };
      let mess = "Города ";
      let res = await arkUtils.dataLoad(url.value, dat, mess);
      if (res.result) {
        rows.value = res.result;
      } else {
        tableNameSting.value = "Выберите справочник";
        rows.value = [];
      }
    }
    // ----------- поля в таблице
    const columns = ref([
      {
        name: "name",
        label: "Наименование",
        align: "left",
        field: "name",
        required: true,
      },
      {
        name: "region_name",
        label: "Регион",
        align: "left",
        field: "region_name",
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
      sprav,
      url,
      tableNameSting,
      rows,
      filter,
      paginationСatalog,
      visibleColumns,
      columns,
      dblClickRow,
      test(val) {
        console.log(val);
      },
      onSaveData,
      addNewEnabled,
      addNew,
      showDialog,
      onDeleteData,
      confirmDelete,
      nameElement,
      rowToDialog,
      onEditData,
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
