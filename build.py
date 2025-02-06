import subprocess


def ask(msg: str) -> bool:
    ryes = ("y", "yes")
    rno = ("n", "no")
    while True:
        reply = input(msg).strip().lower()
        if reply in ryes:
            return True
        elif reply in rno:
            return False
        print(
            f"invalid response, expected '{"', '".join(ryes)}' for yes or '{"', '".join(ryes)}' for no"
        )


def eprin(msg: str):
    print(msg)
    raise SystemExit(1)


def runcom(command: str, *args: str) -> subprocess.CompletedProcess[bytes]:
    return subprocess.run([command, *args])


def build_exe(os: str, arch: str):
    print(f"building {os}/{arch}...")
    res = runcom(
        f"GOOS={os} GOARCH={arch} go",
        "build",
        "-o",
        f"'./build/se_{os}-{arch}" + (".exe" if os == "windows" else "") + "'",
        ".",
    )
    if res.returncode != 0:
        eprin(f"build of {os}/{arch} failed:\n{res.stderr}")


def main():
    if runcom("go", "version").returncode:
        eprin("Go is not installed")
    elif runcom("bytpend", "-h").returncode:
        if ask("do you want to install the bytpend tool? (this is required)"):
            if (
                res := runcom("go", "get", "github.com/voidwyrm-2/bytpend"),
                res.returncode,
            )[1]:
                eprin(f"error installing bytpend:\n{res.stderr}")
        else:
            eprin("build aborted")

    """
    build_exe("darwin", "arm64")
    if (res := runcom("go", "get"), res.returncode)[1]:
        eprin(f"error installing bytpend:\n{res.stderr}")

    build_exe("windows", "amd64")
    if (res := runcom("go", "get", "github.com/voidwyrm-2/bytpend"), res.returncode)[1]:
        eprin(f"error installing bytpend:\n{res.stderr}")
    """

    if runcom("go", "build", "-o", "se", ".").returncode:
        eprin("build failed")
    elif runcom("bytpend", "-o", "zse", "se", "magic.txt", "rsc.zip").returncode:
        eprin("appending failed")


main()
