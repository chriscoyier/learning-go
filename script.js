fetch("/.netlify/functions/get-data")
  .then((response) => response.json())
  .then((data) => {
    let html = ``;
    data.forEach((item) => {
      html += `<div class="item">${item.title}</div>`;
    });
    content.innerHTML = html;
  });
