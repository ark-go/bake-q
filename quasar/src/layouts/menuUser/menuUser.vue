<template>
  <div class="q-pa-none">
    <q-btn flat :label="userInfo.username || userInfo.email || 'Вход'">
      <q-menu>
        <div class="row no-wrap q-pa-sm">
          <div v-if="!userInfo.email" class="column">
            <!-- <div class="text-h6 q-ma-sm">Вход</div> -->
            <q-btn flat v-close-popup @click="modalLoginOpen = true"
              >Логин</q-btn
            >
            <q-separator inset class="q-mx-lg" />
            <q-btn
              flat
              v-close-popup
              @click="router.push({ path: '/registration' })"
              >Регистрация</q-btn
            >
          </div>

          <!-- <q-separator vertical inset class="q-mx-lg" /> -->

          <div v-else class="column items-center">
            <q-avatar size="72px">
              <img src="/public/images/user_avatar.png" />
            </q-avatar>

            <div class="text-subtitle1 q-mt-md q-mb-xs">
              {{ userInfo.username || userInfo.email }}
            </div>

            <q-btn
              flat
              color="red"
              label="Выход"
              push
              size="sm"
              v-close-popup
              @click="loginUnRegister"
            />
          </div>
        </div>
      </q-menu>
    </q-btn>
  </div>
</template>

<script>
import { defineComponent, ref, nextTick } from "vue";
import { useUserStore, storeToRefs } from "stores/userStore.js";
import { useMainStore } from "stores/mainStore.js";
import { useRouter } from "vue-router";
import { useArkUtils } from "src/utils/arkUtils";
export default defineComponent({
  name: "UserMenu",
  props: [],
  emits: [],
  components: {},
  setup() {
    const { userInfo } = storeToRefs(useUserStore());
    const { modalLoginOpen } = storeToRefs(useMainStore());
    const { dataLoad } = useArkUtils();
    const router = useRouter();
    async function loginUnRegister() {
      try {
        // let resp = await axios.post("/api/unLogin", {});
        let resp = await dataLoad("/api/unLogin", {});
        let data = resp.data;
        if (!data.error) {
          console.log("Выход..", data);
          userInfo = {};
        }
        nextTick(() => {
          router.push({ path: "/" });
        });
      } catch (err) {
        console.log("Выход не удался", err);
        nextTick(() => {
          router.push({ path: "/" });
        });
      }
    }
    return { userInfo, modalLoginOpen, router, loginUnRegister };
  },
});
</script>
