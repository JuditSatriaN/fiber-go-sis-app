async function sendVoidTransactionRequest(row) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/void",
    });
}

function processVoidTransaction(row) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendVoidTransactionRequest(row).then(function () {
        alertify.success("Data transaksi berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}