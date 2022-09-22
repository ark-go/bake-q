<template>
  <transition name="component-fade" mode="out-in">
    <keep-alive include="f-products">
      <products></products>
    </keep-alive>
  </transition>
</template>

<script>
// keep-alive include="Products"  имя берется сначала из name-свойства компонента, есть и exclude
// т.е. можно замораживать не все компоненты а только выбранные, если они динамически подключаются
import { defineComponent, ref, computed } from "vue";
import Products from "components/Products/Products.vue";
export default defineComponent({
  name: "PageProducts",
  components: {
    Products,
  },
  setup() {
    return {};
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
