# gofighter_examples

[Stockfighter](https://www.stockfighter.io) client programs using my [Gofighter](https://github.com/fohristiwhirl/gofighter) library.

The library is now complete enough to write programs that can (e.g.) clear Level 4. I can't publish such solutions, but the B.L.S.H. bots, the Manual trader, and the Market and Position printers give a reasonable sample of how stuff works.

**Talking to the GM**

To actually run something on a real server, run `game_start.exe`, which saves the info from the gamemaster in the directory `/gm/` - after which it can be loaded by calling `gofighter.GetUserSelection("known_levels.json")`. The json file contains a list of known levels, and asks for a user choice of which level to load.
