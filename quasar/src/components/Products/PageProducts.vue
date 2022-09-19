<template>
  <q-page class="flex flex-center" style="min-width: 360px">
    <Page-Products-Ext></Page-Products-Ext>
  </q-page>
</template>
<script>
import { defineComponent } from "vue";
import PageProductsExt from "./PageProductsExt.vue";
export default defineComponent({
  name: "PageProductsV",
  components: { PageProductsExt },
});
</script>
