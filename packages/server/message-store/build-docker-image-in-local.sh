#!/bin/sh
rm -rf tmp-sources
mkdir tmp-sources
mkdir tmp-sources/customer-os-common-module
mkdir tmp-sources/message-store

rsync -av --progress --exclude="tmp-sources" * tmp-sources/message-store
cp .env tmp-sources/message-store/.env
cp -r ../customer-os-common-module/* tmp-sources/customer-os-common-module

cp Dockerfile tmp-sources/Dockerfile

docker build -t aa tmp-sources/.

rm -rf tmp-sources