
$json = Join-Path -Path $PSScriptRoot -ChildPath "install.json" 
$cmds = Get-Content $json -Raw
$RestResults = Invoke-RestMethod -Method 'POST' -Uri 'http://localhost:8080/cmd' -Body $cmds
$RestResults | Out-Host
