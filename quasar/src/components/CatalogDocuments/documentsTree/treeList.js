import { ref } from "vue";
import TreeModel from "tree-model";
export function useTreeList() {
  function getList(refTree, key) {
    const treeArr = [
      {
        itemLabel: "Документы",
        itemValue: "documents",
        pathUrl: "/doc/title",
        description: "",
        disabled: false,
        expandable: true,
        children: [
          {
            itemLabel: "Продажи",
            itemValue: "sale",
            pathUrl: "/doc/sale/page",
            description: "Ввод данных о продажах",
            expandable: true,
          },
          {
            itemLabel: "Установка цен",
            itemValue: "price",
            pathUrl: "/doc/price/page",
            description: "Установка цен на товар",
            expandable: true,
          },
          {
            itemLabel: "Продукция ТТК",
            itemValue: "products",
            pathUrl: "/doc/products",
            description: "Установка ТТК",
            disabled: false,
            expandable: true,
            // children: [
            //   {
            //     itemLabel: "Продукция",
            //     itemValue: "products",
            //     // pathUrl: "/doc/products",
            //     description: "",
            //     disabled: false,
            //     expandable: true,
            //     children: [
            //       {
            //         itemLabel: "Товар",
            //         itemValue: "tovar",
            //         pathUrl: "/doc/tovar/card",
            //         description: "Продукция на продажу",
            //         disabled: false,
            //         expandable: true,
            //       },
            //       {
            //         itemLabel: "Ассортимент",
            //         itemValue: "assortament",
            //         pathUrl: "/doc/assortament/card",
            //         description: "",
            //         disabled: false,
            //         expandable: true,
            //       },
            //       {
            //         itemLabel: "Тип продукции",
            //         itemValue: "typeproduct",
            //         pathUrl: "/doc/typeproduct/card",
            //         description: "горячие, холодные и т.п.",
            //         disabled: false,
            //         expandable: true,
            //       },
            //     ],
            //   },
            //   {
            //     itemLabel: "Сырьё",
            //     itemValue: "raw",
            //     pathUrl: "/doc/raw/card",
            //     description: "",
            //     disabled: false,
            //     expandable: true,
            //     children: [
            //       {
            //         itemLabel: "Тип сырья",
            //         itemValue: "typeraw",
            //         pathUrl: "/doc/typeraw/card",
            //         description: "",
            //         disabled: false,
            //         expandable: true,
            //       },
            //     ],
            //   },
            // ],
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
    // удаляем из дерева узлы, например по доступу.. ролям и тп
    // в данном случае тут пример
    rootOriginal.forEach((element) => {
      let root = tree.parse(element);
      // подготовим, если найдем, список для удаления
      var nodesDrop = root.all(function (node) {
        return node.model.permiss && node.model.permiss == "Arkadii@yandex.ru";
      });
      // удаляем то что нашли
      nodesDrop.forEach(function (node) {
        node.drop();
      });
      result.push(root.model);
    });
    let result2 = [];
    // Расставляем всем узлам parent - родительский узел
    // нам нужен parent, чтоб раскрывать узлы в дереве
    // и parentPath массив , будет создан
    result.forEach((element) => {
      let root = tree.parse(element); // в модель
      root.walk({ strategy: "pre" }, (node) => {
        // обход всех элементов
        if (node.model.children && node.model.children.length > 0) {
          // если есть дети то обходим их
          node.model.children.forEach((el, idx) => {
            // тут наш родитель
            node.model.children[idx].parent = node.model.itemValue;
            // и поместим его в список всех родителей узла
            if (node.model.parentPath) {
              node.model.children[idx].parentPath = [
                ...node.model.parentPath,
                node.model.itemValue,
              ];
            } else {
              node.model.children[idx].parentPath = [node.model.itemValue];
            }
          });
          // node.model.children.parent = node.model.itemValue;
        }
      });
      result2.push(root.model);
    });
    console.log("rootres>> parent", result2);
    return result2;
  }
  return { getList };
}
