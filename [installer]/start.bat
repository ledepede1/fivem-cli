@echo off
 
REM Set the name of the executable
set executable=fstruct.exe

NET SESSION >nul 2>&1
IF %ERRORLEVEL% EQU 0 (
    REM Move the executable to a directory in the PATH
    move "%~dp0%executable%" "C:\Program Files\%executable%"

    REM Add the executable's directory to the user's PATH
    setx PATH "%PATH%;C:\Program Files" /M

    echo Installation complete. You can now use '%executable%' from anywhere by typing %executable%.
    REM Prompt user to press a key before closing
    pause
    exit /b 0
) ELSE (
    ECHO You need to open this in admin mode!
    REM Prompt user to press a key before closing
    pause
)

 
