async function sendUpdateStockRequest(row) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/inventory/update_stock",
    })
}

function processUpdateStock(row) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendUpdateStockRequest(row).then(function () {
        alertify.success("Data stock berhasil diubah");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}