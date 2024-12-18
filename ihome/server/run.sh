# !/bin/bash
consul agent -dev &

sleep 4

echo "consul agent -dev..."

dirs=("getCaptcha" "getArea" "user" "house")

for dir in "${dirs[@]}"; do
    echo "$dir 已启动..."
    cd "$dir"
    go run . &
    cd ..
done