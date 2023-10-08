import { expect } from "chai";
import { driver, WebDriver } from "mocha-webdriver";
import { Key, WebElement } from "selenium-webdriver";
import { QuizNoAnswers } from "../quizes";

describe("testChangePassword", function () {
  it("should be relogged", async function () {
    this.timeout(20000);
    await login(driver);
    expect(await driver.find("body > p").getText()).to.include("test");
    await driver.find("body > a").click();
    await driver.find("#username").sendKeys("test");
    await driver.find("#password").sendKeys("test");
    await driver.find("#new-password").sendKeys("test");
    await (
      await driver.find("body > form > input[type=submit]:nth-child(11)")
    ).click();
    expect(
      await (
        await driver.find("body > form > input[type=submit]:nth-child(11)")
      ).getAttribute("value")
    ).to.include("Zaloguj");
  });
});
const isQuiz = (quiz: any): quiz is QuizNoAnswers => {
  return quiz.desc && quiz.questions && quiz.id !== undefined;
};

describe("testTests", function () {
  it("should be in json", async function () {
    this.timeout(20000);
    await login(driver);
    const obj = <any>await driver.executeAsyncScript((callback: any) => {
      fetch("/get_quiz/2").then((res) => {
        res.json().then(callback);
      });
    });
    expect(isQuiz(obj)).to.equal(true);
    if (isQuiz(obj)) {
      const res = <any>await driver.executeAsyncScript((callback: any) => {
        fetch("/post_results/" + 2, {
          method: "POST",
          credentials: "same-origin",
          headers: {
            "Content-Type": "application/json",
            "CSRF-Token":
              document
                ?.querySelector('meta[name="csrf-token"]')
                ?.getAttribute("content") || "",
          },
          body: JSON.stringify({
            times: [0.25, 0.1, 0.35, 0.3],
            answers: [1, 1, 1, 1],
          }),
        }).then((raw) => raw.json().then(callback));
      });
      expect(typeof res === "number").to.equal(true);
      const answers = <any>await driver.executeAsyncScript((callback: any) => {
        fetch("/get_answers/" + 2).then((raw) => raw.json().then(callback));
      });
      expect(answers instanceof Array).to.equal(true);
      for (const ans of answers) {
        expect(ans.answer).to.equal(1);
      }
    }

    await logout(driver);
  });
});

describe("testTryingToSolveQuizMultipleTimes", function () {
  it("should be not possible", async function () {
    this.timeout(20000);
    await login(driver);

    await driver
      .find(
        "#quizes > div:nth-child(2) > div.col-4.d-flex.justify-content-center.align-items-center"
      )
      .click();
    expect(await isDialogPresent(driver)).to.equal(false);
    await driver.find("#answer").click();
    for (let i = 0; i < 4; i++) {
      await driver.sendKeys("1");
      await driver.sendKeys(Key.ENTER);
    }

    const saveBtn = await driver.findWait("#save-with-stats", 2000);

    //wait for possible error while fetching data
    await new Promise((resolve) => setTimeout(resolve, 500));
    await driver.sendKeys(Key.ENTER);
    await driver.sendKeys(Key.ENTER);

    await saveBtn.doClick();
    await (
      await driver.findWait(
        "#quizes > div:nth-child(2) > div.col-4.d-flex.justify-content-center.align-items-center",
        2000
      )
    ).click();
    expect(await isDialogPresent(driver)).to.equal(true);
    await logout(driver);
  });
});

async function login(driver: WebDriver) {
  await driver.get("http://localhost:3000");

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
  await driver.get("http://localhost:3000");
  await (
    await driver.find("body > form > input[type=submit]:nth-child(2)")
  ).click();
  await driver.findWait("#sign-up", 2000);
}

async function isDialogPresent(driver: WebDriver): Promise<boolean> {
  try {
    await driver.getTitle();
    return false;
  } catch (_) {
    // Modal dialog showed
    return true;
  }
}
