<?php

function show_menu($message, $valid, $release, $operation) {
    echo $message;
    $input = trim(fgets(STDIN));

    if (preg_match('/^\d+$/', $input)) {
        $choice = (int)$input;
        if (in_array($choice, $valid)) {
            $operation($choice);
            if ($choice !== $release) {
                show_menu($message, $valid, $release, $operation);
            }
            return null;
        } else {
            echo "You choose the wrong number, please take one of " . implode(', ', $valid) . PHP_EOL;
            show_menu($message, $valid, $release, $operation);
        }
    } else {
        echo "Invalid, your input must be one of " . implode(', ', $valid) . PHP_EOL;
        show_menu($message, $valid, $release, $operation);
    }
}

?>
