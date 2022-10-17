<template>
  <q-dialog
    :model-value="dialogShow"
    @update:model-value="$emit('update:dialogShow', $event)"
    :maximized="maximizedToggle"
    transition-show="slide-up"
    transition-hide="slide-down"
    :style="dialogHeight"
  >
    <q-card>
      <q-resize-observer @resize="sizeDialog" />
      <q-bar>
        <q-resize-observer @resize="sizeBar" />
        <q-btn flat dense round icon="menu" aria-label="Menu" />
        <div>Такая вот</div>
        <q-space />

        <q-btn
          dense
          flat
          icon="minimize"
          @click="maximizedToggle = false"
          :disable="!maximizedToggle"
        >
          <q-tooltip v-if="maximizedToggle" class="bg-white text-primary"
            >Свернуть</q-tooltip
          >
        </q-btn>
        <q-btn
          dense
          flat
          icon="crop_square"
          @click="maximizedToggle = true"
          :disable="maximizedToggle"
        >
          <q-tooltip v-if="!maximizedToggle" class="bg-white text-primary"
            >Развернуть</q-tooltip
          >
        </q-btn>
        <q-btn dense flat icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">Закрыть</q-tooltip>
        </q-btn>
      </q-bar>

      <q-card-section :style="cardBodyHeight" class="scroll">
        <Ark-Pdf-Viewer src="/psql1.pdf" :page="3"></Ark-Pdf-Viewer>
      </q-card-section>
      <div>
        <q-resize-observer @resize="sizeBottom" />

        <q-separator />

        <q-card-actions align="right">
          <q-btn flat dense label="Decline" color="primary" v-close-popup />
          <q-btn flat dense label="Accept" color="primary" v-close-popup />
        </q-card-actions>
      </div>
    </q-card>
  </q-dialog>
</template>
<script>
import {
  defineComponent,
  ref,
  computed,
  onMounted,
  inject,
  nextTick,
} from "vue";
import ArkPdfViewer from "./ArkPdfViewer.vue";
export default defineComponent({
  name: "ArkPdfDialog",
  props: {
    dialogShow: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["update:dialogShow"],
  components: { ArkPdfViewer },
  setup(props, {}) {
    const maximizedToggle = ref(false);
    const sizeBarH = ref(50);
    const sizeBottomH = ref(100);
    const sizeDialogH = ref(500);
    const bus = inject("bus");
    const dialogHeight = computed(() => {
      let MaxH = {
        maxHeight: "70vh",
      };
      if (!maximizedToggle.value) {
        MaxH = {
          maxHeight: "70vh",
        };
      }
      return MaxH;
    });
    const cardBodyHeight = computed(() => {
      let he = sizeDialogH.value - sizeBarH.value - sizeBottomH.value;
      console.log(
        "re resize",
        he,
        sizeDialogH.value,
        sizeBarH.value,
        sizeBottomH.value
      );
      return { maxHeight: he + "px", minHeight: he + "px" };
    });
    return {
      maximizedToggle,
      cardBodyHeight,
      sizeBar(val) {
        sizeBarH.value = val.height;
      },
      sizeBottom(val) {
        sizeBottomH.value = val.height;
      },
      sizeDialog(val) {
        sizeDialogH.value = val.height;
        nextTick(() => {
          console.log("sizeDialog", val.width);
          bus.emit("busResize", val.width);
        });
      },
      dialogHeight,
    };
  },
});
</script>
