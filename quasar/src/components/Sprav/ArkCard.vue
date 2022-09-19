<template>
  <q-card
    flat
    class="ark-card-panel"
    style="overflow: auto; user-select: none"
    @click.right.prevent="$emit('stop')"
  >
    <div>
      <q-resize-observer @resize="(val) => (topSectionSize = val)" />
      <q-card-section
        :class="{ 'bg-red-3': spravStore.historyOn && historyOn }"
      >
        <div class="row items-center no-wrap">
          <div class="col">
            <div
              class="text-h6 row"
              :style="{
                color: spravStore.historyOn && historyOn ? 'red' : 'black',
              }"
            >
              {{ title }}
              <q-space />
              <div class="row items-center" v-if="spravStore.historyOn">
                <div class="row items-center">История:</div>
                <q-toggle
                  v-model="historyOn"
                  :color="historyOn ? 'red' : 'green'"
                />
                <Select-Date-Ext
                  v-if="historyOn"
                  v-model:valueDate="spravStore.historyDate"
                ></Select-Date-Ext>
              </div>
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
      <q-card-section v-if="spravStore.selectedRow.name" style="padding-top: 0">
        <Tab-Button
          v-model:tabModel="tabModel"
          :selectedNode="selectedNode"
        ></Tab-Button>
      </q-card-section>
    </div>
    <div>
      <q-resize-observer @resize="(val) => (bodySectionSize = val)" />
      <q-card-section class="q-py-none" :style="{ maxHeight: maxBodyHeight }">
        <q-tab-panels
          v-model="tabModel"
          animated
          :keep-alive-max="keepAliveMax"
          :keep-alive="keepAlive"
          keep-alive-include="main"
        >
          <q-tab-panel v-if="keepAlive" name="main" style="padding: 0">
            <Splitter-Sprav :maxBodyHeight="maxBodyHeight">
              <template v-slot:before>
                <slot name="before"></slot>
              </template>
              <template v-slot:after>
                <slot name="after"></slot>
              </template>
            </Splitter-Sprav>
          </q-tab-panel>
        </q-tab-panels>
        <component
          v-if="currentTabComponent"
          :is="currentTabComponent"
          v-bind="{
            maxBodyHeight: maxBodyHeight,
            tabModel: tabModel,
          }"
        ></component>
      </q-card-section>
    </div>
    <div style="position: absolute; bottom: 0; width: 100%">
      <q-resize-observer @resize="(val) => (bottomSectionSize = val)" />
    </div>
  </q-card>
  <Page-Setup-Dialog
    v-model:menuDialogShow="menuDialogShow"
  ></Page-Setup-Dialog>
</template>

<script>
// prettier-ignore
import {ref,watch, onMounted, watchEffect, onUpdated, computed, nextTick, onBeforeUnmount, onUnmounted, defineAsyncComponent, } from "vue";
import { useRouter, useRoute, onBeforeRouteLeave } from "vue-router";
import PageSetupDialog from "./PageSetupDialog.vue";
import SplitterSprav from "./SplitterSprav.vue";
//import TabSprav from "./TabSprav.vue";
import TabButton from "./TabButton.vue";
import { useSpravStore } from "stores/spravStore";
import { useQuasar, dom } from "quasar";
import { getComponent } from "./selectComponent.js";
import SelectDateExt from "./SelectDateExt.vue";
//import { usePagesSetupStore, storeToRefs } from "src/stores/pagesSetupStore";
import { useArkCardStore, storeToRefs } from "src/stores/arkCardStore";

// menuObj объект для меню ключ/знчение  значение - показано в меню
// menuClick вернет событие с именем ключа из объекта menuObj
export default {
  // props: ["title", "subTitle", "menuObj"],
  props: {
    title: String,
    styleFn: Object,
    subTitle: String,
    buttonArr: [Object, Boolean],
    menuObj: Object,
    selectedNode: Object, // выбраный пункт Tree дерева
  },
  components: {
    SplitterSprav,
    //    TabSprav,
    TabButton,
    PageSetupDialog,
    SelectDateExt,
  },
  setup(props, { emit }) {
    const $q = useQuasar();
    const $router = useRouter();
    const route = useRoute();
    const spravStore = useSpravStore();
    const { style, height } = dom;
    const cardHeight = ref(400);
    const tabModel = ref("main"); // чтото должно быть на экране
    const splitterModel = ref(30);
    const splitHorizont = ref(false);
    const menuDialogShow = ref(false);
    const currentTabComponent = ref();
    const historyOn = ref(false);
    const keepAlive = ref(true);
    const keepAliveMax = ref(10);
    const {
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,
      maxBodyHeight,
    } = storeToRefs(useArkCardStore());

    console.log("route name", route.name);
    onBeforeUnmount(() => {
      keepAlive.value = false;
      console.log("arkCard sprav before unmount", keepAlive.value);
    });

    onBeforeRouteLeave(async (to, from, next) => {
      keepAliveMax.value = 0;
      keepAlive.value = false;
      console.log("arkCard sprav RouteLeave", keepAlive.value);

      next();
    });
    watch(
      () => tabModel.value,
      () => {
        spravStore.currentTab = tabModel.value; // занесем выбор вкладки в Penia
        console.log("Выбрана вкладка", tabModel.value, props.selectedNode);
      }
    );
    watch(
      // ловим изменение т.е. выбор дерева
      () => props.selectedNode.key,
      () => {
        currentTabComponent.value = getComponent();
      }
    );
    watch(
      () => historyOn.value,
      () => {
        if (!historyOn.value) {
          spravStore.historyDate = null;
        }
      }
    );
    const splitStyleH = {
      background: "rgb(214 214 214)",
      minWidth: "6px",
      minHeight: "50px",
      borderRadius: "3px",
    };
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
    const splitStyleW = {
      background: "rgb(214 214 214)",
      minWidth: "50px",
      minHeight: "6px",
      borderRadius: "3px",
    };
    watchEffect(() => {
      splitHorizont.value = $q.screen.width < $q.screen.height;
    });
    onMounted(() => {
      // reSizeCard();
      console.log("PageSprav size", maxBodyHeight.value);
    });
    function clickSelectButton() {
      $q.dialog({
        title: "Нихренаси, вы жмете!!",
        message: spravStore.selectedRow.name,
        cancel: true,
        persistent: true,
        ok: { label: "vvvm", color: "orange-3" }, // q-btn
        cancel: { label: "Отменить", color: "blue-5" },
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
      //  valueDate,
      //  historyDate,
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,

      historyOn,
      currentTabComponent,
      onClickMenu,
      menuDialogShow,
      tabModel,
      spravStore,
      clickSelectButton,
      maxBodyHeight,
      // maxHeigh,
      splitterModel,
      splitHorizont,
      splitStyleH,
      splitStyleW,
      cardHeight,
      emit,
      onClose() {
        emit("onClose");
      },
      keepAlive,
      keepAliveMax,
      route,
    };
  },
};
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
