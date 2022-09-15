<template>
  <Splitter-Body>
    <template v-slot:before>
      <Buttons-Panel
        :refTable="refTable"
        v-model:activeTab="activeTab"
      ></Buttons-Panel>
    </template>
    <template v-slot:after>
      <q-tab-panels v-model="activeTab" animated swipeable>
        <q-tab-panel name="groupBakery"
          ><table-panel-group
            tableName="regionTerritory"
            title="Территории в текщей группе"
            :parentRow="spravStore.selectedRow"
          ></table-panel-group
        ></q-tab-panel>
        <q-tab-panel name="freeBakery">
          <table-panel-free
            tableName="regionTerritory"
            title="Территории без группы"
            :parentRow="spravStore.selectedRow"
          ></table-panel-free
        ></q-tab-panel>
        <q-tab-panel name="busyBakery">
          <table-panel-busy
            tableName="regionTerritory"
            panelName="busyBakery"
            title="Территории в других группах"
            :parentRow="spravStore.selectedRow"
          ></table-panel-busy>
        </q-tab-panel>
        <q-tab-panel name="allBakery">
          <table-panel-all
            title="Все Территории"
            tableName="regionTerritory"
            :parentRow="spravStore.selectedRow"
          ></table-panel-all>
        </q-tab-panel>
      </q-tab-panels>
    </template>
  </Splitter-Body>
</template>
<style lang="scss">
.q-tab-panel {
  padding: 0;
}
</style>
<script>
import { defineComponent, ref, defineAsyncComponent, onMounted } from "vue";
import SplitterBody from "./SplitterBody.vue";
import { useSpravStore } from "stores/spravStore";

//dataLoad(url, data, logInfo = "")
export default defineComponent({
  name: "TerritoryTableRegion",
  components: {
    TablePanelAll: defineAsyncComponent(() =>
      import("./table_all/TablePanel.vue")
    ),
    TablePanelBusy: defineAsyncComponent(() =>
      import("./table_busy/TablePanel.vue")
    ),
    TablePanelFree: defineAsyncComponent(() =>
      import("./table_free/TablePanel.vue")
    ),
    TablePanelGroup: defineAsyncComponent(() =>
      import("./table_group/TablePanel.vue")
    ),
    SplitterBody,
    ButtonsPanel: defineAsyncComponent(() => import("./ButtonsPanel.vue")),
  },
  props: {
    maxBodyHeight: String,
  },
  emits: ["update:filter"],
  setup(props) {
    const spravStore = useSpravStore();
    const refTable = ref();
    const activeTab = ref("groupBakery");
    return { spravStore, refTable, activeTab };
  },
});
</script>
<style lang="scss" scoped>
.grid-main {
  display: grid;
  gap: 5px;
  grid-template-columns: 29% 70%;
  @media (max-width: 600px) {
    grid-template-columns: 100%;
  }
  background-color: transparent;
  .item-1,
  .item-2 {
    border-color: blue;
  }
}
</style>
