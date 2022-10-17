<template>
  <slot v-if="loading" name="loading">Загрузка</slot>
  <div
    ref="refContainer"
    style="position: relative"
    :style="{ minWidth: pdfWidth + 'px' }"
  >
    <!-- <div id="viewer" class="pdfViewer"></div> -->
  </div>
</template>
<script>
import {
  defineComponent,
  ref,
  watch,
  onMounted,
  onBeforeMount,
  toRaw,
  nextTick,
  resolveDirective,
  inject,
  onUnmounted,
} from "vue";
//import { getDocument } from "pdfjs-dist";
import {
  DefaultAnnotationLayerFactory,
  DefaultTextLayerFactory,
  PDFFindController,
  PDFLinkService,
  PDFPageView,
  EventBus,
  DefaultXfaLayerFactory,
  DefaultStructTreeLayerFactory,
  PDFViewer,
} from "pdfjs-dist/web/pdf_viewer.js";
const CMAP_URL = "pdfjs-dist/cmaps/";
console.log("CMAP_URL", CMAP_URL);
//import "pdfjs-dist/build/pdf.worker.entry";

export default defineComponent({
  name: "PdfName",
  components: {},
  props: {
    // документ в Proxy (а Proxy не пропускает приватные свойства)
    pdfDocumentSt: {
      type: Object,
    },
    pdfWidth: {
      type: Number,
      default: 700,
    },
    // номер страницы
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
    // слой текста
    text: {
      type: Boolean,
      default: true,
    },
  },
  setup(props, { emit }) {
    let pdfViewer = null;
    const loading = ref(true);
    let pdfLinkService = null;
    let pdfFindController = null;
    let pdfOnePage = null; // текущая страница
    let pdfWidthScale1 = ref(null); // ширина страницы при scale = 1
    let textLayer = undefined;

    let pdfDocument = null;
    const refContainer = ref(null);
    let eventBus = null;

    function createEvents() {
      eventBus = new EventBus();
      // pagerender pagesloaded pagesloaded updateviewarea  textlayerrendered  updatetextlayermatches find findbarclose
      eventBus.on("pagerendered", function (e) {
        //emit("pagerendered", e); e.pageNumber  там есть source
        console.log("pagerendered", e);
      });
      eventBus.on("scalechange", (e) => {
        console.log("scale change", e, e.scale);
      });
      eventBus.on("all", (e) => {
        console.log("all eventBus", e);
      });
      eventBus.on("pagesinit", (e) => {
        console.log("page init", e);
      });
    }
    watch(
      () => props.page,
      (val) => {
        pdfDocument.getPage(val).then(function (pdfPage) {
          pdfViewer.setPdfPage(pdfPage);
          pdfViewer.draw();
        });
      }
    );
    watch([() => props.rotate], () => {
      console.log("Изменение rotate");
      if (pdfViewer) {
        pdfViewer.update(props.scale, props.rotate);
        pdfViewer.draw();
      }
    });
    watch(
      () => props.pdfWidth,
      async () => {
        if (pdfViewer && refContainer.value) {
          await drawScaled();
        }
      }
    );
    watch(
      () => props.pdfDocumentSt,
      async (val) => {
        if (val) {
          pdfDocument = toRaw(props.pdfDocumentSt); // получим документ из Proxy
          var pdfInfo = val.pdfInfo || val._pdfInfo; // получим кол-во страниц
          emit("numpages", pdfInfo.numPages); // отправим обратно ?
          console.log("pdfDocument", pdfInfo, pdfDocument);
          createEvents(); // построим прием сообщений
          await createPage(); // создадим страницу
        }
      },
      { immediate: true } // вместо onMounted
    );

    async function createPage() {
      // (Optionally) включить гиперссылки в файлах PDF.
      pdfLinkService = new PDFLinkService({
        eventBus: eventBus,
        externalLinkTarget: 2,
      });

      // (Optionally) включить поиск - контроллер.
      pdfFindController = new PDFFindController({
        eventBus: eventBus,
        linkService: pdfLinkService,
      });
      if (props.text) {
        textLayer = !pdfDocument.isPureXfa
          ? new DefaultTextLayerFactory()
          : null;
      }

      try {
        pdfOnePage = await pdfDocument.getPage(props.page);
        // Создание представления страницы с параметрами по умолчанию.
        pdfViewer = new PDFPageView({
          container: refContainer.value, // куда вставляем
          id: props.page, // номер страницы
          scale: 1, // масштаб
          defaultViewport: pdfOnePage.getViewport({
            scale: 1,
          }),
          eventBus: eventBus,
          textLayerFactory: textLayer,
          xfaLayerFactory: pdfDocument.isPureXfa
            ? new DefaultXfaLayerFactory()
            : null,
          structTreeLayerFactory: new DefaultStructTreeLayerFactory(),
        });

        //pdfViewer.pageLabel = "ark" + props.page;
        // Связывает реальную страницу с представлением и рисует ее.
        pdfViewer.setPdfPage(pdfOnePage);
        // Настройте функцию scrollPageIntoView для наших ссылок
        let viewer = {
          scrollPageIntoView: function (params) {
            // INFO не используем.. это для другого
            // событие при переходе по внутренним ссылкам, чтобы мы могли обрабатывать загрузку/прокрутку на нужную страницу.
            emit("link-clicked", params); // помоему это не link clicked
            console.log("scroll PageIntoView", params);
          },
        };
        pdfLinkService.setDocument(pdfDocument);
        pdfLinkService.setViewer(viewer);
        pdfFindController.setDocument(pdfDocument);
        await drawScaled();

        loading.value = false;
        emit("loading", false);
      } catch (err) {
        emit("error", err);
        loading.value = false;
        emit("loading", false);
        console.log("ERROR", err);
      }
    }

    onBeforeMount(() => {
      if (pdfViewer) {
        pdfViewer.destroy();
        pdfViewer = null;
      }
    });

    function getScaleWith() {
      if (!pdfWidthScale1.value) {
        pdfViewer.update(1, props.rotate); // ниечего не возвращает
        pdfWidthScale1.value = pdfViewer.viewport.width; // получим и сохраним размер при scale = 1
      }
      let width = refContainer.value.offsetWidth; // размер контейнера
      //  return width / pdfViewer.viewport.width;
      console.log(
        "нов расчет",
        width,
        pdfViewer,
        pdfViewer.viewport.width,
        width / pdfWidthScale1.value
      );
      return pdfOnePage.getViewport({
        scale: width / pdfWidthScale1.value, //pdfViewer.viewport.width,
      });
    }

    async function drawScaled(scaleX) {
      let scale = getScaleWith();
      if (pdfViewer) {
        pdfViewer.update(scale, props.rotate);
        console.log(
          "Размерчик height:",
          scale,
          pdfViewer,
          pdfViewer.viewport.height,
          pdfViewer.viewport.width
        );
        // надо передать высоту страницы для intersection
        let H = {};
        H[props.page] = Math.floor(pdfViewer.viewport.height);
        console.log("page height", H[props.page]);

        emit("intersectHeight", H); // там объединится в общий объект
        pdfViewer.draw();
        console.log("page-", props.page);
        //FindController нуждается в том, чтобы текстовый слой был создан в функции Draw(),(или после Draw(),?) поэтому свяжите его сейчас.
        if (props.text) {
          pdfViewer.textLayer.findController = pdfFindController;
        }
        // pdfViewer.draw();
        loading.value = false;
        emit("loading", false);
      }
    }

    return { refContainer, loading };
  },
});
</script>
<style src="pdfjs-dist/web/pdf_viewer.css"></style>
