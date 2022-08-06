// Insert action inventory
function insertAction() {
    // Reset form inventory modal
    clearFormInput();
    setModalFormatter();
    $("#modalUpsert #id").val();
    $("#modalUpsert #product_add").prop('disabled', false);
    $("#modalUpsert #unit_add").prop('disabled', false);
    $('#modalUpsert').modal('show');
}