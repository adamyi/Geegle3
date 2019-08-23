setTimeout(function(){
  sandbox = document.getElementById("ps").contentWindow;
  div = document.createElement("div");
  div.setAttribute("id", "fb");
  div.innerHTML = "Found an issue? Send the link to pasteweb-feedback@geegle.org and we'll take a look!";
  sandbox.document.body.appendChild(div);
  $(sandbox.document.body).prepend(sandbox.fb);
}, 1000);
