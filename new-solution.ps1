Param([String] $Title, [String] $IDE)

. $PSScriptRoot/bin/.venv/Scripts/Activate.ps1
python $PSScriptRoot/bin/new-solution.py $PSScriptRoot $Title $IDE