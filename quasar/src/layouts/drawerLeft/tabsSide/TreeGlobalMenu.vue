<template>
  <q-list>
    <Item-Return
      returnName="main"
      @click="$emit('update:drawerPanel', $event)"
    ></Item-Return>
  </q-list>
  <Tree-Global></Tree-Global>
</template>

<script>
// tabbakeryConfig
import { defineComponent, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import ItemReturn from "./ItemReturn.vue";
import TreeGlobal from "src/components/CatalogAllTables/TreeGlobal/TreeGlobal.vue";
//import MenuSideItems from "./MenuSideItems.vue";
import { useStartPageStore, storeToRefs } from "src/stores/startPageStore";
export default defineComponent({
  name: "TreeGlobalMenu",
  components: { ItemReturn, TreeGlobal },
  props: {
    drawerPanel: {
      type: String,
    },
  },
  emits: ["update:drawerPanel"],
  setup() {
    //>tag="a" :href="link">
    const router = useRouter();
    const { listMenuSide } = storeToRefs(useStartPageStore());
    return {
      listMenuSide,
      router,
      pinCheckbox: ref(false),
    };
  },
});
</script>
