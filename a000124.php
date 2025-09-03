<?php

function lazyCharter($n) {
    return intval((($n * $n) + $n + 2) / 2);
}

$stdin = fopen('php://stdin', 'r');

print("Number: ");
$n = trim(fgets($stdin));

if (!is_numeric($n)) {
    echo "Input Must Be Integer\n";
    exit(1);
}

$n = intval($n);

$a = [];
for ($i = 0; $i < $n; $i++) {
    $a[] = strval(lazyCharter($i));
}

echo "Output: " . implode("-", $a) . "\n";

fclose($stdin);
