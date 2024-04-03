document.addEventListener("DOMContentLoaded", (event) => {
  document.body.addEventListener("htmx:beforeSwap", function (evt) {
    if (evt.detail.xhr.status === 400) {
      // allow 400 responses to swap.
      // 400 responses are used to return validation errors
      // set isErrors to false to avoid logging in console
      evt.detail.shouldSwap = true;
      evt.detail.isError = false;
    }
  });
});
