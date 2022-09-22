<template>
  <q-list>
    <q-item>
      <q-item-section>
        <q-item-label overline>Продукт</q-item-label>
        <q-item-label>{{ nameProduct }}</q-item-label>
        <q-item-label caption>{{ nameFullProduct }}</q-item-label>
      </q-item-section>

      <q-item-section side top>
        <q-item-label caption>{{ shortVidProducts }}</q-item-label>
      </q-item-section>
    </q-item>
    <q-separator v-if="productOne" />
    <q-item v-if="productOne" clickable v-ripple @click="onClickRecept">
      <q-item-section avatar>
        <q-avatar color="primary" text-color="white"> R </q-avatar>
      </q-item-section>

      <q-item-section>
        <q-item-label>Рецепт продукта</q-item-label>
        <q-item-label caption
          >Изменение, редактирование состава продукта</q-item-label
        >
      </q-item-section>
    </q-item>
    <q-item
      v-if="productOne && selectedRow?.count_ingredients > 0"
      clickable
      v-ripple
      @click="$emit('onClickPdf', selectedRow)"
    >
      <q-item-section avatar>
        <q-icon name="picture_as_pdf" size="3em" color="primary"></q-icon>
        <!-- <q-avatar color="primary" text-color="white"> R </q-avatar> -->
      </q-item-section>

      <q-item-section>
        <q-item-label>Просмотр PDF</q-item-label>
        <q-item-label caption>Просмотр/печать состава продукта</q-item-label>
      </q-item-section>
    </q-item>
  </q-list>
</template>

<script>
import { defineComponent, ref, watchEffect } from "vue";
import { arkVuex } from "src/utils/arkVuex.js";
import { useProductsStore, storeToRefs } from "stores/productsStore.js";
import { useRouter, useRoute } from "vue-router";
export default defineComponent({
  name: "ProductsAction",
  components: {},
  props: {
    // selectedNode: Object,
    onSelectedNode: Function,
  },
  emits: ["update:selectedNode", "onSelectedNode", "onClickPdf"],
  setup(props, { emit }) {
    const { selectedRow } = storeToRefs(useProductsStore());
    const nameProduct = ref("");
    const nameFullProduct = ref("");
    const shortVidProducts = ref("");
    const productOne = ref(false);
    const router = useRouter();
    const route = useRoute();
    function onClickRecept() {
      if (route.path.includes("/tbl/")) {
        router.push({
          name: "tbl",
          params: { tblRouteParam: "recept", formFactor: "none" },
        });
      } else {
        router.push({
          name: "recept",
          params: { tblRouteParam: "recept", formFactor: "none" },
        });
      }
    }
    watchEffect(() => {
      nameFullProduct.value = "";
      nameProduct.value = "";
      shortVidProducts.value = "";
      productOne.value = false;
      if (selectedRow.value?.id) {
        productOne.value = true;
        nameProduct.value = selectedRow.value.productvid_name;
        nameFullProduct.value = `${selectedRow.value.productvid_name} ${selectedRow.value.name} `;
        if (selectedRow.value.document_num)
          nameFullProduct.value += ` ТТК№ ${selectedRow.value.document_num}`;
        shortVidProducts.value = selectedRow.value.prefix;
      }
    });
    return {
      //  currentRow,
      //selectedRowsVue1x,
      selectedRow,
      nameProduct,
      nameFullProduct,
      shortVidProducts,
      productOne,
      router,
      onClickRecept,
    };
  },
});
</script>
<!-- <div v-else>
              {{ currentRowText }}
              
            </div> -->
