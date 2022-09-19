<template>
  <q-list>
    <Item-Return
      returnName="main"
      @click="$emit('update:drawerPanel', $event)"
    ></Item-Return>
    <!-- <q-item clickable @click="$emit('update:drawerPanel', 'main')">
      <q-item-section avatar>
        <q-icon name="eva-arrowhead-left" color="blue-grey-4" />
      </q-item-section>

      <q-item-section>
        <q-item-label>Назад</q-item-label>
        <q-item-label caption> "назад" </q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-checkbox
          v-model="pinCheckbox"
          label=""
          checked-icon="mdi-pin"
          unchecked-icon="mdi-pin-off"
          color="green"
          :keep-color="false"
        />
      </q-item-section>
    </q-item>
    <q-separator /> -->
    <!-- <q-item-label header> </q-item-label> -->
    <!-- <Menu-Side-Items
      v-for="link in listMenuSide"
      :key="link.title"
      v-bind="link"
    /> -->
  </q-list>
</template>

<script>
// tabbakeryConfig
import { defineComponent, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import ItemReturn from "./ItemReturn.vue";
//import MenuSideItems from "./MenuSideItems.vue";
import { useStartPageStore, storeToRefs } from "src/stores/startPageStore";
export default defineComponent({
  name: "MenuSideBakeryConfig",
  components: { ItemReturn },
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
    function onClick(lnk) {
      if (lnk == "api/pdf") {
        window.open(lnk, "_blank");
      } else {
        router.push({ path: lnk });
      }
    }
    return {
      onClick,
      listMenuSide,
      router,
      pinCheckbox: ref(false),
    };
  },
});
</script>
