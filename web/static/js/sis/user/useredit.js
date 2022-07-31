function editUser(row) {
    fillFormEdit(row)
    $('#modalUpsert').modal('show');
}

function fillFormEdit(row) {
    let isAdminFmt = row["is_admin"] === true ? "Ya" : "Tidak"
    let userIDObject = $("#modalUpsert #user_id")
    userIDObject.val(row["user_id"]);
    userIDObject.prop('disabled', true);
    $("#modalUpsert #user_name").val(row["user_name"]);
    $("#modalUpsert #full_name").val(row["full_name"]);
    $("#modalUpsert #password").val(row["password"]);
    $("#modalUpsert #is_admin").val(isAdminFmt);
}