@echo off

:menu
echo.
echo Escolha uma opção:
echo 1. Rodar o servidor
echo 2. Rodar as migrações
echo 3. Rodar os testes
echo 4. Sair
echo.

set /p opcao="Digite o número da opção desejada: "

if "%opcao%"=="1" goto run
if "%opcao%"=="2" goto migrate
if "%opcao%"=="3" goto test
if "%opcao%"=="4" goto end

:run
echo Rodando o servidor...
go run main.go
goto end

:migrate
echo Rodando as migrações...
go run main.go migrate
goto end

:test
echo Rodando os testes...
go test
pause
goto end

:end
