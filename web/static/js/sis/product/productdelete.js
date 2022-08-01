async function sendDeleteProductRequest(row) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/product/delete",
    })
}

function deleteProduct(row) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendDeleteProductRequest(row).then(function () {
        alertify.success("Data barang berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}