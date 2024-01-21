# W.I.P
<p>Still in early development and there are many things that needs to be fixed and added! <br></br> Feel free to contribute</p>

## About the Project
A CLI tool made to store and organize your Project Structures (Using Github Repos). And creating new scripts with them. Made with Go.

## Want to try it out?
Start by cloning or downloading the repo and run the `installer.bat` in the `[installer]` folder as an `Administrator` (Its important to run as admin else it cant finde the Program Files)
<p>Now you can use it by typing in `fstruct` in your console</p>

### Usage
Commands right now (root command is `fstruct`)
- **create {structure name} --name {project name} --path {project path}** *(Used for creating a new project)*
- **struct add --name {structure name} --label {structure label} --url {structure github url}** *(Used for creating a new structure in the config.json it works with Github repos)*
- **struct delete --name {structure name}** *(Used for deleting a existing structure)*
- **struct edit {structurename} optional(--name {newname} --label {newlabel} --url {newurl})** Used to edit an existing structure (*doesn't need to use every flags/arguments*)
- **fstruct --help** *(Will give you help if you are lost i will work on making my own help command soon)*
> *More to come in the future!*
