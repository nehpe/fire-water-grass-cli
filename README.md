# go-fwg: Fire / Water / Grass

An alternate to Rock / Paper / Scissors.

You select either fire, water, or grass. The CPU chooses. You win / lose / draw.

## Usage

```
./fwg (fire|water|grass)
```

## Rules

Shorthand Rules:

    - Fire
        - beats Grass (burns it)
        - is beaten by Water (extinguished)
    - Water
        - beats Fire (extinguishes it)
        - is beaten by Grass (gives it growth)
    - Grass
        - beats Water (grows from it)
        - is beaten by Fire (burnt)
