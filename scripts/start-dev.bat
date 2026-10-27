@echo off
cd /d "%~dp0.."

echo ============================================
echo   GoFrame Mall - development server
echo   URL : http://127.0.0.1:8000
echo   Docs: http://127.0.0.1:8000/swagger/
echo   Stop: Ctrl+C
echo ============================================
echo.

REM open browser after 5 seconds, server needs time to boot
start "" cmd /c "timeout /t 5 >nul & start http://127.0.0.1:8000/"

go run .

echo.
echo Server stopped.
pause
