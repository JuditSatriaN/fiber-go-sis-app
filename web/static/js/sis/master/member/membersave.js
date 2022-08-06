// Function to handle button insert / edit member
$("#upsertMember").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        buildErrorPopup(param.err);
        return
    }

    saveMember(param.data);
});

// Function to get parameter value
function getParamValue() {
    let id = parseInt($("#modalUpsert #id").val().trim());
    let name = $("#modalUpsert #name").val().trim();
    let phone = $("#modalUpsert #phone").val().trim();

    if (name === "") {
        return {"data": null, "err": "Nama member tidak boleh kosong !"}
    }

    return {
        "err": null,
        "data": {
            "id": id,
            "name": name,
            "phone": phone,
        }
    }

}

function saveMember(data) {
    let url = $('#baseURL').text() + "api/member/upsert";
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendPostRequest(url, data).then(function () {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data member berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}