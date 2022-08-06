function editUnit(row) {
    fillFormEdit(row)
    $('#modalUpsert').modal('show');
}

function fillFormEdit(row){
    $("#modalUpsert #id").val(row["id"]);
    $("#modalUpsert #name").val(row["name"]);
}