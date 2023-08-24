# clip

a simple copying tool for windows and linux. It takes standard user input through a pipe and copies the text in the clipboard.


## Installation

```sh
go install github.com/saeed0x1/clip@latest
```

### Usage :

for single lines
```bash
echo "some text" | clip
```

for multiple lines
```bash
cat file.txt | clip
```
