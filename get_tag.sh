#!/bin/sh -eu

# 最新タグのパッチバージョンをインクリメントして返す
git tag | sort | tail -1 | awk -F "." "{cnt=\$3+1}{printf \"%s.%s.%d\", \$1, \$2, cnt}"