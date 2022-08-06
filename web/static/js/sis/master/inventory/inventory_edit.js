// Function to edit inventory
function editInventory(row) {
    clearFormInput();
    setModalFormatter();

    fillFormEdit(row);
    $('#modalUpsert').modal('show');
}

// Function to fill form edit
function fillFormEdit(row) {
    // Get previous product data
    let previousProductData = $("<option selected></option>").val(row["plu"]).text(row["name"]);
    let productObj = $("#modalUpsert #product_add");
    productObj.append(previousProductData).trigger('change');
    productObj.prop('disabled', true);

    // Get previous unit data
    let previousUnitData = $("<option selected></option>").val(row["unit_id"]).text(row["unit_name"]);
    let unitObj = $("#modalUpsert #unit_add");
    unitObj.append(previousUnitData).trigger('change');
    unitObj.prop('disabled', true);

    // Get previous other data
    $("#modalUpsert #id").val(row["id"]);
    $("#modalUpsert #multiplier").val(row["multiplier"]);
    $("#modalUpsert #stock").val(row["stock"]);
    $("#modalUpsert #price").val(row["price"]);
    $("#modalUpsert #member_price").val(row["member_price"]);
    $("#modalUpsert #purchase").val(row["purchase"]);
    $("#modalUpsert #discount").val(row["discount"]);
}