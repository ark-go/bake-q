<template>
  <q-dialog
    :model-value="showDialog"
    @update:model-value="(val) => emit('update:showDialog', val)"
  >
    <q-card class="my-card">
      <q-card-section class="bg-primary text-white">
        <div class="text-h6">
          {{ inputForSave.id ? "Редактирование:" : "Новый элемент" }}
        </div>
        <div class="text-subtitle2">{{ oldValue }}</div>
      </q-card-section>

      <q-input
        style="padding: 0 7px"
        v-model="inputForSave.name"
        label="Наименование:"
        color="indigo"
        autocomplete="off"
        autofocus
        clearable
        dense
      />
      <q-separator />
      <q-card-actions align="right">
        <q-btn dense flat color="primary" v-close-popup @click="onSave"
          >Сохранить</q-btn
        >
        <q-btn dense flat color="primary" v-close-popup>Отмена</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
import { ref, watch, defineComponent } from "vue";
export default defineComponent({
  name: "OneValueDialog",
  props: {
    showDialog: Boolean,
    inputArr: Object,
  },
  emits: ["update:showDialog"],
  setup(props, { emit }) {
    const cancelEnabled = ref(true);
    const inputForSave = ref({});
    watch(
      () => props.showDialog,
      (val) => {
        if (val) {
          inputForSave.value = props.inputArr;
        }
      },
      { immediate: true }
    );

    return {
      cancelEnabled,
      emit,
      onSave() {
        emit("inputForSave", inputForSave.value);
      },
      inputForSave,
    };
  },
});
</script>
