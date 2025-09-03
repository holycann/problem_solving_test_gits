<?php

function isValidNumber($s, $index = 0) {
    if ($index >= strlen($s)) {
        return true;
    }
    if ($s[$index] < '0' || $s[$index] > '9') {
        return false;
    }
    return isValidNumber($s, $index + 1);
}

function countMismatch($s, $left, $right) {
    if ($left >= $right) {
        return 0;
    }
    if ($s[$left] !== $s[$right]) {
        return 1 + countMismatch($s, $left + 1, $right - 1);
    }
    return countMismatch($s, $left + 1, $right - 1);
}

function makePalindrome($original, $left, $right, $result, &$changed) {
    if ($left > $right) {
        return [$result, $changed];
    }
    if ($left === $right) {
        $result[$left] = $original[$left];
        return [$result, $changed];
    }

    if ($original[$left] === $original[$right]) {
        $result[$left] = $original[$left];
        $result[$right] = $original[$right];
        return makePalindrome($original, $left + 1, $right - 1, $result, $changed);
    }

    // beda → samakan ke digit lebih besar
    if ($original[$left] > $original[$right]) {
        $result[$left] = $original[$left];
        $result[$right] = $original[$left];
        $changed[$left] = true;
    } else {
        $result[$left] = $original[$right];
        $result[$right] = $original[$right];
        $changed[$right] = true;
    }
    return makePalindrome($original, $left + 1, $right - 1, $result, $changed);
}

function maximizePalindrome($result, $left, $right, $kLeft, &$changed) {
    if ($left > $right) {
        return [$result, $kLeft];
    }
    if ($left === $right) {
        if ($kLeft > 0 && $result[$left] !== '9') {
            $result[$left] = '9';
            $kLeft--;
        }
        return [$result, $kLeft];
    }

    if ($result[$left] === '9' && $result[$right] === '9') {
        return maximizePalindrome($result, $left + 1, $right - 1, $kLeft, $changed);
    }

    if (($changed[$left] || $changed[$right]) && $kLeft >= 1) {
        $result[$left] = '9';
        $result[$right] = '9';
        $kLeft--;
        return maximizePalindrome($result, $left + 1, $right - 1, $kLeft, $changed);
    }

    if ($kLeft >= 2) {
        $result[$left] = '9';
        $result[$right] = '9';
        $kLeft -= 2;
        return maximizePalindrome($result, $left + 1, $right - 1, $kLeft, $changed);
    }

    return maximizePalindrome($result, $left + 1, $right - 1, $kLeft, $changed);
}

function HighestPalindrome($s, $k) {
    if (strlen($s) === 0) {
        return "-1";
    }
    if (!isValidNumber($s)) {
        return "-1";
    }

    $original = str_split($s);
    $n = count($original);

    $mismatch = countMismatch($original, 0, $n - 1);
    if ($mismatch > $k) {
        return "-1";
    }

    $result = array_fill(0, $n, '');
    $changed = array_fill(0, $n, false);
    list($result, $changed) = makePalindrome($original, 0, $n - 1, $result, $changed);

    $kLeft = $k - $mismatch;
    list($result, $_) = maximizePalindrome($result, 0, $n - 1, $kLeft, $changed);

    return implode('', $result);
}

print("Number: ");
$s = trim(fgets(STDIN));

print("Limit: ");
$k = intval(trim(fgets(STDIN)));

$result = HighestPalindrome($s, $k);
echo "Highest Palindrome: " . $result . "\n";
