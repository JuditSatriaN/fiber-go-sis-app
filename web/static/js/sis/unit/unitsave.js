$("#upsertUnit").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        buildErrorPopup(param.err);
        return
    }

    saveUnit(param.data);
});

function getParamValue() {
    let id = parseInt($("#modalUpsert #id").val().trim());
    let name = $("#modalUpsert #name").val().trim();

    if (name === "") {
        return {"data": null, "err": "Nama unit tidak boleh kosong !"}
    }

    return {
        "err": null,
        "data": {
            "id": id,
            "name": name,
        }
    }

}

async function sendSaveUnitRequest(data) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: data,
        method: 'POST',
        url: baseURL + "api/unit/upsert",
    })
}

function saveUnit(data) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendSaveUnitRequest(data).then(function () {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data unit berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}