const ENDPOINT_BASE_URL = '/endpoint';

const endpointAPI = {
  get: async (number) => {
    const url = `${ENDPOINT_BASE_URL}/${number}`;
    const options = { method: 'GET' };

    const response = await fetch(url, options);

    const payload = await response.json();
    if (!response.ok) {
      throw new Error(payload.message);
    }

    return payload;
  },
}

$( '#submitBtn' ).on( 'click', async function(e)  {
  e.preventDefault();

  const resultsDiv = $('.alert-primary');

  resultsDiv.text('Processing...');
  const number = $("#inputNumber").val();

  try {
    const result = await endpointAPI.get(number);
    resultsDiv.text(`Index: ${result.id} Value: ${result.value}`);
  } catch (err) {
    resultsDiv.text(err.message);
  }
});
