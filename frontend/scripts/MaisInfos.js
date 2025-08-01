document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", (event) => {
        if (event.target.classList.contains("moreInfo")) {
            const dataId = event.target.dataset.id;
            console.log("click no botão com data-id:", dataId);
        }
    });
});
