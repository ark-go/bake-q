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
        <Menu-Side v-if="userInfo.email != 'Arkadii@yandex1.ru'"></Menu-Side>
        <q-list v-if="userInfo.email == 'Arkadii@yandex.ru'">
          <q-item-label header> Только мое </q-item-label>

          <EssentialLink
            v-for="link in essentialLinks"
            :key="link.title"
            v-bind="link"
          />
        </q-list>
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
import { defineComponent, ref, onMounted } from "vue";
import { useUserStore, storeToRefs } from "stores/userStore.js";
import EssentialLink from "components/EssentialLink.vue";
import MenuSide from "src/components/MenuSide/MenuSide.vue";
export default defineComponent({
  name: "DrawerLeft",
  props: ["modelValue"],
  emits: ["update:modelValue"],
  components: {
    EssentialLink,
    MenuSide,
  },
  setup() {
    const { userInfo } = storeToRefs(useUserStore());
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
    function handlePan({ evt, ...newInfo }) {
      info.value = newInfo;
      widthDrawer.value = newInfo.position.left;
      console.log("panning", newInfo.position.left, info.value);
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
