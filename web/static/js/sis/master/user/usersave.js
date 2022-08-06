$("#upsertUser").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        buildErrorPopup(param.err);
        return
    }

    saveUser(param.data);
});

function getParamValue() {
    let user_id = $("#modalUpsert #user_id").val().trim()
    let user_name = $("#modalUpsert #user_name").val().trim()
    let full_name = $("#modalUpsert #full_name").val().trim()
    let password = $("#modalUpsert #password").val().trim()
    let is_admin = $("#modalUpsert #is_admin").val().trim() === "Ya"

    if (user_name === "") {
        return {"data": null, "err": "Nama user tidak boleh kosong !"}
    }
    return {
        "err": null,
        "data": {
            "user_id": user_id,
            "user_name": user_name,
            "full_name": full_name,
            "password": password,
            "is_admin": is_admin,
        }
    }

}

async function sendSaveUserRequest(data) {
    let baseURL = $('#baseURL').text();
    return await axios({
        data: data,
        method: 'POST',
        url: baseURL + "api/user/upsert",
    })
}

function saveUser(data) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendSaveUserRequest(data).then(function (results) {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data user berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}