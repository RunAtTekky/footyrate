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

Expected = 1 / (10 ** (ELO DIFFERENCE/400) + 1)

K = 20

ELO_change = K * (1 - expected)

winner_ELO += ELO_change
loser_ELO -= ELO_change
```

## TODO
- [ ] Implement variable K value
- [ ] Database to have persistent ratings
- [ ] Responsive website
- [ ] Reduce image size for faster loading
- [ ] Better images without watermark

