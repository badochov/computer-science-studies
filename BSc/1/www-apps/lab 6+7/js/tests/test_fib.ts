import { fib, fib_sync } from '../fib';
import { expect } from 'chai';
import 'mocha';
const res = [ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89 ];

describe('Fibonacci async', () => {
	for (const i in res) {
		const val = res[i];
		it(`should equal ${val} for call with ${i}`, async () => {
			expect(await fib(parseInt(i))).to.equal(val);
		});
	}
});

describe('Fibonacci sync', () => {
	for (const i in res) {
		const val = res[i];
		it(`should equal ${val} for call with ${i}`, () => {
			expect(fib_sync(parseInt(i))).to.equal(val);
		});
	}
});
