async function sendGetInventoryRequest() {
    let baseURL = $('#baseURL').text();
    const response = await axios.get(baseURL + "api/inventory");
    return response.data
}

function ajaxRequest(params) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendGetInventoryRequest().then(function (results) {
        params.success(results);
    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}