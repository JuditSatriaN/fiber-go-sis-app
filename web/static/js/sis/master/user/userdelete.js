async function sendDeleteUserRequest(row) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/user/delete",
    })
}

function deleteUser(row) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendDeleteUserRequest(row).then(function () {
        alertify.success("Data user berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}