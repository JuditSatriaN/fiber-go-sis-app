$("#upsertMember").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        buildErrorPopup(param.err);
        return
    }

    saveMember(param.data);
});

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

async function sendSaveMemberRequest(data) {
    let baseURL = $('#baseURL').text();
    const response = await axios({
        data: data,
        method: 'POST',
        url: baseURL + "svc/member/upsert",
    });
    return response
}

function saveMember(data) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendSaveMemberRequest(data).then(function () {
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