<template>
  <q-card flat class="ark-card-panel" style="overflow: auto; max-width: 800px">
    <div>
      <q-resize-observer @resize="(val) => (topSectionSize = val)" />
      <q-card-section>
        <div class="row items-center no-wrap">
          <div class="col">
            <div class="text-h6">{{ title }}</div>
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
                    <q-item-section @click="emit('menuClick', nameKey)">{{
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
      <q-card-section class="row q-py-none">
        <div class="col-6 q-pr-xs"><slot name="leftTop"></slot></div>
        <div class="col-6 q-pl-xs"><slot name="rightTop"></slot></div>
      </q-card-section>
    </div>
    <div>
      <q-resize-observer @resize="(val) => (bodySectionSize = val)" />
      <q-card-section class="row maxBodyHeight q-py-none">
        <div class="col-6 q-pr-xs" style="max-height: inherit">
          <slot name="leftCenter"></slot>
        </div>
        <div class="col-6 q-pl-xs" style="max-height: inherit">
          <slot name="rightCenter"></slot>
        </div>
      </q-card-section>
    </div>
    <div style="position: absolute; bottom: 0; width: 100%">
      <q-resize-observer @resize="(val) => (bottomSectionSize = val)" />
      <q-separator v-if="buttonArrProp" />
      <q-card-actions class="q-py-none">
        <slot v-if="buttonArrProp" name="buttons">
          <q-btn
            flat
            icon="eva-arrowhead-left"
            :key="item.key"
            v-for="item in buttonArrProp"
            @click="emit('buttonClick', item.key)"
          >
            ?????????? ?? ??????????????????
          </q-btn>
        </slot>
        <!-- <q-btn flat @click="clickNewRecept">??????????????</q-btn>
        <q-btn flat @click="emit('closeModal')">????????????</q-btn> -->
      </q-card-actions>
    </div>
  </q-card>
</template>

<script>
import {
  ref,
  watch,
  onMounted,
  computed,
  watchEffect,
  onUpdated,
  nextTick,
  onActivated,
} from "vue";
import { useQuasar, Screen } from "quasar";
import { dom } from "quasar";
import { useArkCardStore, storeToRefs } from "src/stores/arkCardStore";
// menuObj ???????????? ?????? ???????? ????????/??????????????  ???????????????? - ???????????????? ?? ????????
// menuClick ???????????? ?????????????? ?? ???????????? ?????????? ???? ?????????????? menuObj
export default {
  // props: ["title", "subTitle", "menuObj"],
  props: {
    title: String,
    subTitle: String,
    buttonArr: Object,
    menuObj: Object,
    maxWidth: String,
  },
  emits: ["onBodyResize"],
  setup(props, { emit }) {
    const { style, height } = dom;
    const buttonArrProp = ref([]);
    // const keyReset = ref(Date.now());
    const {
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,
      maxBodyHeight,
      maxBodyHeightCss,
    } = storeToRefs(useArkCardStore());
    // ?????????????????? ???????????? ?????????????? ??????????????
    // watch(
    //   [() => maxBodyHeight.value, () => infoSectionSize.value.height],
    //   () => {
    //     nextTick(() => {
    //       reResize();
    //     });
    //     console.log(
    //       "??????????????...",
    //       "top " + topSectionSize.value.height,
    //       "info " + infoSectionSize.value.height,
    //       "body " + bodySectionSize.value.height,
    //       "bottom " + bottomSectionSize.value.height,
    //       "maxbody " + maxBodyHeight.value,
    //       "css " + maxBodyHeightCss.value
    //     );
    //   },
    //   { immediate: true }
    // );
    // function reResize() {
    //   tableLeftHeight.value = `calc( ${maxBodyHeight.value} - ${infoSectionSize.value.height}px )`;
    //   console.log(
    //     "MMMMMMMMMMMMMMMMME",
    //     tableLeftHeight.value,
    //     infoSectionSize.value
    //   );
    // }
    // -----------------------------------^^^^^^^--------------------
    onMounted(() => {
      // keyReset.value = Date.now();
      buttonArrProp.value = props?.buttonArr;
    });
    // onActivated(() => {
    //   keyReset.value = Date.now();
    // });
    watch(props, () => {
      buttonArrProp.value = props.buttonArr;
      console.log("????????????:", buttonArrProp.value);
    });
    return {
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,
      maxBodyHeight,
      maxBodyHeightCss,
      Screen,
      buttonArrProp,
      emit,
      onClose() {
        emit("onClose");
      },
      //  keyReset,
    };
  },
};
</script>
<style lang="scss" scoped>
.ark-grid {
  display: grid;
  //  padding: 6px;
  gap: 10px;
  grid-template-columns: 1fr 1fr;
  overflow: auto; // ?????? ????????????
  max-height: inherit; // ???????????? ????????????
  .ark-grid-left {
    overflow: auto;
    max-height: inherit;
  }
  .ark-grid-right {
    overflow: auto;
    max-height: inherit;
  }
  @media (max-width: 600px) {
    grid-template-columns: 100%;
    min-width: 94vw;
    .ark-grid-left {
      overflow: unset;
    }
    .ark-grid-right {
      overflow: unset;
    }
  }
}
</style>
