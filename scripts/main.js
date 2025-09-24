let nameInput = document.getElementById("name");
nameInput.addEventListener("keyup", function(evt){
    let display = document.getElementById("display");
    display.innerHTML = evt.target.value
})