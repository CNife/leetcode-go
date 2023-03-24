#! /usr/bin/env bash
set -euo pipefail

print-usage() {
    cat <<EOF
    $0 <solutionTitle>
EOF
}

title=""
if [ "$1" ]; then
    title="$1"
else
    echo >&2 "no solution provided"
    print-usage
    exit 1
fi

# 将 - 替换为 _
title="${title//-/_}"

declare -A COUNTS
COUNTS['_i']='_1'
COUNTS['_ii']='_2'
COUNTS['_iii']='_3'
COUNTS['_iv']='_4'
COUNTS['_v']='_5'
# 替换结尾的罗马数字为阿拉伯数字
for key in "${!COUNTS[@]}"; do
    title="${title/%$key/${COUNTS[$key]}}"
done

projectDir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"
solutionDir="${projectDir}/code/${title}"
mkdir -p "$solutionDir"
echo "package ${title}" | tee "${solutionDir}/solution.go" >"${solutionDir}/solution_test.go"
