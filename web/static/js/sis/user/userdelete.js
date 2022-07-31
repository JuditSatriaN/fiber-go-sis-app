async function sendDeleteUserRequest(row) {
    let baseURL = $('#baseURL').text();
    const response = await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/user/delete",
    });
    return response
}

function deleteUser(row) {
    let baseURL = $('#baseURL').text();
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendDeleteUserRequest(row).then(function (results) {
        alertify.success("Data user berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}