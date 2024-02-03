@echo off

REM Set the name of the executable
set executable=fstruct.exe
set config=config.json
set userconfig=userconfig.json
set toolFolder=ForgeStruct-CLI
set installPath=C:\Program Files\%toolFolder%

REM Create a dedicated folder for the tool
mkdir "%installPath%" >nul 2>&1

REM Move the executable to the tool folder
move "%~dp0%executable%" "%installPath%\%executable%"

REM Move the config to the tool folder
move "%~dp0%config%" "%installPath%\%config%"

REM Move the user config to the tool folder
move "%~dp0%userconfig%" "%installPath%\%userconfig%"

REM Add the tool's directory to the system PATH
setx PATH "%installPath%;%PATH%" /M >nul 2>&1

echo Installation complete. You can now use '%executable%' from anywhere by typing %toolFolder%.bat.
REM Prompt user to press a key before closing
pause
