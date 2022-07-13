import subprocess
from enum import Enum
from pathlib import Path

import typer

COUNTS = {
    "_i": "_1",
    "_ii": "_2",
    "_iii": "_3",
    "_iv": "_4",
    "_v": "_5",
}


class IDE(Enum):
    GOLAND = "goland"
    VSCODE = "vscode"


IDE_COMMANDS = {
    IDE.GOLAND: "goland",
    IDE.VSCODE: "code",
}


def main(directory: Path, title: str, ide: IDE):
    title = title.strip()
    title = title.replace("-", "_")
    for roma, india in COUNTS.items():
        if title.endswith(roma):
            title.replace(roma, india)

    subdir = directory / "code" / title
    subdir.mkdir(exist_ok=True)
    firstLine = f"package {title}\n"
    codeFile = subdir / "solution.go"
    codeFile.write_text(firstLine)
    testFile = subdir / "solution_test.go"
    testFile.write_text(firstLine)

    subprocess.run(
        [IDE_COMMANDS[ide], str(directory), str(codeFile)],
        shell=True,
        check=True,
    )
    subprocess.run(
        [IDE_COMMANDS[ide], str(directory), str(testFile)],
        shell=True,
        check=True,
    )


if __name__ == "__main__":
    typer.run(main)
