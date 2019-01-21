#!/usr/bin/env bash

git tag | while read line
    do
       if [[ ! ${line} =~ ^v[0-9]+.[0-9]+.[0-9]+$ ]] ;then
        echo "版本号不符合规范：" ${line}
        continue
       fi

#      获取当前版本号的数组
       current=($(echo $line |tr -d "v"|tr "." "\n"|awk '{print $1}'))

       [ -f temp ] || echo 'v0.0.0' > temp

       # 读取 最新版本号
       version=$(cat temp |awk 'END {print}')

       # 获取最新版本号的数组
       max=($(echo $version |tr -d "v"|tr "." "\n"|awk '{print $1}'))

       # 进行版本比较
       if [[ ${current[0]} -gt ${max[0]} || ${current[0]} -eq  ${max[0]} && ${current[1]} -gt ${max[1]} || ${current[0]} -eq ${max[0]} && ${current[1]} -eq ${max[1]} && ${current[2]} -gt ${max[2]} ]]
          then
            ## 用读写文件解决子进程作用域的问题
            echo $line > temp
       fi
    done

