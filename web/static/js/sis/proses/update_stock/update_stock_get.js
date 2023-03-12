$(function () {
    initTable();
})

function actionFormatter() {
    return [
        '<a class="save" href="javascript:void(0)" title="Save"><i class="fa fas fa-save"></i></a>',
    ].join('')
}

function stockFormatter(value, data, index) {
    return '<input data-index="' + index +'" class="input-qty-number" min="1" max="10000" style="width: 80px;vertical-align: center;horiz-align: center;" type="number" value="' + value + '">'
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
                    formatter: stockFormatter,
                },
                {
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

async function sendGetUpdateStockRequest(params) {
    let page = 1;
    let req = params.data;
    let baseURL = $('#baseURL').text();
    if (params.data["offset"] !== 0) {
        page = (params.data["offset"] / params.data["limit"]) + 1
    }

    const response = await axios({
        method: 'GET',
        url: baseURL + "api/inventory",
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

    sendGetUpdateStockRequest(params).then(function (results) {
        params.success(results);
    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}

$(document).on('change', "input.input-qty-number", function () {
    // initialize variable
    let qty_value = parseInt($(this).val());
    let minValue = parseInt($(this).attr('min'));
    let maxValue = parseInt($(this).attr('max'));
    let dataIndex = parseInt($(this).attr('data-index'));

    // check if current value is higher than minimum value
    if (qty_value < minValue) {
        $(this).val($(this).data('oldValue'));
        alert('Maaf, stock melebihi batas minimal stock. Min stock : ' + minValue.toString());
        return false
    }

    // check if current value is smaller than maximum value
    if (qty_value > maxValue) {
        $(this).val($(this).data('oldValue'));
        alert('Maaf, stock melebihi batas maksimum stock. Max stock : ' + maxValue.toString());
        return false
    }

    // set the new value
    $(this).val($(this).val().replace(/^0+/, ''));
    $('#table').bootstrapTable('updateRow', {
        index: dataIndex,
        row: {
            stock: qty_value
        }
    })
});