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
        :class="[
          {
            'text-weight-bold':
              prop.node.itemValue === route.params.tblRouteParam,
          },
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
import { defineComponent, onMounted, ref, watch, nextTick } from "vue";
import { useRouter, useRoute } from "vue-router";
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
    const route = useRoute();
    const expanded = ref(["config"]);
    onMounted(() => {
      treeNodes.value = treeList.getList();
    });
    watch(
      () => route.params, //tblRouteParam
      (val) => {
        if (val.tblRouteParam) {
          console.log("ПАРАМЕТР РОУТА:", val.tblRouteParam);
          nextTick(async () => {
            if (refTree.value) {
              //  refTree.value.setExpanded(val.tblRouteParam, true);
              let nodeKey = refTree.value.getNodeByKey(val.tblRouteParam);

              if (nodeKey && nodeKey.parentPath) {
                // console.log(">>>>> nodeKey:", nodeKey.parentPath);
                for (let el of nodeKey.parentPath) {
                  //nodeKey.parentPath.forEach((el) => {
                  console.log("раскроем:", el);
                  await refTree.value.setExpanded(el, true);
                }
              } else {
                // если нет, в элементе parentPath, то попробуем открыть по tblRouteParam
                let nodeKey = refTree.value.getNodeByKey("title");
                if (nodeKey) await refTree.value.setExpanded("title", true);
              }
              // console.log(
              //   "node.tree",
              //   refTree.value.getNodeByKey(val.tblRouteParam)
              // );
            }
            //   selectedKey.value = val.tblRouteParam;
          });
        }
      },
      { immediate: true }
    );
    function clickTreeNode(node) {
      let isExp = refTree.value.isExpanded(node.itemValue);
      refTree.value.setExpanded(node.itemValue, !isExp);
      if (node.pathUrl) {
        router.push({ path: node.pathUrl });
      }
    }
    // function onDblClick(event, node) {
    //   //let nodeCl = refTree.value.getNodeByKey(node.id);
    //   if (refTree.value.isExpanded(node.itemValue)) {
    //     refTree.value.setExpanded(node.itemValue, false);
    //   } else {
    //     refTree.value.setExpanded(node.itemValue, true);
    //   }
    // }
    async function rightClick(event, node) {}
    return {
      refTree,
      treeNodes,
      selectedKey,
      clickTreeNode,
      rightClick,
      //  onDblClick,
      expanded,
      route,
    };
  },
});
</script>
