$(function () {
    initTable();
})

function actionFormatter() {
    return [
        '<a class="save" href="javascript:void(0)" title="Save"><i class="fa fas fa-save"></i></a>',
    ].join('')
}

window.eventActions = {
    'click .save': function (e, value, row, index) {
        buildDeleteDataPopup("Apakah anda yakin ingin mengubah stock product ini?", function () {
            processUpdateStock(row);
        });
    }
}

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
                    field: 'unit_name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 100,
                    title: 'Stock',
                    field: 'stock',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 200,
                    title: 'Price',
                    field: 'price',
                    align: 'left',
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
                // {
                //     title: 'Action',
                //     align: 'center',
                //     clickToSelect: false,
                //     formatter: actionFormatter,
                //     events: window.eventActions,
                // }
            ],
        ]
    });
}

async function sendGetInventoryRequest() {
    let baseURL = $('#baseURL').text();
    const response = await axios.get(baseURL + "api/inventory");
    return response.data
}

function ajaxRequest(params) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendGetInventoryRequest().then(function (results) {
        params.success(results);
    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}