@echo off
setlocal

REM ============================================================
REM  MySQL 8.4 one-click setup
REM  Usage: right-click this file -> "Run as administrator"
REM  NOTE: keep this file ASCII-only. Chinese text in .bat gets
REM        garbled under GBK console and breaks the syntax.
REM ============================================================

set "MYSQL_BIN=C:\Program Files\MySQL\MySQL Server 8.4\bin"
set "MY_INI=C:\ProgramData\MySQL\MySQL Server 8.4\my.ini"
set "SVC=MySQL84"
set "DBNAME=gfstudy"

REM Default root password. Override with: set ROOT_PWD=yourpass
if "%ROOT_PWD%"=="" set "ROOT_PWD=123456"

echo.
echo ============================================
echo   MySQL 8.4 service setup
echo ============================================
echo.

net session >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Administrator privileges required.
    echo         Right-click this file, choose "Run as administrator".
    echo.
    pause
    exit /b 1
)
echo [OK] Running as administrator.

if not exist "%MYSQL_BIN%\mysqld.exe" (
    echo [ERROR] mysqld.exe not found at %MYSQL_BIN%
    pause
    exit /b 1
)
if not exist "%MY_INI%" (
    echo [ERROR] Config file not found: %MY_INI%
    pause
    exit /b 1
)

sc query %SVC% >nul 2>&1
if not errorlevel 1 (
    echo [INFO] Removing existing service %SVC% ...
    net stop %SVC% >nul 2>&1
    sc delete %SVC% >nul 2>&1
    timeout /t 2 >nul
)

echo [1/4] Installing service %SVC% ...
"%MYSQL_BIN%\mysqld.exe" --install %SVC% --defaults-file="%MY_INI%"
if errorlevel 1 goto FAIL

echo [2/4] Setting auto start ...
sc config %SVC% start= auto >nul

echo [3/4] Starting service ...
net start %SVC%
if errorlevel 1 goto FAIL
timeout /t 5 >nul

echo [4/4] Setting root password and creating database ...
"%MYSQL_BIN%\mysql.exe" -u root --skip-password -e "ALTER USER 'root'@'localhost' IDENTIFIED BY '%ROOT_PWD%';"
if errorlevel 1 (
    echo [WARN] Failed to set password. root may already have one.
) else (
    echo        root password set to: %ROOT_PWD%
)
"%MYSQL_BIN%\mysql.exe" -u root -p%ROOT_PWD% -e "CREATE DATABASE IF NOT EXISTS %DBNAME% DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;"
if errorlevel 1 (
    echo [WARN] Failed to create database %DBNAME%.
) else (
    echo        database created: %DBNAME%
)

echo.
echo ============================================
echo   DONE
echo   Host     : 127.0.0.1
echo   Port     : 3306
echo   User     : root
echo   Password : %ROOT_PWD%
echo   Database : %DBNAME%
echo ============================================
echo.
echo Verify with:
echo   "%MYSQL_BIN%\mysql.exe" -u root -p%ROOT_PWD% -e "SHOW DATABASES;"
echo.
pause
exit /b 0

:FAIL
echo.
echo [ERROR] Setup failed. See messages above.
pause
exit /b 1
