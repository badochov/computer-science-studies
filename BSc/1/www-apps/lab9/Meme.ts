export interface Price {
	price: number;
	date: Date;
}

export interface MemeDesc {
	id: number;
	name: string;
	price: number;
	url: string;
}

export class Meme {
	private prices: Price[] = [];
	constructor(private id: number, private name: string, price: number, private url: string) {
		this.set_price(price);
	}

	public get price(): number {
		return this.prices[this.prices.length - 1].price;
	}

	private set_price(price: number) {
		this.prices.push({ price: price, date: new Date() });
	}

	public get desc(): MemeDesc {
		return {
			id: this.id,
			name: this.name,
			price: this.price,
			url: this.url
		};
	}

	public change_price(raw_price: any): void {
		const price =
			typeof raw_price === 'number' ? raw_price : typeof raw_price === 'string' ? parseFloat(raw_price) : null;
		if (price !== null && !isNaN(price)) {
			this.set_price(price);
		}
	}

	public get price_history(): Price[] {
		return this.prices.slice().reverse();
	}
}
