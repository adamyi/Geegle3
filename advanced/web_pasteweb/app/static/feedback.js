setTimeout(function(){
  div = document.createElement("div");
  div.setAttribute("id", "ad");
  div.innerHTML = "Like what you see? Go tell your Geegler friends to start using PasteWeb!";
  document.body.insertBefore(div, document.body.firstChild);
  setTimeout(function(){
    div = document.createElement("div");
    div.setAttribute("id", "fb");
    div.innerHTML = "Found an issue? Send the link to pasteweb-feedback@geegle.org and we'll take a look!";
    document.body.appendChild(div);
    $("#ad").after($("#fb"));
  }, 1000);
}, 2000);
