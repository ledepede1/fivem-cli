# W.I.P
<p>Still not done, i have some more ideas that i plan on adding in the future! <br></br> Feel free to contribute</p>

## About the Project
A CLI tool made to store and organize your Project Structures (Using Github Repos). And creating new scripts/projects with your stored structures. Made in Go.

## Want to try it out?
Start by cloning or downloading the repo and run the `installer.bat` in the `[installer]` folder as an `Administrator` (Its important to run as admin else it cant finde the Program Files)
<p>Now you can use it by typing in `fstruct` in your console</p>

### Usage
Commands right now (root command is `fstruct`)
- **create {structure name} --name {project name} --path {project path}** *(Used for creating a new project)*
- **struct add --name {structure name} --label {structure label} --url {structure github url}** *(Used for creating a new structure in the config.json it works with Github repos)*
- **struct delete --name {structure name}** *(Used for deleting a existing structure)*
- **struct edit {structurename} optional(--name {newname} --label {newlabel} --url {newurl})** Used to edit an existing structure (*doesn't need to use every flags/arguments*)
- **user username --username {username}** Used to edit the username in the userconfig.json that is used to write the author name to the `fxmanifest.lua` or `__resource.lua`
- **user usename --bool={true/false}** Used to select wether to use the Username or not in the `fxmanifest.lua` or `__resource.lua`
- **fstruct --help** *(Will give you help if you are lost i will work on making my own help command soon)*
> *More to come in the future!*

This software is licensed under the [MIT License](https://github.com/ledepede1/forgestruct-cli/master/LICENSE)
