import { MemeDesc, Meme } from './Meme';

export const get_most_expensive = (amount: number = 3): MemeDesc[] => {
	return memes.sort((a, b) => a.price - b.price).slice(0, amount).map((meme) => meme.desc);
};

export const memes: Meme[] = [
	new Meme(10, 'Gold', 1000, 'https://i.redd.it/h7rplf9jt8y21.png'),
	new Meme(
		9,
		'Platinum',
		1100,
		'http://www.quickmeme.com/img/90/90d3d6f6d527a64001b79f4e13bc61912842d4a5876d17c1f011ee519d69b469.jpg'
	),
	new Meme(11, 'Elite', 1200, 'https://i.imgflip.com/30zz5g.jpg')
];

export const get_meme = (raw_id: string): Meme | undefined => {
	const id = parseInt(raw_id);

	const meme = memes.find((meme) => meme.desc.id === id);

	return meme;
};
