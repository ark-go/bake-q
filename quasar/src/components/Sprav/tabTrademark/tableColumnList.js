export const columns = [
  {
    name: "name",
    label: "Торговая сеть",
    align: "left",
    field: "name",
    required: true,
    sortable: true,
  },
  {
    name: "brand_name",
    label: "Бренд",
    align: "left",
    field: "brand_name",
    required: true,
  },

  {
    name: "bakery_count",
    label: "Пекарни шт.",
    align: "right",
    field: "bakery_count",
    sortable: true,
  },
  {
    name: "user_email",
    label: "Е-Mail",
    align: "left",
    field: "user_email",
    hidden: true,
  },
  {
    name: "user_date",
    label: "Дата",
    align: "left",
    field: "user_date",
    hidden: true,
  },
];
