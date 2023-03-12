// Function to handle button insert / edit system conf
$("#upsertSystemConf").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        buildErrorPopup(param.err);
        return
    }

    saveSystemConf(param.data);
});

// Function to get parameter value
function getParamValue() {
    let id = $("#modalUpsert #id").val().trim();
    let value = $("#modalUpsert #value").val().trim();

    if (id === "") {
        return {"data": null, "err": "ID tidak boleh kosong !"}
    }

    if (value === "") {
        return {"data": null, "err": "Value tidak boleh kosong !"}
    }

    return {
        "err": null,
        "data": {
            "id": id,
            "value": value,
        }
    }

}

function saveSystemConf(data) {
    let url = $('#baseURL').text() + "api/system_conf/upsert";
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendPostRequest(url, data).then(function () {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data system config berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}