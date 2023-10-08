/*
 * Cryptographic API.
 *
 * Adler32 Digest Algorithm.
 *
 * Copyright (c) Cryptoapi developers.
 * Copyright (c) 2023 Hubert Badocha <h.badocha@student.uw.edu.pl>
 *
 * This program is free software; you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation; either version 2 of the License, or (at your option)
 * any later version.
 *
 */
#include <crypto/internal/hash.h>
#include <linux/init.h>
#include <linux/module.h>
#include <linux/string.h>
#include <linux/types.h>
#include <asm/byteorder.h>

#include "adler32.h"

#define MOD_ADLER32 65521

static int adler32_init(struct shash_desc *desc)
{
	struct adler32_state *mctx = shash_desc_ctx(desc);

	mctx->a = 1;
	mctx->b = 0;

	return 0;
}

static int adler32_update(struct shash_desc *desc, const u8 *data, unsigned int len)
{
	struct adler32_state *mctx = shash_desc_ctx(desc);

	for (unsigned int i = 0; i < len; i++)
	{
		mctx->a = (mctx->a + data[i]) % MOD_ADLER32;
		mctx->b = (mctx->b + mctx->a) % MOD_ADLER32;
	}

	return 0;
}

static int adler32_final(struct shash_desc *desc, u8 *out)
{
	struct adler32_state *mctx = shash_desc_ctx(desc);
	u32 res = mctx->b << 16 | mctx->a;

	for (int i = 0; i < sizeof(res); i++)
	{
		out[i] = (res >> (8 * (sizeof(res) - i - 1))) & 0xff;
	}

	memset(mctx, 0, sizeof(*mctx));

	return 0;
}

static int adler32_export(struct shash_desc *desc, void *out)
{
	struct adler32_state *ctx = shash_desc_ctx(desc);

	memcpy(out, ctx, sizeof(*ctx));
	return 0;
}

static int adler32_import(struct shash_desc *desc, const void *in)
{
	struct adler32_state *ctx = shash_desc_ctx(desc);

	memcpy(ctx, in, sizeof(*ctx));
	return 0;
}

static struct shash_alg alg = {
	.digestsize = sizeof(u32),
	.init = adler32_init,
	.update = adler32_update,
	.final = adler32_final,
	.export = adler32_export,
	.import = adler32_import,
	.descsize = sizeof(struct adler32_state),
	.statesize = sizeof(struct adler32_state),
	.base = {
		.cra_name = "adler32",
		.cra_driver_name = "adler32-generic",
		.cra_blocksize = 1,
		.cra_module = THIS_MODULE,
	}};

static int __init adler32_mod_init(void)
{
	return crypto_register_shash(&alg);
}

static void __exit adler32_mod_fini(void)
{
	crypto_unregister_shash(&alg);
}

subsys_initcall(adler32_mod_init);
module_exit(adler32_mod_fini);

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Adler32 Digest Algorithm");
MODULE_ALIAS_CRYPTO("adler32");
