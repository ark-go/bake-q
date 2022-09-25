export const columns = [
  {
    name: "name",
    label: "Регион",
    align: "left",
    field: "name",
    required: true,
    sortable: true,
  },
  // {
  //   name: "date_start_region",
  //   label: "С даты",
  //   align: "right",
  //   field: "date_start_region",
  //   required: true,
  //   sortable: true,
  // },

  {
    name: "region_manager_name",
    label: "Рег.менеджер",
    align: "left",
    field: "region_manager_name",
    sortable: true,
  },
  {
    name: "territory_count",
    label: "Территория шт.",
    align: "right",
    field: "territory_count",
    sortable: true,
  },
  {
    name: "bakery_count",
    label: "Пекарни шт.",
    align: "right",
    field: "bakery_count",
    sortable: true,
  },
];
