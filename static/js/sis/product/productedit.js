function editProduct(row) {
    fillFormEdit(row)
    $('#modalUpsert').modal('show');
}

function fillFormEdit(row){
    let ppnFmt = row["ppn"] === true ? "Ya" : "Tidak"
    $("#modalUpsert #plu").val(row["plu"]);
    $("#modalUpsert #name").val(row["name"]);
    $("#modalUpsert #barcode").val(row["barcode"]);
    $("#modalUpsert #ppn").val(ppnFmt);
}