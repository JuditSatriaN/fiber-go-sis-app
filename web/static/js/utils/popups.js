function buildErrorPopup(text) {
    Swal.fire({
        icon: 'error',
        title: 'Oops...',
        text: text.charAt(0).toUpperCase() + text.slice(1),
    });
}

function buildSuccessPopup(text) {
    Swal.fire({
        icon: 'success',
        showConfirmButton: false,
        text: text.charAt(0).toUpperCase() + text.slice(1),
    });
}

function buildDeleteDataPopup(text, successCallback) {
    Swal.fire({
        text: text,
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: '#d33',
        confirmButtonText: "Yes",
    }).then(function (result) {
        if (result.isConfirmed) {
            successCallback();
        }
    });
}