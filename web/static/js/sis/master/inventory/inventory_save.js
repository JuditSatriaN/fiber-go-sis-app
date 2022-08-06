// Function to handle button upsert member
$("#upsertInventory").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        buildErrorPopup(param.err);
        return
    }

    saveInventory(param.data);
});

// Function to fill param value
function getParamValue() {
    let id = parseInt($("#modalUpsert #id").val().trim());
    let product = $("#modalUpsert #product_add").select2('data');
    let unit = $("#modalUpsert #unit_add").select2('data');
    let multiplier = $("#modalUpsert #multiplier").val().trim();
    let stock = $("#modalUpsert #stock").val().trim();
    let price = $("#modalUpsert #price").val().trim();
    let member_price = $("#modalUpsert #member_price").val().trim();
    let purchase = $("#modalUpsert #purchase").val().trim();
    let discount = $("#modalUpsert #discount").val().trim();

    if (product.length === 0) {
        return {"data": null, "err": "Silahkan pilih barang terlebih dahulu !"}
    }

    if (unit.length === 0) {
        return {"data": null, "err": "Silahkan pilih unit terlebih dahulu !"}
    }

    if (multiplier === "") {
        return {"data": null, "err": "Harap masukkan inputan pcs terlebih dahulu !"}
    }

    if (stock === "") {
        return {"data": null, "err": "Harap masukkan inputan stok terlebih dahulu !"}
    }

    if (price === "") {
        return {"data": null, "err": "Harap masukkan inputan harga jual terlebih dahulu !"}
    }

    if (member_price === "") {
        return {"data": null, "err": "Harap masukkan inputan harga jual member terlebih dahulu !"}
    }

    if (purchase === "") {
        return {"data": null, "err": "Harap masukkan inputan harga beli terlebih dahulu !"}
    }

    if (discount === "") {
        return {"data": null, "err": "Harap masukkan inputan diskon terlebih dahulu !"}
    }


    return {
        "err": null,
        "data": {
            "id": id,
            "plu": product[0].id,
            "unit_id": unit[0].id,
            "stock": parseInt(stock),
            "price": parseInt(price),
            "purchase": parseInt(purchase),
            "discount": parseInt(discount),
            "multiplier": parseInt(multiplier),
            "member_price": parseInt(member_price),
        },
    }

}

// Function handle save inventory
function saveInventory(data) {
    let url = $('#baseURL').text() + "api/inventory/upsert";
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendPostRequest(url, data).then(function () {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data inventory berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (err) {
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}
