Param([String] $Title)

$Title = $Title.Trim()
$Title = $Title -replace '-', '_'

if ($Title -match '_(i{1,3})$')
{
    $Title = $Title -replace '_(i{1,3})$', "_$( $Matches[1].Length )"
}
$Title = $Title -replace '_iv$', '_4'
$Title = $Title -replace '_v$', '_5'

$SolutionDir = "$PSScriptRoot\code\$Title"
New-Item $SolutionDir -ItemType Directory -ErrorAction Ignore

$SolutionSrc = "$SolutionDir\solution.go"
$SolutionTest = "$SolutionDir\solution_test.go"
Write-Output "package $Title`n" | Tee-Object $SolutionSrc | Tee-Object $SolutionTest
goland $PSScriptRoot
goland $SolutionSrc
goland $SolutionTest
