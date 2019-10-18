<?php
error_reporting(E_ALL);

// stolen from https://github.com/munpanel/MUNPANEL_v1/blob/master/app/Http/Controllers/ImageController.php
/**
 * Add text to an ImagickDraw in order to draw text on an Imagick graph with
 * fonts options for both Chinese and English fonts as well as automatic
 * spacing options for generating badge. Auto wrap is enabled in this function,
 * so if text is too long it will automatically be wrapped from below to top.
 * From top to below wrapping mode will be developed in the future.
 *
 * @param Imagick $image the image to be drawed on (It will not draw on that image, but on $draw instead)
 * @param ImagickDraw $draw the instance will the annotation will be applied
 * @param int $x the X position of the center of the text
 * @param int $y the Y position of the center of the text
 * @param string $str the text to be annotated
 * @param int $size the size of the text to be annotated (in pt)
 * @param string $fontCN the name of the font used for Chinese characters
 * @param string $fontEN the name of the font used for English characters
 * @param bool $space whether to automatically add space between characters
 * @return ImagickDraw the instance with text annotated
 */
function addText($image, $draw, $x, $y, $str, $size, $color, $fontCN, $fontEN, $space = false)
{
  if (preg_match("/[\x7f-\xff]/", $str)) // if contains Chinese
  {
    // $draw->setFont(storage_path('app/images/templates/fonts/' . $fontCN));
    $words = preg_split('/(?<!^)(?!$)/u', $str);
    if ($space) {
      switch(count($words))
      {
        // Excel Version by Jiazhao Xu:
        //=IF(LEN(A8)=2,MID(A8,1,1)&"      "&MID(A8,2,1),IF(LEN(A8)=3,MID(A8,1,1)&"  "&MID(A8,2,1)&"  "&MID(A8,3,1),IF(LEN(A8)=4,MID(A8,1,1)&" "&MID(A8,2,1)&" "&MID(A8,3,1)&" "&MID(A8,4,1),IF(LEN(A8)>4,A8))))
      case 2: $space = "      "; break;
      case 3: $space = " "; break;
      case 4: $space = " "; break;
      default: $space= "";
      }
    }
    else $space = '';
  }
  else
  {
    // $draw->setFont(storage_path('app/images/templates/fonts/' . $fontEN));
    $words = preg_split('% %', $str);
    $space = ' ';
  }
  $draw->setFont($fontEN);
  $draw->setFillColor($color);
  $draw->setFontSize($size * 25 / 6); // for 300 ppi

  $font_size = $size * 25 / 6;
  $max_height = 99999;

  $max_width = $image->getImageWidth() / 131 * 125;

  // Holds calculated height of lines with given font, font size
  $total_height = 0;

  // Run until we find a font size that doesn't exceed $max_height in pixels
  while ( 0 == $total_height || $total_height > $max_height ) {
    if ( $total_height > 0 ) $font_size--; // we're still over height, decrease font size and try again

    $draw->setFontSize($font_size);

    // Calculate number of lines / line height
    // Props users Sarke / BMiner: http://stackoverflow.com/questions/5746537/how-can-i-wrap-text-using-imagick-in-php-so-that-it-is-drawn-as-multiline-text
    //$words = preg_split('%\s%', $str);//, PREG_SPLIT_NO_EMPTY);
    $lines = array();
    $l = count($words);
    $i = $l;
    $line_height_ratio = 1;

    $line_height = 0;

    while ( $l > 0 ) {
      $metrics = $image->queryFontMetrics( $draw, implode($space, array_slice($words, --$i - 1) ) );
      $line_height = max( $metrics['textHeight'], $line_height );
      if ( $metrics['textWidth'] > $max_width || $i < 1 ) {
        $lines[] = implode($space, array_slice($words, ++$i - 1) );
        if ($i == 1)
          break;
        $words = array_slice( $words, 0, --$i);
        $l = $i ;
      }
    }

    $total_height = count($lines) * $line_height * $line_height_ratio;



    if ( $total_height === 0 ) return false; // don't run endlessly if something goes wrong
  }

  // Writes text to image
  for( $i = 0; $i < count($lines); $i++ ) {
    $draw->annotation($x, $y - ($i * $line_height * $line_height_ratio), $lines[$i] );
  }
  //        $draw->annotation($x, $y, $str);
  return $draw;
}

$target_dir = "uploads/";
$target_file = $target_dir . basename($_FILES["fileToUpload"]["name"]);
$target_meme_file = $target_dir . "meme_" . basename($_FILES["fileToUpload"]["name"]);
$uploadOk = 1;
// $imageFileType = strtolower(pathinfo($target_file,PATHINFO_EXTENSION));
$imageFileType = mime_content_type($_FILES["fileToUpload"]["tmp_name"]);
// Check if image file is a actual image or fake image
if(isset($_POST["submit"])) {
  $check = getimagesize($_FILES["fileToUpload"]["tmp_name"]);
  if($check == false) {
    die("File is not an image. (getimagesize failed)");
  }
}
// Check if file already exists
if (file_exists($target_file)) {
  die("Sorry, file already exists.");
}
// Check file size
if ($_FILES["fileToUpload"]["size"] > 500000) {
  die("Sorry, your file is too large.");
}
// Allow certain file formats
$formats = [
  "image/gif",
  "image/png",
  "image/jpeg",
];
if (!in_array($imageFileType, $formats)) {
  die("only image/gif image/png image/jpeg allowed. Yours is " . $imageFileType);
}
// Check if $uploadOk is set to 0 by an error
if ($uploadOk == 0) {
  echo "Sorry, your file was not uploaded.";
  // if everything is ok, try to upload file
} else {
  if (move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $target_file)) {
    echo "The file ". basename( $_FILES["fileToUpload"]["name"]). " has been uploaded.";
    $img = new Imagick();
    $img->readImage($target_file);
    $color = "#000000";
    $img->scaleImage(1,1);
    $pixels = $img->getImageHistogram();
    if ($pixels) {
      $rgb = $pixels[0]->getColor();
      $color = sprintf('#%02X%02X%02X', 255 - $rgb['r'], 255 - $rgb['g'], 255 - $rgb['b']);
    }
    $img->readImage($target_file);
    $draw = new ImagickDraw();
    $draw->setTextAlignment(Imagick::ALIGN_CENTER);
    $w = $img->getImageWidth() / 2;
    $h = $img->getImageHeight() - 20;
    addText($img, $draw, $w, $h, $_POST["text"], 12, $color, "", "Roboto-Regular.ttf");
    $img->drawImage($draw);
    $img->writeImage($target_meme_file);

  } else {
    echo "Sorry, there was an error uploading your file.";
  }
}
echo "<img src=\"$target_file\"><img src=\"$target_meme_file\">";
?>
