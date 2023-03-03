Server = ./game/api/server

Server:
	cd $(Server) && goctl api go -api server.api -dir .;
	@echo "Done generate BrowserQuest api";
