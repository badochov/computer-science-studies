import { Builder, Capabilities } from 'selenium-webdriver';
import { expect } from 'chai';
import { driver } from 'mocha-webdriver';

export const TS_NODE_COMPILER_OPTIONS = '{"lib": ["ES2015"]}';

describe('testDrugi', function () {
    it('should say something', async function () {
        this.timeout(20000);
        await driver.get('file:///home/badochov/Desktop/WWW/lab 6+7/index.html');

        expect(await driver.find('#booking-to').getText()).to.include('Łódź');
        await driver.find('input[type=text]').sendKeys('Jan Woreczko');
        await driver.find('input[type=submit]').doClick();
    });
})