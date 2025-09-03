<?php

function inputTotalPlayer() {
    $input = trim(fgets(STDIN));
    if (!is_numeric($input)) {
        throw new Exception("Invalid input: not a number");
    }
    return intval($input);
}

function inputListScore($totalPlayer) {
    $input = trim(fgets(STDIN));
    $listScoreStr = preg_split('/\s+/', $input);

    if (count($listScoreStr) !== $totalPlayer) {
        throw new Exception("Input mismatch: expected {$totalPlayer} scores, but received " . count($listScoreStr));
    }

    $listScore = [];
    foreach ($listScoreStr as $index => $value) {
        if (!is_numeric($value)) {
            throw new Exception("Invalid score format at position " . ($index + 1) . ": {$value}");
        }
        $listScore[] = intval($value);
    }

    return $listScore;
}

function denseRanking($listScore, $gitsScore) {
    $uniqueScore = array_values(array_unique(array_map('intval', $listScore), SORT_NUMERIC));
    rsort($uniqueScore);

    $result = [];
    foreach ($gitsScore as $gs) {
        $idx = count($uniqueScore) + 1; // Default peringkat terendah
        foreach ($uniqueScore as $index => $score) {
            if ($gs >= $score) {
                $idx = $index + 1;
                break;
            }
        }
        $result[] = $idx;
    }

    return $result;
}

try {
    print("Total Player: ");
    $totalPlayer = inputTotalPlayer();

    print("List Score (Separate By Space): ");
    $listScore = inputListScore($totalPlayer);

    print("Total GITS Player: ");
    $totalGitsPlayer = inputTotalPlayer();

    print("List GITS Score (Separate By Space): ");
    $listGitsScore = inputListScore($totalGitsPlayer);

    $gitsPlayerRanking = denseRanking($listScore, $listGitsScore);

    $gitsPlayerRankingStr = array_map('strval', $gitsPlayerRanking);
    echo "GITS Player Ranking: " . implode(" ", $gitsPlayerRankingStr) . "\n";

} catch (Exception $e) {
    echo "Error: " . $e->getMessage() . "\n";
}
