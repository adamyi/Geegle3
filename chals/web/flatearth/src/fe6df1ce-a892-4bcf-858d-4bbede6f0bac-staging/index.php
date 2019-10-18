<?php
if (isset($_GET['src']))
  die(highlight_string(file_get_contents(__FILE__)));
if (isset($_GET["search"])) {
  $search = $_GET["search"];
  if (isset($_GET["type"]))
    $file = "arguments/" . $_GET["type"] . ".php";
  else
    $file = "arguments/";
} else {
  if (isset($_GET["type"]))
    $type = $_GET["type"];
  else
    $type = "proof"; // or counterproof
  $file = "arguments/" . $type . ".php";
}

// code review says there was lfi
// so i fixed it
assert("strpos('$file', '..') === false") or die("you bad bad");

// code review says there was xss
// so i fixed it
echo "<h1>Is the Earth Flat?</h1>";
echo "<h3>" . htmlspecialchars($type) . "</h3>";
echo "<pre>";
if (isset($search)) {
  // code review says there was command line injection
  // so i fixed it
  system("grep -r " . escapeshellarg($search) . " " . escapeshellarg($file));
} else {
  require_once($file);
}
echo "</pre>";
echo "<a href=\"/about.html\">About</a>&nbsp;<a href=\"?src\">Sauce</a>";
