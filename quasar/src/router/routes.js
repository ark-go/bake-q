const routes = [
  // {
  //   path: "/new",
  //   component: () => import("layouts/MainLayoutMin.vue"),
  //   children: [
  //     {
  //       path: "",
  //       component: () => import("components/PageStart/PageStart.vue"),
  //       meta: { checkAccess: true, title: "Главная." },
  //     },
  //     {
  //       path: "registration2/:id?/:code?",
  //       name: "registration2",
  //       props: true,
  //       component: () => import("components/Registration/PageRegistration.vue"),
  //       meta: { title: "Регистрация" },
  //     },
  //     {
  //       path: "products",
  //       component: () => import("components/Products/PageProducts.vue"),
  //       meta: { checkAccess: true, title: "Продукция" },
  //     },
  //     {
  //       path: "bakeryconf",
  //       name: "bakeryconf",
  //       component: () => import("components/Sprav/PageSprav.vue"),
  //       meta: { checkAccess: true, title: "Конфигурация пекарен" },
  //     },
  //   ],
  // },
  {
    path: "/",
    component: () => import("layouts/MainLayoutMin.vue"),
    // component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        path: "",
        component: () => import("components/PageStart/PageStart.vue"),
        meta: { checkAccess: true, title: "Главная" },
      },
      {
        path: "test",
        component: () => import("pages/PageTest.vue"),
      },
      {
        path: "departments",
        component: () => import("components/Users/PageUsersTree.vue"),
        meta: { checkAccess: true, title: "Пользователи" },
      },
      {
        path: "registration/:id?/:code?",
        name: "registration",
        props: true,
        component: () => import("components/Registration/PageRegistration.vue"),
        meta: { checkAccess: true, title: "Регистрация" },
      },
      {
        path: "chartsold",
        component: () => import("pages/PageCharts.vue"),
      },
      { path: "tables2", component: () => import("pages/Tables.vue") },
      {
        path: "login",
        component: () => import("pages/PageLogin.vue"),
      },
      // {
      //   path: "nomencl",
      //   component: () => import("pages/PageNomencl.vue"),
      // },
      {
        path: "photo",
        component: () => import("pages/PagePhoto.vue"),
      },
      {
        path: "prodaja",
        component: () => import("pages/PageProdaja.vue"),
        meta: { checkAccess: true, title: "Продажи" },
      },
      {
        path: "specstore",
        component: () => import("pages/PageSpecStore.vue"),
      },
      // {
      //   path: "bakeryconf",
      //   name: "bakeryconf",
      //   component: () => import("components/Sprav/PageSprav.vue"),
      //   meta: { checkAccess: true, title: "Конфигурация пекарен" },
      // },
      {
        path: "spravochnik",
        component: () =>
          import(
            /* webpackChunkName: "spravochnik" */ "components/Sprav/PageSprav.vue"
          ),
        meta: {
          checkAccess: true,
          title: "Справочник",
          replace: true,
          yesReplace: false,
        },
      },
      {
        path: "price",
        component: () => import("components/Price/PagePrice.vue"),
        meta: { checkAccess: true, title: "Прайсы" },
      },
      {
        path: "sale",
        component: () => import("components/Sale/PageSale.vue"),
        meta: { checkAccess: true, title: "Продажи" },
      },
      {
        path: "kagent",
        component: () => import("components/Kagent/PageKagent.vue"),
        meta: { checkAccess: true, title: "Контрагенты" },
      },
      {
        path: "bakery",
        component: () => import("components/Bakery/PageBakery.vue"),
        meta: { checkAccess: true, title: "Пекарни" },
      },
      {
        path: "products",
        component: () => import("components/Products/PageProducts.vue"),
        meta: { checkAccess: true, title: "Продукция" },
      },
      {
        path: "docprice",
        component: () => import("components/Docprice/PageDocprice.vue"),
        meta: { checkAccess: true, title: "Прайс" },
      },
      {
        path: "bakeryttk",
        component: () => import("components/BakeryTTK/PageBakeryTTK.vue"),
        meta: { checkAccess: true, title: "Продукты" },
      },
      {
        path: "xls",
        component: () => import("pages/PageXls.vue"),
      },
      {
        path: "parsery",
        component: () => import("components/testParser/PageParser.vue"),
      },
      {
        path: "charts/:id",
        component: () => import("components/Chart/PageCharts.vue"),
      },
      // {
      //   path: "tbl",
      //   props: true,
      //   component: () => import("components/CatalogAllTables/PageTbl.vue"),
      //   meta: { checkAccess: true, title: "Конфигурация пекарен" },
      // },
      {
        path: "tbl/:tblRouteParam/:formFactor",
        props: true,
        component: () =>
          import("components/CatalogAllTables/SpravPageTables.vue"),
      },
      {
        path: "tbl/:tblRouteParam",
        props: true,
        component: () =>
          import("components/CatalogAllTables/SpravPageTables.vue"),
      },
    ],
  },
  {
    path: "/charts2/:id",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        path: "",
        component: () => import("components/Chart/PageCharts.vue"),
      },
      {
        path: "chart1",
        component: () => import("components/Chart/Chart1.vue"),
      },
      {
        path: "chart2",
        component: () => import("components/Chart/Chart1.vue"),
      },
    ],
  },

  // {
  //   path: "/price",
  //   component: () => import("layouts/MainLayout.vue"),
  //   children: [
  //     {
  //       path: "",
  //       component: () => import("components/Price/PagePrice.vue"),
  //       meta: { checkAccess: true, title: "Прайсы" },
  //     },
  //   ],
  // },
  {
    path: "/tables",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      { path: "", component: () => import("pages/PageTabs.vue") },
      { path: "users", component: () => import("pages/PageTabUsers.vue") },
      {
        path: "bakehouses",
        component: () => import("pages/PageTabBakehouses.vue"),
      },
      {
        path: "nomencl",
        component: () => import("pages/PageTabNomenclature.vue"),
      },
      {
        path: "partners",
        component: () => import("pages/PageTabPartners.vue"),
      },
      {
        path: "csv",
        component: () => import("pages/PageConvertCSV.vue"),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: "/:catchAll(.*)*",
    component: () => import("pages/Error404.vue"),
  },
];
// {
//   path: "/departments",
//   component: () => import("layouts/MainLayout.vue"),
//   children: [
//     {
//       path: "",
//       component: () => import("components/Departments/PageDepartments.vue"),
//     },
//   ],
// },
// {
//   path: "/pdf",
//   component: () => import("layouts/MainLayout.vue"),
//   children: [
//     {
//       path: "",
//       component: () => import("components/PDF/PagePdf.vue"),
//       children: [
//         {
//           path: "pricelist/:cmd",
//           props: true,
//           component: () => import("components/PDF/PriceList.vue"),
//         },
//         {
//           path: "pricelist",
//           name: "pricelist",
//           props: true,
//           component: () => import("components/PDF/PriceListMenu.vue"),
//         },
//         {
//           path: "products/:id",
//           props: true,
//           component: () => import("components/PDF/ProductsPdf.vue"),
//         },
//         {
//           path: "products",
//           name: "products",
//           props: true,
//           component: () => import("components/PDF/ProductsPdf.vue"),
//         },
//       ],
//     },
//   ],
// },
// {
//   path: "/pdf",
//   name: "pdf",
//   props: true,
//   component: () => import("pages/PagePdf.vue"),
// },
export default routes;
