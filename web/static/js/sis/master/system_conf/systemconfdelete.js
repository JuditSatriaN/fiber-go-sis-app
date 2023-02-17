// Function to handle delete system conf
function deleteSystemConf(row) {
    let url = $('#baseURL').text() + "api/system_conf/delete";
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendPostRequest(url, row).then(function () {
        alertify.success("Data system config berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });

}