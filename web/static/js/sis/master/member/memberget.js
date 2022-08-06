$(function () {
    initTable();
    $("#modalUpsert #phone").numeric({decimal: false, negative: false});
})

window.eventActions = {
    'click .edit': function (e, value, row, index) {
        editMember(row);
    },
    'click .remove': function (e, value, row, index) {
        buildDeleteDataPopup("Apakah anda yakin ingin menghapus data ini?", function () {
            deleteMember(row);
        });
    }
}

function initTable() {
    $('#table').bootstrapTable({
        locale: $('#locale').val(),
        columns: [
            [
                {
                    width: 150,
                    field: 'id',
                    title: 'ID',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 350,
                    title: 'Nama',
                    field: 'name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 250,
                    align: 'left',
                    title: 'No HP',
                    field: 'phone',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    title: 'Action',
                    align: 'center',
                    clickToSelect: false,
                    events: window.eventActions,
                    formatter: actionEditDeleteFormatter,
                }
            ],
        ]
    });
}

function ajaxRequest(params) {
    let url = $('#baseURL').text() + "api/dt_members";
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendGetRequest(url).then(function (results) {
        params.success(results);
    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}

function clearFormInput() {
    $("#modalUpsert #id").val("");
    $("#modalUpsert #name").val("");
    $("#modalUpsert #phone").val("");
}