# ROT-N Reader
A customizable ROT-N cipher reader implementing io.Reader in Go.

Features
Custom ROT-N: Any rotation value (0-25), not just ROT13

Stream Processing: Real-time transformation of any io.Reader

Preserves Format: Non-alphabet characters remain unchanged

# Usage Examples
``` go
// ROT13 (standard)
rot13 := rotReader{strings.NewReader("Hello"), 13}

// ROT1 (Caesar cipher)  
rot1 := rotReader{strings.NewReader("ABC"), 1}

// File decryption
file, _ := os.Open("secret.txt")
rotFile := rotReader{file, 13}
io.Copy(os.Stdout, &rotFile)
How It Works
Uppercase: A → (A + N) % 26

Lowercase: a → (a + N) % 26

```

Non-letters: Unchanged

# Security Note
⚠️ ROT-N is not encryption - use only for education and simple obfuscation.