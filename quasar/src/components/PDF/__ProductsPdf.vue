<template>
  <q-pdfviewer
    v-model="showPdf"
    type="html5"
    :src="pdfData"
    content-class="absolute"
    title="Тест2"
  />
</template>
<script>
import { ref, onMounted, onUnmounted } from "vue";
import { useArkUtils } from "src/utils/arkUtils"; // const arkUtils = useArkUtils();
export default {
  name: "ProductsPdf", // pdfMainLoad
  props: ["cmd"],
  setup(props) {
    const arkUtils = useArkUtils();
    const pdfData = ref("");
    const showPdf = ref(true);
    onMounted(() => {
      // if (!props.cmd) {
      //   mess.value = "нет данных";
      // } else {
      //   mess.value = props.cmd;
      // }
      loadTable();
    });
    async function loadTable() {
      let mess = "Загрузка прайса";
      let res = await arkUtils.dataLoad(
        "/api/pdfMainLoad",
        { command: "priselist" },
        mess
      );
      if (res.result) {
        pdfData.value = res.result;
      } else {
        pdfData.value = null;
      }
    }
    return { pdfData, showPdf };
  },
};
</script>
