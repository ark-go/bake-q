<template>
  <div class="column no-wrap" style="max-height: inherit">
    <Table-Template
      flat
      title="Товар"
      :rows="rows"
      :columns="columns"
      :tableFunc="tableFunc"
      yesBtnEdit
      yesBtnDelete
      @onInfoRow="onInfoRow"
      @onBtnDelete="onDelete"
      @onBtnEdit="onEdit"
      @onRowClick="onRowClick"
      @onAdd="onAdd"
      @onRowDblClick="dblClickRow"
      :currentRow="currentRow"
      noExpandPanel
      :noEditTable="false"
      :store="store"
      :rowsPerPage="0"
    >
    </Table-Template>
  </div>
  <form-dialog
    :rowData="currentRow"
    :allSprav="allSprav"
    v-model:showDialog="showDialog"
    @onSave="onSave"
    :title="tablabel"
  ></form-dialog>
</template>

<script>
import {
  defineComponent,
  ref,
  onMounted,
  computed,
  watch,
  watchEffect,
  unref,
} from "vue";
import { useArkUtils } from "src/utils/arkUtils"; // const arkUtils = useArkUtils();
import FormDialog from "./FormDialog.vue";
import { useQuasar } from "quasar";
import TableTemplate from "src/components/template/table/TableTemplate.vue";

export default defineComponent({
  name: "SpravTable",
  components: {
    TableTemplate,
    FormDialog,
  },
  props: {
    tabname: { type: String, default: "productvid" },
    tablabel: { type: String, default: "Товар" },
  },
  setup(props, { emit }) {
    const $q = useQuasar();
    const arkUtils = useArkUtils();
    const rows = ref([]);
    const visibleColumns = ref([]);
    const showDialog = ref(false);
    const currentRow = ref({});
    const allSprav = ref();
    //const visibleOffDefault = ref([]);
    //const columns = ref([]);
    onMounted(async () => {
      await loadTable();
      columnFilter();
    });
    async function onSave(row) {
      if (row?.id) console.log("Готов записывать Обновления", row);
      else console.log("Готов записывать Новый объект", row);
      console.log("SAVE1: ", row);
      await addTable(row);
    }
    function columnFilter() {
      visibleColumns.value = [];
      // columns.forEach((item, index, array) => {
      //   if (visibleOffDefault.includes(item.name)) return;
      //   visibleColumns.value.push(item.name);
      // });
    }

    async function loadTable() {
      let mess = "Загрузка Видов продукции";
      console.log("load: ", props.tabname);
      let res = await arkUtils.dataLoad(
        "/api/products",
        { cmd: "load", tabname: props.tabname },
        mess
      );
      if (res.result) {
        rows.value = res.result;
      } else {
        rows.value = [];
      }
    }
    async function deleteTable(row) {
      let mess = "Удаление";
      let res = await arkUtils.dataLoad(
        "/api/products",
        { id: row.id, cmd: "delete", tabname: props.tabname },
        mess
      );
      if (res.result) {
        rows.value = res.result;
        try {
          let idx = rows.value.findIndex((val) => {
            return val.id == row.id;
          });
          if (idx != -1) {
            // если нашли то что удаляли, то удалим из нашего Array таблицы
            //  rows.value.splice(idx, 1);
            await loadTable();
            //console.log("посде удал", rows.value);
          }
        } catch {
          console.log("тут ошибка № 2410");
        }
      } else {
        // Мы не удалили ничего
        //await loadTable();
      }
    }
    async function addTable(row) {
      let mess = "Добавление вида сырья";
      console.log("SAVE-0: ", row);
      let res = await arkUtils.dataLoad(
        "/api/products",
        { ...row, ...{ cmd: "add", tabname: props.tabname } },
        mess
      );
      if (res.result) {
        // получим строку в []
        console.log("Записали", res.result);
        //  console.log("return one row", res.result[0]);
        if (row.id) {
          // если передавали id значит был Update
          // найдем в нашем Array таблице id обновляемой строки
          let idx = rows.value.findIndex((val) => {
            return val.id == row.id;
          });
          if (idx != -1) {
            // если нашли то что обновляли, то заменим новой строкой
            rows.value[idx] = res.result[0];
          } else {
            //! Запустить полное перечитывание
            // не может произойти
            console.log("Аварийное перечитывание");
            await loadTable();
          }
        } else {
          // передавали без Id значит была новая строка
          // вставим полученную новую строку в наш Array таблицы, вперед
          rows.value.unshift(res.result[0]);
        }
        showDialog.value = false;
      } else {
        //! не обнуляем таблицу, у нас открыто окно ввода данных
        // rows.value = [];
      }
    }
    async function onDelete(val) {
      //------------- Dialog
      $q.dialog({
        title: "Удалить запись?",
        message: val.name,
        cancel: true,
        persistent: true,
        ok: { label: "Удалить", color: "red-3" }, // q-btn
        cancel: { label: "Отменить", color: "blue-5" },
        focus: "cancel",
      })
        .onOk(async () => {
          // console.log('>>>> OK')
          await deleteTable(val);
        })
        .onOk(() => {
          console.log(">>>> second OK catcher");
        })
        .onCancel(() => {
          // console.log('>>>> Cancel')
        })
        .onDismiss(() => {
          // console.log('I am triggered on both OK and Cancel')
        });

      //--------------------
    }
    async function loadAllSprav() {
      let res = await arkUtils.dataLoad(
        "/api/products",
        { cmd: "allSprav", tabname: props.tabname },
        "Чтение справочников для вида продукции"
      );
      return res?.result || [];
    }
    async function showDialogStart(row) {
      allSprav.value = await loadAllSprav();
      currentRow.value = row;
      showDialog.value = true;
    }
    const selectedRow = ref([]);
    function onRowClick(row) {
      console.log("clicked on", row);

      if (selectedRow.value.indexOf(row) == -1) {
        selectedRow.value.length = 0;
        selectedRow.value.push(row);
      } else {
        selectedRow.value.splice(selectedRow.value.indexOf(row), 1);
      }
    }
    function onRowClick(row) {
      currentRow.value = row;
      emit("selectedRow", row);
    }
    return {
      onRowClick,
      selectedRow,
      showDialog,
      showDialogStart,
      allSprav,
      onDelete,
      onSave,
      currentRow,
      rows,
      filter: ref(""),
      paginationСatalog: ref({
        rowsPerPage: 10,
      }),
      columns,
      visibleColumns,
      visibleOffDefault,
      async onAdd(row) {
        await showDialogStart(row);
      },
      async onEdit(row) {
        await showDialogStart(row);
      },
    };
  },
});
let visibleOffDefault = ["user_email", "user_date"];
let columns = [
  {
    name: "producttype_prefix",
    label: "Тип",
    align: "left",
    field: "producttype_prefix",
  },
  {
    name: "productassortment_name",
    label: "Ассортимент",
    align: "left",
    field: "productassortment_name",
    required: true,
  },
  {
    name: "name_display",
    label: "Наименование",
    align: "left",
    field: "name_display",
    required: true,
  },
  // {
  //   name: "fullname",
  //   label: "Полное название",
  //   align: "left",
  //   field: "fullname",
  // },
  {
    name: "unit_name",
    label: "ед.изм.",
    align: "left",
    field: "unit_name",
    style: "width: 50px",
    //required: true,
  },
  {
    name: "description",
    label: "Примечание",
    align: "left",
    field: "description",
    hidden: true,
  },
];
</script>
<style lang="scss" scoped>
.select-row {
  color: red;
}
</style>
