#!/bin/bash
$1/minify $2/assets/style.css -o $2/assets/style.min.css
$1/minify $2/assets/script.js -o $2/assets/script.min.js
