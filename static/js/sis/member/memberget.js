$(function () {
    initTable();
})

function actionFormatter() {
    return [
        '<a class="edit" href="javascript:void(0)" title="Edit"><i class="fa fa-edit"></i></a>',
        '<a class="remove" href="javascript:void(0)" title="Remove"><i class="fa fa-trash"></i></a>'
    ].join('')
}

window.eventActions = {
    'click .edit': function (e, value, row, index) {
        editMember(row);
    }, 'click .remove': function (e, value, row, index) {
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
                    formatter: actionFormatter,
                    events: window.eventActions,
                }
            ],
        ]
    });
}

async function sendGetMemberRequest() {
    let baseURL = $('#baseURL').text();
    const response = await axios.get(baseURL + "svc/dt_members");
    return response.data
}

function ajaxRequest(params) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendGetMemberRequest().then(function (results) {
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