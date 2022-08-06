// Function to send get axios request without param
async function sendGetRequest(url) {
    const response = await axios.get(url);
    return response.data
}

// Function to send post axios request
async function sendPostRequest(url, data) {
    return await axios({
        url: url,
        data: data,
        method: 'POST',
    })
}