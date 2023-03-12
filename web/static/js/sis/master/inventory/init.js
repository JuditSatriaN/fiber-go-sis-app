// Initialize table when document is ready
$(function () {
    initTable();
    addNewUnitForm();
    addNewProductForm();
});

// Add new product form (to handle fill product_id)
function addNewProductForm() {
    let baseURL = $('#baseURL').text();
    $("#modalUpsert #product_add").select2({
        placeholder: "Silahkan pilih product",
        ajax: {
            type: "GET",
            url: baseURL + "api/products",
            data: function (params) {
                return {
                    search: params.term,
                };
            },
            processResults: function (results) {
                return {
                    results: $.map(results, function (result) {
                        return {
                            id: result.plu,
                            text: result.name
                        }
                    })
                };
            }
        },
    });
}

// Add new unit form (to handle fill unit_id)
function addNewUnitForm() {
    let unitURL = $('#baseURL').text() + "api/units";
    sendGetRequest(unitURL).then(function (results) {
        let resultData = []
        results.forEach(function (result) {
            resultData.push({
                id: result.id,
                text: result.name
            })
        });

        $("#modalUpsert #unit_add").select2({
            data: resultData,
            placeholder: "Silahkan pilih unit",
        });

    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    });
}

function clearFormInput() {
    $("#modalUpsert #id").val("0");
    $('#modalUpsert #product_add').val(null).trigger('change');
    $('#modalUpsert #unit_add').val(null).trigger('change');
    $("#modalUpsert #multiplier").val("0");
    $("#modalUpsert #stock").val("0");
    $("#modalUpsert #price").val("0");
    $("#modalUpsert #member_price").val("0");
    $("#modalUpsert #purchase").val("0");
    $("#modalUpsert #discount").val("0");
    $("#modalUpsert #current_price_form").hide();
}

function setModalFormatter() {
    $('#modalUpsert #multiplier').number(true, 0, ',', '.');
    $('#modalUpsert #stock').number(true, 0, ',', '.');
    $('#modalUpsert #current_price').number(true, 0, ',', '.');
    $('#modalUpsert #price').number(true, 0, ',', '.');
    $('#modalUpsert #member_price').number(true, 0, ',', '.');
    $('#modalUpsert #purchase').number(true, 0, ',', '.');
    $('#modalUpsert #discount').number(true, 0, ',', '.');
    $('#modalUpsert #persentage_price').numeric({decimal: false, negative: false});
}

$('#modalUpsert #persentage_price').on('input', function () {
    let currentValue = $(this).val();
    let valueInt = parseInt(currentValue);
    if (valueInt > 100) {
        $(this).val(currentValue.slice(0, -1))
        valueInt = currentValue.slice(0, -1)
    }
    let oldPrice = parseInt($('#modalUpsert #current_price').val())
    let newPrice = oldPrice + (oldPrice * valueInt / 100)
    $('#modalUpsert #price').val(newPrice)
});
