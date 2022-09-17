<template>
  <q-layout view="hHh Lpr fFf">
    <q-header elevated style="height: 35px">
      <q-toolbar style="min-height: 30px">
        <div class="title-grid">
          <div class="title-grid-item-left">
            <q-btn
              v-if="userInfo.email"
              flat
              dense
              round
              icon="menu"
              aria-label="Menu"
              @click="isLeftDrawer = !isLeftDrawer"
            />

            <q-toolbar-title>
              <router-link to="/" style="color: white; text-decoration: none">
                <!-- <q-img
                  v-if="!platform.is.mobile"
                  src="/public/images/logo.png"
                  spinner-color="silver"
                  height="40px"
                  width="156px"
                  fit="fill"
                />
                <div v-else>ХиТ</div> -->
                <div>Хлеб и Тандыр</div>
              </router-link>
            </q-toolbar-title>
          </div>
          <div class="title-grid-item-center">
            <q-btn
              flat
              dense
              round
              icon="restaurant_menu"
              to="/"
              style="font-weight: bold"
            />
          </div>
          <div class="title-grid-item-right">
            <!-- <div>{{ userInfo.username || userInfo.email }}</div> -->
            <Menu-User></Menu-User>
            <!-- <q-btn
              dense
              flat
              round
              icon="menu"
              @click="rightDrawerOpen = !rightDrawerOpen"
            /> -->
          </div>
        </div>
      </q-toolbar>
    </q-header>
    <drawer-left v-if="userInfo.email" v-model="isLeftDrawer"></drawer-left>

    <!-- <q-drawer
      side="right"
      v-model="rightDrawerOpen"
      overlay
      bordered
      behavior="mobile"
      elevated
    >
      <right-items></right-items>
    </q-drawer> -->
    <q-page-container @click="onContainer" class="main-background">
      <!-- <router-view /> -->
      <router-view v-slot="{ Component }">
        <transition name="mode-fade">
          <component :is="Component" v-bind="{ styleFn: pageFnHeight }" />
        </transition>
      </router-view>
    </q-page-container>
    <q-footer
      ref="footer"
      id="footer"
      v-if="footerVisible && !platform.is.mobile"
    >
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          :color="ioSocket.onLine ? 'green' : 'red'"
          icon="star_rate"
          @click="getAllUsers"
        />
        <q-toolbar-title>
          <span style="font-size: 14px; color: rgb(104, 104, 235)">{{
            ioSocket.timeServer
          }}</span>
          <q-avatar> </q-avatar>
        </q-toolbar-title>
        <div style="font-size: 20px; color: green">
          {{ ioSocket.versionSite }}
        </div>
      </q-toolbar>
    </q-footer>
    <Test-Move></Test-Move>
  </q-layout>
  <!-- @click.self="modalLoginOpen = modalLoginOpen" -->
  <teleport to="body">
    <Login-Dialog v-model="modalLoginOpen"></Login-Dialog>
    <!-- <div
      v-if="modalLoginOpen"
      class="modal1 flex flex-center ark-popup shadow-6"
    >
      <form-login class="ark-popup-body"></form-login>
    </div> -->
  </teleport>
  <teleport to="body" v-if="pdfModal">
    <pdf-dialog></pdf-dialog>
  </teleport>
</template>
<style lang="scss" scoped>
.main-background {
  background-image: none;
}
</style>
<style lang="scss" scoped>
.title-grid {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  grid-template-rows: 1fr;
  width: 100%;
  .title-grid-item-left {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    font-weight: bold;
  }
  .title-grid-item-center {
    display: flex;
    flex-direction: row;
  }
  .title-grid-item-right {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
  }
}
</style>
<style lang="scss" scoped>
#footer {
  // background-color: rgb(161, 158, 158);
  background-color: rgb(233, 229, 229);
}
.modal1 {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  overflow: auto;
  z-index: 3500;
}

.ark-popup-body {
  background-color: white;
}
.ark-popup {
  background-color: rgba(123, 123, 124, 0.5);
}

.mode-fade-enter-active,
.mode-fade-leave-active {
  transition: opacity 0.5s ease;
}

.mode-fade-enter-from,
.mode-fade-leave-to {
  opacity: 0;
}
</style>
<script>
import { Notify } from "quasar";
//import EssentialLink from "components/EssentialLink.vue";
import DrawerLeft from "./drawerLeft/DrawerLeft.vue";
//import RightItems from "components/mainPage/rightDrawer/rightItems.vue";
//import Vue from "vue";
import {
  defineComponent,
  ref,
  onMounted,
  nextTick,
  watchEffect,
  onBeforeMount,
} from "vue";
import FormLogin from "components/Registration/FormLogin.vue";
import pdfDialog from "components/PDF/PdfDialog.vue";
//import { arkVuex } from "src/utils/arkVuex"; // const { pdfWindow } = createArkVuex();
import { useArkUtils } from "src/utils/arkUtils"; // const arkUtils = useArkUtils();
import {
  onBeforeRouteLeave,
  onBeforeRouteUpdate,
  useRoute,
  useRouter,
} from "vue-router";
//import { useTest } from "stores/test.js";
import { useUserStore } from "stores/userStore.js";
import { useMainStore } from "stores/mainStore.js";
import { usePagesSetupStore, storeToRefs } from "stores/pagesSetupStore.js";
//import { storeToRefs } from "pinia";
import { useIoSocket } from "stores/ioSocket.js";
import { useQuasar } from "quasar";
import TestMove from "./TestMove.vue";
import MenuUser from "./menuUser/menuUser.vue";
import LoginDialog from "./loginDialog/LoginDialog.vue";
//import MenuSide from "src/components/MenuSide/MenuSide.vue";
export default defineComponent({
  name: "MainLayout",

  components: {
    DrawerLeft,
    // EssentialLink,
    FormLogin,
    pdfDialog,
    //    RightItems,
    TestMove,
    MenuUser,
    LoginDialog,
    //   MenuSide,
  },

  setup() {
    const arkUtils = useArkUtils();
    const titleBrand = ref("");
    onMounted(async () => {
      titleBrand.value = platform.is.mobile ? "ХиТ" : "Хлеб и Тандыр";
    });
    const { rightDrawerOpen, isLeftDrawer, modalLoginOpen } = storeToRefs(
      useMainStore()
    );
    const footerVisible = ref(false);
    const route = useRoute();
    const router = useRouter();
    const pageSetup = usePagesSetupStore();
    const maxHeight = ref(300);
    const $q = useQuasar();
    const { notify, platform } = useQuasar();
    const ioSocket = useIoSocket();
    const { userInfo } = storeToRefs(useUserStore());
    const pdfModal = ref(false);
    // const { pdfWindow } = arkVuex();
    // const command = ref(pdfWindow.command);
    // const pdfModal = ref(pdfWindow.show);
    const essentialLinks = ref([]);
    const username = ref("");
    onMounted(async () => {
      // if (!(await checkAccess(route.path, route.meta?.title))) {
      //   // если мы заходим по URL проверяем доступность первый раз
      //   router.push("/");
      // }
      console.log("User rol", userInfo.value.roles);
      essentialLinks.value = linkList(
        userInfo.value.roles,
        userInfo.value.email
      );
    });
    // TODO Глобальные размеры окна
    function pageFnHeight(offset, height) {
      console.log("global ofset - height", offset, height);
      pageSetup.pageOffset = offset;
      pageSetup.pageHeight = height;
      pageSetup.screenWidth = $q.screen.width;
      if ($q.platform.is.mobile) pageSetup.pagePaddingY = 2;
      else pageSetup.pagePaddingY = 4;
      maxHeight.value = height - offset - pageSetup.pagePaddingY + 26 + "px"; // добавляем на глаз padding 16, не хватает милиметра убрать скрола

      console.log("pageSetup.arkCardHeight", pageSetup.arkCardHeight);

      // let heightCss = `calc(100vh - ${offset}px)`;
      let heightCss = `${height - offset - pageSetup.pagePaddingY}px`;
      //cardMain.value.width.max = $q.screen.width - 20;
      // возвращаем рабочий размер окна, прокрутка?
      return {
        minHeight: heightCss,
        maxHeight: heightCss,
        minWidth: "100px",
      };
      return {};
    }

    router.onError((err, to, from) => {
      // ловим ошибки роута или при пропадании сети
      console.error("MainLayout route error", err);
      let caption = "";
      let captionErr = "";
      if (err.toString().indexOf("fetch dynamically imported") > 0) {
        caption =
          "Ошибка загрузки модуля<br>проверьте подключение интернета.<br>перезагрузите страницу";
      } else {
        captionErr = err.toString();
      }
      if (err.toString().indexOf("Unable to preload CSS") > 0) {
        caption =
          "Ошибка загрузки модуля css<br>проверьте подключение интернета.<br>перезагрузите страницу";
      } else {
        captionErr = err.toString();
      }
      Notify.create({
        classes: "notify-error-top",
        position: "top",
        // icon: "done", // we add an icon
        spinner: false, // we reset the spinner setting so the icon can be displayed
        message: caption ? "Ошибка интернета!" : "Ошибка 1423",
        caption: caption ? caption : captionErr,
        progress: true,
        multiLine: true,
        timeout: 1000 * 30, // we will timeout it in 2.5s
        color: "deep-orange",
        html: caption ? true : false,
        // textColor: "white",
        //closeBtn: true,
        actions: [
          {
            label: "Закрыть",
            color: "yellow",
            textColor: "white",
            handler: () => {
              /* ... */
            },
          },
        ],
      });
    });

    // onBeforeMount(() => {
    //   // Вроде не используем нигде
    //   onBeforeRouteLeave(async (to, from, next) => {
    //     console.log("onBeforeRouteLeave mainLayout.vue", to);
    //     let check = await arkUtils.dataLoad("/api/accessCheck", { path: to.path }, "");
    //     if (check.error) {
    //       // выдали ошибку
    //       console.log("запрет accessCheck :", to.path, check.error);
    //       return next(false);
    //     }
    //     if (check.result) {
    //       // запрос прошел, и без ошибки
    //       return next(true);
    //     }
    //     // не поняли ничего, стоим на месте
    //     return next(false);
    //   });
    //   onBeforeRouteUpdate(async (to, from) => {
    //     console.log("onBeforeRouteUpdate mainLayout.vue", to);
    //   });
    // });
    // onBeforeRouteLeave(async (to, from, next) => {
    //   currentPage.value = "";
    //   console.log("Переход то", to);
    //   if (to.name == "registration") {
    //     return next(true);
    //   }

    //   let check = await arkUtils.dataLoad("/api/accessCheck", { path: to.path }, "");
    //   if (check.error) {
    //     console.log("запрет accessCheck :", to.path);
    //     return next(false);
    //   }

    //   return next(true);
    //   // https://router.vuejs.org/guide/advanced/composition-api.html#navigation-guards
    //   // const answer = window.confirm("Переходим ?");
    //   // // cancel the navigation and stay on the same page
    //   // if (!answer) return false;
    // });

    return {
      titleBrand,
      footerVisible,
      onContainer() {
        //  isLeftDrawer.value = false;
      },
      platform,
      rightDrawerOpen,
      isLeftDrawer,
      ioSocket,
      userInfo, //
      pdfModal,
      modalLoginOpen,
      username,
      essentialLinks,
      pageFnHeight,
      maxHeight,
      getAllUsers() {
        ioSocket.socket
          .timeout(5000)
          .emit("getAllUsers", "Хочу всех пользователей", (err, val) => {
            if (err) {
              console.log("IO emit timeout err", err);
              notify({
                color: "red-6",
                textColor: "white",
                icon: "bolt",
                message: "Нет ответа",
                //caption: "Перегрузить окно",
                position: "center",
                // avatar,
                multiLine: true,
                timeout: 0,
                closeBtn: "Ok",
                html: true,
              });
              // другая сторона не подтвердила событие в данную задержку
            } else {
              notify({
                color: "brown",
                classes: "glossy",
                textColor: "white",
                icon: "bolt",
                message: "Пользователи OnLine",
                caption: val.allUsers,
                position: "center",
                // avatar,
                multiLine: true,
                timeout: 0,
                closeBtn: "Ok",
                html: true,
              });
            }
          });
        // ioSocket.socket.emit("getAllUsers");
      },
    };
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
<style lang="scss" scoped>
:deep(.ark-card-panel) {
  width: 100%;
  max-width: 80%;
  min-height: v-bind(maxHeight);
  max-height: v-bind(maxHeight);
  @media (max-width: 1200px) {
    max-width: 90%;
  }
  @media (max-width: 1000px) {
    max-width: 99vw;
  }
}
</style>
