function showError(event) {
  console.log("Error event", event);
  const errorDiv = document.getElementById("error");

  if (event.detail.xhr.status >= 300) {
    errorDiv.textContent = event.detail.xhr.responseText; // Expecting error HTML from server
    errorDiv.classList.add("error-visible"); // Example class for styling
    // Start timer for automatic error clearing (optional)
    clearTimeout(errorTimeout);
    errorTimeout = setTimeout(clearError, 5000);
    console.log("Error message", event.detail.xhr.responseText);
  }
}

function clearError() {
  console.log("Clearing error");
  const errorDiv = document.getElementById("error");
  errorDiv.textContent = "";
  errorDiv.classList.remove("error-visible");
}
