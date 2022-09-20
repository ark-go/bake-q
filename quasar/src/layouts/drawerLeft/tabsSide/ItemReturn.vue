<template>
  <q-item clickable @click="onClick">
    <q-item-section avatar>
      <q-icon name="eva-arrowhead-left" color="blue-grey-4" />
    </q-item-section>

    <q-item-section>
      <q-item-label>Назад</q-item-label>
      <q-item-label caption>В главное меню</q-item-label>
    </q-item-section>
  </q-item>
  <q-separator />
</template>

<script>
// tabbakeryConfig
import { defineComponent, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
//import MenuSideItems from "./MenuSideItems.vue";
import { useStartPageStore, storeToRefs } from "src/stores/startPageStore";
export default defineComponent({
  name: "MenuSideBakeryConfig",
  components: {},
  props: {
    drawerPanel: {
      type: String,
    },
    returnName: {
      type: String,
      default: "main",
    },
  },
  emits: ["update:drawerPanel"],
  setup(props, { emit }) {
    //>tag="a" :href="link">
    const router = useRouter();
    const { listMenuSide } = storeToRefs(useStartPageStore());

    return {
      listMenuSide,
      router,
      pinCheckbox: ref(false),
      onClick() {
        emit("click", props.returnName);
      },
    };
  },
});
</script>
