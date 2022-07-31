async function sendDeleteProductRequest(row) {
    let baseURL = $('#baseURL').text();
    const response = await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/product/delete",
    });
    return response
}

function deleteProduct(row) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendDeleteProductRequest(row).then(function (results) {
        alertify.success("Data barang berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}