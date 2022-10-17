<template>
  <q-btn @click="onBtnClick">Жми</q-btn>
  <q-page class="flex justify-center">
    <Ark-Pdf-Viewer src="/psql1.pdf" :page="3"></Ark-Pdf-Viewer>
    <Pdf-Card v-model:dialogShow="dialogShow"></Pdf-Card>
  </q-page>
</template>

<script>
import { defineComponent, ref } from "vue";
import PdfPage from "./PdfPage.vue";
import ArkPdfViewer from "./ArkPdfViewer.vue";
import PdfCard from "./PdfCard.vue";
//import "pdfjs-dist/build/pdf.worker.entry";
export default defineComponent({
  name: "PagePDF",
  components: { ArkPdfViewer, PdfCard },
  setup() {
    const dialogShow = ref(false);
    function onBtnClick() {
      dialogShow.value = true;
    }
    return { dialogShow, onBtnClick };
  },
});
</script>
