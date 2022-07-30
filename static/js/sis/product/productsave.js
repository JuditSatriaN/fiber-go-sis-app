$("#upsertProduct").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        buildErrorPopup(param.err);
        return
    }

    saveProduct(param.data);
});

function getParamValue() {
    let plu = $("#modalUpsert #plu").val().trim()
    let name = $("#modalUpsert #name").val().trim()
    let barcode = $("#modalUpsert #barcode").val().trim()
    let ppn = $("#modalUpsert #ppn").val().trim().trim() === "Ya"

    if (plu === "") {
        return {"data": null, "err": "PLU tidak boleh kosong !"}
    }

    if (name === "") {
        return {"data": null, "err": "Nama barang tidak boleh kosong !"}
    }

    return {
        "err": null,
        "data": {
            "plu": plu,
            "name": name,
            "barcode": barcode,
            "ppn": ppn,
        }
    }

}

async function sendSaveProductRequest(data) {
    let baseURL = $('#baseURL').text();
    const response = await axios({
        data: data,
        method: 'POST',
        url: baseURL + "svc/product/upsert",
    });
    return response
}

function saveProduct(data) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendSaveProductRequest(data).then(function () {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data barang berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}