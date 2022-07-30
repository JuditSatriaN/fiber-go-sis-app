function insertAction() {
    clearFormInput()
    $('#modalUpsert').modal('show');
    $("#modalUpsert #user_id").prop('disabled', false);
}