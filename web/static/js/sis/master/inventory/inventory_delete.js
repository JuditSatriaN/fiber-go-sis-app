// Function to delete inventory
function deleteInventory(row) {
    let url = $('#baseURL').text() + "api/inventory/delete"
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendPostRequest(url, row).then(function () {
        alertify.success("Data inventory berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });

}