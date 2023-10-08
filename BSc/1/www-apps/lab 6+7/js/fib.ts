export async function fib(n: number): Promise<number> {
	if (n < 2) {
		return Promise.resolve(n);
	}
	return (await fib(n - 1)) + (await fib(n - 2));
}
export function fib_sync(n: number): number {
	if (n < 2) {
		return n;
	}
	return fib_sync(n - 1) + fib_sync(n - 2);
}
