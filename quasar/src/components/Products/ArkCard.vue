<template>
  <q-card flat class="ark-card-panel" style="overflow: auto">
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
      <q-resize-observer @resize="(val) => (bodySectionSize = val)" />

      <q-card-section class="q-py-none maxBodyHeight">
        <slot></slot>
        <!-- <q-input v-model="nameRecept" label="Наименование (выход/вес)" /> -->
      </q-card-section>
    </div>
    <div style="position: absolute; bottom: 0; width: 100%">
      <q-resize-observer @resize="(val) => (bottomSectionSize = val)" />
    </div>
  </q-card>
</template>

<script>
import { ref, watch, onMounted, computed, watchEffect, onUpdated } from "vue";
import { useQuasar, Screen } from "quasar";
import { dom } from "quasar";
import { useArkCardStore, storeToRefs } from "src/stores/arkCardStore";
// menuObj объект для меню ключ/знчение  значение - показано в меню
// menuClick вернет событие с именем ключа из объекта menuObj
export default {
  // props: ["title", "subTitle", "menuObj"],
  props: {
    title: String,
    subTitle: String,
    buttonArr: Object,
    menuObj: Object,
    maxWidth: String,
  },
  setup(props, { emit }) {
    const { style, height } = dom; //import { dom } from "quasar";
    const {
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,
      maxBodyHeight,
      maxBodyHeightCss,
    } = storeToRefs(useArkCardStore());
    return {
      maxBodyHeight,
      Screen,
      emit,
      topSectionSize,
      infoSectionSize,
      bodySectionSize,
      bottomSectionSize,
      onClose() {
        emit("onClose");
      },
    };
  },
};
</script>
