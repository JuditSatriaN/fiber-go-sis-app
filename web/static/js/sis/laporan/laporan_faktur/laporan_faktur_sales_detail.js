$(function () {
    initTblSalesDetailModal();
})

function initTblSalesDetailModal() {
    $('#table-get-sales-detail-modal').bootstrapTable({
        locale: $('#locale').val(),
        formatNoMatches: function () {
            return 'No data found';
        },
        columns: [
            [
                {
                    width: 200,
                    field: 'plu',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Product ID',
                },
                {
                    width: 300,
                    field: 'name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Product Name',
                },
                {
                    width: 150,
                    field: 'unit_name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Unit Name',
                },
                {
                    width: 150,
                    field: 'qty',
                    align: 'right',
                    widthUnit: "px",
                    title: 'Jumlah',
                    valign: 'middle',
                },
                {
                    width: 200,
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'purchase',
                    title: 'Purchase',
                    formatter: priceFormatter,
                },
                {
                    width: 200,
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'discount',
                    title: 'Discount',
                    formatter: priceFormatter,
                },
                {
                    width: 200,
                    field: 'price',
                    title: 'Price',
                    align: 'right',
                    widthUnit: "px",
                    valign: 'middle',
                    formatter: priceFormatter,
                },
            ],
        ]
    });
}

// your custom ajax request here
function ajaxGetSalesDetailModalRequest(params) {
    let baseURL = $('#baseURL').text();
    let invoice = $('#modalGetLaporanSalesDetail #invoiceModal').val();

    if (invoice === "") {
        return
    }

    $.ajax({
        'type': "GET",
        'data': {'invoice': invoice},
        'url': baseURL + "api/list_sales_detail_by_invoice",
    }).done(function (data) {
        $('#table-get-sales-detail-modal').bootstrapTable('resetView');
        params.success(data);
    }).fail(function (xhr, textStatus, errorThrown) {
        buildErrorPopup(xhr.responseJSON.message)
    });
}