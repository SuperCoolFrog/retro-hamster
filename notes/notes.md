# Notes

## Why 0 Angle Spawns to the Right

- In standard 2D polar coordinates, which use radians:
- Angle 0 points directly to the right (along the positive X-axis).
- Angle π/2 points straight down (positive Y-axis).
- Angle π points left.
- Angle 3π/2 or -π/2 points up.
- This is due to:
- The unit circle.
- The screen coordinate system in most 2D engines where (0,0) is top-left, and Y increases downward.

```
     ↑
      |       π/2 (up)
      |
← π ------- 0 → (right)
      |
      |       3π/2 or -π/2 (down)
      ↓
```

## Web Assembly Build

### Build wasm 

```
go build -o bin/retro-hamster.exe ./cmd
```

### create html file

```
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Your Ebitengine Game</title>
    <style>
        body { margin: 0; overflow: hidden; } /* Optional: Remove default margins and scrollbars */
        canvas { display: block; margin: 0 auto; } /* Optional: Center the canvas */
    </style>
</head>
<body>
    <script src="wasm_exec.js"></script>
    <script>
        // Polyfill for browsers that don't support WebAssembly.instantiateStreaming
        if (!WebAssembly.instantiateStreaming) {
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }

        const go = new Go(); // Create a new Go instance
        WebAssembly.instantiateStreaming(fetch("retro-hamster.wasm"), go.importObject)
            .then((result) => {
                go.run(result.instance); // Run the Go program
            })
            .catch((err) => {
                console.error("Error loading WebAssembly:", err);
            });
    </script>
</body>
</html
```

### Include wasm_exec.js from go lib

```
echo $(go env GOROOT)/lib/wasm/wasm_exec.js
```


### Test
```
npx http-server .
```



## Credits

https://www.dafont.com/sunny-spells.charmap