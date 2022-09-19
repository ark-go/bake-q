<template>
  <q-tree
    ref="refTree"
    :nodes="treeNodes"
    node-key="itemValue"
    label-key="itemLabel"
    selected-color="primary"
    no-nodes-label="Нет данных"
    no-selection-unset
    v-model:selected="selectedKey"
    no-connectors
    :accordion="false"
    @click.right.stop.prevent
    v-model:expanded="expanded"
  >
    <template v-slot:default-header="prop">
      <div
        :id="prop.node.itemValue"
        @click.right="rightClick($event, prop.node)"
        @click="clickTreeNode(prop.node)"
        @dblclick="onDblClick($event, prop.node)"
        :class="[
          { 'text-weight-bold': prop.node.itemValue === selectedKey },
          'text-black',
          'text-no-wrap',
          'non-selectable',
          'cursor-pointer',
        ]"
      >
        {{ prop.node.itemLabel }}
        <q-tooltip
          :delay="3000"
          anchor="bottom right"
          self="bottom start"
          max-width="250px"
          style="max-width: 250px; font-size: 14px"
        >
          <span style="color: rgb(127, 209, 127)">{{
            prop.node.itemLabel
          }}</span>
          <br />
          {{ prop.node.description }}
        </q-tooltip>
      </div>
    </template>
  </q-tree>
</template>

<script>
import { defineComponent, onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";
import { useTreeList } from "./treeList.js";
export default defineComponent({
  name: "TreeGlobal",
  components: {},
  setup() {
    const refTree = ref();
    const treeList = useTreeList();
    const treeNodes = ref([]);
    const selectedKey = ref("");
    const router = useRouter();
    const expanded = ref(["config"]);
    onMounted(() => {
      treeNodes.value = treeList.getList();
    });
    // watch(
    //   () => selectedKey.value,
    //   (item) => {
    //     console.log("werwrwr", item);
    //     if (item.pathUrl) {
    //       router.push({ path: item.pathUrl });
    //     }
    //   }
    // );
    function clickTreeNode(node) {
      let isExp = refTree.value.isExpanded(node.itemValue);
      refTree.value.setExpanded(node.itemValue, !isExp);
      if (node.pathUrl) {
        router.push({ path: node.pathUrl });
      }
    }
    function onDblClick(event, node) {
      //let nodeCl = refTree.value.getNodeByKey(node.id);
      if (refTree.value.isExpanded(node.itemValue)) {
        refTree.value.setExpanded(node.itemValue, false);
      } else {
        refTree.value.setExpanded(node.itemValue, true);
      }
    }
    async function rightClick(event, node) {}
    return {
      refTree,
      treeNodes,
      selectedKey,
      clickTreeNode,
      rightClick,
      onDblClick,
      expanded,
    };
  },
});
</script>
