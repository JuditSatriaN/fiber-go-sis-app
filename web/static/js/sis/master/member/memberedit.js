// Function to handle edit member
function editMember(row) {
    fillFormEdit(row);
    $('#modalUpsert').modal('show');
}

// Function to fill form edit member
function fillFormEdit(row){
    $("#modalUpsert #id").val(row["id"]);
    $("#modalUpsert #name").val(row["name"]);
    $("#modalUpsert #phone").val(row["phone"]);
}