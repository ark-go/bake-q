<template>
  <q-page class="flex justify-center ark"> <menu-grid></menu-grid> </q-page>
</template>

<script>
import { defineComponent } from "vue";
import MenuGrid from "./MenuGrid.vue";
export default defineComponent({
  name: "PageIndex",
  components: { MenuGrid },
});
</script>
