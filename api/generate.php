<?php
declare(strict_types = 1);

$sentences = json_decode(file_get_contents("data.json"), true);

function generate (int $count): string {
	global $sentences;
	
	$beginnings = $sentences["beginnings"];
	$expressions = $sentences["expressions"];
	$endings = $sentences["endings"];

    $beginningsIdx = range(0, count($beginnings) - 1);
    $expressionsIdx = range(0, count($expressions) - 1);
    $endingsIdx = range(0, count($endings) - 1);
    shuffle($beginningsIdx);
    shuffle($expressionsIdx);
    shuffle($endingsIdx);

    $p = "";

    for ($i = 0; $i < $count; $i++) {
        $p .= $beginnings[$beginningsIdx[$i % count($beginnings)]];
        $p .= " ";
        $p .= $expressions[$expressionsIdx[$i % count($expressions)]];

        $ending = $endings[$endingsIdx[$i % count($endings)]];

        if ($ending[0] !== ","){
            $p .= " ";
        }

        $p .= $endings[$endingsIdx[$i % count($endings)]];

        if ($i !== $count - 1) {
            $p .= " ";
        }
    }

    return $p;
}

function generateParagraphs (int $count, string $size): array {
    $result = Array();
    $sentences = 7;
    switch ($size) {
        case 'small':
            $sentences = 3;   
        break;
        case 'medium':
            $sentences = 7;
        break;
        case "large":
            $sentences = 12;
        break;
    }

    for ($i = 0; $i < $count; $i++) {
        $result[] = generate($sentences);
    }

    return $result;
}



header('Access-Control-Allow-Origin: *');
header("Content-Type: application/json; charset=UTF-8");

$allowedMethods = array('GET');
 
$requestMethod = strtoupper($_SERVER['REQUEST_METHOD']);
 
if(!in_array($requestMethod, $allowedMethods)){
	http_response_code(405);
    exit;
}
 
$length = min(intval($_GET["length"]), 100);
$size = strval(htmlspecialchars($_GET["size"]));

if (!isset($size) || ! isset($length) || $length <= 0) {
	http_response_code(400);
	exit;
}


$response = Array('data' => generateParagraphs($length, $size));

echo json_encode($response);

?>
