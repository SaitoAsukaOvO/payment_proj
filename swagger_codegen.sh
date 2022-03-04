#swagger_py_codegen -s ./swagger/swagger.json -p server --ui --validate  --spec  ./
name="server"
sudo chown `whoami` /var/run/docker.sock
docker run -v `pwd`:/tmp openapitools/openapi-generator-cli   generate -t /tmp/swagger-template/go-gin-server -g go-gin-server  -i /tmp/swagger-template/swagger.yaml -o /tmp/swagger-generated
sudo chown -R `whoami` swagger-template
#for file in `find swagger/generated -type f `
#do
#    sed -i "s/swagger_server/$name/g" $file
#done

#mv swagger/generated/swagger_server swagger/generated/server
