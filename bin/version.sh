#!/usr/bin/env bash
## 解决check mismatch的问题
rm -rf ../go.sum

## 获取git tag最后一个的版本号
./get_new_tag.sh

# 如果temp文件不存在，则创建一个
# 覆盖 git tag没有结果的情况
[ -f temp ] || echo 'v0.0.0' > temp

# 获取tag
version=$(cat temp |awk 'END {print}')
[ ! -f temp ] || rm temp

# 把tag写入版本文件
[ ! -f version.toml ] || rm version.toml
echo "no='$version'" > version.toml
time=$(date "+%Y-%m-%d %H:%M:%S")
echo "date='$time'" >> version.toml
cat version.toml