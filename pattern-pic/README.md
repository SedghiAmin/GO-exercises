# 🧩 Go Tour Picture Generator

This Go program is based on the **"Pic" exercise** from [A Tour of Go](https://go.dev/tour/moretypes/18).  
It generates a grayscale image using mathematical patterns, represented as a 2D slice of `uint8` values.

---

## 🚀 What It Does

The function `Pic(dx, dy int) [][]uint8` returns a 2D slice (rows × columns),  
where each number represents a pixel intensity between **0–255** (grayscale).

The helper function `pic.Show(Pic)` (from `golang.org/x/tour/pic`) uses this data  
to display an image when run inside the **Go Tour environment**.

---

## 🧠 Pattern Options

Inside the code, you can modify the line that generates pixel values:

| Option | Formula | Description |
|:-------:|:--------|:-------------|
| 1 | `(x + y) / 2` | Simple gradient |
| 2 | `x * y` | Interesting geometric pattern *(default)* |
| 3 | `x ^ y` | XOR pattern |
| 4 | `(x * y) / 2` | Smooth combined pattern |
| 5 | `x % 256` | Wave pattern |
| 6 | `(x/10 + y/10) % 2 * 255` | Checkerboard pattern |

You can comment/uncomment these lines to experiment with different effects.

---

## 🧩 How to Run the Program

### 🟢 Option 1 — Run in the **Go Tour website** (recommended)

1. Visit: [https://go.dev/tour/moretypes/18](https://go.dev/tour/moretypes/18)  
2. Copy and paste your full code into the code editor.
3. Click **“Run”**.  
   → The image will appear directly in the browser.

This is the easiest and official way to visualize your `pic.Show(Pic)` output.

---

### 🧩 Option 2 — Run a local Go Tour server

If you want to run Go Tour locally on your machine:

```bash
# Install Go Tour
go install golang.org/x/website/tour@latest

# Start the tour server
tour
Then open the address shown in your terminal (usually
http://127.0.0.1:3999) in your browser.

Find the “More types → Exercise: Slices” section,
replace the example with your code, and click “Run” to see the result.

⚙️ Option 3 — Running with go run
If you execute the program with:

bash
Copy code
go run main.go
You’ll only see long text output (base64 image data) in your terminal —
that’s normal, since pic.Show is meant to render the image only inside Go Tour.

To actually visualize it, use Option 1 or Option 2 above.