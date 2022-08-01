async function sendDeleteMemberRequest(row) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: row,
        method: 'POST',
        url: baseURL + "api/member/delete",
    })
}

function deleteMember(row) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendDeleteMemberRequest(row).then(function () {
        alertify.success("Data member berhasil dihapus");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}