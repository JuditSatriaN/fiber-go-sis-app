$(function () {
    initTable();
})

function actionFormatter() {
    return [
        '<a class="save" href="javascript:void(0)" title="Edit"><i class="fa fas fa-trash"></i></a>',
    ].join('')
}

window.eventActions = {
    'click .save': function (e, value, row, index) {
        buildDeleteDataPopup("Apakah anda yakin ingin menghapus transaksi ini?", function () {
            processVoidTransaction(row);
        });
    }
}

function initTable() {
    $('#table').bootstrapTable({
        locale: $('#locale').val(),
        columns: [
            [
                {
                    width: 250,
                    align: 'left',
                    widthUnit: "px",
                    field: 'invoice',
                    title: 'Invoice',
                    valign: 'middle',
                },
                {
                    width: 200,
                    align: 'left',
                    widthUnit: "px",
                    title: 'User ID',
                    field: 'user_id',
                    valign: 'middle',
                },
                {
                    width: 210,
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'total_purchase',
                    title: 'Total Beli',
                    formatter: priceFormatter,
                },
                {
                    width: 210,
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'total_price',
                    title: 'Total Harga',
                    formatter: priceFormatter,
                },
                {
                    width: 210,
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'total_discount',
                    title: 'Total Discount',
                    formatter: priceFormatter,
                },
                {
                    width: 210,
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'total_tax',
                    title: 'Total PPN',
                    formatter: priceFormatter,
                },
                {
                    width: 210,
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'total_pay',
                    title: 'Total Bayar',
                    formatter: priceFormatter,
                },
                {
                    width: 150,
                    title: 'Action',
                    align: 'center',
                    clickToSelect: false,
                    formatter: actionFormatter,
                    events: window.eventActions,
                }
            ],
        ]
    });
}

async function sendGetVoidTransactionRequest(params) {
    let page = 1;
    let req = params.data;
    let baseURL = $('#baseURL').text();
    if (params.data["offset"] !== 0) {
        page = (params.data["offset"] / params.data["limit"]) + 1
    }

    const response = await axios({
        method: 'GET',
        url: baseURL + "api/list_sales_head",
        params: {
            "page": page,
            "limit": req["limit"],
            "search": req["search"],
        },
    });
    return response.data
}

function ajaxRequest(params) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendGetVoidTransactionRequest(params).then(function (results) {
        params.success(results);
    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}