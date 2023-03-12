async function sendGetInventoryRequest(params) {
    let page = 1;
    let req = params.data;
    let baseURL = $('#baseURL').text();
    if (params.data["offset"] !== 0) {
        page = (params.data["offset"] / params.data["limit"]) + 1
    }

    const response = await axios({
        method: 'GET',
        url: baseURL + "api/inventory",
        params: {
            "page": page,
            "limit": req["limit"],
            "search": req["search"],
        },
    });
    return response.data
}

function ajaxRequest(params) {
    let loadingIndicator = $('body').loadingIndicator().data("loadingIndicator");

    sendGetInventoryRequest(params).then(function (results) {
        params.success(results);
    }).catch(function (err) {
        params.error(err);
        buildErrorPopup(err.response.data.message);
    }).finally(function () {
        loadingIndicator.hide();
    });
}