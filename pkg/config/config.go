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
	'Configs/Shared.lua',
}

client_scripts {
	'Configs/Cl_Config.lua',
	'Client/Main.lua',
}

server_scripts {
	'Configs/Sv_Config.lua',
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
