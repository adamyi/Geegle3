<?php
if (isset($_GET['src']))
  die(highlight_string(file_get_contents(__FILE__)));
if (isset($_GET["type"]))
  $type = $_GET["type"];
else
  $type = 1;

switch ($type) {
case 1:
  $htype = "Proof";
  break;
case 2:
  $htype = "Counter Proof";
  break;
/*
case 3:
  $htype = "Dev notes";
  break;
 */
default:
  $type = 4;
  $htype = "Shakespeare";
}

echo "<h1>Is the Earth Flat?</h1>";
echo "<h3>$htype</h3>";
$sql = "SELECT text FROM arguments WHERE type=$type";
$db = new SQLite3("/db");
$result = $db->query($sql);
while($arg = $result->fetchArray()) {
  echo $arg[0] . "<br>";
}
echo "<a href=\"/about.html\">About</a>&nbsp;<a href=\"?src\">Sauce</a>";
