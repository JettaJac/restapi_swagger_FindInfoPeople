


swagger: 
	oapi-codegen --config configs/cfgswagger.yaml ../swagger/swagger.yaml
# генерирует другую реализацию сваггера на echo
#	oapi-codegen --generate types,server --config configs/cfgswagger.yaml  --package swapi ../swagger/swagger.yaml 
	sh pkg/script/script_change_version.sh




run: 
	go run cmd/main.go --config=./configs/local.env

.DEFAULT_GOAL := run