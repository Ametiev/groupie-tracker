document.addEventListener("DOMContentLoaded", function() {
  let dyelesho = document.querySelector("strong[data-gif]");
  let gifSrc = dyelesho.dataset.gif;
  let gifContainer = document.createElement("div");
  let gifImg = document.createElement("img");
  gifImg.src = gifSrc;
  gifContainer.classList.add("gif-container");
  gifContainer.appendChild(gifImg);
  dyelesho.addEventListener("mouseenter", function() {
    setTimeout(function() {
      document.body.appendChild(gifContainer);
    }, 1000);
  });
  dyelesho.addEventListener("mouseleave", function() {
    document.body.removeChild(gifContainer);
  });
});