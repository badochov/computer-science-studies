import { expect } from "chai";
import { driver, WebDriver } from "mocha-webdriver";
import webDriver from "selenium-webdriver";
const URL = "http://localhost:3001";
describe("Test 5 entries on main page", function () {
  it("should be present", async function () {
    await driver.get(URL);

    await driver.findWait("#ostatnie-wpisy > ul > li:nth-child(1)", 1000);
    const ul = await driver.find("#ostatnie-wpisy > ul");
    const li = await ul.findElements(webDriver.By.tagName("li"));
    expect(li.length).to.equal(5);
  });
});
describe("Test login", function () {
  it("should redirect to users and the logout", async function () {
    await driver.get(URL);

    await driver.find("[name=login]").sendKeys("nauczyciel_1");
    await driver.find("[name=haslo]").sendKeys("haslo_1");
    await (await driver.find("[type=submit]")).click();
    await expect(await driver.getCurrentUrl()).to.equal(URL + "/users");
    await (
      await driver.find("body > form > input[type=submit]:nth-child(2)")
    ).click();

    await expect(await driver.getCurrentUrl()).to.equal(URL + "/");
  });
  it("should show error message", async function () {
    await driver.get(URL);

    await (await driver.find("[type=submit]")).click();
    await driver.findWait("#error-message", 1000);
  });
});

async function login(driver: WebDriver) {
  await driver.get("http://localhost:3001");

  try {
    await driver.find("#username").sendKeys("test");
  } catch (_) {
    expect(await (await driver.find("body > p")).getText()).to.equal(
      "Witaj user1!"
    );
    return;
  }
  await driver.find("#password").sendKeys("test");
  await driver.find("body > form > input[type=submit]:nth-child(11)").doClick();
  expect(await driver.find("body > p").getText()).to.include("test");
}

async function logout(driver: WebDriver) {
  await driver.get(URL);
  await (
    await driver.find("body > form > input[type=submit]:nth-child(2)")
  ).click();
  await driver.findWait("#sign-up", 2000);
}
