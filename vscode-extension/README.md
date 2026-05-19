# R-Lang Syntax

VS Code syntax highlighting for `.rlang` files.

## Local Install

Copy or symlink this folder into your VS Code extensions directory, then reload VS Code.

macOS/Linux:

```sh
mkdir -p ~/.vscode/extensions
ln -s "$(pwd)/vscode-extension" ~/.vscode/extensions/rlang-syntax
```

Windows PowerShell:

```powershell
New-Item -ItemType Directory -Force "$env:USERPROFILE\.vscode\extensions"
New-Item -ItemType SymbolicLink -Path "$env:USERPROFILE\.vscode\extensions\rlang-syntax" -Target "$PWD\vscode-extension"
```

Open a `.rlang` file and choose `R-Lang` as the language mode if VS Code does not pick it automatically.
