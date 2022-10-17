import { ref } from "vue";
import TreeModel from "tree-model";
export function useTreeList() {
  function getList(refTree, key) {
    const treeArr = [
      {
        itemLabel: "Пекарни",
        itemValue: "bakery",
        pathUrl: "/tbl/bakery/card",
        description: "Все пекарни",
        expandable: true,
      },
      {
        itemLabel: "Продукция",
        itemValue: "productsX",
        // pathUrl: "/tbl/products",
        description: "",
        disabled: false,
        expandable: true,
        children: [
          {
            itemLabel: "Продукты",
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
      // -----------------------
      // {
      //   itemLabel: "Справочники",
      //   itemValue: "title",
      //   description: " настройка установка назначение размещение управление",
      //   pathUrl: "/tbl/title",
      //   expandable: true,
      //   children: [
      // {
      //   itemLabel: "Отчеты",
      //   itemValue: "otchets",
      //   // pathUrl: "/tbl/title",
      //   description: "",
      //   disabled: false,
      //   expandable: true,
      //   children: [
      //     {
      //       itemLabel: "Отчетов нет",
      //       itemValue: "otchet1",
      //       // pathUrl: "/tbl/sale/page",
      //       description: "",
      //       expandable: true,
      //     },
      //   ],
      // },
      // {
      //   itemLabel: "Документы",
      //   itemValue: "documents",
      //   pathUrl: "/tbl/title",
      //   description: "",
      //   disabled: false,
      //   expandable: true,
      //   children: [
      //     {
      //       itemLabel: "Продажи",
      //       itemValue: "sale",
      //       pathUrl: "/tbl/sale/page",
      //       description: "Ввод данных о продажах",
      //       expandable: true,
      //     },
      //     {
      //       itemLabel: "Установка цен",
      //       itemValue: "price",
      //       pathUrl: "/tbl/price/page",
      //       description: "Установка цен на товар",
      //       expandable: true,
      //     },
      //     {
      //       itemLabel: "Рецептура ТТК",
      //       itemValue: "products",
      //       pathUrl: "/tbl/products",
      //       description: "Установка ТТК",
      //       disabled: false,
      //       expandable: true,
      //       children: [
      //         {
      //           itemLabel: "Продукция",
      //           itemValue: "productsX",
      //           // pathUrl: "/tbl/products",
      //           description: "",
      //           disabled: false,
      //           expandable: true,
      //           children: [
      //             {
      //               itemLabel: "Продукты",
      //               itemValue: "tovar",
      //               pathUrl: "/tbl/tovar/card",
      //               description: "Продукция на продажу",
      //               disabled: false,
      //               expandable: true,
      //             },
      //             {
      //               itemLabel: "Ассортимент",
      //               itemValue: "assortament",
      //               pathUrl: "/tbl/assortament/card",
      //               description: "",
      //               disabled: false,
      //               expandable: true,
      //             },
      //             {
      //               itemLabel: "Тип продукции",
      //               itemValue: "typeproduct",
      //               pathUrl: "/tbl/typeproduct/card",
      //               description: "горячие, холодные и т.п.",
      //               disabled: false,
      //               expandable: true,
      //             },
      //           ],
      //         },
      //         {
      //           itemLabel: "Сырьё",
      //           itemValue: "raw",
      //           pathUrl: "/tbl/raw/card",
      //           description: "",
      //           disabled: false,
      //           expandable: true,
      //           children: [
      //             {
      //               itemLabel: "Тип сырья",
      //               itemValue: "typeraw",
      //               pathUrl: "/tbl/typeraw/card",
      //               description: "",
      //               disabled: false,
      //               expandable: true,
      //             },
      //           ],
      //         },
      //       ],
      //     },
      //   ],
      // },

      // {
      //   itemLabel: "Продукция",
      //   itemValue: "products",
      //   pathUrl: "/tbl/products",
      //   description: "",
      //   disabled: false,
      //   expandable: true,
      //   children: [
      //     {
      //       itemLabel: "Товар",
      //       itemValue: "tovar",
      //       pathUrl: "/tbl/tovar/card",
      //       description: "Продукция на продажу",
      //       disabled: false,
      //       expandable: true,
      //     },
      //     {
      //       itemLabel: "Ассортимент",
      //       itemValue: "assortament",
      //       pathUrl: "/tbl/assortament/card",
      //       description: "",
      //       disabled: false,
      //       expandable: true,
      //     },
      //     {
      //       itemLabel: "Тип продукции",
      //       itemValue: "typeproduct",
      //       pathUrl: "/tbl/typeproduct/card",
      //       description: "горячие, холодные и т.п.",
      //       disabled: false,
      //       expandable: true,
      //     },
      //   ],
      // },

      {
        itemLabel: "Контрагенты",
        itemValue: "kagent",
        pathUrl: "/tbl/kagent/card",
        description: "Контрагенты",
        disabled: false,
        expandable: true,
        children: [
          {
            itemLabel: "Виды контрагентов",
            itemValue: "kagentvid",
            pathUrl: "/tbl/kagentvid/card",
            description: "",
            disabled: false,
            expandable: true,
          },
          {
            itemLabel: "Виды регистрации контрагентов",
            itemValue: "kagentreg",
            pathUrl: "/tbl/kagentreg/card",
            description: "",
            disabled: false,
            expandable: true,
          },
          {
            itemLabel: "Группы контрагентов",
            itemValue: "kagentgroups",
            pathUrl: "/tbl/kagentgroups/card",
            description: "",
            disabled: false,
            expandable: true,
          },
        ],
      },
      {
        itemLabel: "Торговые сети",
        itemValue: "trademark",
        pathUrl: "/tbl/trademark/card",
        description: "",
        disabled: false,
        expandable: true,
        children: [
          {
            itemLabel: "Бренды",
            itemValue: "brand",
            pathUrl: "/tbl/brand/card",
            description: "Бренды торговых сетей",
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
      //   ], //TODO
      // },
      // {
      //   itemLabel: "Пользователи",
      //   itemValue: "department",
      //   pathUrl: "/tbl/department/page",
      //   description: "",
      //   disabled: false,
      //   expandable: true,
      // },

      // {
      //   itemLabel: "Конфигурация пекарен",
      //   itemValue: "bakeryconfig",
      //   pathUrl: "/tbl/bakeryconfig/page",
      //   description: "",
      //   disabled: false,
      //   expandable: true,
      // },
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
