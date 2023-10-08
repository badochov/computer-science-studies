import express from 'express';
import { get_most_expensive, get_meme } from './memes';
const app = express();
const PORT = 3000;

app.set('view engine', 'pug');
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) =>
	res.render('index', { title: 'Meme market', message: 'Hello there!', memes: get_most_expensive() })
);
app.get('/meme/:memeId', (req, res) => {
	const meme = get_meme(req.params.memeId);
	if (meme) {
		res.render('meme', { title: 'Meme price history', meme: meme });
	} else {
		res.status(404).send("Requested page doesn't exists!");
	}
});
app.post('/meme/:memeId', (req, res) => {
	const meme = get_meme(req.params.memeId);
	const price = req.body.price;
	if (meme) {
		meme.change_price(price);
	}
	res.redirect(req.url);
});

app.listen(PORT, () => console.log(`Example app listening at http://localhost:${PORT}`));
