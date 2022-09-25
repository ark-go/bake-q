<template>
  <q-list>
    <Item-Return
      returnName="main"
      @click="$emit('update:drawerPanel', $event)"
    ></Item-Return>
  </q-list>
  <Document-Tree></Document-Tree>
</template>

<script>
// tabbakeryConfig
import { defineComponent, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import ItemReturn from "./ItemReturn.vue";
//import TreeGlobal from "src/components/CatalogAllTables/TreeGlobal/TreeGlobal.vue";
import DocumentTree from "components/CatalogDocuments/documentsTree/DocumentsTree.vue";
//import MenuSideItems from "./MenuSideItems.vue";
import { useStartPageStore, storeToRefs } from "src/stores/startPageStore";
export default defineComponent({
  name: "TreeDocumetnsMenu",
  components: { ItemReturn, DocumentTree },
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
