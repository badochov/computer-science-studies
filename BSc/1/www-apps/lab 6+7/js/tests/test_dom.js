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
const chai_1 = require("chai");
const mocha_webdriver_1 = require("mocha-webdriver");
exports.TS_NODE_COMPILER_OPTIONS = '{"lib": ["ES2015"]}';
describe('testDrugi', function () {
    it('should say something', function () {
        return __awaiter(this, void 0, void 0, function* () {
            this.timeout(20000);
            yield mocha_webdriver_1.driver.get('file:///home/badochov/Desktop/WWW/lab 6+7/index.html');
            chai_1.expect(yield mocha_webdriver_1.driver.find('#booking-to').getText()).to.include('Łódź');
            yield mocha_webdriver_1.driver.find('input[type=text]').sendKeys('Jan Woreczko');
            yield mocha_webdriver_1.driver.find('input[type=submit]').doClick();
        });
    });
});
