$(function () {
    initTable();
})


function actionFormatter() {
    return [
        '<a class="edit" href="javascript:void(0)" title="Edit"><i class="fa fa-edit"></i></a>',
        '<a class="remove" href="javascript:void(0)" title="Remove"><i class="fa fa-trash"></i></a>'
    ].join('')
}

window.eventActions = {
    'click .edit': function (e, value, row, index) {
        editProduct(row);
    },
    'click .remove': function (e, value, row, index) {
        buildDeleteDataPopup("Apakah anda yakin ingin menghapus data ini?", function () {
            deleteProduct(row);
        });
    }
}

function initTable() {
    $('#table').bootstrapTable({
        fixedNumber: 2,
        fixedColumns: true,
        locale: $('#locale').val(),
        columns: [
            [
                {
                    width: 150,
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'ID Barang',
                    field: 'product_id',
                },
                {
                    width: 350,
                    field: 'name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Nama Barang',
                },
                {
                    width: 250,
                    align: 'left',
                    widthUnit: "px",
                    title: 'Barcode',
                    field: 'barcode',
                    valign: 'middle',
                },
                {
                    width: 150,
                    title: 'Stok',
                    field: 'stock',
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 100,
                    title: 'PPN',
                    field: 'ppn',
                    widthUnit: "px",
                    align: 'center',
                    valign: 'middle',
                    formatter: checkboxFormatter
                },
                {
                    width: 200,
                    title: 'Harga',
                    field: 'price',
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    formatter: priceFormatter
                },
                {
                    width: 200,
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Harga Member',
                    field: 'member_price',
                    formatter: priceFormatter
                },
                {
                    width: 200,
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Diskon',
                    field: 'discount',
                    formatter: priceFormatter
                },
                {
                    width: 200,
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

async function sendGetProductRequest(params) {
    let page = 1;
    let req = params.data;
    let baseURL = $('#baseURL').text();
    if (params.data["offset"] !== 0) {
        page = (params.data["offset"] / params.data["limit"]) + 1
    }

    const response = await axios({
        method: 'GET',
        url: baseURL + "svc/dt_products",
        data: {
            "page": page,
            "limit": req["limit"],
            "search": req["search"],
        },
    });
    return response.data
}

function ajaxRequest(params) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendGetProductRequest(params).then(function (results) {
        params.success(results);
    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}

function clearFormInput() {
    $("#modalUpsert #product_id").val("");
    $("#modalUpsert #name").val("");
    $("#modalUpsert #barcode").val("0");
    $("#modalUpsert #stock").val("0");
    $("#modalUpsert #ppn").val("Ya");
    $("#modalUpsert #price").val("0");
    $("#modalUpsert #member_price").val("0");
    $("#modalUpsert #discount").val("0");
}