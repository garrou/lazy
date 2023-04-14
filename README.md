# lazy

## Usage

```sh
lazy add series D:/dev/seriesmanager
lazy add mines D:/dev/minesweeper
```

In .config/lazy.json, 2 shortcuts was created

```json
[
    {
        "name": "mines",
        "path": "D:\\dev\\minesweeper"
    },
    {
        "name": "series",
        "path": "D:\\dev\\seriesmanager"
    }
]
```

```sh
code (lazy get mines)               # open vscode at D:\\dev\\minesweeper
cd (lazy get series)                # move at D:\\dev\\seriesmanager
git -C (lazy get series) status     # git status at D:\\dev\\seriesmanager
```
