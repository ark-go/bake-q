<template>
  <q-list>
    <q-item clickable @click="router.push({ path: '/' })">
      <q-item-section avatar>
        <q-icon name="home" color="blue-grey-4" />
      </q-item-section>

      <q-item-section>
        <q-item-label>Главная</q-item-label>
        <q-item-label caption> "в начало" </q-item-label>
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
    <q-item
      v-if="userInfo.email != '1Arkadii@yandex.ru'"
      clickable
      @click="$emit('update:drawerPanel', 'treeMenu')"
    >
      <q-item-section>
        <q-item-label>?</q-item-label>
      </q-item-section>
    </q-item>
    <q-separator />
    <!-- <q-item-label header> </q-item-label> -->
    <Menu-Side-Items
      v-for="link in listMenuSide"
      :key="link.title"
      v-bind="link"
    />
  </q-list>
</template>

<script>
// tab main
import { defineComponent, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import MenuSideItems from "./MenuSideItems.vue";
import { useStartPageStore, storeToRefs } from "src/stores/startPageStore";
import { useUserStore } from "src/stores/userStore";
export default defineComponent({
  name: "MenuSide",
  components: { MenuSideItems },
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
    const { userInfo } = storeToRefs(useUserStore());

    return {
      listMenuSide,
      router,
      pinCheckbox: ref(false),
      userInfo,
    };
  },
});
</script>
