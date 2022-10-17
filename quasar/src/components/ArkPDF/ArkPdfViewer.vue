<style lang="scss" scoped>
.max-width {
  width: 100%;
  // div {
  //   max-width: 100%;
  // }
}
</style>
<template>
  <div v-if="numPages > 0" class="max-width">
    <q-resize-observer @resize="resizeMainDiv" :debounce="300" />

    <q-intersection
      v-for="i in numPages"
      :key="i"
      :style="{ height: intersectHeight[i] + 'px' }"
    >
      <div
        v-if="i > 1"
        :id="'page-number-' + i"
        style="
          width: 100%;
          height: 15px;
          border-bottom: 1px solid gray;
          border-bottom-style: dashed;
          text-align: center;
          display: inline-block;
        "
      >
        <span
          style="
            font-size: 20px;
            background-color: #f3f5f6;
            padding: 0 10px;
            border-radius: 15px;
          "
        >
          {{ i }} стр.
        </span>
      </div>
      <pdf-page
        ref="refPdfPage"
        :src="src"
        :key="i"
        :id="i"
        :page="i"
        :scale="scale"
        style="margin: 0"
        @link-clicked="onPdfLink"
        :pdfDocumentSt="pdfDocument"
        :pdfWidth="pdfWidth"
        @intersectHeight="
          (val) => (intersectHeight = { ...intersectHeight, ...val })
        "
      >
        <template v-slot:loading> Загружаем... </template>
      </pdf-page>
    </q-intersection>
  </div>
</template>

<script>
import "pdfjs-dist/build/pdf.worker.entry"; // загружаем GlobalWorkerOptions.workerSrc
import { getDocument } from "pdfjs-dist";
import {
  defineComponent,
  ref,
  watch,
  onMounted,
  onUnmounted,
  onBeforeMount,
  nextTick,
  inject,
} from "vue";
import { debounce } from "quasar";
import PdfPage from "./PdfPage.vue";
//const bus = inject("bus");
export default defineComponent({
  name: "ArkPdfViewer",
  props: {
    src: {
      type: [String, Object, Promise],
      default: "",
    },
    page: {
      type: Number,
      default: 1,
    },
    rotate: {
      type: Number,
      default: 0,
    },
    scale: {
      type: [Number, String],
      default: "",
    },
  },
  components: { PdfPage },
  setup(props) {
    const refPdfPage = ref(null); //array for компоненты
    const pdfdata = ref("");
    const numPages = ref(0);
    const pdfDocument = ref();
    const pdfWidth = ref(100);
    const bus = inject("bus");
    function busResizes(val) {
      pdfWidth.value = val;
    }
    onMounted(() => {
      bus.on("busResize", busResizes);
    });
    onUnmounted(() => {
      bus.off("busResize", busResizes);
    });
    // const scale = ref(1);
    function onPdfLink() {}
    const intersectHeight = ref({ 1: 2000 });
    onBeforeMount(async () => {
      intersectHeight.value = {};
      let pdfDoc = await getDocument({ url: props.src, enableXfa: true })
        .promise;
      // loadingTask.onPassword = options.onPassword; ???
      // loadingTask.onProgress = options.onProgress;  ????
      // let pdfDoc = await loadingTask; // Proxy объекты не пропускают приватные свойства _ххх поэтому гднто не используем Proxy обекты
      var pdfInfo = pdfDoc.pdfInfo || pdfDoc._pdfInfo; // не понятно что там будет
      numPages.value = pdfInfo.numPages; // кол-во страниц
      pdfDocument.value = pdfDoc; // засунем в Proxy объект и передадим в компоненты там для работы используем toRaw
      for (let i = 1; i <= numPages.value; i++) {
        intersectHeight.value[i] = 2000;
      }
      console.log("pdfDocument.value", pdfDocument.value);
    });
    function resizeMainDiv(val) {
      console.log("resize но ", val);
      nextTick(() => {
        pdfWidth.value = val.width;
      });
    }
    return {
      refPdfPage,
      pdfdata,
      pdfDocument,
      numPages,
      onPdfLink,
      intersectHeight,
      pdfWidth,
      resizeMainDiv,
    };
  },
});
</script>
