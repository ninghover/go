# !/bin/bash
consul agent -dev &

sleep 5

echo "consul agent -dev..."

dirs=$(find . -maxdepth 1 -type d -not -path '*/\.*' | grep -v '^.$')

for dir in $dirs; do
    echo "$dir 已启动..."
    cd "$dir"
    go run . &
    cd ..
done