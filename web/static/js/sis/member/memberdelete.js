async function sendDeleteMemberRequest(row) {
    let baseURL = $('#baseURL').text();
    const response = await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/member/delete",
    });
    return response
}

function deleteMember(row) {
    let baseURL = $('#baseURL').text();
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendDeleteMemberRequest(row).then(function (results) {
        alertify.success("Data member berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}