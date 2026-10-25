# Wordle (Go)

A terminal implementation of the classic [Wordle](https://www.nytimes.com/games/wordle/index.html) game, written in Go as a single-file console game.

## How to run

```sh
go run main.go
```

Requires Go 1.22+ (uses `math/rand/v2`).

## How the game works

1. On start, the game picks a random secret 4-letter word from `WordsDB` using `rand.IntN`.
2. The player gets **6 attempts** to guess it.
3. After each guess, every letter is evaluated and printed on its own line:
   - **"is in the correct place"** — the letter matches the secret at the same index (green).
   - **"is in the word but not in the correct place"** — the letter exists elsewhere in the secret (yellow).
   - **"is not in the word"** — the letter is absent (gray).
4. If the guess equals the secret, the player wins. After 6 wrong guesses, the secret word is revealed.

### Repeated-letter logic

The feedback follows the real Wordle rules for duplicate letters. A `Used []bool` slice tracks which positions of the secret have already been "claimed":

- Exact (green) matches are checked **first** and mark their position as used.
- A yellow match only claims an **unclaimed** position of the secret.

So for secret `eggs` and guess `gggg`: the first `g` is green (claims index 1), the second `g` is yellow (claims index 2), and the remaining `g`s are reported as not in the word — matching official Wordle behavior.

### Code structure

| Part | Responsibility |
|---|---|
| `WordsDB` | Fixed array of valid 4-letter words |
| `engine()` | Full game loop: pick secret, read guesses, validate length, give feedback, decide win/loss |
| `main()` | Prints welcome/farewell messages and runs one game |

## Known issues

These were identified in a review and are **intentionally not fixed yet**:

- **Byte-based length check** — `len(guess)` counts bytes, not characters. A guess with non-ASCII letters (e.g. accented characters) has a wrong length and gets rejected or misprocessed; the game uses `rune` conversions elsewhere but compares length in bytes.
- **No case normalization** — typing `HELP` never matches the secret `help`; uppercase guesses are treated as wrong and every letter is reported as "not in the word".
- **No alphabet validation** — digits or symbols (`1234`) pass the length check and are evaluated as letters.
- **Guesses aren't checked against `WordsDB`** — any 4-character token is accepted, unlike real Wordle which only allows known words.
- **Feedback format differs from classic Wordle** — one line per letter in guess order, instead of a colored per-position row with greens shown before yellows.
- **Misleading message on the last attempt** — "Wrong guess. Try another word." is printed even on attempt 6, right before the loss is announced.
- **Questionable dictionary entries** — e.g. `"chan"`, `"yoyo"` are not standard Wordle answers.

## Note

I am aware of all the bugs and issues listed above. The core game logic (including the repeated-letter handling) works as intended for normal input. I may come back later to complete the fixes — this project is a abanded.
