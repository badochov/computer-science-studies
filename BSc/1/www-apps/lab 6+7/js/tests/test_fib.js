"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const main_1 = require("../main");
const chai_1 = require("chai");
require("mocha");
const res = [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89];
describe('Fibonacci async', () => {
    for (const i in res) {
        const val = res[i];
        it(`should equal ${val} for call with ${i}`, () => __awaiter(void 0, void 0, void 0, function* () {
            chai_1.expect(yield main_1.fib(parseInt(i))).to.equal(val);
        }));
    }
});
describe('Fibonacci sync', () => {
    for (const i in res) {
        const val = res[i];
        it(`should equal ${val} for call with ${i}`, () => {
            chai_1.expect(main_1.fib_sync(parseInt(i))).to.equal(val);
        });
    }
});
