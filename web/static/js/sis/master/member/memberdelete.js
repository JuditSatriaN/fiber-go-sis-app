// Function to handle delete member
function deleteMember(row) {
    let url = $('#baseURL').text() + "api/member/delete";
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendPostRequest(url, row).then(function () {
        alertify.success("Data member berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });

}