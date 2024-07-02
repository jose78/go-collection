function copyToClipboard(button) {
  let codeContainer = button.nextElementSibling;
  let code = codeContainer.innerText;

  navigator.clipboard.writeText(code).then(() => {
      button.innerText = 'Copied!';
      setTimeout(() => {
          button.innerText = 'Copy';
      }, 2000);
  }).catch(err => {
      console.error('Error copying text: ', err);
  });
}

document.addEventListener('DOMContentLoaded', (event) => {
  // Ensure the function is available globally
  window.copyToClipboard = copyToClipboard;
});


