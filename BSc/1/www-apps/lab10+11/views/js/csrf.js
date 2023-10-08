const xhttp = new XMLHttpRequest();
xhttp.onreadystatechange = function () {
  if (this.readyState == 4 && this.status == 200) {
    document.getElementById("demo").innerHTML = this.responseText;
  }
};
xhttp.open("POST", "/meme/1" + Math.random(), true);
xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
xhttp.send("price=" + Math.round(Math.random()));
