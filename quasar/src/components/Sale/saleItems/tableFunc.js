import { ref, nextTick } from "vue";
import { useArkUtils } from "src/utils/arkUtils"; // const arkUtils = useArkUtils();
import { useSaleStore, storeToRefs } from "stores/saleStore";
import { useTableFunc as useTableFuncParent } from "../tableFunc.js";
import { date } from "quasar";
export function useTableFunc(tabUrl) {
  const { exportPriceExcel } = useTableFuncParent();
  const {
    articleBakeryRows,
    bakerySelectedRow,
    selectedArticleBakeryRow,
    trademarkId,
    showHiddenArticle,
    showDoobleArticle,
    selectedDateBetweenBakery,
    currentDateSale,
    checkDateSale,
  } = storeToRefs(useSaleStore());
  const dateFormat = ref("DD.MM.YYYY");
  const arkUtils = useArkUtils();

  async function loadBakeryArticle() {
    let isGetCount = false;
    let aId = selectedArticleBakeryRow.value?.id; // запомним если было выбрано
    let dateTo = dateToDateUnix(selectedDateBetweenBakery.value.to); // до этой даты из диапазона
    // if (isGetCount) {
    // если запрашиваем по одному дню
    if (checkDateSale.value && currentDateSale.value) {
      // если есть галка по дате, и выбрана дата
      dateTo = dateToDateUnix(currentDateSale.value); // берем единичную дату
      isGetCount = true;
    }
    //}
    let command = {
      cmd: isGetCount ? "loadBakeryArticleOneDay" : "loadBakeryArticle",
      bakery_id: bakerySelectedRow.value.id,
      trademark_id: trademarkId.value,
      showHiddenArticle: showHiddenArticle.value,
      showDoobleArticle: showDoobleArticle.value && !checkDateSale.value, // и если дата снята
      dateBetween: {
        from: dateToDateUnix(selectedDateBetweenBakery.value.from),
        to: dateTo,
      },
    };
    // command.historyDate = dateToDateUnix(spravStore.historyDate);
    let mess = "Артикулы";
    let url = `/api/${tabUrl}`;
    let res = await arkUtils.dataLoad(url, command, mess);
    console.log("Артикулы", url, res);
    if (res.result) {
      articleBakeryRows.value = res.result;
    } else {
      articleBakeryRows.value = [];
    }
    if (aId) {
      // если пропало то ставим пустой объект
      selectedArticleBakeryRow.value =
        articleBakeryRows.value.find((val) => val.id == aId) || {};
    }
    if (
      isGetCount &&
      !selectedArticleBakeryRow.value.id &&
      articleBakeryRows.value.length > 0
    ) {
      // таблица по одному дню, не выбрано ничего, и есть записи, - выбираем первую (а можно первую по сортировке таблицы ??)
      selectedArticleBakeryRow.value = articleBakeryRows.value[0];
    }
  }
  async function addBakeryArticleOneDay(countsale) {
    if (!checkDateSale.value || !currentDateSale.value) {
      throw new Error("Нужна дата!");
    }
    if (!selectedArticleBakeryRow.value?.id) {
      throw new Error("Нужно выбрать артикул!");
    }
    if (
      bakerySelectedRow.value.id != selectedArticleBakeryRow.value.bakery_id
    ) {
      throw new Error(
        "Внимание ошибка программы! 7735<br>пекарня не совпадает с имеющейся в базе"
      ); // как это?
    }
    if (trademarkId.value != selectedArticleBakeryRow.value.trademark_id) {
      throw new Error(
        "Внимание ошибка программы! 7735<br>Торговая сеть не совпадает с имеющейся в базе"
      ); // как это?
    }
    // если есть галка по дате, и выбрана дата
    let dateTo = dateToDateUnix(currentDateSale.value); // берем единичную дату
    let command = {
      cmd: "addBakeryArticleOneDay",
      datesale: dateTo,
      bakery_id: selectedArticleBakeryRow.value.bakery_id,
      price_value_id: selectedArticleBakeryRow.value.id,
      trademark_id: trademarkId.value,
      kagent_id: selectedArticleBakeryRow.value.kagent_id,
      kagent_own_id: selectedArticleBakeryRow.value.kagent_own_id,
      kagent_franch_id: selectedArticleBakeryRow.value.kagent_franch_id,
      filename: "",
      countsale: countsale,
      price_id: selectedArticleBakeryRow.value.price_id,
    };
    // command.historyDate = dateToDateUnix(spravStore.historyDate);
    let mess = "Артикул на запись";
    let url = `/api/${tabUrl}`;
    let res = await arkUtils.dataLoad(url, command, mess);
    console.log("Артикулы", url, res);
    if (res.result) {
      return res.result;
    } else {
      return null;
    }
  }

  async function toggleHiddenArticle(id, hidden, dateBetween) {
    let command = {
      cmd: "toggleHiddenArticle",
      bakery_id: bakerySelectedRow.value.id,
      price_value_id: id,
      hidden: hidden,
    };
    // command.historyDate = dateToDateUnix(spravStore.historyDate);
    let mess = "Артикулы видимость";
    let url = `/api/${tabUrl}`;
    let res = await arkUtils.dataLoad(url, command, mess);
    console.log("Артикулы видимость", url, res);
    if (res.result) {
      await loadBakeryArticle();
    } else {
      console.log("Артикулы видимость ошибка", res?.error);
    }
  }
  async function sendTextToTable(data, lastMulti) {
    let command = {
      cmd: "sendTextToTable",
      data: data,
      lastMulti: lastMulti, // стоит ли размножать если пустые
      bakery_id: bakerySelectedRow.value.id,
      trademark_id: trademarkId.value,
    };
    // command.historyDate = dateToDateUnix(spravStore.historyDate);
    let mess = "текст";
    let url = `/api/${tabUrl}`;
    let res = await arkUtils.dataLoad(url, command, mess);
    console.log("sendTextToTable", url, res);
    if (res.result) {
      return res.result;
    } else {
      console.log("sendTextToTable ошибка", res?.error);
      return [];
    }
  }
  // async function exportPriceExcel(dateBetween, trademark_id, bakery_id) {
  //   let command = { cmd: "exportPriceExcel" };
  //   command.dateBetween = {
  //     from: dateToDateUnix(dateBetween.from),
  //     to: dateToDateUnix(dateBetween.to),
  //   };
  //   command.excelFrom = dateBetween.from;
  //   command.excelTo = dateBetween.to;
  //   command.trademark_id = trademark_id;
  //   command.bakery_id = bakery_id;
  //   console.log("exportSaleExcel comand");
  //   //   command.historyDate = dateToDateUnix(spravStore.historyDate);
  //   let mess = "Подготовка Excel";
  //   let url = "/api/exportSaleExcel";
  //   let res = await arkUtils.dataLoad(url, command, mess);
  //   console.log("exportPriceExcel возврат", res);
  //   if (res.result) {
  //     onClickGetExcel(res.result);
  //     return true;
  //   } else {
  //     return null;
  //   }
  // }
  // async function insertBufferToSale(rows) {
  //   let command = {
  //     cmd: "insertBufferToSale",
  //     dateBetween: {
  //       from: dateToDateUnix(selectedDateBetweenBakery.value.from),
  //       to: dateToDateUnix(selectedDateBetweenBakery.value.to),
  //     },
  //     trademark_id: trademarkId.value,
  //     bakery_id: bakerySelectedRow.value.id,
  //     rows: rows,
  //   };

  //   console.log("insertBufferToSale comand");
  //   //   command.historyDate = dateToDateUnix(spravStore.historyDate);
  //   let mess = "Подготовка Excel";
  //   let url = `/api/${tabUrl}`;
  //   let res = await arkUtils.dataLoad(url, command, mess);
  //   console.log("insertBufferToSale возврат", res);
  //   if (res.result) {
  //     return res.result;
  //   } else {
  //     return null;
  //   }
  // }
  function dateToDateUnix(dat) {
    if (!dat) {
      return null;
    }
    return date.extractDate(dat, dateFormat.value);
  }
  function dateFormatDate(dat) {
    // dat - timeStamp
    return date.formatDate(dat, dateFormat.value);
  }
  // function onClickGetExcel(data) {
  //   if (data) {
  //     const blob = b64toBlob(
  //       data.bufferExcel,
  //       "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
  //     );
  //     const blobUrl = URL.createObjectURL(blob);
  //     const anchorTag = document.createElement("a");
  //     anchorTag.href = blobUrl;

  //     anchorTag.download = data.fileName; //"My PDF File.xlsx"
  //     anchorTag.click();
  //     URL.revokeObjectURL(blobUrl);
  //   }
  // }
  // function b64toBlob(b64Data, contentType = "", sliceSize = 512) {
  //   const byteCharacters = atob(b64Data);
  //   const byteArrays = [];

  //   for (let offset = 0; offset < byteCharacters.length; offset += sliceSize) {
  //     const slice = byteCharacters.slice(offset, offset + sliceSize);

  //     const byteNumbers = new Array(slice.length);
  //     for (let i = 0; i < slice.length; i++) {
  //       byteNumbers[i] = slice.charCodeAt(i);
  //     }

  //     const byteArray = new Uint8Array(byteNumbers);
  //     byteArrays.push(byteArray);
  //   }

  //   const blob = new Blob(byteArrays, { type: contentType });
  //   return blob;
  // }
  return {
    dateToDateUnix,
    dateFormatDate,
    loadBakeryArticle,
    toggleHiddenArticle,
    addBakeryArticleOneDay,
    sendTextToTable,
    exportPriceExcel,
    // insertBufferToSale,
  };
}
