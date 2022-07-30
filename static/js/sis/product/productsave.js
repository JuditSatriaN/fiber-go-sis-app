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
    let product_id = $("#modalUpsert #product_id").val().trim()
    let name = $("#modalUpsert #name").val().trim()
    let barcode = $("#modalUpsert #barcode").val().trim()
    let stock = parseInt($("#modalUpsert #stock").val().trim())
    let ppn = $("#modalUpsert #ppn").val().trim().trim() === "Ya"
    let price = parseFloat($("#modalUpsert #price").val().trim())
    let member_price = parseFloat($("#modalUpsert #member_price").val().trim())
    let discount = parseFloat($("#modalUpsert #discount").val().trim())

    if (product_id === "") {
        return {"data": null, "err": "ID barang tidak boleh kosong !"}
    }

    if (name === "") {
        return {"data": null, "err": "Nama barang tidak boleh kosong !"}
    }

    return {
        "err": null,
        "data": {
            "product_id": product_id,
            "name": name,
            "barcode": barcode,
            "stock": stock,
            "ppn": ppn,
            "price": price,
            "member_price": member_price,
            "discount": discount
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

    sendSaveProductRequest(data).then(function (results) {
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