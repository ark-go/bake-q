<template>
  <div style="padding: 5px">
    <q-card flat bordered class="ark-card-panel">
      <q-card-section>
        <div class="row items-center no-wrap">
          <div class="col">
            <div class="text-h6">{{ title }}</div>
            <div v-if="!!subTitle" class="text-subtitle2">{{ subTitle }}</div>
          </div>

          <div class="col-auto">
            <q-btn v-if="menuObj" color="grey-7" round flat icon="more_vert">
              <q-menu cover auto-close>
                <q-list>
                  <q-item
                    clickable
                    :key="nameKey"
                    v-for="(value, nameKey) in menuObj"
                  >
                    <q-item-section @click="emit('menuClick', nameKey)">{{
                      value
                    }}</q-item-section>
                  </q-item>
                </q-list>
              </q-menu>
            </q-btn>
          </div>
        </div>
      </q-card-section>

      <q-card-section style="padding: 0 16px 16px 16px">
        <slot></slot>
        <!-- <q-input v-model="nameRecept" label="Наименование (выход/вес)" /> -->
      </q-card-section>

      <q-separator v-if="buttonArrProp" />

      <q-card-actions>
        <slot v-if="buttonArrProp" name="buttons">
          <q-btn
            flat
            :key="item.key"
            v-for="item in buttonArrProp"
            @click="emit('buttonClick', item.key)"
          >
            {{ item.name }}
          </q-btn>
        </slot>
        <!-- <q-btn flat @click="clickNewRecept">Создать</q-btn>
        <q-btn flat @click="emit('closeModal')">Отмена</q-btn> -->
      </q-card-actions>
    </q-card>
    <!-- </q-card> -->
  </div>
</template>

<script>
import { ref, watch, onMounted } from "vue";
import { useQuasar, Screen } from "quasar";
// menuObj объект для меню ключ/знчение  значение - показано в меню
// menuClick вернет событие с именем ключа из объекта menuObj
export default {
  // props: ["title", "subTitle", "menuObj"],
  props: {
    title: String,
    subTitle: String,
    buttonArr: [Object, Boolean],
    menuObj: Object,
    maxWidth: String,
  },
  setup(props, { emit }) {
    const buttonArrProp = ref();
    onMounted(() => {
      console.log("Чтото тут", props?.buttonArr);
      //   if (buttonArrProp.value.length > 0) {
      buttonArrProp.value = props.buttonArr; // кнопки пришли
      //   }
      //  else {
      //   buttonArrProp.value = [{ key: "back", name: "Назад" }]; // кнопки забыли
      // }
    });
    watch(props, () => {
      buttonArrProp.value = props.buttonArr;
      console.log("Кнопки:", buttonArrProp.value);
    });
    return {
      Screen,
      buttonArrProp,
      emit,
      onClose() {
        emit("onClose");
      },
    };
  },
};
</script>
