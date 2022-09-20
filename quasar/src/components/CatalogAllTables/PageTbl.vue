<template>
  <q-page class="flex flex-center" style="min-width: 360px">
    <Ark-Card> Справочники </Ark-Card>
  </q-page>
</template>

<script>
import {
  defineComponent,
  defineAsyncComponent,
  ref,
  watch,
  computed,
} from "vue";
import { useRoute } from "vue-router";
//import Sprav from "components/Sprav/Sprav.vue";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
import ArkCard from "./ArkCard.vue";
export default defineComponent({
  name: "PageTbl",
  components: { ArkCard },
  props: [],
  setup(props) {},
});
</script>
