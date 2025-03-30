# FootyRate
A website to compare footballers, inspired by [The Social Network](https://www.imdb.com/title/tt1285016) movie.

Backend: [go](https://github.com/golang/go)

Backend URL: [Render](https://footyrate.onrender.com)

Frontend: [Svelte](https://github.com/sveltejs/svelte)

Frontend URL: [Vercel](https://footyrate.vercel.app)

## How it works?
A base ELO rating of 1400 is given to all the players.

When two players are compared, their ratings change according to the given formula.

```python
ELO_DIFFERENCE = WINNER_CURRENT_ELO - LOSER_CURRENT_ELO

Expected = 1 / (10 ** (ELO_DIFFERENCE/400) + 1)

K = (10 if ROUNDS > 30, 20 if ROUNDS > 20, 30 if ROUNDS > 10, 40 if ROUNDS <= 10)

ELO_change = K * (1 - expected)

winner_ELO += ELO_change
loser_ELO -= ELO_change
```

## Onrender Problem
The backend server shuts down after inactivity and takes a long time to start back up.

So when you open the website after the server has shut down, it will take some time before showing you the images.

## TODO
- [x] Implement variable K value
- [x] Standings
- [x] Add styling to standings
- [x] Left and Right button to select players
- [ ] Database to have persistent ratings
- [x] Responsive website
- [x] Reduce image size for faster loading
- [ ] Better images without watermark

