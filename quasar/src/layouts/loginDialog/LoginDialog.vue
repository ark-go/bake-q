<template>
  <q-dialog
    :modelValue="modelValue"
    @update:modelValue="$emit('update:modelValue', $event)"
    persistent
  >
    <q-card>
      <q-bar>
        <q-icon name="account_circle" />
        <div>Вход</div>

        <q-space />

        <q-btn dense flat icon="close" v-close-popup>
          <q-tooltip>Закрыть</q-tooltip>
        </q-btn>
      </q-bar>

      <q-card-section class="q-pa-none">
        <Form-Login-Dialog></Form-Login-Dialog>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, ref } from "vue";
import FormLoginDialog from "./FormLoginDialog.vue";

export default defineComponent({
  name: "UserMenu",
  props: ["modelValue"],
  emits: ["update:modelValue"],
  components: { FormLoginDialog },
  setup() {},
});
</script>
