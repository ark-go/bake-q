<template>
  <q-drawer
    :width="widthDrawer"
    side="left"
    :model-value="modelValue"
    @update:model-value="$emit('update:model-value', $event)"
    bordered
    class="scrollmini"
  >
    <q-tab-panels v-model="drawerPanel" animated class="rounded-borders">
      <q-tab-panel name="main" style="padding: 0">
        <Menu-Side
          v-if="userInfo.email != 'Arkadii@yandex1.ru'"
          v-model:drawerPanel="drawerPanel"
        ></Menu-Side>
        <q-list v-if="userInfo.email == 'Arkadii@yandex.ru'">
          <q-item-label header> Только мое </q-item-label>

          <EssentialLink
            v-for="link in essentialLinks"
            :key="link.title"
            v-bind="link"
          />
        </q-list>
      </q-tab-panel>
      <q-tab-panel name="bakeryConfig" style="padding: 0">
        <Menu-Side-Bakery-Config
          v-model:drawerPanel="drawerPanel"
        ></Menu-Side-Bakery-Config>
      </q-tab-panel>
      <q-tab-panel name="treeMenu" style="padding: 0">
        <Tree-Global-Menu v-model:drawerPanel="drawerPanel"></Tree-Global-Menu>
      </q-tab-panel>
      <q-tab-panel name="documentsMenu" style="padding: 0">
        <Tree-Documetns-Menu
          v-model:drawerPanel="drawerPanel"
        ></Tree-Documetns-Menu>
      </q-tab-panel>
    </q-tab-panels>
    <div
      v-touch-pan.horizontal.prevent.mouse="handlePan"
      class="u-u-u"
      style="
        background: transparent;
        width: 3px;
        top: 0px;
        position: absolute;
        right: 0;
        height: 100%;
        cursor: e-resize;
      "
    ></div>
  </q-drawer>
</template>

<script>
import { defineComponent, ref, onMounted, watch } from "vue";
import { useUserStore, storeToRefs } from "stores/userStore.js";
import EssentialLink from "components/EssentialLink.vue";
import MenuSide from "./MenuSide/MenuSide.vue";
import MenuSideBakeryConfig from "./tabsSide/MenuSideBakeryConfig.vue";
import TreeGlobalMenu from "./tabsSide/TreeGlobalMenu.vue";
import TreeDocumetnsMenu from "./tabsSide/TreeDocumetnsMenu.vue";
import { useRoute } from "vue-router";

export default defineComponent({
  name: "DrawerLeft",
  props: ["modelValue"],
  emits: ["update:modelValue"],
  components: {
    EssentialLink,
    MenuSide,
    TreeGlobalMenu,
    MenuSideBakeryConfig,
    TreeDocumetnsMenu,
  },
  setup() {
    const { userInfo } = storeToRefs(useUserStore());
    const route = useRoute();
    const essentialLinks = ref([]);
    const info = ref(null);
    const panning = ref(false);
    const widthDrawer = ref(300);
    const drawerPanel = ref("main");
    onMounted(async () => {
      console.log("User rol", userInfo.value.roles);
      essentialLinks.value = linkList(
        userInfo.value.roles,
        userInfo.value.email
      );
    });
    watch(
      () => route.path,
      (val, valOld) => {
        console.log(" переход роута", val);
        if (val.includes("/tbl/")) {
          drawerPanel.value = "treeMenu";
        } else if (val.includes("/doc/")) {
          drawerPanel.value = "documentsMenu";
        } else {
          drawerPanel.value = "main";
        }
      },
      { immediate: true }
    );
    function handlePan({ evt, ...newInfo }) {
      info.value = newInfo;
      widthDrawer.value = newInfo.position.left;
      //  console.log("panning", newInfo.position.left, info.value);
      // native Javascript event
      // console.log(evt)

      if (newInfo.isFirst) {
        panning.value = true;
      } else if (newInfo.isFinal) {
        panning.value = false;
      }
    }
    return { userInfo, essentialLinks, handlePan, widthDrawer, drawerPanel };
  },
});
function linkList(roles = [], email) {
  return [
    {
      title: "HOME",
      caption: "На начальную страницу",
      icon: "school",
      link: "/",
      visible: roles.includes("USER"),
    },
    {
      title: "Вход",
      caption: "Вход на сайт",
      icon: "school",
      link: "/login",
      visible: !email,
    },

    {
      title: "Контрагенты",
      caption: "Партнеры",
      icon: "code",
      link: "/kagent",
      visible: roles.includes("USER"),
    },
    {
      title: "Пекарни",
      caption: "Все пекарни",
      icon: "home_work",
      link: "/bakery",
      visible: roles.includes("USER"),
    },
    {
      title: "Продукция",
      caption: "Номенклатура продукции",
      icon: "home_work",
      link: "/products",
      visible: roles.includes("USER"),
    },
    {
      title: "Цены",
      caption: "Докумены для изменения цен",
      icon: "attach_money",
      link: "/docprice",
      visible: roles.includes("USER"),
    },
    {
      title: "Справочники",
      //  caption: "Технические справочники",
      icon: "code",
      link: "/spravochnik",
      visible: roles.includes("USER"),
    },
    // {
    //   title: "Номенклатура",
    //   caption: "номенклатура ввод",
    //   icon: "dynamic_form",
    //   link: "/nomencl",
    // },
    // {
    //   title: "Спецификации магазина",
    //   caption: "Спецификации магазинов",
    //   icon: "code",
    //   link: "/specstore",
    //   visible: roles.includes("ADMIN"),
    // },
    // {
    //   title: "Продажи",
    //   caption: "продажи, регистрация",
    //   icon: "dynamic_form",
    //   link: "/prodaja",
    // },
    {
      title: "Все",
      caption: " таблицы",
      icon: "dynamic_form",
      link: "/tables",
      visible: roles.includes("MODERATOR"),
    },
    {
      title: "Пользователи",
      caption: "что-то про таблицы",
      icon: "code",
      link: "/tables/users",
      visible: roles.includes("MODERATOR"),
    },
    // {
    //   title: "Пекарни",
    //   caption: "Список",
    //   icon: "code",
    //   link: "/tables/bakehouses",
    // },

    {
      title: "Фото",
      caption: "Тест фотографии",
      icon: "close",
      link: "/photo",
      visible: roles.includes("ADMIN"),
    },
    {
      title: "Загрузка XLS",
      caption: "Тест",
      icon: "close",
      link: "/xls",
      visible: roles.includes("ADMIN"),
    },
    {
      title: "Получим PDF",
      caption: "pdf",
      icon: "close",
      link: "api/pdf",
      visible: roles.includes("ADMIN"),
    },
  ];
}
</script>
