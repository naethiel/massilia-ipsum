<?php
declare(strict_types = 1);

$beginnings = [
    "Vé,",
	"Hey,",
	"Minot,",
	"Peuchère,",
	"Ma foi,",
	"Fada,",
	"Ho Gàrri,",
	"Oh tronche d'Àpi,",
	"Doumé,",
	"Parles meilleur,",
	"Fatche de…,",
	"Zou,",
	"Méfi,",
	"Tronche-plate,",
	"Stàssi,",
	"Arretes de marronner de longue,",
	"Bonne mère,",
];

$expressions = [
    "tu as l’air tchouche",
	"tu en as perdu la tchatche",
	"tu aimes mettre le ouaï",
	"tu es un mastre",
	"t'engatse-pas",
	"tu t’es fais chourer ton jaune",
	"t'es tout blanquinas",
	"arrete de faire la bèbe",
	"je t’escagass",
	"on va manger des panisses",
	"il est un peu calu",
	"Je crains dégun",
	"tu boulègues",
	"tu vas t’estramasser",
	"tu me gaves ",
	"elle a la scoumoune",
	"tu me nifles!",
	"il a une figuane de gobi",
	"j'ai eu nibe",
	"il a pris un taquet",
	"j’ai quillé le ballon",
	"je me suis gagué",
	"j'ai passé la pièce car c'était tout pègant",
];

$endings = [
    "avec ton straou.",
	"au vélodrome.",
	"à Endoume.",
	"au cabanon.",
	"dans le teston.",
	"du jaune.",
	"du pastaga.",
	"dans le cabestron.",
	"ça sent l'aïoli.",
	"avec ta figure de poulpe.",
	"avec tes oursins dans les poches.",
	", c'est une belle de cagade.",
	"une soupe d'esques et te jeter aux goudes. ",
	"dans la Gineste.",
	"comme ce pébron de papé.",
	"sur la Corniche.",
	"devant tous ses collègues.",
	", c’est une trompette.",
	", c'est le ouaille.",
	", c'est une vraie arapède.",
	", c’est une radasse.",
	", c'est une bordille",
	", c'est une cagole.",
	", c'est une bouche.",
];

function generate (int $count): string {

    global $beginnings;
    global $expressions;
    global $endings;

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

function generateParagraphs (int $count = 3, string $size = "medium"): array {
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
$length = intval($_GET["length"]);
$size = strval($_GET["size"]);

$response = Array('data' => generateParagraphs($length, $size));

echo json_encode($response);

?>