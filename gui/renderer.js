function getData() {
    alert("foo");
}

document.querySelector('#click-me').addEventListener('click', () => {
    getData()
})
