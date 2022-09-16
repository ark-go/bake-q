<template>
  <q-drawer
    side="left"
    :model-value="modelValue"
    @update:model-value="$emit('update:model-value', $event)"
    bordered
    class="scrollmini"
  >
    <Menu-Side v-if="userInfo.email != 'Arkadii@yandex1.ru'"></Menu-Side>
    <q-list v-if="userInfo.email == 'Arkadii@yandex.ru'">
      <q-item-label header> Только мое </q-item-label>

      <EssentialLink
        v-for="link in essentialLinks"
        :key="link.title"
        v-bind="link"
      />
    </q-list>
    <div
      class="u-u-u"
      style="
        background: teal;
        width: 2px;
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
export default defineComponent({
  name: "DrawerLeft",
  props: ["modelValue"],
  emits: ["update:modelValue"],
});
</script>
