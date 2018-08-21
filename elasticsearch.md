# elasticsearch 
采用了Docker部署，Docker只有elasticsearch 5
## 获得配置文件
```$xslt
 docker run -d --name elastic -e transport.host=0.0.0.0 -p 9200:9200 -p 9300:9300 elasticsearch:5.6.10
 docker cp elastic:/usr/share/elasticsearch/config ./
 docker rm -f elastic
```
## 修改配置文件,放到`/mnt/docker_container/elasticsearch/config`目录下
```$xslt
 docker run -d --name elastic -e transport.host=0.0.0.0\
  -v /mnt/docker_container/elasticsearch/config:/usr/share/elasticsearch/config\
  -v /mnt/docker_container/elasticsearch/esdata:/usr/share/elasticsearch/data\
  elasticsearch:5.6.10
```
修改`config.yml`中的信息
```$xslt
elastic:
  enable: true
  server_address:
    - http://127.0.0.1:9200
  sniffer_enabled: false
  auth: false
  auth_username:
  auth_password:
```