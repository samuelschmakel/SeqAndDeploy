const button = document.querySelector('.analysis-button');

function init() {
    console.log(`button: ${button}`);
    button.addEventListener("click", handleClick);
}

function handleClick(e) {
    console.log(`This event: ${e} happened.`);
}

init();