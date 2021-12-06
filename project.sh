#!/bin/sh
# docker 版本为19.03.0+ 安装最新的即可
docker_compose_install_version="1.29.2"

docker_compose_install_dir="/usr/local/bin/docker-compose"

echo "开始安装docker-compose(国内源)"
sudo curl -L https://get.daocloud.io/docker/compose/releases/download/$docker_compose_install_version/docker-compose-`uname -s`-`uname -m` > $docker_compose_install_dir
default_install_dir="/usr/local/bin/docker-compose"
if [ "$docker_compose_install_dir" != "$default_install_dir" ]
then
        echo"安装目录非默认目录"
        sudo ln -s $docker_compose_install_dir $default_install_dir
fi
sudo chmod +x $docker_compose_install_dir

docker_compose_install_success=`docker-compose -v|grep -o version`
if [ $"$docker_compose_install_success" ]
then
        echo "docker-compose已经成功安装"
else
        echo "docker-compose未安装成功,请检查执行过程"
        exit
fi

docker-compose up -d

./js_statistics