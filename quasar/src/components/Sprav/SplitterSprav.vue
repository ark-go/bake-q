<template>
  <q-splitter
    v-model="splitterModel"
    :limits="[20, 70]"
    :style="{ maxHeight: maxBodyHeight, height: maxBodyHeight }"
  >
    <template v-slot:separator>
      <div :style="splitStyleH"></div>
    </template>
    <template v-slot:before>
      <div :style="{ maxHeight: maxBodyHeight, overflow: 'auto' }">
        <slot name="before"></slot>
      </div>
    </template>
    <template v-slot:after>
      <div
        :style="{
          maxHeight: maxBodyHeight,
          overflow: 'auto',
          margin: '0 6px 0 6px',
        }"
      >
        <slot
          name="after"
          :style="{ maxHeight: maxBodyHeight, overflow: 'auto' }"
        ></slot>
      </div>
    </template>
  </q-splitter>
</template>

<script>
import { defineComponent, ref } from "vue";
//dataLoad(url, data, logInfo = "")
export default defineComponent({
  name: "SplitterSprav",
  components: {},
  props: {
    maxBodyHeight: String,
  },
  // emits: ["update:filter"],
  setup(props) {
    const splitterModel = ref(30);
    const splitStyleH = {
      background: "rgb(214 214 214)",
      minWidth: "6px",
      minHeight: "50px",
      borderRadius: "3px",
    };
    const splitStyleW = {
      background: "rgb(214 214 214)",
      minWidth: "50px",
      minHeight: "6px",
      borderRadius: "3px",
    };
    return {
      splitStyleH,
      splitStyleW,
      splitterModel,
    };
  },
});
</script>
