$remote_command = @{
    key = '088d7646-3e16-11e9-b6e5-af6f2bafb279'
    command = 'Get-Host'
}

$remote_command_json = $remote_command | ConvertTo-Json

$RestResults = Invoke-RestMethod -Method 'POST' -Uri 'http://localhost:8080/powershell' -Body $remote_command_json
$RestResults | Out-Host
