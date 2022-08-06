// Function to handle if user click edit or delete on web
window.eventActions = {
    'click .edit': function (e, value, row, index) {
        editInventory(row);
    },
    'click .remove': function (e, value, row, index) {
        buildDeleteDataPopup("Apakah anda yakin ingin menghapus data ini?", function () {
            deleteInventory(row);
        });
    }
}

// Function to initialize inventory table
function initTable() {
    $('#table').bootstrapTable({
        locale: $('#locale').val(),
        columns: [
            [
                {
                    width: 150,
                    field: 'plu',
                    title: 'PLU',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 350,
                    title: 'Nama',
                    field: 'name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 200,
                    title: 'Unit',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'unit_name',
                },
                {
                    width: 100,
                    align: 'left',
                    title: 'Stock',
                    field: 'stock',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 200,
                    align: 'left',
                    title: 'Price',
                    field: 'price',
                    widthUnit: "px",
                    valign: 'middle',
                    formatter: priceFormatter,
                },
                {
                    width: 200,
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'member_price',
                    title: 'Member Price',
                    formatter: priceFormatter,
                },
                {
                    width: 200,
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'discount',
                    title: 'Discount',
                    formatter: priceFormatter,
                },
                {
                    title: 'Action',
                    align: 'center',
                    clickToSelect: false,
                    events: window.eventActions,
                    formatter: actionEditDeleteFormatter
                }
            ],
        ]
    });
}