#!/bin/bash
mkdir bard
cd bard

git clone https://github.com/GuoYuefei/bard.git
cd bard 
git checkout v0.4.4
git submodule init
git submodule update

go build -o bard-client client/client.go
mv bard-client ../
cp -rf client ../

cd bard-plugin
go build -buildmode=plugin -o get_des_cfb_client.so base/client/get_des_cfb_client.go
mv get_des_cfb_client.so ../../client/debug/plugins/
cd ../..

rm -rf bard

vim client/debug/config/config.yml

