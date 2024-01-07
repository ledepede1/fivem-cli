package config

type DefaultConfig struct {
	FxManifest   []byte
	ClientMain   []byte
	ServerMain   []byte
	ClConfig     []byte
	SvConfig     []byte
	SharedConfig []byte
}

var DefaultStructure = DefaultConfig{
	FxManifest: []byte(`fx_version 'adamant'

game 'gta5'

author 'your_name'
description 'project_name'

version '1.0'

lua54 'yes'

shared_scripts {
  'Shared/Config.lua',
}

client_scripts {
  'Client/Main.lua',
}

server_scripts {
  'Server/Server.lua',
}
`),

	ClientMain: []byte(`RegisterCommand("ping", function()
	TriggerServerEvent("ping:command")
end, false)`),

	ServerMain: []byte(`RegisterNetEvent('ping:command', function()
	print('Pong')
end)
`),

	ClConfig:     []byte(`Config = {}`),
	SvConfig:     []byte(`Config = {}`),
	SharedConfig: []byte(`Shared = {}`),
}
