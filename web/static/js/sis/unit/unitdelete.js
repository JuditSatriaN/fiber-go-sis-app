async function sendDeleteUnitRequest(row) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/unit/delete",
    })
}

function deleteUnit(row) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendDeleteUnitRequest(row).then(function () {
        alertify.success("Data unit berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}