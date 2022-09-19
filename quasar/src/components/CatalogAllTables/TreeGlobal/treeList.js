import { ref } from "vue";
import TreeModel from "tree-model";
export function useTreeList() {
  function getList(refTree, key) {
    const treeArr = [
      {
        itemLabel: "Справочники",
        itemValue: "config",
        description: " настройка установка назначение размещение управление",
        // pathUrl: "/tbl/bakery",

        expandable: true,
        children: [
          {
            itemLabel: "Пекарни",
            itemValue: "bakery",
            pathUrl: "/tbl/bakery/card",
            description: "Все пекарни",
            expandable: true,
          },
          {
            itemLabel: "Контрагенты",
            itemValue: "kagents",
            pathUrl: "/tbl/kagent/card",
            description: "Контрагенты",
            disabled: false,
            expandable: true,
          },
          {
            itemLabel: "Бренды",
            itemValue: "brand",
            pathUrl: "/tbl/brand/card",
            description: "Бренды торговых сетей",
            disabled: false,
            expandable: true,
          },

          {
            itemLabel: "Продукция",
            itemValue: "products",
            pathUrl: "/tbl/products",
            description: "",
            disabled: false,
            expandable: true,
            children: [
              {
                itemLabel: "Товар",
                itemValue: "tovar",
                pathUrl: "/tbl/tovar/card",
                description: "Продукция на продажу",
                disabled: false,
                expandable: true,
              },
              {
                itemLabel: "Ассортимент",
                itemValue: "assortament",
                pathUrl: "/tbl/assortament/card",
                description: "",
                disabled: false,
                expandable: true,
              },
              {
                itemLabel: "Тип продукции",
                itemValue: "typeproduct",
                pathUrl: "/tbl/typeproduct/card",
                description: "горячие, холодные и т.п.",
                disabled: false,
                expandable: true,
              },
            ],
          },

          {
            itemLabel: "Сырьё",
            itemValue: "raw",
            pathUrl: "/tbl/raw/card",
            description: "",
            disabled: false,
            expandable: true,
            children: [
              {
                itemLabel: "Тип сырья",
                itemValue: "typeraw",
                pathUrl: "/tbl/typeraw/card",
                description: "",
                disabled: false,
                expandable: true,
              },
            ],
          },

          {
            itemLabel: "Города",
            itemValue: "city",
            pathUrl: "/tbl/city/card",
            description: "",
            disabled: false,
            expandable: true,
          },
        ],
      },
    ];
    return dataTreeCheck(treeArr);
  }
  function dataTreeCheck(rootOriginal) {
    let result = [];
    let tree = new TreeModel();
    // если много корневых узлов, перебираем все
    rootOriginal.forEach((element) => {
      let root = tree.parse(element);
      // подготовим, если найдем, список для удаления
      var nodesDrop = root.all(function (node) {
        return node.model.permiss && node.model.permiss == "Arkadii@yandex.ru";
      });
      // удаляем то что нашли
      nodesDrop.forEach(function (node) {
        console.log("node>>", node);
        node.drop();
      });
      console.log("rootres>>", root);
      result.push(root.model);
    });
    return result;
  }
  return { getList };
}
