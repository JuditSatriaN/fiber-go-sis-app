// Function to handle edit system config
function editSystemConf(row) {
    fillFormEdit(row);
    $('#modalUpsert').modal('show');
}

// Function to fill form edit member
function fillFormEdit(row){
    $("#modalUpsert #id").val(row["id"]);
    $("#modalUpsert #value").val(row["value"]);
}