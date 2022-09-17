<template>
  <q-page class="flex flex-center" style="min-width: 360px">
    <transition name="component-fade" mode="out-in">
      <keep-alive include="f-products">
        <products
          v-if="isMain"
          @on-to-recept="onToRecept"
          :pageMaxHeight="pageMaxHeight"
        ></products>
        <recepts
          v-else
          @on-to-main="onToMain"
          :pageMaxHeight="pageMaxHeight"
          :heightRabZone="heightRabZone"
        ></recepts>
      </keep-alive>
    </transition>
  </q-page>
</template>

<script>
// keep-alive include="Products"  имя берется сначала из name-свойства компонента, есть и exclude
// т.е. можно замораживать не все компоненты а только выбранные, если они динамически подключаются
import { defineComponent, ref, computed } from "vue";
import Products from "components/Products/Products.vue";
import Recepts from "components/Products/Recepts/Recepts.vue";
import { arkVuex } from "src/utils/arkVuex.js";
import { usePagesSetupStore } from "src/stores/pagesSetupStore";
export default defineComponent({
  name: "PageProducts",
  components: {
    Products,
    Recepts,
  },
  setup() {
    const isMain = ref(true);
    const { selectedRowsVuex } = arkVuex();
    const pageSetup = usePagesSetupStore();
    const heightRabZone = ref(0);
    const pageMaxHeight = computed(() => {
      return {
        minHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
        maxHeight: `calc(100vh - ${pageSetup.pageOffset}px)`,
      };
    });
    // function panelFnHeight(offset, height2) {
    //   heightRabZone.value = height2 - offset;
    //   let height = `calc(100vh - ${offset}px)`;
    //   let heightChild = `calc(100vh - ${offset}px)`;
    //   pageMaxHeight.value = { minHeight: heightChild, maxHeight: heightChild };
    //   console.log("ну ка размерчик мне ", offset, pageMaxHeight.value);
    //   return { minHeight: height, maxHeight: height };
    // }
    return {
      pageMaxHeight,
      // panelFnHeight,
      heightRabZone,
      selectedRowsVuex,
      isMain,
      onToRecept() {
        isMain.value = false;
      },
      onToMain() {
        isMain.value = true;
      },
    };
  },
});
</script>
<style lang="scss" scoped>
.component-fade-enter-active,
.component-fade-leave-active {
  transition: opacity 0.3s ease;
}

.component-fade-enter-from,
.component-fade-leave-to {
  opacity: 0;
}
</style>
