<template>
  <Tab-Panel-Split>
    <Side-Doc
      @onClickEdit="onClickEdit"
      @onClickNew="onClickNew"
      @onDelete="onDelete"
      v-model:trademarkId="trademarkId"
      @onAddDay="(val) => onAddDay(val)"
      @onKeyEnterCount="onKeyEnterCount"
      @refSelectCard="setRefSelectCard"
      @onSaveManualData="onSaveManualData"
      @onClipboardError="onClipboardError"
      @onGetExcel="onGetExcel"
      v-model:showDialogBuffer="showDialogBuffer"
      :rows="rowsManualData"
      :visibleSendSave="visibleSendSave"
    ></Side-Doc>
    <template v-slot:after>
      <Table-Panel
        :checkSave="checkSave"
        @onRowDblClick="onClickEdit"
        @onHideArticle="onHideArticle"
        @onShowArticle="onShowArticle"
      ></Table-Panel>
    </template>
  </Tab-Panel-Split>
  <!-- <form-doc v-model:showDialog="showDialog" @onSave="onSave"></form-doc> -->
</template>
<script>
import { ref, defineComponent, watch, onMounted } from "vue";
import TabPanelSplit from "../TabPanelSplit.vue";
import TablePanel from "./table/TablePanel.vue";
//import FormDoc from "./FormDoc.vue";
import SideDoc from "./side/SideDoc.vue";
import { useSaleStore, storeToRefs } from "stores/saleStore";
import { useTableFunc } from "./tableFunc";
import { date } from "quasar";
export default defineComponent({
  name: "TabSaleItems",
  components: { TabPanelSplit, TablePanel, SideDoc },
  setup() {
    const saleStore = useSaleStore();
    const {
      bakerySelectedRow,
      showHiddenArticle,
      showDoobleArticle,
      tabModel,
      trademarkId,
      checkDateSale,
      currentDateSale,
      selectedDateBetweenBakery,
      selectedArticleBakeryRow,
    } = storeToRefs(useSaleStore());
    const dateFormat = ref("DD.MM.YYYY");
    const tableFunc = useTableFunc("tabSale");
    const showDialog = ref(false);
    const showDialogBuffer = ref(false);
    const checkSave = ref(false);
    const refSelectCard = ref(null);
    const rowsManualData = ref([]);
    const visibleSendSave = ref(false);
    // const trademarkId = ref(null);
    onMounted(async () => {
      await tableFunc.loadBakeryArticle();
    });
    watch(
      [
        () => bakerySelectedRow.value,
        () => showHiddenArticle.value,
        () => showDoobleArticle.value,
        () => checkDateSale.value,
        () => currentDateSale.value,
      ],
      async (newVal, oldVal) => {
        // if (newVal != oldVal) articleBakeryRows.value = []; // ?????????? ?????????????????????? ??????????????????
        // console.log("trade", trademarkId.value);
        if (bakerySelectedRow.value.id) {
          await tableFunc.loadBakeryArticle();
          if (refSelectCard.value) refSelectCard.value.correctInputCount();
        }
      }
    );
    async function onHideArticle(menuObj) {
      await tableFunc.toggleHiddenArticle(menuObj.row.id, true);
    }
    async function onShowArticle(menuObj) {
      await tableFunc.toggleHiddenArticle(menuObj.row.id, false);
    }
    function onClipboardError() {
      rowsManualData.value = [];
    }
    async function onSaveManualData(val) {
      visibleSendSave.value = false;
      // ?????????????? ???????????? ???? ???????????? ???????????? ????????
      let resMan = await tableFunc.sendTextToTable(val);
      rowsManualData.value = resMan;
      if (val?.column && resMan.length > 0) {
        // columns - ?????????? ??????????
        visibleSendSave.value = true;
      }
    }
    // async function onSaveData(val) {
    //   console.log("save date table --", val);
    //   let resS = await tableFunc.insertBufferToSale(val.rows);
    //   if (resS) {
    //     showDialogBuffer.value = false;
    //     visibleSendSave.value = false;
    //     await tableFunc.loadBakeryArticle();
    //   }
    // }
    function onClickEdit() {
      // showDialog.value = true;
    }
    function onClickNew() {
      // bakerySelectedRow.value = {};
      // showDialog.value = true;
    }
    function onSave() {
      checkSave.value = !checkSave.value;
    }
    async function onGetExcel() {
      await tableFunc.exportPriceExcel();
    }
    function onAddDay(countDays) {
      console.log("onAddDay", countDays);
      saleStore.debonceArticle(() => {
        let minimumDate = date.extractDate(
          selectedDateBetweenBakery.value.from,
          dateFormat.value
        );
        let maximumDate = date.extractDate(
          selectedDateBetweenBakery.value.to,
          dateFormat.value
        );
        let currentDate = date.extractDate(
          currentDateSale.value,
          dateFormat.value
        );
        currentDate = date.addToDate(currentDate, { days: countDays });
        if (currentDate <= maximumDate && currentDate >= minimumDate) {
          currentDateSale.value = date.formatDate(
            currentDate,
            dateFormat.value
          );
        }
      });
    }
    function setRefSelectCard(val) {
      // ?????????????? ???????????????? ref ???? ??????????????????
      refSelectCard.value = val;
    }
    //
    let pressBloked = false;
    // ?????????????????? ?????????????? ENTER
    async function onKeyEnterCount(val) {
      if (pressBloked) return;
      console.log("??????-????:", val);
      if (!val) val = Number(val); // ?????????????????? ?? ??????????. ?????? ???????????????? ?????????? ????????????????????
      if (val == Number(selectedArticleBakeryRow.value.count_sale)) {
        console.log("???????????? ???? ????????????????, ???? ?????????? ????????????????????");
        refSelectCard.value.onClickRowMove(false); // ?????????????????? ???????????? ????????????
        //refSelectCard.value.inputBgColor = "yellow-2";
        return;
      }
      //  saleStore.debonceAddCount(async () => {
      let bgColor = ""; // refSelectCard.value?.inputBgColor; // ???????????????? ?????? ????????????
      refSelectCard.value.inputBgColor = "blue-6"; // ???????????? ?????? ????????????
      let addResult = await tableFunc.addBakeryArticleOneDay(val); // ???????????????? ??????????
      console.log("????????????????:::::", addResult);
      if (Array.isArray(addResult) && addResult.length == 1) {
        // ???????????? ????????????????????
        // await tableFunc.loadBakeryArticle(); // ???????????????????????? ??????????????.. ??????????????
        // ???? ?????????????? ?????? ???????????? ?????????? ????????
        selectedArticleBakeryRow.value.count_sale = addResult[0].countsale;
        refSelectCard.value.onClickRowMove(false); // ?????????????????? ???????????? ????????????
        refSelectCard.value.inputBgColor = bgColor; // ???????????????????? ???????? ???????????? ???? ??????????
      } else if (typeof addResult === "number" && addResult == 1) {
        // ???????? ?????????????? ????????????????
        selectedArticleBakeryRow.value.count_sale = ""; // ?????????????? ?? ????????
        refSelectCard.value.onClickRowMove(false); // ?????????????????? ???????????? ????????????
        refSelectCard.value.inputBgColor = bgColor; // ???????????????????? ???????? ???????????? ???? ??????????
      } else if (addResult === null) {
        console.log("???????????? ??????????????, ?????? ???????? ???????????? ?????? ????????????????????");
        refSelectCard.value.inputBgColor = "red-4";
      } else if (typeof addResult === "number" && addResult == 0) {
        console.log("???? ???????????????? ????????????????");
        refSelectCard.value.inputBgColor = "pink-4"; // ???? ???????????? ???????? ???????????? ?? ??????????????
      } else {
        console.log("???? ?????????????? ?????? ???????????????????? ??????-???? ", addResult);
        refSelectCard.value.inputBgColor = "pink-4"; // ???? ???????????? ???????? ???????????? ?? ??????????????
      }
      //  });
    }
    return {
      showDialog,
      onClickEdit,
      onClickNew,
      onSave,
      checkSave,
      tabModel,
      trademarkId,
      onHideArticle,
      onShowArticle,
      onAddDay,
      onKeyEnterCount,
      setRefSelectCard,
      onSaveManualData,
      onClipboardError,
      rowsManualData,
      visibleSendSave,
      //  onSaveData,
      onGetExcel,
      showDialogBuffer,
    };
  },
});
</script>
