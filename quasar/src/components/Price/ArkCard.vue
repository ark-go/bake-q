<template>
  <q-card
    flat
    class="ark-card-panel"
    style="overflow: auto; user-select: none"
    :style="maxHeigh"
    @click.right.prevent="$emit('stop')"
  >
    <div>
      <q-resize-observer @resize="(val) => (topSectionSize = val)" />
      <q-card-section class="q-pb-xs q-pt-xs">
        <div class="row items-center no-wrap">
          <div class="col">
            <div class="text-h6 row">
              {{ title }}
            </div>
            <div v-if="!!subTitle" class="text-subtitle2">{{ subTitle }}</div>
          </div>

          <div class="col-auto">
            <q-btn v-if="menuObj" color="grey-7" round flat icon="more_vert">
              <q-menu cover auto-close>
                <q-list>
                  <q-item
                    clickable
                    :key="nameKey"
                    v-for="(value, nameKey) in menuObj"
                  >
                    <q-item-section @click="onClickMenu(nameKey)">{{
                      value
                    }}</q-item-section>
                  </q-item>
                </q-list>
              </q-menu>
            </q-btn>
          </div>
        </div>
      </q-card-section>
    </div>
    <div>
      <q-resize-observer @resize="(val) => (infoSectionSize = val)" />
      <q-card-section class="q-py-xs">
        <Tab-Button></Tab-Button>
      </q-card-section>
    </div>
    <div>
      <q-card-section
        style="padding: 0 16px 16px 16px"
        :style="{ maxHeight: maxBodyHeight }"
      >
        <slot> </slot>
      </q-card-section>
    </div>
    <div style="position: absolute; bottom: 0; width: 100%">
      <q-resize-observer @resize="(val) => (bottomSectionSize = val)" />
      <!-- <q-separator v-if="buttonArrProp" />

        <q-card-actions>
          <slot v-if="buttonArrProp" name="buttons">
            <q-btn
              disable
              flat
              :key="item.key"
              v-for="item in buttonArrProp"
              @click="emit('buttonClick', item.key)"
            >
              {{ item.name }}
            </q-btn>
          </slot>
        </q-card-actions> -->
      <div id="tabTeleport"></div>
    </div>
    <!-- <slot name="bottomSlot"> </slot> -->
  </q-card>
  <Page-Setup-Dialog
    v-model:menuDialogShow="menuDialogShow"
  ></Page-Setup-Dialog>
</template>

<script>
// prettier-ignore
import {ref,watch, onMounted, watchEffect, onUpdated, computed, nextTick, defineAsyncComponent, defineComponent  } from "vue";
import { useRouter } from "vue-router";
import PageSetupDialog from "./PageSetupDialog.vue";
import TabButton from "./TabButton.vue";
import { useQuasar, dom } from "quasar";
import { usePriceStore, storeToRefs } from "src/stores/priceStore";
import { useArkCardStore } from "src/stores/arkCardStore";

//import { getComponent } from "./selectComponent.js"; //defineComponents
//import SelectDateExt from "./SelectDateExt.vue";

// menuObj ???????????? ?????? ???????? ????????/??????????????  ???????????????? - ???????????????? ?? ????????
// menuClick ???????????? ?????????????? ?? ???????????? ?????????? ???? ?????????????? menuObj
export default defineComponent({
  // props: ["title", "subTitle", "menuObj"],
  props: {
    title: String,
    styleFn: Object,
    subTitle: String,
    buttonArr: [Object, Boolean],
    menuObj: Object,
    pageMaxHeight: Object,
    fullScreenTr: Boolean,
  },
  components: {
    //   SplitterSprav,
    //   TabMobil,
    TabButton,
    PageSetupDialog,

    //   SelectDateExt,
  },
  setup(props, { emit }) {
    const $q = useQuasar();
    const $router = useRouter();
    const priceStore = usePriceStore();

    const { style, height } = dom;
    // const refTopSection = ref();
    // const refBodySection = ref();
    // const refInfoSection = ref();
    // const refBottomSection = ref();
    // const refAllForm = ref(null);
    // const cardHeight = ref(400);
    const buttonArrProp = ref();
    const splitterModel = ref(30);
    const splitHorizont = ref(false);
    const menuDialogShow = ref(false);
    const currentTabComponent = ref();
    const { tabModel, maxBodyHeightResize } = storeToRefs(usePriceStore());
    const {
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,
      maxBodyHeight,
    } = storeToRefs(useArkCardStore());
    priceStore.watchStore(() => {
      return watch(
        () => tabModel.value,
        () => {
          console.log("?????????????? ??????????????", tabModel.value);
        }
      );
    });
    // priceStore.watchStore(() => {
    //   return watch(
    //     () => props.fullScreenTr,
    //     () => {
    //       if (!$q.fullscreen.isActive) {
    //         $q.fullscreen
    //           .request(refAllForm.value)
    //           .then(() => {
    //             // success!
    //           })
    //           .catch((err) => {
    //             console.log("fullscreen1", err);
    //           });
    //       } else {
    //         $q.fullscreen
    //           .exit()
    //           .then(() => {
    //             // success!
    //           })
    //           .catch((err) => {
    //             // oh, no!!!
    //           });
    //       }
    //     }
    //   );
    // });
    function onClickMenu(nameKey) {
      emit("menuClick", nameKey);
      if (nameKey == "sizeForm") {
        console.log("menu");
        menuDialogShow.value = true;
        nextTick(() => {
          console.log("menu", menuDialogShow.value);
        });
      }
    }
    priceStore.watchStore(() => {
      return watchEffect(() => {
        splitHorizont.value = $q.screen.width < $q.screen.height;
      });
    });
    // priceStore.watchStore(() => {
    //   return watch(
    //     [() => priceStore.selectedRowDoc.name, () => maxBodyHeightResize.value],
    //     () => {
    //       reSizeCard();
    //     }
    //   );
    // });
    // const maxHeigh = computed(() => {
    //   return props.pageMaxHeight;
    // });
    // const maxBodyHeight = ref("");
    // function reSizeCard() {
    //   // console.log("maxHeighmaxHeighmaxHeigh", maxHeigh.value);
    //   // if (!$q.fullscreen.isActive) {
    //   try {
    //     let topBottom =
    //       height(refTopSection.value) +
    //       height(refInfoSection.value) +
    //       height(refBottomSection.value);
    //     // if (maxHeigh.value?.maxHeight) {
    //     // ???????? ???????? ????????????
    //     maxBodyHeight.value = `calc(${maxHeigh.value.maxHeight} - ${topBottom}px)`;
    //     // } else {
    //     //   // ???????? ??????, ??.??. ???????????? ??????????
    //     //   maxBodyHeight.value = `${topBottom}px`;
    //     // }
    //   } catch (e) {
    //     console.log("?????? ????????????????. ??????????????", e);
    //     maxBodyHeight.value = maxHeigh.value; //! ???? ??????????????????
    //   }
    //   // }else{
    //   // if($q.platform.is.electron){

    //   // }
    //   // }
    //   console.log("test size", maxBodyHeight.value, props.pageMaxHeight);
    // }

    // onUpdated(() => {
    //   // ?????????????? ?????????????????????? ????????
    //   reSizeCard();
    // });
    onMounted(() => {
      // reSizeCard();
      console.log("?????????? ??????", props?.buttonArr);
      //   if (buttonArrProp.value.length > 0) {
      buttonArrProp.value = props.buttonArr; // ???????????? ????????????
      //   }
      //  else {
      //   buttonArrProp.value = [{ key: "back", name: "??????????" }]; // ???????????? ????????????
      // }
    });
    priceStore.watchStore(() => {
      return watch(props, () => {
        buttonArrProp.value = props.buttonArr;
        console.log("????????????:", buttonArrProp.value);
      });
    });
    function clickSelectButton() {
      $q.dialog({
        title: "??????????????????, ???? ??????????!!",
        message: priceStore.selectedRowDoc.name,
        cancel: true,
        persistent: true,
        ok: { label: "??????", color: "orange-3" }, // q-btn
        cancel: { label: "????????????????", color: "blue-5" },
        focus: "cancel",
      })
        .onOk(() => {
          // console.log('>>>> OK')
          //
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
    }
    return {
      $router,
      currentTabComponent,
      onClickMenu,
      menuDialogShow,
      tabModel,
      priceStore,
      clickSelectButton,
      // refAllForm,
      maxBodyHeight,
      // maxHeigh,
      splitterModel,
      splitHorizont,
      //  cardHeight,
      buttonArrProp,
      emit,
      onClose() {
        emit("onClose");
      },
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,
    };
  },
});
</script>

<style lang="scss" scoped>
.maxAutoHeight {
  max-height: v-bind("maxBodyHeight");
  overflow: auto;
}
</style>
<style lang="scss" scoped>
// :deep(.maxBodyHeightAndHeight) {
//   height: v-bind(maxBodyHeight);
//   max-height: v-bind(maxBodyHeight);
//   overflow: auto;
// }
:deep(.maxBodyHeight) {
  height: v-bind(maxBodyHeight);
  max-height: v-bind(maxBodyHeight);
  overflow: auto;
}
:deep(.arkadii-sticky-header-table) {
  /* height or max-height is important */
  // height: 310px;

  // .q-table__top,
  // .q-table__bottom,
  thead tr:first-child th {
    /* bg color is important for th; just specify one */
    background-color: #f2f8fd;
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
:deep(.arkadii-table-header) {
  text-align: center;
}
</style>
