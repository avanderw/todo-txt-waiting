# Set the URL to post to
$url = "https://avanderw.tplinkdns.com:31025/daily"

# Set the path to the file to write the output to
$outputFile = "./waiting.txt"

# Set the path to the binary to execute
$binary = "./todo-txt-waiting.exe"
$params = "./config.txt"

# Execute the binary and redirect the output to the file
& $binary $params | Out-File $outputFile

# Read the contents of the output file
$content = Get-Content $outputFile -Raw

# Set the headers for the request
$headers = @{
    "Content-Type" = "text/plain"
}

# Make the POST request with the output contents
Invoke-RestMethod -Uri $url -Method Post -Headers $headers -Body $content
